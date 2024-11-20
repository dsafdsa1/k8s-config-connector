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

package container

import (
	krm "github.com/GoogleCloudPlatform/k8s-config-connector/apis/container/v1beta1"
	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/container/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func int64_FromProto(mapCtx *direct.MapContext, in *pb.MaxPodsConstraint) *int64 {
	if in == nil {
		return nil
	}
	return direct.LazyPtr(in.GetMaxPodsPerNode())
}

func int64_ToProto(mapCtx *direct.MapContext, in *int64) *pb.MaxPodsConstraint {
	if in == nil {
		return nil
	}
	out := &pb.MaxPodsConstraint{
		MaxPodsPerNode: direct.ValueOf(in),
	}
	return out
}

func BlueGreenSettings_StandardRolloutPolicy_BatchPercentage_ToProto(mapCtx *direct.MapContext, in *float32) *pb.BlueGreenSettings_StandardRolloutPolicy_BatchPercentage {
	if in == nil {
		return nil
	}
	out := &pb.BlueGreenSettings_StandardRolloutPolicy_BatchPercentage{
		BatchPercentage: direct.ValueOf(in),
	}
	return out
}

func BlueGreenSettings_StandardRolloutPolicy_BatchNodeCount_ToProto(mapCtx *direct.MapContext, in *int32) *pb.BlueGreenSettings_StandardRolloutPolicy_BatchNodeCount {
	if in == nil {
		return nil
	}
	out := &pb.BlueGreenSettings_StandardRolloutPolicy_BatchNodeCount{
		BatchNodeCount: direct.ValueOf(in),
	}
	return out
}

func NodePool_UpgradeSettings_Strategy_ToProto(mapCtx *direct.MapContext, in *string) *pb.NodePoolUpdateStrategy {
	if in == nil {
		return nil
	}
	out := direct.Enum_ToProto[pb.NodePoolUpdateStrategy](mapCtx, in)
	return &out
}

func BoolValue_ToProto(mapCtx *direct.MapContext, in *krm.BoolValue) *wrapperspb.BoolValue {
	if in == nil {
		return nil
	}
	out := &wrapperspb.BoolValue{
		Value: direct.ValueOf(in.Value),
	}
	return out
}

func BoolValue_FromProto(mapCtx *direct.MapContext, in *wrapperspb.BoolValue) *krm.BoolValue {
	if in == nil {
		return nil
	}
	out := &krm.BoolValue{
		Value: direct.LazyPtr(in.Value),
	}
	return out
}

func HostMaintenancePolicy_MaintenanceInterval_ToProto(mapCtx *direct.MapContext, in *string) *pb.HostMaintenancePolicy_MaintenanceInterval {
	if in == nil {
		return nil
	}
	out := direct.Enum_ToProto[pb.HostMaintenancePolicy_MaintenanceInterval](mapCtx, in)
	return &out
}

func GPUDriverInstallationConfig_GpuDriverVersion_ToProto(mapCtx *direct.MapContext, in *string) *pb.GPUDriverInstallationConfig_GPUDriverVersion {
	if in == nil {
		return nil
	}
	out := direct.Enum_ToProto[pb.GPUDriverInstallationConfig_GPUDriverVersion](mapCtx, in)
	return &out
}

func GPUSharingConfig_GpuSharingStrategy_ToProto(mapCtx *direct.MapContext, in *string) *pb.GPUSharingConfig_GPUSharingStrategy {
	if in == nil {
		return nil
	}
	out := direct.Enum_ToProto[pb.GPUSharingConfig_GPUSharingStrategy](mapCtx, in)
	return &out
}

func NodeAffinities_FromProto(mapCtx *direct.MapContext, in *pb.SoleTenantConfig_NodeAffinity) *krm.SoleTenantConfig_NodeAffinity {
	if in == nil {
		return nil
	}
	out := &krm.SoleTenantConfig_NodeAffinity{
		Key:      direct.LazyPtr(in.GetKey()),
		Operator: direct.Enum_FromProto(mapCtx, in.GetOperator()),
	}

	if in.GetValues() != nil {
		out.Values = make([]string, len(in.GetValues()))
		copy(out.Values, in.GetValues())
	}

	return out
}

func NodeAffinity_ToProto(mapCtx *direct.MapContext, in *krm.SoleTenantConfig_NodeAffinity) *pb.SoleTenantConfig_NodeAffinity {
	if in == nil {
		return nil
	}
	out := &pb.SoleTenantConfig_NodeAffinity{
		Key:      direct.ValueOf(in.Key),
		Operator: direct.Enum_ToProto[pb.SoleTenantConfig_NodeAffinity_Operator](mapCtx, in.Operator),
	}
	if in.Values != nil {
		out.Values = make([]string, len(in.Values))
		copy(out.Values, in.Values)
	}
	return out
}
