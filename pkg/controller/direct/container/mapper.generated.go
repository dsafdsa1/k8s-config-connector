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
	"google.golang.org/genproto/googleapis/rpc/code"

	pb "github.com/GoogleCloudPlatform/k8s-config-connector/mockgcp/generated/mockgcp/container/v1beta1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/controller/direct"
)

func AdditionalNodeNetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.AdditionalNodeNetworkConfig) *krm.AdditionalNodeNetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.AdditionalNodeNetworkConfig{}
	if in.GetNetwork() != "" {
		out.NetworkRef = &v1alpha1.ResourceRef{External: in.GetNetwork()}
	}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &v1alpha1.ResourceRef{External: in.GetSubnetwork()}
	}
	return out
}
func AdditionalNodeNetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.AdditionalNodeNetworkConfig) *pb.AdditionalNodeNetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.AdditionalNodeNetworkConfig{}
	if in.NetworkRef != nil {
		out.Network = in.NetworkRef.External
	}
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	return out
}
func AdditionalPodNetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.AdditionalPodNetworkConfig) *krm.AdditionalPodNetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.AdditionalPodNetworkConfig{}
	if in.GetSubnetwork() != "" {
		out.SubnetworkRef = &v1alpha1.ResourceRef{External: in.GetSubnetwork()}
	}
	out.SecondaryPodRange = direct.LazyPtr(in.GetSecondaryPodRange())
	out.MaxPodsPerNode = int64_FromProto(mapCtx, in.GetMaxPodsPerNode())
	return out
}
func AdditionalPodNetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.AdditionalPodNetworkConfig) *pb.AdditionalPodNetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.AdditionalPodNetworkConfig{}
	if in.SubnetworkRef != nil {
		out.Subnetwork = in.SubnetworkRef.External
	}
	out.SecondaryPodRange = direct.ValueOf(in.SecondaryPodRange)
	if oneof := int64_ToProto(mapCtx, in.MaxPodsPerNode); oneof != nil {
		out.MaxPodsPerNode = oneof
	}
	return out
}
func AdvancedMachineFeatures_FromProto(mapCtx *direct.MapContext, in *pb.AdvancedMachineFeatures) *krm.AdvancedMachineFeatures {
	if in == nil {
		return nil
	}
	out := &krm.AdvancedMachineFeatures{}
	out.ThreadsPerCore = in.ThreadsPerCore
	// MISSING: EnableNestedVirtualization - REMOVED_TO_MATCH_CRD
	return out
}
func AdvancedMachineFeatures_ToProto(mapCtx *direct.MapContext, in *krm.AdvancedMachineFeatures) *pb.AdvancedMachineFeatures {
	if in == nil {
		return nil
	}
	out := &pb.AdvancedMachineFeatures{}
	out.ThreadsPerCore = in.ThreadsPerCore
	// MISSING: EnableNestedVirtualization - REMOVED_TO_MATCH_CRD
	return out
}
func BestEffortProvisioning_FromProto(mapCtx *direct.MapContext, in *pb.BestEffortProvisioning) *krm.BestEffortProvisioning {
	if in == nil {
		return nil
	}
	out := &krm.BestEffortProvisioning{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	out.MinProvisionNodes = direct.LazyPtr(in.GetMinProvisionNodes())
	return out
}
func BestEffortProvisioning_ToProto(mapCtx *direct.MapContext, in *krm.BestEffortProvisioning) *pb.BestEffortProvisioning {
	if in == nil {
		return nil
	}
	out := &pb.BestEffortProvisioning{}
	out.Enabled = direct.ValueOf(in.Enabled)
	out.MinProvisionNodes = direct.ValueOf(in.MinProvisionNodes)
	return out
}
func BlueGreenSettings_FromProto(mapCtx *direct.MapContext, in *pb.BlueGreenSettings) *krm.BlueGreenSettings {
	if in == nil {
		return nil
	}
	out := &krm.BlueGreenSettings{}
	out.StandardRolloutPolicy = BlueGreenSettings_StandardRolloutPolicy_FromProto(mapCtx, in.GetStandardRolloutPolicy())
	// MISSING: AutoscaledRolloutPolicy - REMOVED_TO_MATCH_CRD
	out.NodePoolSoakDuration = direct.StringDuration_FromProto(mapCtx, in.GetNodePoolSoakDuration())
	return out
}
func BlueGreenSettings_ToProto(mapCtx *direct.MapContext, in *krm.BlueGreenSettings) *pb.BlueGreenSettings {
	if in == nil {
		return nil
	}
	out := &pb.BlueGreenSettings{}
	if oneof := BlueGreenSettings_StandardRolloutPolicy_ToProto(mapCtx, in.StandardRolloutPolicy); oneof != nil {
		out.RolloutPolicy = &pb.BlueGreenSettings_StandardRolloutPolicy_{StandardRolloutPolicy: oneof}
	}
	// MISSING: AutoscaledRolloutPolicy - REMOVED_TO_MATCH_CRD
	if oneof := direct.StringDuration_ToProto(mapCtx, in.NodePoolSoakDuration); oneof != nil {
		out.NodePoolSoakDuration = oneof
	}
	return out
}
func BlueGreenSettings_StandardRolloutPolicy_FromProto(mapCtx *direct.MapContext, in *pb.BlueGreenSettings_StandardRolloutPolicy) *krm.BlueGreenSettings_StandardRolloutPolicy {
	if in == nil {
		return nil
	}
	out := &krm.BlueGreenSettings_StandardRolloutPolicy{}
	out.BatchPercentage = direct.LazyPtr(in.GetBatchPercentage())
	out.BatchNodeCount = direct.LazyPtr(in.GetBatchNodeCount())
	out.BatchSoakDuration = direct.StringDuration_FromProto(mapCtx, in.GetBatchSoakDuration())
	return out
}
func BlueGreenSettings_StandardRolloutPolicy_ToProto(mapCtx *direct.MapContext, in *krm.BlueGreenSettings_StandardRolloutPolicy) *pb.BlueGreenSettings_StandardRolloutPolicy {
	if in == nil {
		return nil
	}
	out := &pb.BlueGreenSettings_StandardRolloutPolicy{}
	if oneof := BlueGreenSettings_StandardRolloutPolicy_BatchPercentage_ToProto(mapCtx, in.BatchPercentage); oneof != nil {
		out.UpdateBatchSize = oneof
	}
	if oneof := BlueGreenSettings_StandardRolloutPolicy_BatchNodeCount_ToProto(mapCtx, in.BatchNodeCount); oneof != nil {
		out.UpdateBatchSize = oneof
	}
	if oneof := direct.StringDuration_ToProto(mapCtx, in.BatchSoakDuration); oneof != nil {
		out.BatchSoakDuration = oneof
	}
	return out
}
func ConfidentialNodes_FromProto(mapCtx *direct.MapContext, in *pb.ConfidentialNodes) *krm.ConfidentialNodes {
	if in == nil {
		return nil
	}
	out := &krm.ConfidentialNodes{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}
func ConfidentialNodes_ToProto(mapCtx *direct.MapContext, in *krm.ConfidentialNodes) *pb.ConfidentialNodes {
	if in == nil {
		return nil
	}
	out := &pb.ConfidentialNodes{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
func ContainerNodePoolSpec_FromProto(mapCtx *direct.MapContext, in *pb.NodePool) *krm.ContainerNodePoolSpec {
	if in == nil {
		return nil
	}
	out := &krm.ContainerNodePoolSpec{}
	// MISSING: Name
	// MISSING: Config
	out.InitialNodeCount = direct.LazyPtr(int64(in.GetInitialNodeCount()))
	// MISSING: Locations
	out.NetworkConfig = NodeNetworkConfig_FromProto(mapCtx, in.GetNetworkConfig())
	// MISSING: SelfLink
	out.Version = direct.LazyPtr(in.GetVersion())
	// MISSING: InstanceGroupUrls
	// MISSING: Status
	// MISSING: StatusMessage
	out.Autoscaling = NodePoolAutoscaling_FromProto(mapCtx, in.GetAutoscaling())
	out.Management = NodeManagement_FromProto(mapCtx, in.GetManagement())
	// MISSING: MaxPodsConstraint
	// MISSING: Conditions
	// MISSING: PodIpv4CidrSize
	out.UpgradeSettings = NodePool_UpgradeSettings_FromProto(mapCtx, in.GetUpgradeSettings())
	out.PlacementPolicy = NodePool_PlacementPolicy_FromProto(mapCtx, in.GetPlacementPolicy())
	// MISSING: UpdateInfo
	// MISSING: Etag
	// MISSING: QueuedProvisioning
	// MISSING: BestEffortProvisioning
	return out
}
func ContainerNodePoolSpec_ToProto(mapCtx *direct.MapContext, in *krm.ContainerNodePoolSpec) *pb.NodePool {
	if in == nil {
		return nil
	}
	out := &pb.NodePool{}
	// MISSING: Name
	// MISSING: Config
	out.InitialNodeCount = int32(direct.ValueOf(in.InitialNodeCount))
	// MISSING: Locations
	out.NetworkConfig = NodeNetworkConfig_ToProto(mapCtx, in.NetworkConfig)
	// MISSING: SelfLink
	out.Version = direct.ValueOf(in.Version)
	// MISSING: InstanceGroupUrls
	// MISSING: Status
	// MISSING: StatusMessage
	out.Autoscaling = NodePoolAutoscaling_ToProto(mapCtx, in.Autoscaling)
	out.Management = NodeManagement_ToProto(mapCtx, in.Management)
	// MISSING: MaxPodsConstraint
	// MISSING: Conditions
	// MISSING: PodIpv4CidrSize
	out.UpgradeSettings = NodePool_UpgradeSettings_ToProto(mapCtx, in.UpgradeSettings)
	out.PlacementPolicy = NodePool_PlacementPolicy_ToProto(mapCtx, in.PlacementPolicy)
	// MISSING: UpdateInfo
	// MISSING: Etag
	// MISSING: QueuedProvisioning
	// MISSING: BestEffortProvisioning
	return out
}
func EphemeralStorageConfig_FromProto(mapCtx *direct.MapContext, in *pb.EphemeralStorageConfig) *krm.EphemeralStorageConfig {
	if in == nil {
		return nil
	}
	out := &krm.EphemeralStorageConfig{}
	out.LocalSsdCount = direct.LazyPtr(in.GetLocalSsdCount())
	return out
}
func EphemeralStorageConfig_ToProto(mapCtx *direct.MapContext, in *krm.EphemeralStorageConfig) *pb.EphemeralStorageConfig {
	if in == nil {
		return nil
	}
	out := &pb.EphemeralStorageConfig{}
	out.LocalSsdCount = direct.ValueOf(in.LocalSsdCount)
	return out
}
func EphemeralStorageLocalSsdConfig_FromProto(mapCtx *direct.MapContext, in *pb.EphemeralStorageLocalSsdConfig) *krm.EphemeralStorageLocalSsdConfig {
	if in == nil {
		return nil
	}
	out := &krm.EphemeralStorageLocalSsdConfig{}
	out.LocalSsdCount = direct.LazyPtr(in.GetLocalSsdCount())
	return out
}
func EphemeralStorageLocalSsdConfig_ToProto(mapCtx *direct.MapContext, in *krm.EphemeralStorageLocalSsdConfig) *pb.EphemeralStorageLocalSsdConfig {
	if in == nil {
		return nil
	}
	out := &pb.EphemeralStorageLocalSsdConfig{}
	out.LocalSsdCount = direct.ValueOf(in.LocalSsdCount)
	return out
}
func FastSocket_FromProto(mapCtx *direct.MapContext, in *pb.FastSocket) *krm.FastSocket {
	if in == nil {
		return nil
	}
	out := &krm.FastSocket{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}
func FastSocket_ToProto(mapCtx *direct.MapContext, in *krm.FastSocket) *pb.FastSocket {
	if in == nil {
		return nil
	}
	out := &pb.FastSocket{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
func GPUDriverInstallationConfig_FromProto(mapCtx *direct.MapContext, in *pb.GPUDriverInstallationConfig) *krm.GPUDriverInstallationConfig {
	if in == nil {
		return nil
	}
	out := &krm.GPUDriverInstallationConfig{}
	out.GpuDriverVersion = direct.Enum_FromProto(mapCtx, in.GetGpuDriverVersion())
	return out
}
func GPUDriverInstallationConfig_ToProto(mapCtx *direct.MapContext, in *krm.GPUDriverInstallationConfig) *pb.GPUDriverInstallationConfig {
	if in == nil {
		return nil
	}
	out := &pb.GPUDriverInstallationConfig{}
	if oneof := GPUDriverInstallationConfig_GpuDriverVersion_ToProto(mapCtx, in.GpuDriverVersion); oneof != nil {
		out.GpuDriverVersion = oneof
	}
	return out
}
func GPUSharingConfig_FromProto(mapCtx *direct.MapContext, in *pb.GPUSharingConfig) *krm.GPUSharingConfig {
	if in == nil {
		return nil
	}
	out := &krm.GPUSharingConfig{}
	out.MaxSharedClientsPerGpu = direct.LazyPtr(in.GetMaxSharedClientsPerGpu())
	out.GpuSharingStrategy = direct.Enum_FromProto(mapCtx, in.GetGpuSharingStrategy())
	return out
}
func GPUSharingConfig_ToProto(mapCtx *direct.MapContext, in *krm.GPUSharingConfig) *pb.GPUSharingConfig {
	if in == nil {
		return nil
	}
	out := &pb.GPUSharingConfig{}
	out.MaxSharedClientsPerGpu = direct.ValueOf(in.MaxSharedClientsPerGpu)
	if oneof := GPUSharingConfig_GpuSharingStrategy_ToProto(mapCtx, in.GpuSharingStrategy); oneof != nil {
		out.GpuSharingStrategy = oneof
	}
	return out
}
func GcfsConfig_FromProto(mapCtx *direct.MapContext, in *pb.GcfsConfig) *krm.GcfsConfig {
	if in == nil {
		return nil
	}
	out := &krm.GcfsConfig{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}
func GcfsConfig_ToProto(mapCtx *direct.MapContext, in *krm.GcfsConfig) *pb.GcfsConfig {
	if in == nil {
		return nil
	}
	out := &pb.GcfsConfig{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
func HostMaintenancePolicy_FromProto(mapCtx *direct.MapContext, in *pb.HostMaintenancePolicy) *krm.HostMaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &krm.HostMaintenancePolicy{}
	out.MaintenanceInterval = direct.Enum_FromProto(mapCtx, in.GetMaintenanceInterval())
	// MISSING: OpportunisticMaintenanceStrategy
	return out
}
func HostMaintenancePolicy_ToProto(mapCtx *direct.MapContext, in *krm.HostMaintenancePolicy) *pb.HostMaintenancePolicy {
	if in == nil {
		return nil
	}
	out := &pb.HostMaintenancePolicy{}
	if oneof := HostMaintenancePolicy_MaintenanceInterval_ToProto(mapCtx, in.MaintenanceInterval); oneof != nil {
		out.MaintenanceInterval = oneof
	}
	// MISSING: OpportunisticMaintenanceStrategy
	return out
}
func LinuxNodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.LinuxNodeConfig) *krm.LinuxNodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.LinuxNodeConfig{}
	out.Sysctls = in.Sysctls
	out.CgroupMode = direct.Enum_FromProto(mapCtx, in.GetCgroupMode())
	// MISSING: Hugepages
	return out
}
func LinuxNodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.LinuxNodeConfig) *pb.LinuxNodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.LinuxNodeConfig{}
	out.Sysctls = in.Sysctls
	out.CgroupMode = direct.Enum_ToProto[pb.LinuxNodeConfig_CgroupMode](mapCtx, in.CgroupMode)
	// MISSING: Hugepages
	return out
}
func LocalNvmeSsdBlockConfig_FromProto(mapCtx *direct.MapContext, in *pb.LocalNvmeSsdBlockConfig) *krm.LocalNvmeSsdBlockConfig {
	if in == nil {
		return nil
	}
	out := &krm.LocalNvmeSsdBlockConfig{}
	out.LocalSsdCount = direct.LazyPtr(in.GetLocalSsdCount())
	return out
}
func LocalNvmeSsdBlockConfig_ToProto(mapCtx *direct.MapContext, in *krm.LocalNvmeSsdBlockConfig) *pb.LocalNvmeSsdBlockConfig {
	if in == nil {
		return nil
	}
	out := &pb.LocalNvmeSsdBlockConfig{}
	out.LocalSsdCount = direct.ValueOf(in.LocalSsdCount)
	return out
}
func LoggingVariantConfig_FromProto(mapCtx *direct.MapContext, in *pb.LoggingVariantConfig) *krm.LoggingVariantConfig {
	if in == nil {
		return nil
	}
	out := &krm.LoggingVariantConfig{}
	out.Variant = direct.Enum_FromProto(mapCtx, in.GetVariant())
	return out
}
func LoggingVariantConfig_ToProto(mapCtx *direct.MapContext, in *krm.LoggingVariantConfig) *pb.LoggingVariantConfig {
	if in == nil {
		return nil
	}
	out := &pb.LoggingVariantConfig{}
	out.Variant = direct.Enum_ToProto[pb.LoggingVariantConfig_Variant](mapCtx, in.Variant)
	return out
}
func NodeConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeConfig) *krm.NodeConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeConfig{}
	out.MachineType = direct.LazyPtr(in.GetMachineType())
	out.DiskSizeGb = direct.LazyPtr(in.GetDiskSizeGb())
	out.OauthScopes = in.OauthScopes
	if in.GetServiceAccount() != "" {
		out.ServiceAccountRef = &v1alpha1.ResourceRef{External: in.GetServiceAccount()}
	}
	out.Metadata = in.Metadata
	out.ImageType = direct.LazyPtr(in.GetImageType())
	out.Labels = in.Labels
	out.LocalSsdCount = direct.LazyPtr(in.GetLocalSsdCount())
	out.Tags = in.Tags
	out.Preemptible = direct.LazyPtr(in.GetPreemptible())
	// MISSING: Accelerators
	out.SandboxConfig = SandboxConfig_FromProto(mapCtx, in.GetSandboxConfig())
	if in.GetNodeGroup() != "" {
		out.NodeGroupRef = &v1alpha1.ResourceRef{External: in.GetNodeGroup()}
	}
	out.ReservationAffinity = ReservationAffinity_FromProto(mapCtx, in.GetReservationAffinity())
	out.DiskType = direct.LazyPtr(in.GetDiskType())
	out.MinCpuPlatform = direct.LazyPtr(in.GetMinCpuPlatform())
	out.WorkloadMetadataConfig = WorkloadMetadataConfig_FromProto(mapCtx, in.GetWorkloadMetadataConfig())
	// MISSING: Taints
	// MISSING: BootDiskKmsKey
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_FromProto(mapCtx, in.GetShieldedInstanceConfig())
	out.LinuxNodeConfig = LinuxNodeConfig_FromProto(mapCtx, in.GetLinuxNodeConfig())
	out.KubeletConfig = NodeKubeletConfig_FromProto(mapCtx, in.GetKubeletConfig())
	out.EphemeralStorageConfig = EphemeralStorageConfig_FromProto(mapCtx, in.GetEphemeralStorageConfig())
	out.GcfsConfig = GcfsConfig_FromProto(mapCtx, in.GetGcfsConfig())
	out.AdvancedMachineFeatures = AdvancedMachineFeatures_FromProto(mapCtx, in.GetAdvancedMachineFeatures())
	out.Gvnic = VirtualNIC_FromProto(mapCtx, in.GetGvnic())
	out.Spot = direct.LazyPtr(in.GetSpot())
	out.ConfidentialNodes = ConfidentialNodes_FromProto(mapCtx, in.GetConfidentialNodes())
	out.FastSocket = FastSocket_FromProto(mapCtx, in.GetFastSocket())
	out.ResourceLabels = in.ResourceLabels
	// MISSING: LoggingConfig
	// MISSING: WindowsNodeConfig
	out.LocalNvmeSsdBlockConfig = LocalNvmeSsdBlockConfig_FromProto(mapCtx, in.GetLocalNvmeSsdBlockConfig())
	out.EphemeralStorageLocalSsdConfig = EphemeralStorageLocalSsdConfig_FromProto(mapCtx, in.GetEphemeralStorageLocalSsdConfig())
	out.SoleTenantConfig = SoleTenantConfig_FromProto(mapCtx, in.GetSoleTenantConfig())
	// MISSING: ContainerdConfig
	out.HostMaintenancePolicy = HostMaintenancePolicy_FromProto(mapCtx, in.GetHostMaintenancePolicy())
	// MISSING: ResourceManagerTags
	// MISSING: EnableConfidentialStorage
	// MISSING: SecondaryBootDisks
	// MISSING: SecondaryBootDiskUpdateStrategy
	return out
}
func NodeConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeConfig) *pb.NodeConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeConfig{}
	out.MachineType = direct.ValueOf(in.MachineType)
	out.DiskSizeGb = direct.ValueOf(in.DiskSizeGb)
	out.OauthScopes = in.OauthScopes
	if in.ServiceAccountRef != nil {
		out.ServiceAccount = in.ServiceAccountRef.External
	}
	out.Metadata = in.Metadata
	out.ImageType = direct.ValueOf(in.ImageType)
	out.Labels = in.Labels
	out.LocalSsdCount = direct.ValueOf(in.LocalSsdCount)
	out.Tags = in.Tags
	out.Preemptible = direct.ValueOf(in.Preemptible)
	// MISSING: Accelerators
	out.SandboxConfig = SandboxConfig_ToProto(mapCtx, in.SandboxConfig)
	if in.NodeGroupRef != nil {
		out.NodeGroup = in.NodeGroupRef.External
	}
	out.ReservationAffinity = ReservationAffinity_ToProto(mapCtx, in.ReservationAffinity)
	out.DiskType = direct.ValueOf(in.DiskType)
	out.MinCpuPlatform = direct.ValueOf(in.MinCpuPlatform)
	out.WorkloadMetadataConfig = WorkloadMetadataConfig_ToProto(mapCtx, in.WorkloadMetadataConfig)
	// MISSING: Taints
	// MISSING: BootDiskKmsKey
	out.ShieldedInstanceConfig = ShieldedInstanceConfig_ToProto(mapCtx, in.ShieldedInstanceConfig)
	out.LinuxNodeConfig = LinuxNodeConfig_ToProto(mapCtx, in.LinuxNodeConfig)
	out.KubeletConfig = NodeKubeletConfig_ToProto(mapCtx, in.KubeletConfig)
	out.EphemeralStorageConfig = EphemeralStorageConfig_ToProto(mapCtx, in.EphemeralStorageConfig)
	out.GcfsConfig = GcfsConfig_ToProto(mapCtx, in.GcfsConfig)
	out.AdvancedMachineFeatures = AdvancedMachineFeatures_ToProto(mapCtx, in.AdvancedMachineFeatures)
	out.Gvnic = VirtualNIC_ToProto(mapCtx, in.Gvnic)
	out.Spot = direct.ValueOf(in.Spot)
	out.ConfidentialNodes = ConfidentialNodes_ToProto(mapCtx, in.ConfidentialNodes)
	if oneof := FastSocket_ToProto(mapCtx, in.FastSocket); oneof != nil {
		out.FastSocket = oneof
	}
	out.ResourceLabels = in.ResourceLabels
	// MISSING: LoggingConfig
	// MISSING: WindowsNodeConfig
	out.LocalNvmeSsdBlockConfig = LocalNvmeSsdBlockConfig_ToProto(mapCtx, in.LocalNvmeSsdBlockConfig)
	out.EphemeralStorageLocalSsdConfig = EphemeralStorageLocalSsdConfig_ToProto(mapCtx, in.EphemeralStorageLocalSsdConfig)
	out.SoleTenantConfig = SoleTenantConfig_ToProto(mapCtx, in.SoleTenantConfig)
	// MISSING: ContainerdConfig
	out.HostMaintenancePolicy = HostMaintenancePolicy_ToProto(mapCtx, in.HostMaintenancePolicy)
	// MISSING: ResourceManagerTags
	// MISSING: EnableConfidentialStorage
	// MISSING: SecondaryBootDisks
	// MISSING: SecondaryBootDiskUpdateStrategy
	return out
}
func NodeKubeletConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeKubeletConfig) *krm.NodeKubeletConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeKubeletConfig{}
	out.CpuManagerPolicy = direct.LazyPtr(in.GetCpuManagerPolicy())
	out.CpuCfsQuota = BoolValue_FromProto(mapCtx, in.GetCpuCfsQuota())
	out.CpuCfsQuotaPeriod = direct.LazyPtr(in.GetCpuCfsQuotaPeriod())
	out.PodPidsLimit = direct.LazyPtr(in.GetPodPidsLimit())
	// MISSING: InsecureKubeletReadonlyPortEnabled
	return out
}
func NodeKubeletConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeKubeletConfig) *pb.NodeKubeletConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeKubeletConfig{}
	out.CpuManagerPolicy = direct.ValueOf(in.CpuManagerPolicy)
	out.CpuCfsQuota = BoolValue_ToProto(mapCtx, in.CpuCfsQuota)
	out.CpuCfsQuotaPeriod = direct.ValueOf(in.CpuCfsQuotaPeriod)
	out.PodPidsLimit = direct.ValueOf(in.PodPidsLimit)
	// MISSING: InsecureKubeletReadonlyPortEnabled
	return out
}
func NodeManagement_FromProto(mapCtx *direct.MapContext, in *pb.NodeManagement) *krm.NodeManagement {
	if in == nil {
		return nil
	}
	out := &krm.NodeManagement{}
	out.AutoUpgrade = direct.LazyPtr(in.GetAutoUpgrade())
	out.AutoRepair = direct.LazyPtr(in.GetAutoRepair())
	// MISSING: UpgradeOptions
	return out
}
func NodeManagement_ToProto(mapCtx *direct.MapContext, in *krm.NodeManagement) *pb.NodeManagement {
	if in == nil {
		return nil
	}
	out := &pb.NodeManagement{}
	out.AutoUpgrade = direct.ValueOf(in.AutoUpgrade)
	out.AutoRepair = direct.ValueOf(in.AutoRepair)
	// MISSING: UpgradeOptions
	return out
}
func NodeNetworkConfig_FromProto(mapCtx *direct.MapContext, in *pb.NodeNetworkConfig) *krm.NodeNetworkConfig {
	if in == nil {
		return nil
	}
	out := &krm.NodeNetworkConfig{}
	out.CreatePodRange = direct.LazyPtr(in.GetCreatePodRange())
	out.PodRange = direct.LazyPtr(in.GetPodRange())
	out.PodIpv4CidrBlock = direct.LazyPtr(in.GetPodIpv4CidrBlock())
	out.EnablePrivateNodes = in.EnablePrivateNodes
	// MISSING: NetworkPerformanceConfig
	out.PodCidrOverprovisionConfig = PodCIDROverprovisionConfig_FromProto(mapCtx, in.GetPodCidrOverprovisionConfig())
	out.AdditionalNodeNetworkConfigs = direct.Slice_FromProto(mapCtx, in.AdditionalNodeNetworkConfigs, AdditionalNodeNetworkConfig_FromProto)
	out.AdditionalPodNetworkConfigs = direct.Slice_FromProto(mapCtx, in.AdditionalPodNetworkConfigs, AdditionalPodNetworkConfig_FromProto)
	// MISSING: PodIpv4RangeUtilization
	return out
}
func NodeNetworkConfig_ToProto(mapCtx *direct.MapContext, in *krm.NodeNetworkConfig) *pb.NodeNetworkConfig {
	if in == nil {
		return nil
	}
	out := &pb.NodeNetworkConfig{}
	out.CreatePodRange = direct.ValueOf(in.CreatePodRange)
	out.PodRange = direct.ValueOf(in.PodRange)
	out.PodIpv4CidrBlock = direct.ValueOf(in.PodIpv4CidrBlock)
	out.EnablePrivateNodes = in.EnablePrivateNodes
	// MISSING: NetworkPerformanceConfig
	out.PodCidrOverprovisionConfig = PodCIDROverprovisionConfig_ToProto(mapCtx, in.PodCidrOverprovisionConfig)
	out.AdditionalNodeNetworkConfigs = direct.Slice_ToProto(mapCtx, in.AdditionalNodeNetworkConfigs, AdditionalNodeNetworkConfig_ToProto)
	out.AdditionalPodNetworkConfigs = direct.Slice_ToProto(mapCtx, in.AdditionalPodNetworkConfigs, AdditionalPodNetworkConfig_ToProto)
	// MISSING: PodIpv4RangeUtilization
	return out
}
func NodePoolAutoscaling_FromProto(mapCtx *direct.MapContext, in *pb.NodePoolAutoscaling) *krm.NodePoolAutoscaling {
	if in == nil {
		return nil
	}
	out := &krm.NodePoolAutoscaling{}
	// MISSING: Enabled
	out.MinNodeCount = direct.LazyPtr(in.GetMinNodeCount())
	out.MaxNodeCount = direct.LazyPtr(in.GetMaxNodeCount())
	// MISSING: Autoprovisioned
	out.LocationPolicy = direct.Enum_FromProto(mapCtx, in.GetLocationPolicy())
	out.TotalMinNodeCount = direct.LazyPtr(in.GetTotalMinNodeCount())
	out.TotalMaxNodeCount = direct.LazyPtr(in.GetTotalMaxNodeCount())
	return out
}
func NodePoolAutoscaling_ToProto(mapCtx *direct.MapContext, in *krm.NodePoolAutoscaling) *pb.NodePoolAutoscaling {
	if in == nil {
		return nil
	}
	out := &pb.NodePoolAutoscaling{}
	// MISSING: Enabled
	out.MinNodeCount = direct.ValueOf(in.MinNodeCount)
	out.MaxNodeCount = direct.ValueOf(in.MaxNodeCount)
	// MISSING: Autoprovisioned
	out.LocationPolicy = direct.Enum_ToProto[pb.NodePoolAutoscaling_LocationPolicy](mapCtx, in.LocationPolicy)
	out.TotalMinNodeCount = direct.ValueOf(in.TotalMinNodeCount)
	out.TotalMaxNodeCount = direct.ValueOf(in.TotalMaxNodeCount)
	return out
}
func NodePool_PlacementPolicy_FromProto(mapCtx *direct.MapContext, in *pb.NodePool_PlacementPolicy) *krm.NodePool_PlacementPolicy {
	if in == nil {
		return nil
	}
	out := &krm.NodePool_PlacementPolicy{}
	out.Type = direct.Enum_FromProto(mapCtx, in.GetType())
	out.TpuTopology = direct.LazyPtr(in.GetTpuTopology())
	if in.GetPolicyName() != "" {
		out.PolicyNameRef = &v1alpha1.ResourceRef{External: in.GetPolicyName()}
	}
	return out
}
func NodePool_PlacementPolicy_ToProto(mapCtx *direct.MapContext, in *krm.NodePool_PlacementPolicy) *pb.NodePool_PlacementPolicy {
	if in == nil {
		return nil
	}
	out := &pb.NodePool_PlacementPolicy{}
	out.Type = direct.Enum_ToProto[pb.NodePool_PlacementPolicy_Type](mapCtx, in.Type)
	out.TpuTopology = direct.ValueOf(in.TpuTopology)
	if in.PolicyNameRef != nil {
		out.PolicyName = in.PolicyNameRef.External
	}
	return out
}
func NodePool_QueuedProvisioning_FromProto(mapCtx *direct.MapContext, in *pb.NodePool_QueuedProvisioning) *krm.NodePool_QueuedProvisioning {
	if in == nil {
		return nil
	}
	out := &krm.NodePool_QueuedProvisioning{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}
func NodePool_QueuedProvisioning_ToProto(mapCtx *direct.MapContext, in *krm.NodePool_QueuedProvisioning) *pb.NodePool_QueuedProvisioning {
	if in == nil {
		return nil
	}
	out := &pb.NodePool_QueuedProvisioning{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
func NodePool_UpdateInfo_FromProto(mapCtx *direct.MapContext, in *pb.NodePool_UpdateInfo) *krm.NodePool_UpdateInfo {
	if in == nil {
		return nil
	}
	out := &krm.NodePool_UpdateInfo{}
	out.BlueGreenInfo = NodePool_UpdateInfo_BlueGreenInfo_FromProto(mapCtx, in.GetBlueGreenInfo())
	return out
}
func NodePool_UpdateInfo_ToProto(mapCtx *direct.MapContext, in *krm.NodePool_UpdateInfo) *pb.NodePool_UpdateInfo {
	if in == nil {
		return nil
	}
	out := &pb.NodePool_UpdateInfo{}
	out.BlueGreenInfo = NodePool_UpdateInfo_BlueGreenInfo_ToProto(mapCtx, in.BlueGreenInfo)
	return out
}
func NodePool_UpdateInfo_BlueGreenInfo_FromProto(mapCtx *direct.MapContext, in *pb.NodePool_UpdateInfo_BlueGreenInfo) *krm.NodePool_UpdateInfo_BlueGreenInfo {
	if in == nil {
		return nil
	}
	out := &krm.NodePool_UpdateInfo_BlueGreenInfo{}
	out.Phase = direct.Enum_FromProto(mapCtx, in.GetPhase())
	out.BlueInstanceGroupUrls = in.BlueInstanceGroupUrls
	out.GreenInstanceGroupUrls = in.GreenInstanceGroupUrls
	out.BluePoolDeletionStartTime = direct.LazyPtr(in.GetBluePoolDeletionStartTime())
	out.GreenPoolVersion = direct.LazyPtr(in.GetGreenPoolVersion())
	return out
}
func NodePool_UpdateInfo_BlueGreenInfo_ToProto(mapCtx *direct.MapContext, in *krm.NodePool_UpdateInfo_BlueGreenInfo) *pb.NodePool_UpdateInfo_BlueGreenInfo {
	if in == nil {
		return nil
	}
	out := &pb.NodePool_UpdateInfo_BlueGreenInfo{}
	out.Phase = direct.Enum_ToProto[pb.NodePool_UpdateInfo_BlueGreenInfo_Phase](mapCtx, in.Phase)
	out.BlueInstanceGroupUrls = in.BlueInstanceGroupUrls
	out.GreenInstanceGroupUrls = in.GreenInstanceGroupUrls
	out.BluePoolDeletionStartTime = direct.ValueOf(in.BluePoolDeletionStartTime)
	out.GreenPoolVersion = direct.ValueOf(in.GreenPoolVersion)
	return out
}
func NodePool_UpgradeSettings_FromProto(mapCtx *direct.MapContext, in *pb.NodePool_UpgradeSettings) *krm.NodePool_UpgradeSettings {
	if in == nil {
		return nil
	}
	out := &krm.NodePool_UpgradeSettings{}
	out.MaxSurge = direct.LazyPtr(in.GetMaxSurge())
	out.MaxUnavailable = direct.LazyPtr(in.GetMaxUnavailable())
	out.Strategy = direct.Enum_FromProto(mapCtx, in.GetStrategy())
	out.BlueGreenSettings = BlueGreenSettings_FromProto(mapCtx, in.GetBlueGreenSettings())
	return out
}
func NodePool_UpgradeSettings_ToProto(mapCtx *direct.MapContext, in *krm.NodePool_UpgradeSettings) *pb.NodePool_UpgradeSettings {
	if in == nil {
		return nil
	}
	out := &pb.NodePool_UpgradeSettings{}
	out.MaxSurge = direct.ValueOf(in.MaxSurge)
	out.MaxUnavailable = direct.ValueOf(in.MaxUnavailable)
	if oneof := NodePool_UpgradeSettings_Strategy_ToProto(mapCtx, in.Strategy); oneof != nil {
		out.Strategy = oneof
	}
	if oneof := BlueGreenSettings_ToProto(mapCtx, in.BlueGreenSettings); oneof != nil {
		out.BlueGreenSettings = oneof
	}
	return out
}
func NodeTaint_FromProto(mapCtx *direct.MapContext, in *pb.NodeTaint) *krm.NodeTaint {
	if in == nil {
		return nil
	}
	out := &krm.NodeTaint{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Value = direct.LazyPtr(in.GetValue())
	out.Effect = direct.Enum_FromProto(mapCtx, in.GetEffect())
	return out
}
func NodeTaint_ToProto(mapCtx *direct.MapContext, in *krm.NodeTaint) *pb.NodeTaint {
	if in == nil {
		return nil
	}
	out := &pb.NodeTaint{}
	out.Key = direct.ValueOf(in.Key)
	out.Value = direct.ValueOf(in.Value)
	out.Effect = direct.Enum_ToProto[pb.NodeTaint_Effect](mapCtx, in.Effect)
	return out
}
func PodCIDROverprovisionConfig_FromProto(mapCtx *direct.MapContext, in *pb.PodCIDROverprovisionConfig) *krm.PodCIDROverprovisionConfig {
	if in == nil {
		return nil
	}
	out := &krm.PodCIDROverprovisionConfig{}
	// MISSING: Disable
	return out
}
func PodCIDROverprovisionConfig_ToProto(mapCtx *direct.MapContext, in *krm.PodCIDROverprovisionConfig) *pb.PodCIDROverprovisionConfig {
	if in == nil {
		return nil
	}
	out := &pb.PodCIDROverprovisionConfig{}
	// MISSING: Disable
	return out
}
func ReservationAffinity_FromProto(mapCtx *direct.MapContext, in *pb.ReservationAffinity) *krm.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &krm.ReservationAffinity{}
	out.ConsumeReservationType = direct.Enum_FromProto(mapCtx, in.GetConsumeReservationType())
	out.Key = direct.LazyPtr(in.GetKey())
	out.Values = in.Values
	return out
}
func ReservationAffinity_ToProto(mapCtx *direct.MapContext, in *krm.ReservationAffinity) *pb.ReservationAffinity {
	if in == nil {
		return nil
	}
	out := &pb.ReservationAffinity{}
	out.ConsumeReservationType = direct.Enum_ToProto[pb.ReservationAffinity_Type](mapCtx, in.ConsumeReservationType)
	out.Key = direct.ValueOf(in.Key)
	out.Values = in.Values
	return out
}
func SandboxConfig_FromProto(mapCtx *direct.MapContext, in *pb.SandboxConfig) *krm.SandboxConfig {
	if in == nil {
		return nil
	}
	out := &krm.SandboxConfig{}
	out.SandboxType = direct.LazyPtr(in.GetSandboxType())
	// MISSING: Type
	return out
}
func SandboxConfig_ToProto(mapCtx *direct.MapContext, in *krm.SandboxConfig) *pb.SandboxConfig {
	if in == nil {
		return nil
	}
	out := &pb.SandboxConfig{}
	out.SandboxType = direct.ValueOf(in.SandboxType)
	// MISSING: Type
	return out
}
func ShieldedInstanceConfig_FromProto(mapCtx *direct.MapContext, in *pb.ShieldedInstanceConfig) *krm.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &krm.ShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.LazyPtr(in.GetEnableSecureBoot())
	out.EnableIntegrityMonitoring = direct.LazyPtr(in.GetEnableIntegrityMonitoring())
	return out
}
func ShieldedInstanceConfig_ToProto(mapCtx *direct.MapContext, in *krm.ShieldedInstanceConfig) *pb.ShieldedInstanceConfig {
	if in == nil {
		return nil
	}
	out := &pb.ShieldedInstanceConfig{}
	out.EnableSecureBoot = direct.ValueOf(in.EnableSecureBoot)
	out.EnableIntegrityMonitoring = direct.ValueOf(in.EnableIntegrityMonitoring)
	return out
}
func SoleTenantConfig_FromProto(mapCtx *direct.MapContext, in *pb.SoleTenantConfig) *krm.SoleTenantConfig {
	if in == nil {
		return nil
	}
	out := &krm.SoleTenantConfig{}

	// MISSING: NodeAffinities - RENAMED_TO_MATCH_CRD
	out.NodeAffinity = direct.Slice_FromProto(mapCtx, in.NodeAffinities, NodeAffinities_FromProto)
	return out
}
func SoleTenantConfig_ToProto(mapCtx *direct.MapContext, in *krm.SoleTenantConfig) *pb.SoleTenantConfig {
	if in == nil {
		return nil
	}
	out := &pb.SoleTenantConfig{}

	// MISSING: NodeAffinities - RENAMED_TO_MATCH_CRD
	out.NodeAffinities = direct.Slice_ToProto(mapCtx, in.NodeAffinity, NodeAffinity_ToProto)
	return out
}
func SoleTenantConfig_NodeAffinity_FromProto(mapCtx *direct.MapContext, in *pb.SoleTenantConfig_NodeAffinity) *krm.SoleTenantConfig_NodeAffinity {
	if in == nil {
		return nil
	}
	out := &krm.SoleTenantConfig_NodeAffinity{}
	out.Key = direct.LazyPtr(in.GetKey())
	out.Operator = direct.Enum_FromProto(mapCtx, in.GetOperator())
	out.Values = in.Values
	return out
}
func SoleTenantConfig_NodeAffinity_ToProto(mapCtx *direct.MapContext, in *krm.SoleTenantConfig_NodeAffinity) *pb.SoleTenantConfig_NodeAffinity {
	if in == nil {
		return nil
	}
	out := &pb.SoleTenantConfig_NodeAffinity{}
	out.Key = direct.ValueOf(in.Key)
	out.Operator = direct.Enum_ToProto[pb.SoleTenantConfig_NodeAffinity_Operator](mapCtx, in.Operator)
	out.Values = in.Values
	return out
}
func StatusCondition_FromProto(mapCtx *direct.MapContext, in *pb.StatusCondition) *krm.StatusCondition {
	if in == nil {
		return nil
	}
	out := &krm.StatusCondition{}
	out.Code = direct.Enum_FromProto(mapCtx, in.GetCode())
	out.Message = direct.LazyPtr(in.GetMessage())
	out.CanonicalCode = direct.Enum_FromProto(mapCtx, in.GetCanonicalCode())
	return out
}
func StatusCondition_ToProto(mapCtx *direct.MapContext, in *krm.StatusCondition) *pb.StatusCondition {
	if in == nil {
		return nil
	}
	out := &pb.StatusCondition{}
	out.Code = direct.Enum_ToProto[pb.StatusCondition_Code](mapCtx, in.Code)
	out.Message = direct.ValueOf(in.Message)
	out.CanonicalCode = direct.Enum_ToProto[code.Code](mapCtx, in.CanonicalCode)
	return out
}
func VirtualNIC_FromProto(mapCtx *direct.MapContext, in *pb.VirtualNIC) *krm.VirtualNIC {
	if in == nil {
		return nil
	}
	out := &krm.VirtualNIC{}
	out.Enabled = direct.LazyPtr(in.GetEnabled())
	return out
}
func VirtualNIC_ToProto(mapCtx *direct.MapContext, in *krm.VirtualNIC) *pb.VirtualNIC {
	if in == nil {
		return nil
	}
	out := &pb.VirtualNIC{}
	out.Enabled = direct.ValueOf(in.Enabled)
	return out
}
func WorkloadMetadataConfig_FromProto(mapCtx *direct.MapContext, in *pb.WorkloadMetadataConfig) *krm.WorkloadMetadataConfig {
	if in == nil {
		return nil
	}
	out := &krm.WorkloadMetadataConfig{}
	out.NodeMetadata = direct.Enum_FromProto(mapCtx, in.GetNodeMetadata())
	out.Mode = direct.Enum_FromProto(mapCtx, in.GetMode())
	return out
}
func WorkloadMetadataConfig_ToProto(mapCtx *direct.MapContext, in *krm.WorkloadMetadataConfig) *pb.WorkloadMetadataConfig {
	if in == nil {
		return nil
	}
	out := &pb.WorkloadMetadataConfig{}
	out.NodeMetadata = direct.Enum_ToProto[pb.WorkloadMetadataConfig_NodeMetadata](mapCtx, in.NodeMetadata)
	out.Mode = direct.Enum_ToProto[pb.WorkloadMetadataConfig_Mode](mapCtx, in.Mode)
	return out
}
