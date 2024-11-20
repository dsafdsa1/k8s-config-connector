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

package v1beta1

import "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"

/*
// +kcc:proto=google.container.v1beta1.AcceleratorConfig
type AcceleratorConfig struct {
	// The number of the accelerator cards exposed to an instance.
	AcceleratorCount *int64 `json:"acceleratorCount,omitempty"`

	// The accelerator type resource name. List of supported accelerators
	//  [here](https://cloud.google.com/compute/docs/gpus)
	AcceleratorType *string `json:"acceleratorType,omitempty"`

	// Size of partitions to create on the GPU. Valid values are described in the
	//  NVIDIA [mig user
	//  guide](https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning).
	GpuPartitionSize *string `json:"gpuPartitionSize,omitempty"`

	// The number of time-shared GPU resources to expose for each physical GPU.
	MaxTimeSharedClientsPerGpu *int64 `json:"maxTimeSharedClientsPerGpu,omitempty"`

	// The configuration for GPU sharing options.
	GpuSharingConfig *GPUSharingConfig `json:"gpuSharingConfig,omitempty"`

	// The configuration for auto installation of GPU driver.
	GpuDriverInstallationConfig *GPUDriverInstallationConfig `json:"gpuDriverInstallationConfig,omitempty"`
}
*/

type NodepoolGuestAccelerator struct {
	/* ADDED_TO_MATCH_CRD: */

	/* Immutable. The number of the accelerator cards exposed to an instance. */
	Count int64 `json:"count"`

	/* Immutable. Configuration for auto installation of GPU driver. */
	// +optional
	GpuDriverInstallationConfig *GPUDriverInstallationConfig `json:"gpuDriverInstallationConfig,omitempty"`

	/* Immutable. Size of partitions to create on the GPU. Valid values are described in the NVIDIA mig user guide (https://docs.nvidia.com/datacenter/tesla/mig-user-guide/#partitioning). */
	// +optional
	GpuPartitionSize *string `json:"gpuPartitionSize,omitempty"`

	/* Immutable. Configuration for GPU sharing. */
	// +optional
	GpuSharingConfig *GPUSharingConfig `json:"gpuSharingConfig,omitempty"`

	/* Immutable. The accelerator type resource name. */
	Type string `json:"type"`
}

// +kcc:proto=google.container.v1beta1.AdditionalNodeNetworkConfig
type AdditionalNodeNetworkConfig struct {
	// Name of the VPC where the additional interface belongs
	NetworkRef *v1alpha1.ResourceRef `json:"networkRef,omitempty"`
	/* ADDED_TO_MATCH_CRD:

	/* REMOVED_TO_MATCH_CRD:
	Network *string `json:"network,omitempty"`
	*/

	// Name of the subnetwork where the additional interface belongs
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`
	/* ADDED_TO_MATCH_CRD:

	/* REMOVED_TO_MATCH_CRD:
	Subnetwork *string `json:"subnetwork,omitempty"`
	*/
}

// +kcc:proto=google.container.v1beta1.AdditionalPodNetworkConfig
type AdditionalPodNetworkConfig struct {
	// Name of the subnetwork where the additional pod network belongs.
	SubnetworkRef *v1alpha1.ResourceRef `json:"subnetworkRef,omitempty"`
	/* ADDED_TO_MATCH_CRD:

	/* REMOVED_TO_MATCH_CRD:
	Subnetwork *string `json:"subnetwork,omitempty"`
	*/

	// The name of the secondary range on the subnet which provides IP address for
	//  this pod range.
	SecondaryPodRange *string `json:"secondaryPodRange,omitempty"`

	// The maximum number of pods per node which use this pod network.
	MaxPodsPerNode *int64 `json:"maxPodsPerNode,omitempty"`
	/* ADDED_TO_MATCH_CRD:

	/* REMOVED_TO_MATCH_CRD:
	MaxPodsPerNode *MaxPodsConstraint `json:"maxPodsPerNode,omitempty"`
	*/
}

// +kcc:proto=google.container.v1beta1.AdvancedMachineFeatures
type AdvancedMachineFeatures struct {
	// The number of threads per physical core. To disable simultaneous
	//  multithreading (SMT) set this to 1. If unset, the maximum number of threads
	//  supported per core by the underlying processor is assumed.
	ThreadsPerCore *int64 `json:"threadsPerCore,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// Whether or not to enable nested virtualization (defaults to false).
	EnableNestedVirtualization *bool `json:"enableNestedVirtualization,omitempty"`
	*/
}

/* REMOVED_TO_MATCH_CRD:
// +kcc:proto=google.container.v1beta1.AutoUpgradeOptions
type AutoUpgradeOptions struct {
	// Output only. This field is set when upgrades are about to commence
	//  with the approximate start time for the upgrades, in
	//  [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	AutoUpgradeStartTime *string `json:"autoUpgradeStartTime,omitempty"`

	// Output only. This field is set when upgrades are about to commence
	//  with the description of the upgrade.
	Description *string `json:"description,omitempty"`
}
*/

// +kcc:proto=google.container.v1beta1.BestEffortProvisioning
type BestEffortProvisioning struct {
	// When this is enabled, cluster/node pool creations will ignore non-fatal
	//  errors like stockout to best provision as many nodes as possible right now
	//  and eventually bring up all target number of nodes
	Enabled *bool `json:"enabled,omitempty"`

	// Minimum number of nodes to be provisioned to be considered as succeeded,
	//  and the rest of nodes will be provisioned gradually and eventually when
	//  stockout issue has been resolved.
	MinProvisionNodes *int32 `json:"minProvisionNodes,omitempty"`
}

// +kcc:proto=google.container.v1beta1.BlueGreenSettings
type BlueGreenSettings struct {
	// Standard policy for the blue-green upgrade.
	StandardRolloutPolicy *BlueGreenSettings_StandardRolloutPolicy `json:"standardRolloutPolicy,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// Autoscaled policy for cluster autoscaler enabled blue-green upgrade.
	AutoscaledRolloutPolicy *BlueGreenSettings_AutoscaledRolloutPolicy `json:"autoscaledRolloutPolicy,omitempty"`
	*/

	// Time needed after draining entire blue pool. After this period, blue pool
	//  will be cleaned up.
	NodePoolSoakDuration *string `json:"nodePoolSoakDuration,omitempty"`
}

/* REMOVED_TO_MATCH_CRD:
// +kcc:proto=google.container.v1beta1.BlueGreenSettings.AutoscaledRolloutPolicy
type BlueGreenSettings_AutoscaledRolloutPolicy struct {
}
*/

// +kcc:proto=google.container.v1beta1.BlueGreenSettings.StandardRolloutPolicy
type BlueGreenSettings_StandardRolloutPolicy struct {
	// Percentage of the blue pool nodes to drain in a batch.
	//  The range of this field should be (0.0, 1.0].
	BatchPercentage *float32 `json:"batchPercentage,omitempty"`

	// Number of blue nodes to drain in a batch.
	BatchNodeCount *int32 `json:"batchNodeCount,omitempty"`

	// Soak time after each batch gets drained. Default to zero.
	BatchSoakDuration *string `json:"batchSoakDuration,omitempty"`
}

// +kcc:proto=google.container.v1beta1.ConfidentialNodes
type ConfidentialNodes struct {
	// Whether Confidential Nodes feature is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

/* REMOVED_TO_MATCH_CRD:
// +kcc:proto=google.container.v1beta1.ContainerdConfig
type ContainerdConfig struct {
	// PrivateRegistryAccessConfig is used to configure access configuration
	//  for private container registries.
	PrivateRegistryAccessConfig *ContainerdConfig_PrivateRegistryAccessConfig `json:"privateRegistryAccessConfig,omitempty"`
}

// +kcc:proto=google.container.v1beta1.ContainerdConfig.PrivateRegistryAccessConfig
type ContainerdConfig_PrivateRegistryAccessConfig struct {
	// Private registry access is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Private registry access configuration.
	CertificateAuthorityDomainConfig []ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig `json:"certificateAuthorityDomainConfig,omitempty"`
}

// +kcc:proto=google.container.v1beta1.ContainerdConfig.PrivateRegistryAccessConfig.CertificateAuthorityDomainConfig
type ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig struct {
	// List of fully qualified domain names (FQDN).
	//  Specifying port is supported.
	//  Wilcards are NOT supported.
	//  Examples:
	//  - my.customdomain.com
	//  - 10.0.1.2:5000
	Fqdns []string `json:"fqdns,omitempty"`

	// Google Secret Manager (GCP) certificate configuration.
	GcpSecretManagerCertificateConfig *ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig_GCPSecretManagerCertificateConfig `json:"gcpSecretManagerCertificateConfig,omitempty"`
}

// +kcc:proto=google.container.v1beta1.ContainerdConfig.PrivateRegistryAccessConfig.CertificateAuthorityDomainConfig.GCPSecretManagerCertificateConfig
type ContainerdConfig_PrivateRegistryAccessConfig_CertificateAuthorityDomainConfig_GCPSecretManagerCertificateConfig struct {
	// Secret URI, in the form
	//  "projects/$PROJECT_ID/secrets/$SECRET_NAME/versions/$VERSION".
	//  Version can be fixed (e.g. "2") or "latest"
	SecretUri *string `json:"secretUri,omitempty"`
}
*/

// +kcc:proto=google.container.v1beta1.EphemeralStorageConfig
type EphemeralStorageConfig struct {
	// Number of local SSDs to use to back ephemeral storage. Uses NVMe
	//  interfaces. The limit for this value is dependent upon the maximum number
	//  of disk available on a machine per zone. See:
	//  https://cloud.google.com/compute/docs/disks/local-ssd
	//  for more information.
	//
	//  A zero (or unset) value has different meanings depending on machine type
	//  being used:
	//  1. For pre-Gen3 machines, which support flexible numbers of local ssds,
	//  zero (or unset) means to disable using local SSDs as ephemeral storage.
	//  2. For Gen3 machines which dictate a specific number of local ssds, zero
	//  (or unset) means to use the default number of local ssds that goes with
	//  that machine type. For example, for a c3-standard-8-lssd machine, 2 local
	//  ssds would be provisioned. For c3-standard-8 (which doesn't support local
	//  ssds), 0 will be provisioned. See
	//  https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	//  for more info.
	LocalSsdCount *int32 `json:"localSsdCount,omitempty"`
}

// +kcc:proto=google.container.v1beta1.EphemeralStorageLocalSsdConfig
type EphemeralStorageLocalSsdConfig struct {
	// Number of local SSDs to use to back ephemeral storage. Uses NVMe
	//  interfaces.
	//
	//  A zero (or unset) value has different meanings depending on machine type
	//  being used:
	//  1. For pre-Gen3 machines, which support flexible numbers of local ssds,
	//  zero (or unset) means to disable using local SSDs as ephemeral storage. The
	//  limit for this value is dependent upon the maximum number of disk
	//  available on a machine per zone. See:
	//  https://cloud.google.com/compute/docs/disks/local-ssd
	//  for more information.
	//  2. For Gen3 machines which dictate a specific number of local ssds, zero
	//  (or unset) means to use the default number of local ssds that goes with
	//  that machine type. For example, for a c3-standard-8-lssd machine, 2 local
	//  ssds would be provisioned. For c3-standard-8 (which doesn't support local
	//  ssds), 0 will be provisioned. See
	//  https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	//  for more info.
	LocalSsdCount *int32 `json:"localSsdCount,omitempty"`
}

// +kcc:proto=google.container.v1beta1.FastSocket
type FastSocket struct {
	// Whether Fast Socket features are enabled in the node pool.
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1beta1.GPUDriverInstallationConfig
type GPUDriverInstallationConfig struct {
	// Mode for how the GPU driver is installed.
	GpuDriverVersion *string `json:"gpuDriverVersion,omitempty"`
}

// +kcc:proto=google.container.v1beta1.GPUSharingConfig
type GPUSharingConfig struct {
	// The max number of containers that can share a physical GPU.
	MaxSharedClientsPerGpu *int64 `json:"maxSharedClientsPerGpu,omitempty"`

	// The type of GPU sharing strategy to enable on the GPU node.
	GpuSharingStrategy *string `json:"gpuSharingStrategy,omitempty"`
}

// +kcc:proto=google.container.v1beta1.GcfsConfig
type GcfsConfig struct {
	// Whether to use GCFS.
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1beta1.HostMaintenancePolicy
type HostMaintenancePolicy struct {
	// Specifies the frequency of planned maintenance events.
	MaintenanceInterval *string `json:"maintenanceInterval,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// Strategy that will trigger maintenance on behalf of the customer.
	OpportunisticMaintenanceStrategy *HostMaintenancePolicy_OpportunisticMaintenanceStrategy `json:"opportunisticMaintenanceStrategy,omitempty"`
	*/
}

/* REMOVED_TO_MATCH_CRD:
// +kcc:proto=google.container.v1beta1.HostMaintenancePolicy.OpportunisticMaintenanceStrategy
type HostMaintenancePolicy_OpportunisticMaintenanceStrategy struct {
	// The amount of time that a node can remain idle (no customer owned
	//  workloads running), before triggering maintenance.
	NodeIdleTimeWindow *string `json:"nodeIdleTimeWindow,omitempty"`

	// The window of time that opportunistic maintenance can run. Example: A
	//  setting of 14 days implies that opportunistic maintenance can only be ran
	//  in the 2 weeks leading up to the scheduled maintenance date. Setting 28
	//  days allows opportunistic maintenance to run at any time in the scheduled
	//  maintenance window (all `PERIODIC` maintenance is set 28 days in
	//  advance).
	MaintenanceAvailabilityWindow *string `json:"maintenanceAvailabilityWindow,omitempty"`

	// The minimum nodes required to be available in a pool. Blocks maintenance
	//  if it would cause the number of running nodes to dip below this value.
	MinNodesPerPool *int64 `json:"minNodesPerPool,omitempty"`
}
*/

// +kcc:proto=google.container.v1beta1.LinuxNodeConfig
type LinuxNodeConfig struct {
	// The Linux kernel parameters to be applied to the nodes and all pods running
	//  on the nodes.
	//
	//  The following parameters are supported.
	//
	//  net.core.busy_poll
	//  net.core.busy_read
	//  net.core.netdev_max_backlog
	//  net.core.rmem_max
	//  net.core.wmem_default
	//  net.core.wmem_max
	//  net.core.optmem_max
	//  net.core.somaxconn
	//  net.ipv4.tcp_rmem
	//  net.ipv4.tcp_wmem
	//  net.ipv4.tcp_tw_reuse
	Sysctls map[string]string `json:"sysctls,omitempty"`

	// cgroup_mode specifies the cgroup mode to be used on the node.
	CgroupMode *string `json:"cgroupMode,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// Optional. Amounts for 2M and 1G hugepages
	Hugepages *LinuxNodeConfig_HugepagesConfig `json:"hugepages,omitempty"`
	*/
}

/* REMOVED_TO_MATCH_CRD:
// +kcc:proto=google.container.v1beta1.LinuxNodeConfig.HugepagesConfig
type LinuxNodeConfig_HugepagesConfig struct {
	// Optional. Amount of 2M hugepages
	HugepageSize2m *int32 `json:"hugepageSize2m,omitempty"`

	// Optional. Amount of 1G hugepages
	HugepageSize1g *int32 `json:"hugepageSize1g,omitempty"`
}
*/

// +kcc:proto=google.container.v1beta1.LocalNvmeSsdBlockConfig
type LocalNvmeSsdBlockConfig struct {
	// Number of local NVMe SSDs to use.  The limit for this value is dependent
	//  upon the maximum number of disk available on a machine per zone. See:
	//  https://cloud.google.com/compute/docs/disks/local-ssd
	//  for more information.
	//
	//  A zero (or unset) value has different meanings depending on machine type
	//  being used:
	//  1. For pre-Gen3 machines, which support flexible numbers of local ssds,
	//  zero (or unset) means to disable using local SSDs as ephemeral storage.
	//  2. For Gen3 machines which dictate a specific number of local ssds, zero
	//  (or unset) means to use the default number of local ssds that goes with
	//  that machine type. For example, for a c3-standard-8-lssd machine, 2 local
	//  ssds would be provisioned. For c3-standard-8 (which doesn't support local
	//  ssds), 0 will be provisioned. See
	//  https://cloud.google.com/compute/docs/disks/local-ssd#choose_number_local_ssds
	//  for more info.
	LocalSsdCount *int32 `json:"localSsdCount,omitempty"`
}

// +kcc:proto=google.container.v1beta1.LoggingVariantConfig
type LoggingVariantConfig struct {
	// Logging variant deployed on nodes.
	Variant *string `json:"variant,omitempty"`
}

/*
// +kcc:proto=google.container.v1beta1.MaxPodsConstraint
type MaxPodsConstraint struct {
	// Constraint enforced on the max num of pods per node.
	MaxPodsPerNode *int64 `json:"maxPodsPerNode,omitempty"`
}
*/

// +kcc:proto=google.container.v1beta1.NodeConfig
type NodeConfig struct {
	// The name of a Google Compute Engine [machine
	//  type](https://cloud.google.com/compute/docs/machine-types).
	//
	//  If unspecified, the default machine type is
	//  `e2-medium`.
	MachineType *string `json:"machineType,omitempty"`

	// Size of the disk attached to each node, specified in GB.
	//  The smallest allowed disk size is 10GB.
	//
	//  If unspecified, the default disk size is 100GB.
	DiskSizeGb *int32 `json:"diskSizeGb,omitempty"`

	// The set of Google API scopes to be made available on all of the
	//  node VMs under the "default" service account.
	//
	//  The following scopes are recommended, but not required, and by default are
	//  not included:
	//
	//  * `https://www.googleapis.com/auth/compute` is required for mounting
	//  persistent storage on your nodes.
	//  * `https://www.googleapis.com/auth/devstorage.read_only` is required for
	//  communicating with **gcr.io**
	//  (the [Google Container
	//  Registry](https://cloud.google.com/container-registry/)).
	//
	//  If unspecified, no scopes are added, unless Cloud Logging or Cloud
	//  Monitoring are enabled, in which case their required scopes will be added.
	OauthScopes []string `json:"oauthScopes,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// The Google Cloud Platform Service Account to be used by the node VMs.
	//  Specify the email address of the Service Account; otherwise, if no Service
	//  Account is specified, the "default" service account is used.
	ServiceAccount *string `json:"serviceAccount,omitempty"`
	*/

	// +optional
	ServiceAccountRef *v1alpha1.ResourceRef `json:"serviceAccountRef,omitempty"`
	/* ADDED_TO_MATCH_CRD: */

	// The metadata key/value pairs assigned to instances in the cluster.
	//
	//  Keys must conform to the regexp `[a-zA-Z0-9-_]+` and be less than 128 bytes
	//  in length. These are reflected as part of a URL in the metadata server.
	//  Additionally, to avoid ambiguity, keys must not conflict with any other
	//  metadata keys for the project or be one of the reserved keys:
	//
	//   - "cluster-location"
	//   - "cluster-name"
	//   - "cluster-uid"
	//   - "configure-sh"
	//   - "containerd-configure-sh"
	//   - "enable-oslogin"
	//   - "gci-ensure-gke-docker"
	//   - "gci-metrics-enabled"
	//   - "gci-update-strategy"
	//   - "instance-template"
	//   - "kube-env"
	//   - "startup-script"
	//   - "user-data"
	//   - "disable-address-manager"
	//   - "windows-startup-script-ps1"
	//   - "common-psm1"
	//   - "k8s-node-setup-psm1"
	//   - "install-ssh-psm1"
	//   - "user-profile-psm1"
	//
	//  Values are free-form strings, and only have meaning as interpreted by
	//  the image running in the instance. The only restriction placed on them is
	//  that each value's size must be less than or equal to 32 KB.
	//
	//  The total size of all keys and values must be less than 512 KB.
	Metadata map[string]string `json:"metadata,omitempty"`

	// The image type to use for this node. Note that for a given image type,
	//  the latest version of it will be used. Please see
	//  https://cloud.google.com/kubernetes-engine/docs/concepts/node-images for
	//  available image types.
	ImageType *string `json:"imageType,omitempty"`

	// The map of Kubernetes labels (key/value pairs) to be applied to each node.
	//  These will added in addition to any default label(s) that
	//  Kubernetes may apply to the node.
	//  In case of conflict in label keys, the applied set may differ depending on
	//  the Kubernetes version -- it's best to assume the behavior is undefined
	//  and conflicts should be avoided.
	//  For more information, including usage and the valid values, see:
	//  https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	Labels map[string]string `json:"labels,omitempty"`

	// The number of local SSD disks to be attached to the node.
	//
	//  The limit for this value is dependent upon the maximum number of
	//  disks available on a machine per zone. See:
	//  https://cloud.google.com/compute/docs/disks/local-ssd
	//  for more information.
	LocalSsdCount *int32 `json:"localSsdCount,omitempty"`

	// The list of instance tags applied to all nodes. Tags are used to identify
	//  valid sources or targets for network firewalls and are specified by
	//  the client during cluster or node pool creation. Each tag within the list
	//  must comply with RFC1035.
	Tags []string `json:"tags,omitempty"`

	// Whether the nodes are created as preemptible VM instances. See:
	//  https://cloud.google.com/compute/docs/instances/preemptible for more
	//  information about preemptible VM instances.
	Preemptible *bool `json:"preemptible,omitempty"`

	/* Immutable. List of the type and count of accelerator cards attached to the instance. */
	// +optional
	GuestAccelerator []NodepoolGuestAccelerator `json:"guestAccelerator,omitempty"`
	/* ADDED_TO_MATCH_CRD: */

	/* REMOVED_TO_MATCH_CRD:
	// A list of hardware accelerators to be attached to each node.
	//  See https://cloud.google.com/compute/docs/gpus for more information about
	//  support for GPUs.
	Accelerators []AcceleratorConfig `json:"accelerators,omitempty"`
	*/

	// Sandbox configuration for this node.
	SandboxConfig *SandboxConfig `json:"sandboxConfig,omitempty"`

	/* Immutable. Setting this field will assign instances
	of this pool to run on the specified node group. This is useful
	for running workloads on sole tenant nodes. */
	// +optional
	NodeGroupRef *v1alpha1.ResourceRef `json:"nodeGroupRef,omitempty"`
	/* ADDED_TO_MATCH_CRD: */

	/* REMOVED_TO_MATCH_CRD:
	// Setting this field will assign instances of this
	//  pool to run on the specified node group. This is useful for running
	//  workloads on [sole tenant
	//  nodes](https://cloud.google.com/compute/docs/nodes/sole-tenant-nodes).
	NodeGroup *string `json:"nodeGroup,omitempty"`
	*/

	// The optional reservation affinity. Setting this field will apply
	//  the specified [Zonal Compute
	//  Reservation](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources)
	//  to this node pool.
	ReservationAffinity *ReservationAffinity `json:"reservationAffinity,omitempty"`

	// Type of the disk attached to each node (e.g. 'pd-standard', 'pd-ssd' or
	//  'pd-balanced')
	//
	//  If unspecified, the default disk type is 'pd-standard'
	DiskType *string `json:"diskType,omitempty"`

	// Minimum CPU platform to be used by this instance. The instance may be
	//  scheduled on the specified or newer CPU platform. Applicable values are the
	//  friendly names of CPU platforms, such as
	//  `minCpuPlatform: "Intel Haswell"` or
	//  `minCpuPlatform: "Intel Sandy Bridge"`. For more
	//  information, read [how to specify min CPU
	//  platform](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform).
	MinCpuPlatform *string `json:"minCpuPlatform,omitempty"`

	// The workload metadata configuration for this node.
	WorkloadMetadataConfig *WorkloadMetadataConfig `json:"workloadMetadataConfig,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// List of kubernetes taints to be applied to each node.
	//
	//  For more information, including usage and the valid values, see:
	//  https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
	Taints []NodeTaint `json:"taints,omitempty"`
	*/

	// +optional
	Taint []NodeTaint `json:"taint,omitempty"`
	/* added_TO_MATCH_CRD:*/

	// +optional
	BootDiskKMSCryptoKeyRef *v1alpha1.ResourceRef `json:"bootDiskKMSCryptoKeyRef,omitempty"`
	/* ADDED_TO_MATCH_CRD: */

	/* REMOVED_TO_MATCH_CRD:
	// The Customer Managed Encryption Key used to encrypt the boot disk attached
	//  to each node in the node pool. This should be of the form
	//  projects/[KEY_PROJECT_ID]/locations/[LOCATION]/keyRings/[RING_NAME]/cryptoKeys/[KEY_NAME].
	//  For more information about protecting resources with Cloud KMS Keys please
	//  see:
	//  https://cloud.google.com/compute/docs/disks/customer-managed-encryption
	BootDiskKmsKey *string `json:"bootDiskKmsKey,omitempty"`
	*/

	// Shielded Instance options.
	ShieldedInstanceConfig *ShieldedInstanceConfig `json:"shieldedInstanceConfig,omitempty"`

	// Parameters that can be configured on Linux nodes.
	LinuxNodeConfig *LinuxNodeConfig `json:"linuxNodeConfig,omitempty"`

	// Node kubelet configs.
	KubeletConfig *NodeKubeletConfig `json:"kubeletConfig,omitempty"`

	// Parameters for the ephemeral storage filesystem.
	//  If unspecified, ephemeral storage is backed by the boot disk.
	EphemeralStorageConfig *EphemeralStorageConfig `json:"ephemeralStorageConfig,omitempty"`

	// GCFS (Google Container File System) configs.
	GcfsConfig *GcfsConfig `json:"gcfsConfig,omitempty"`

	// Advanced features for the Compute Engine VM.
	AdvancedMachineFeatures *AdvancedMachineFeatures `json:"advancedMachineFeatures,omitempty"`

	// Enable or disable gvnic on the node pool.
	Gvnic *VirtualNIC `json:"gvnic,omitempty"`

	// Spot flag for enabling Spot VM, which is a rebrand of
	//  the existing preemptible flag.
	Spot *bool `json:"spot,omitempty"`

	// Confidential nodes config.
	//  All the nodes in the node pool will be Confidential VM once enabled.
	ConfidentialNodes *ConfidentialNodes `json:"confidentialNodes,omitempty"`

	// Enable or disable NCCL fast socket for the node pool.
	FastSocket *FastSocket `json:"fastSocket,omitempty"`

	// The resource labels for the node pool to use to annotate any related
	//  Google Compute Engine resources.
	ResourceLabels map[string]string `json:"resourceLabels,omitempty"`

	/* Type of logging agent that is used as the default value for node pools in the cluster. Valid values include DEFAULT and MAX_THROUGHPUT. */
	// +optional
	LoggingVariant *string `json:"loggingVariant,omitempty"`
	/* ADDED_TO_MATCH_CRD:

	/* REMOVED_TO_MATCH_CRD:
	// Logging configuration.
	LoggingConfig *NodePoolLoggingConfig `json:"loggingConfig,omitempty"`
	*/

	/* REMOVED_TO_MATCH_CRD:
	// Parameters that can be configured on Windows nodes.
	WindowsNodeConfig *WindowsNodeConfig `json:"windowsNodeConfig,omitempty"`
	*/

	// Parameters for using raw-block Local NVMe SSDs.
	LocalNvmeSsdBlockConfig *LocalNvmeSsdBlockConfig `json:"localNvmeSsdBlockConfig,omitempty"`

	// Parameters for the node ephemeral storage using Local SSDs.
	//  If unspecified, ephemeral storage is backed by the boot disk.
	//  This field is functionally equivalent to the ephemeral_storage_config
	EphemeralStorageLocalSsdConfig *EphemeralStorageLocalSsdConfig `json:"ephemeralStorageLocalSsdConfig,omitempty"`

	// Parameters for node pools to be backed by shared sole tenant node groups.
	SoleTenantConfig *SoleTenantConfig `json:"soleTenantConfig,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// Parameters for containerd customization.
	ContainerdConfig *ContainerdConfig `json:"containerdConfig,omitempty"`
	*/

	// HostMaintenancePolicy contains the desired maintenance policy for the
	//  Google Compute Engine hosts.
	HostMaintenancePolicy *HostMaintenancePolicy `json:"hostMaintenancePolicy,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// A map of resource manager tag keys and values to be attached to the nodes.
	ResourceManagerTags *ResourceManagerTags `json:"resourceManagerTags,omitempty"`
	*/

	/* REMOVED_TO_MATCH_CRD:
	// Optional. Reserved for future use.
	EnableConfidentialStorage *bool `json:"enableConfidentialStorage,omitempty"`
	*/

	/* REMOVED_TO_MATCH_CRD:
	// List of secondary boot disks attached to the nodes.
	SecondaryBootDisks []SecondaryBootDisk `json:"secondaryBootDisks,omitempty"`
	*/

	/* REMOVED_TO_MATCH_CRD:
	// Secondary boot disk update strategy.
	SecondaryBootDiskUpdateStrategy *SecondaryBootDiskUpdateStrategy `json:"secondaryBootDiskUpdateStrategy,omitempty"`
	*/
}

// +kcc:proto=google.container.v1beta1.NodeKubeletConfig
type NodeKubeletConfig struct {
	// Control the CPU management policy on the node.
	//  See
	//  https://kubernetes.io/docs/tasks/administer-cluster/cpu-management-policies/
	//
	//  The following values are allowed.
	//  * "none": the default, which represents the existing scheduling behavior.
	//  * "static": allows pods with certain resource characteristics to be granted
	//  increased CPU affinity and exclusivity on the node.
	//  The default value is 'none' if unspecified.
	CpuManagerPolicy *string `json:"cpuManagerPolicy,omitempty"`

	// Enable CPU CFS quota enforcement for containers that specify CPU limits.
	//
	//  This option is enabled by default which makes kubelet use CFS quota
	//  (https://www.kernel.org/doc/Documentation/scheduler/sched-bwc.txt) to
	//  enforce container CPU limits. Otherwise, CPU limits will not be enforced at
	//  all.
	//
	//  Disable this option to mitigate CPU throttling problems while still having
	//  your pods to be in Guaranteed QoS class by specifying the CPU limits.
	//
	//  The default value is 'true' if unspecified.
	CpuCfsQuota *BoolValue `json:"cpuCfsQuota,omitempty"`

	// Set the CPU CFS quota period value 'cpu.cfs_period_us'.
	//
	//  The string must be a sequence of decimal numbers, each with optional
	//  fraction and a unit suffix, such as "300ms".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	//  The value must be a positive duration.
	CpuCfsQuotaPeriod *string `json:"cpuCfsQuotaPeriod,omitempty"`

	// Set the Pod PID limits. See
	//  https://kubernetes.io/docs/concepts/policy/pid-limiting/#pod-pid-limits
	//
	//  Controls the maximum number of processes allowed to run in a pod. The value
	//  must be greater than or equal to 1024 and less than 4194304.
	PodPidsLimit *int64 `json:"podPidsLimit,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// Enable or disable Kubelet read only port.
	InsecureKubeletReadonlyPortEnabled *bool `json:"insecureKubeletReadonlyPortEnabled,omitempty"`
	*/
}

// +kcc:proto=google.container.v1beta1.NodeManagement
type NodeManagement struct {
	// Whether the nodes will be automatically upgraded.
	AutoUpgrade *bool `json:"autoUpgrade,omitempty"`

	// Whether the nodes will be automatically repaired.
	AutoRepair *bool `json:"autoRepair,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// Specifies the Auto Upgrade knobs for the node pool.
	UpgradeOptions *AutoUpgradeOptions `json:"upgradeOptions,omitempty"`
	*/
}

// +kcc:proto=google.container.v1beta1.NodeNetworkConfig
type NodeNetworkConfig struct {
	// Input only. Whether to create a new range for pod IPs in this node pool.
	//  Defaults are provided for `pod_range` and `pod_ipv4_cidr_block` if they
	//  are not specified.
	//
	//  If neither `create_pod_range` or `pod_range` are specified, the
	//  cluster-level default (`ip_allocation_policy.cluster_ipv4_cidr_block`) is
	//  used.
	//
	//  Only applicable if `ip_allocation_policy.use_ip_aliases` is true.
	//
	//  This field cannot be changed after the node pool has been created.
	CreatePodRange *bool `json:"createPodRange,omitempty"`

	// The ID of the secondary range for pod IPs.
	//  If `create_pod_range` is true, this ID is used for the new range.
	//  If `create_pod_range` is false, uses an existing secondary range with this
	//  ID.
	//
	//  Only applicable if `ip_allocation_policy.use_ip_aliases` is true.
	//
	//  This field cannot be changed after the node pool has been created.
	PodRange *string `json:"podRange,omitempty"`

	// The IP address range for pod IPs in this node pool.
	//
	//  Only applicable if `create_pod_range` is true.
	//
	//  Set to blank to have a range chosen with the default size.
	//
	//  Set to /netmask (e.g. `/14`) to have a range chosen with a specific
	//  netmask.
	//
	//  Set to a
	//  [CIDR](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing)
	//  notation (e.g. `10.96.0.0/14`) to pick a specific range to use.
	//
	//  Only applicable if `ip_allocation_policy.use_ip_aliases` is true.
	//
	//  This field cannot be changed after the node pool has been created.
	PodIpv4CidrBlock *string `json:"podIpv4CidrBlock,omitempty"`

	// Whether nodes have internal IP addresses only.
	//  If enable_private_nodes is not specified, then the value is derived from
	//  [cluster.privateClusterConfig.enablePrivateNodes][google.container.v1beta1.PrivateClusterConfig.enablePrivateNodes]
	EnablePrivateNodes *bool `json:"enablePrivateNodes,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// Network bandwidth tier configuration.
	NetworkPerformanceConfig *NodeNetworkConfig_NetworkPerformanceConfig `json:"networkPerformanceConfig,omitempty"`
	*/

	// [PRIVATE FIELD]
	//  Pod CIDR size overprovisioning config for the nodepool.
	//
	//  Pod CIDR size per node depends on max_pods_per_node. By default, the value
	//  of max_pods_per_node is rounded off to next power of 2 and we then double
	//  that to get the size of pod CIDR block per node.
	//  Example: max_pods_per_node of 30 would result in 64 IPs (/26).
	//
	//  This config can disable the doubling of IPs (we still round off to next
	//  power of 2)
	//  Example: max_pods_per_node of 30 will result in 32 IPs (/27) when
	//  overprovisioning is disabled.
	PodCidrOverprovisionConfig *PodCIDROverprovisionConfig `json:"podCidrOverprovisionConfig,omitempty"`

	// We specify the additional node networks for this node pool using this list.
	//  Each node network corresponds to an additional interface
	AdditionalNodeNetworkConfigs []AdditionalNodeNetworkConfig `json:"additionalNodeNetworkConfigs,omitempty"`

	// We specify the additional pod networks for this node pool using this list.
	//  Each pod network corresponds to an additional alias IP range for the node
	AdditionalPodNetworkConfigs []AdditionalPodNetworkConfig `json:"additionalPodNetworkConfigs,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// Output only. The utilization of the IPv4 range for the pod.
	//  The ratio is Usage/[Total number of IPs in the secondary range],
	//  Usage=numNodes*numZones*podIPsPerNode.
	PodIpv4RangeUtilization *float64 `json:"podIpv4RangeUtilization,omitempty"`
	*/
}

/* REMOVED_TO_MATCH_CRD:
// +kcc:proto=google.container.v1beta1.NodeNetworkConfig.NetworkPerformanceConfig
type NodeNetworkConfig_NetworkPerformanceConfig struct {
	// Specifies the total network bandwidth tier for the NodePool.
	TotalEgressBandwidthTier *string `json:"totalEgressBandwidthTier,omitempty"`

	// Specifies the network bandwidth tier for the NodePool for traffic to
	//  external/public IP addresses.
	ExternalIpEgressBandwidthTier *string `json:"externalIpEgressBandwidthTier,omitempty"`
}
*/

// +kcc:proto=google.container.v1beta1.NodePool.PlacementPolicy
type NodePool_PlacementPolicy struct {
	// The type of placement.
	Type *string `json:"type,omitempty"`

	// TPU placement topology for pod slice node pool.
	//  https://cloud.google.com/tpu/docs/types-topologies#tpu_topologies
	TpuTopology *string `json:"tpuTopology,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// If set, refers to the name of a custom resource policy supplied by the
	//  user. The resource policy must be in the same project and region as the
	//  node pool. If not found, InvalidArgument error is returned.
	PolicyName *string `json:"policyName,omitempty"`
	*/

	/* Immutable. If set, refers to the name of a custom resource policy supplied by the user. The resource policy must be in the same project and region as the node pool. If not found, InvalidArgument error is returned. */
	// +optional
	PolicyNameRef *v1alpha1.ResourceRef `json:"policyNameRef,omitempty"`
	/* ADDED_TO_MATCH_CRD: */
}

// +kcc:proto=google.container.v1beta1.NodePool.QueuedProvisioning
type NodePool_QueuedProvisioning struct {
	// Denotes that this nodepool is QRM specific, meaning nodes can be only
	//  obtained through queuing via the Cluster Autoscaler ProvisioningRequest
	//  API.
	Enabled *bool `json:"enabled,omitempty"`
}

// +kcc:proto=google.container.v1beta1.NodePool.UpdateInfo
type NodePool_UpdateInfo struct {
	// Information of a blue-green upgrade.
	BlueGreenInfo *NodePool_UpdateInfo_BlueGreenInfo `json:"blueGreenInfo,omitempty"`
}

// +kcc:proto=google.container.v1beta1.NodePool.UpdateInfo.BlueGreenInfo
type NodePool_UpdateInfo_BlueGreenInfo struct {
	// Current blue-green upgrade phase.
	Phase *string `json:"phase,omitempty"`

	// The resource URLs of the [managed instance groups]
	//  (/compute/docs/instance-groups/creating-groups-of-managed-instances)
	//  associated with blue pool.
	BlueInstanceGroupUrls []string `json:"blueInstanceGroupUrls,omitempty"`

	// The resource URLs of the [managed instance groups]
	//  (/compute/docs/instance-groups/creating-groups-of-managed-instances)
	//  associated with green pool.
	GreenInstanceGroupUrls []string `json:"greenInstanceGroupUrls,omitempty"`

	// Time to start deleting blue pool to complete blue-green upgrade,
	//  in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	BluePoolDeletionStartTime *string `json:"bluePoolDeletionStartTime,omitempty"`

	// Version of green pool.
	GreenPoolVersion *string `json:"greenPoolVersion,omitempty"`
}

// +kcc:proto=google.container.v1beta1.NodePool.UpgradeSettings
type NodePool_UpgradeSettings struct {
	// The maximum number of nodes that can be created beyond the current size
	//  of the node pool during the upgrade process.
	MaxSurge *int32 `json:"maxSurge,omitempty"`

	// The maximum number of nodes that can be simultaneously unavailable during
	//  the upgrade process. A node is considered available if its status is
	//  Ready.
	MaxUnavailable *int32 `json:"maxUnavailable,omitempty"`

	// Update strategy of the node pool.
	Strategy *string `json:"strategy,omitempty"`

	// Settings for blue-green upgrade strategy.
	BlueGreenSettings *BlueGreenSettings `json:"blueGreenSettings,omitempty"`
}

// +kcc:proto=google.container.v1beta1.NodePoolAutoscaling
type NodePoolAutoscaling struct {
	/* REMOVED_TO_MATCH_CRD:
	// Is autoscaling enabled for this node pool.
	Enabled *bool `json:"enabled,omitempty"`
	*/

	// Minimum number of nodes for one location in the NodePool. Must be >= 1 and
	//  <= max_node_count.
	MinNodeCount *int32 `json:"minNodeCount,omitempty"`

	// Maximum number of nodes for one location in the NodePool. Must be >=
	//  min_node_count. There has to be enough quota to scale up the cluster.
	MaxNodeCount *int32 `json:"maxNodeCount,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// Can this node pool be deleted automatically.
	Autoprovisioned *bool `json:"autoprovisioned,omitempty"`
	*/

	// Location policy used when scaling up a nodepool.
	LocationPolicy *string `json:"locationPolicy,omitempty"`

	// Minimum number of nodes in the node pool. Must be greater than 1 less than
	//  total_max_node_count.
	//  The total_*_node_count fields are mutually exclusive with the *_node_count
	//  fields.
	TotalMinNodeCount *int32 `json:"totalMinNodeCount,omitempty"`

	// Maximum number of nodes in the node pool. Must be greater than
	//  total_min_node_count. There has to be enough quota to scale up the cluster.
	//  The total_*_node_count fields are mutually exclusive with the *_node_count
	//  fields.
	TotalMaxNodeCount *int32 `json:"totalMaxNodeCount,omitempty"`
}

/* REMOVED_TO_MATCH_CRD:
// +kcc:proto=google.container.v1beta1.NodePoolLoggingConfig
type NodePoolLoggingConfig struct {
	// Logging variant configuration.
	VariantConfig *LoggingVariantConfig `json:"variantConfig,omitempty"`
}
*/

// +kcc:proto=google.container.v1beta1.NodeTaint
type NodeTaint struct {
	// Key for taint.
	Key *string `json:"key,omitempty"`

	// Value for taint.
	Value *string `json:"value,omitempty"`

	// Effect for taint.
	Effect *string `json:"effect,omitempty"`
}

// +kcc:proto=google.container.v1beta1.PodCIDROverprovisionConfig
type PodCIDROverprovisionConfig struct {
	// Whether Pod CIDR overprovisioning is disabled.
	//  Note: Pod CIDR overprovisioning is enabled by default.
	Disabled bool `json:"disabled"`
	/* ADDED_TO_MATCH_CRD:

	/* REMOVED_TO_MATCH_CRD:
	Disable *bool `json:"disable,omitempty"`
	*/
}

// +kcc:proto=google.container.v1beta1.ReservationAffinity
type ReservationAffinity struct {
	// Corresponds to the type of reservation consumption.
	ConsumeReservationType *string `json:"consumeReservationType,omitempty"`

	// Corresponds to the label key of a reservation resource. To target a
	//  SPECIFIC_RESERVATION by name, specify
	//  "compute.googleapis.com/reservation-name" as the key and specify the name
	//  of your reservation as its value.
	Key *string `json:"key,omitempty"`

	// Corresponds to the label value(s) of reservation resource(s).
	Values []string `json:"values,omitempty"`
}

/* REMOVED_TO_MATCH_CRD:
// +kcc:proto=google.container.v1beta1.ResourceManagerTags
type ResourceManagerTags struct {
	// Tags must be in one of the following formats ([KEY]=[VALUE])
	//  1. `tagKeys/{tag_key_id}=tagValues/{tag_value_id}`
	//  2. `{org_id}/{tag_key_name}={tag_value_name}`
	//  3. `{project_id}/{tag_key_name}={tag_value_name}`
	Tags map[string]string `json:"tags,omitempty"`
}
*/

// +kcc:proto=google.container.v1beta1.SandboxConfig
type SandboxConfig struct {
	// Type of the sandbox to use for the node (e.g. 'gvisor')
	SandboxType *string `json:"sandboxType,omitempty"`

	/* REMOVED_TO_MATCH_CRD:
	// Type of the sandbox to use for the node.
	Type *string `json:"type,omitempty"`
	*/
}

/* REMOVED_TO_MATCH_CRD:
// +kcc:proto=google.container.v1beta1.SecondaryBootDisk
type SecondaryBootDisk struct {
	// Disk mode (container image cache, etc.)
	Mode *string `json:"mode,omitempty"`

	// Fully-qualified resource ID for an existing disk image.
	DiskImage *string `json:"diskImage,omitempty"`
}
*/

/* REMOVED_TO_MATCH_CRD:
// +kcc:proto=google.container.v1beta1.SecondaryBootDiskUpdateStrategy
type SecondaryBootDiskUpdateStrategy struct {
}
*/

// +kcc:proto=google.container.v1beta1.ShieldedInstanceConfig
type ShieldedInstanceConfig struct {
	// Defines whether the instance has Secure Boot enabled.
	//
	//  Secure Boot helps ensure that the system only runs authentic software by
	//  verifying the digital signature of all boot components, and halting the
	//  boot process if signature verification fails.
	EnableSecureBoot *bool `json:"enableSecureBoot,omitempty"`

	// Defines whether the instance has integrity monitoring enabled.
	//
	//  Enables monitoring and attestation of the boot integrity of the instance.
	//  The attestation is performed against the integrity policy baseline. This
	//  baseline is initially derived from the implicitly trusted boot image when
	//  the instance is created.
	EnableIntegrityMonitoring *bool `json:"enableIntegrityMonitoring,omitempty"`
}

// +kcc:proto=google.container.v1beta1.SoleTenantConfig
type SoleTenantConfig struct {
	/* RENAMED_TO_MATCH_CRD:
	// NodeAffinities used to match to a shared sole tenant node group.
	NodeAffinities []SoleTenantConfig_NodeAffinity `json:"nodeAffinities,omitempty"`
	*/

	/* Immutable. . */
	NodeAffinity []SoleTenantConfig_NodeAffinity `json:"nodeAffinity"`
}

// +kcc:proto=google.container.v1beta1.SoleTenantConfig.NodeAffinity
type SoleTenantConfig_NodeAffinity struct {
	// Key for NodeAffinity.
	Key *string `json:"key,omitempty"`

	// Operator for NodeAffinity.
	Operator *string `json:"operator,omitempty"`

	// Values for NodeAffinity.
	Values []string `json:"values,omitempty"`
}

// +kcc:proto=google.container.v1beta1.StatusCondition
type StatusCondition struct {
	// Machine-friendly representation of the condition
	//  Deprecated. Use canonical_code instead.
	Code *string `json:"code,omitempty"`

	// Human-friendly representation of the condition
	Message *string `json:"message,omitempty"`

	// Canonical code of the condition.
	CanonicalCode *string `json:"canonicalCode,omitempty"`
}

// +kcc:proto=google.container.v1beta1.VirtualNIC
type VirtualNIC struct {
	// Whether gVNIC features are enabled in the node pool.
	Enabled *bool `json:"enabled,omitempty"`
}

/* REMOVED_TO_MATCH_CRD:
// +kcc:proto=google.container.v1beta1.WindowsNodeConfig
type WindowsNodeConfig struct {
	// OSVersion specifies the Windows node config to be used on the node
	OsVersion *string `json:"osVersion,omitempty"`
}
*/

// +kcc:proto=google.container.v1beta1.WorkloadMetadataConfig
type WorkloadMetadataConfig struct {
	// NodeMetadata is the configuration for how to expose metadata to the
	//  workloads running on the node.
	NodeMetadata *string `json:"nodeMetadata,omitempty"`

	// Mode is the configuration for how to expose metadata to workloads running
	//  on the node pool.
	Mode *string `json:"mode,omitempty"`
}

// +kcc:proto=google.protobuf.BoolValue
type BoolValue struct {
	// The bool value.
	Value *bool `json:"value,omitempty"`
}
