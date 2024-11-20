// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var ContainerNodePoolGVK = GroupVersion.WithKind("ContainerNodePool")

// ContainerNodePoolSpec defines the desired state of ContainerNodePool
// +kcc:proto=google.container.v1beta1.NodePool
type ContainerNodePoolSpec struct {
	/* Configuration required by cluster autoscaler to adjust the size of the node pool to the current cluster usage. To disable autoscaling, set minNodeCount and maxNodeCount to 0. */
	// +optional
	Autoscaling *NodePoolAutoscaling `json:"autoscaling,omitempty"`

	ClusterRef v1alpha1.ResourceRef `json:"clusterRef"`

	/* Immutable. The initial number of nodes for the pool. In regional or multi-zonal clusters, this is the number of nodes per zone. Changing this will force recreation of the resource. */
	// +optional
	InitialNodeCount *int64 `json:"initialNodeCount,omitempty"`

	/* Immutable. The location (region or zone) of the cluster. */
	Location string `json:"location"`

	/* Node management configuration, wherein auto-repair and auto-upgrade is configured. */
	// +optional
	Management *NodeManagement `json:"management,omitempty"`

	/* Immutable. The maximum number of pods per node in this node pool. Note that this does not work on node pools which are "route-based" - that is, node pools belonging to clusters that do not have IP Aliasing enabled. */
	// +optional
	MaxPodsPerNode *int64 `json:"maxPodsPerNode,omitempty"`

	/* Immutable. Creates a unique name for the node pool beginning with the specified prefix. Conflicts with name. */
	// +optional
	NamePrefix *string `json:"namePrefix,omitempty"`

	/* Networking configuration for this NodePool. If specified, it overrides the cluster-level defaults. */
	// +optional
	NetworkConfig *NodeNetworkConfig `json:"networkConfig,omitempty"`

	/* Immutable. The configuration of the nodepool. */
	// +optional
	NodeConfig *NodeConfig `json:"nodeConfig,omitempty"`

	/* The number of nodes per instance group. This field can be used to update the number of nodes per instance group but should not be used alongside autoscaling. */
	// +optional
	NodeCount *int64 `json:"nodeCount,omitempty"`

	/* The list of zones in which the node pool's nodes should be located. Nodes must be in the region of their regional cluster or in the same region as their cluster's zone for zonal clusters. If unspecified, the cluster-level node_locations will be used. */
	// +optional
	NodeLocations []string `json:"nodeLocations,omitempty"`

	/* Immutable. Specifies the node placement policy. */
	// +optional
	PlacementPolicy *NodePool_PlacementPolicy `json:"placementPolicy,omitempty"`

	// +kubebuilder:validation:XValidation:rule="self == oldSelf",message="ResourceID field is immutable"
	// Immutable.
	// The ContainerNodePool name. If not given, the metadata.name will be used.
	ResourceID *string `json:"resourceID,omitempty"`

	/* Specify node upgrade settings to change how many nodes GKE attempts to upgrade at once. The number of nodes upgraded simultaneously is the sum of max_surge and max_unavailable. The maximum number of nodes upgraded simultaneously is limited to 20. */
	// +optional
	UpgradeSettings *NodePool_UpgradeSettings `json:"upgradeSettings,omitempty"`

	// +optional
	Version *string `json:"version,omitempty"`
}

// ContainerNodePoolStatus defines the config connector machine state of ContainerNodePool
type ContainerNodePoolStatus struct {
	/* Conditions represent the latest available observations of the
	   object's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`

	// ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource.
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* NOTYET:
	// A unique specifier for the ContainerNodePool resource in GCP.
	ExternalRef *string `json:"externalRef,omitempty"`
	*/

	// ObservedState is the state of the resource as most recently observed in GCP.
	ObservedState *ContainerNodePoolObservedState `json:"observedState,omitempty"`

	/* The resource URLs of the managed instance groups associated with this node pool. */
	// +optional
	InstanceGroupUrls []string `json:"instanceGroupUrls,omitempty"`
	/* List of instance group URLs which have been assigned to this node pool. */
	// +optional
	ManagedInstanceGroupUrls []string `json:"managedInstanceGroupUrls,omitempty"`
	// +optional
	Operation *string `json:"operation,omitempty"`
}

// ContainerNodePoolObservedState is the state of the ContainerNodePool resource as most recently observed in GCP.
type ContainerNodePoolObservedState struct {
	// +optional
	Version *string `json:"version,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpcontainernodepool;gcpcontainernodepools
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// ContainerNodePool is the Schema for the ContainerNodePool API
// +k8s:openapi-gen=true
type ContainerNodePool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec   ContainerNodePoolSpec   `json:"spec,omitempty"`
	Status ContainerNodePoolStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ContainerNodePoolList contains a list of ContainerNodePool
type ContainerNodePoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerNodePool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerNodePool{}, &ContainerNodePoolList{})
}
