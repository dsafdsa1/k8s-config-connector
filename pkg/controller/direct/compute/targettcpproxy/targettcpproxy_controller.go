// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package targettcpproxy

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/api/option"

	gcp "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/compute/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/config"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/directbase"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct/registry"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const ctrlName = "firewallpolicyrule-controller"

func init() {
	registry.RegisterModel(krm.ComputeTargetTCPProxyGVK, NewtargetTCPProxyModel)
}

func NewtargetTCPProxyModel(ctx context.Context, config *config.ControllerConfig) (directbase.Model, error) {
	return &targetTCPProxyModel{config: config}, nil
}

type targetTCPProxyModel struct {
	config *config.ControllerConfig
}

// model implements the Model interface.
var _ directbase.Model = &targetTCPProxyModel{}

type targetTCPProxyAdapter struct {
	id                     *krm.ComputeTargetTCPProxyRef
	targetTcpProxiesClient *gcp.TargetTcpProxiesClient
	desired                *krm.ComputeTargetTCPProxy
	actual                 *computepb.TargetTcpProxy
	reader                 client.Reader
}

var _ directbase.Adapter = &targetTCPProxyAdapter{}

func (m *targetTCPProxyModel) client(ctx context.Context) (*gcp.TargetTcpProxiesClient, error) {
	var opts []option.ClientOption
	opts, err := m.config.RESTClientOptions()
	if err != nil {
		return nil, err
	}
	gcpClient, err := gcp.NewTargetTcpProxiesRESTClient(ctx, opts...)
	if err != nil {
		return nil, fmt.Errorf("building TargetTcpProxy client: %w", err)
	}
	return gcpClient, err
}

func (m *targetTCPProxyModel) AdapterForObject(ctx context.Context, reader client.Reader, u *unstructured.Unstructured) (directbase.Adapter, error) {
	obj := &krm.ComputeTargetTCPProxy{}
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, &obj); err != nil {
		return nil, fmt.Errorf("error converting to %T: %w", obj, err)
	}

	// Set label managed-by-cnrm: true
	obj.ObjectMeta.Labels["managed-by-cnrm"] = "true"

	computeTargetTCPProxyRef, err := krm.NewComputeTargetTCPProxyRef(ctx, reader, obj, u)
	if err != nil {
		return nil, err
	}

	targetTCPProxyAdapter := &targetTCPProxyAdapter{
		id:      computeTargetTCPProxyRef,
		desired: obj,
		reader:  reader,
	}

	// Get GCP client
	gcpClient, err := m.client(ctx)
	if err != nil {
		return nil, fmt.Errorf("building gcp client: %w", err)
	}
	targetTCPProxyAdapter.targetTcpProxiesClient = gcpClient

	return targetTCPProxyAdapter, nil
}

func (m *targetTCPProxyModel) AdapterForURL(ctx context.Context, url string) (directbase.Adapter, error) {
	// TODO: Support URLs
	return nil, nil
}

func (a *targetTCPProxyAdapter) Find(ctx context.Context) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("getting ComputeTargetTCPProxy", "name", a.id.External)

	targetTCPProxy, err := a.get(ctx)
	if err != nil {
		if direct.IsNotFound(err) {
			return false, nil
		}
		return false, fmt.Errorf("getting ComputeTargetTCPProxy %q: %w", a.id.External, err)
	}
	a.actual = targetTCPProxy
	return true, nil
}

func (a *targetTCPProxyAdapter) Create(ctx context.Context, createOp *directbase.CreateOperation) error {
	var err error

	err = resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("creating ComputeTargetTCPProxy", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()

	targetTCPProxy := ComputeTargetTCPProxySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	parent, err := a.id.Parent()
	if err != nil {
		return fmt.Errorf("get ComputeTargetTCPProxy parent %s: %w", a.id.External, err)
	}

	tokens := strings.Split(a.id.External, "/")
	targetTCPProxy.Name = direct.LazyPtr(tokens[len(tokens)-1])

	req := &computepb.InsertTargetTcpProxyRequest{
		Project:                parent.ProjectID,
		TargetTcpProxyResource: targetTCPProxy,
	}
	op, err := a.targetTcpProxiesClient.Insert(ctx, req)

	if err != nil {
		return fmt.Errorf("creating ComputeTargetTCPProxy %s: %w", a.id.External, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeTargetTCPProxy %s create failed: %w", a.id.External, err)
		}
	}
	log.V(2).Info("successfully created ComputeTargetTCPProxy", "name", a.id.External)

	// Get the created resource
	created := &computepb.TargetTcpProxy{}
	created, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeTargetTCPProxy %s: %w", a.id.External, err)
	}

	status := &krm.ComputeTargetTCPProxyStatus{}
	status = ComputeTargetTCPProxyStatus_FromProto(mapCtx, created)

	parent, err = a.id.Parent()
	if err != nil {
		return err
	}

	externalRef := parent.String() + "/targetTcpProxies/" + direct.ValueOf(created.Name)
	status.ExternalRef = &externalRef
	return createOp.UpdateStatus(ctx, status, nil)
}

func (a *targetTCPProxyAdapter) Update(ctx context.Context, updateOp *directbase.UpdateOperation) error {
	var err error

	err = resolveDependencies(ctx, a.reader, a.desired)
	if err != nil {
		return err
	}

	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("updating ComputeTargetTCPProxy", "name", a.id.External)
	mapCtx := &direct.MapContext{}

	desired := a.desired.DeepCopy()
	targetTCPProxy := ComputeTargetTCPProxySpec_ToProto(mapCtx, &desired.Spec)
	if mapCtx.Err() != nil {
		return mapCtx.Err()
	}

	op := &gcp.Operation{}
	updated := &computepb.TargetTcpProxy{}

	parent, err := a.id.Parent()
	if err != nil {
		return fmt.Errorf("get ComputeTargetTCPProxy parent %s: %w", a.id.External, err)
	}
	tokens := strings.Split(a.id.External, "/")
	targetTCPProxy.Name = direct.LazyPtr(tokens[len(tokens)-1])

	if !reflect.DeepEqual(targetTCPProxy.Service, a.actual.Service) {
		setBackendServiceReq := &computepb.SetBackendServiceTargetTcpProxyRequest{
			Project: parent.ProjectID,
			TargetTcpProxiesSetBackendServiceRequestResource: &computepb.TargetTcpProxiesSetBackendServiceRequest{Service: targetTCPProxy.Service},
			TargetTcpProxy: tokens[len(tokens)-1],
		}
		op, err = a.targetTcpProxiesClient.SetBackendService(ctx, setBackendServiceReq)
		if err != nil {
			return fmt.Errorf("updating ComputeTargetTCPProxy backend service %s: %w", a.id.External, err)
		}
	}
	if !reflect.DeepEqual(targetTCPProxy.ProxyHeader, a.actual.ProxyHeader) {
		setProxyHeaderReq := &computepb.SetProxyHeaderTargetTcpProxyRequest{
			Project: parent.ProjectID,
			TargetTcpProxiesSetProxyHeaderRequestResource: &computepb.TargetTcpProxiesSetProxyHeaderRequest{ProxyHeader: targetTCPProxy.ProxyHeader},
			TargetTcpProxy: tokens[len(tokens)-1],
		}
		op, err = a.targetTcpProxiesClient.SetProxyHeader(ctx, setProxyHeaderReq)
		if err != nil {
			return fmt.Errorf("updating ComputeTargetTCPProxy proxy header %s: %w", a.id.External, err)
		}
	}

	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return fmt.Errorf("waiting ComputeTargetTCPProxy %s update failed: %w", a.id.External, err)
		}
	}
	log.V(2).Info("successfully updated ComputeTargetTCPProxy", "name", a.id.External)

	// Get the updated resource
	updated, err = a.get(ctx)
	if err != nil {
		return fmt.Errorf("getting ComputeTargetTCPProxy %s: %w", a.id.External, err)
	}

	status := &krm.ComputeTargetTCPProxyStatus{}
	status = ComputeTargetTCPProxyStatus_FromProto(mapCtx, updated)
	return updateOp.UpdateStatus(ctx, status, nil)
}

func (a *targetTCPProxyAdapter) Export(ctx context.Context) (*unstructured.Unstructured, error) {
	if a.actual == nil {
		return nil, fmt.Errorf("targetTcpProxy %s not found", a.id.External)
	}

	mc := &direct.MapContext{}
	spec := ComputeTargetTCPProxySpec_FromProto(mc, a.actual)
	specObj, err := runtime.DefaultUnstructuredConverter.ToUnstructured(spec)
	if err != nil {
		return nil, fmt.Errorf("error converting targetTcpProxy spec to unstructured: %w", err)
	}

	u := &unstructured.Unstructured{
		Object: make(map[string]interface{}),
	}
	u.SetGroupVersionKind(krm.ComputeTargetTCPProxyGVK)

	if err := unstructured.SetNestedField(u.Object, specObj, "spec"); err != nil {
		return nil, fmt.Errorf("setting spec: %w", err)
	}

	return u, nil
}

// Delete implements the Adapter interface.
func (a *targetTCPProxyAdapter) Delete(ctx context.Context, deleteOp *directbase.DeleteOperation) (bool, error) {
	log := klog.FromContext(ctx).WithName(ctrlName)
	log.V(2).Info("deleting ComputeTargetTcpProxy", "name", a.id.External)

	parent, err := a.id.Parent()
	if err != nil {
		return false, fmt.Errorf("get ComputeTargetTcpProxy parent %s: %w", a.id.External, err)
	}

	tokens := strings.Split(a.id.External, "/")

	delReq := &computepb.DeleteTargetTcpProxyRequest{
		Project:        parent.ProjectID,
		TargetTcpProxy: tokens[len(tokens)-1],
	}
	op, err := a.targetTcpProxiesClient.Delete(ctx, delReq)

	if err != nil {
		return false, fmt.Errorf("deleting ComputeTargetTcpProxy %s: %w", a.id.External, err)
	}
	if !op.Done() {
		err = op.Wait(ctx)
		if err != nil {
			return false, fmt.Errorf("waiting ComputeTargetTcpProxy %s delete failed: %w", a.id.External, err)
		}
	}
	log.V(2).Info("successfully deleted ComputeTargetTcpProxy", "name", a.id.External)
	return true, nil
}

func (a *targetTCPProxyAdapter) get(ctx context.Context) (*computepb.TargetTcpProxy, error) {
	parent, err := a.id.Parent()
	if err != nil {
		return nil, fmt.Errorf("get ComputeTargetTcpProxy parent %s: %w", a.id.External, err)
	}

	tokens := strings.Split(a.id.External, "/")

	getReq := &computepb.GetTargetTcpProxyRequest{
		Project:        parent.ProjectID,
		TargetTcpProxy: tokens[len(tokens)-1],
	}
	return a.targetTcpProxiesClient.Get(ctx, getReq)
}
