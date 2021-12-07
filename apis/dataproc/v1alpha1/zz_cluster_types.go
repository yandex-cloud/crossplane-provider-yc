/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AutoscalingConfigObservation struct {
}

type AutoscalingConfigParameters struct {

	// +kubebuilder:validation:Optional
	CPUUtilizationTarget *float64 `json:"cpuUtilizationTarget,omitempty" tf:"cpu_utilization_target,omitempty"`

	// +kubebuilder:validation:Optional
	DecommissionTimeout *int64 `json:"decommissionTimeout,omitempty" tf:"decommission_timeout,omitempty"`

	// +kubebuilder:validation:Required
	MaxHostsCount *int64 `json:"maxHostsCount" tf:"max_hosts_count,omitempty"`

	// +kubebuilder:validation:Optional
	MeasurementDuration *int64 `json:"measurementDuration,omitempty" tf:"measurement_duration,omitempty"`

	// +kubebuilder:validation:Optional
	Preemptible *bool `json:"preemptible,omitempty" tf:"preemptible,omitempty"`

	// +kubebuilder:validation:Optional
	StabilizationDuration *int64 `json:"stabilizationDuration,omitempty" tf:"stabilization_duration,omitempty"`

	// +kubebuilder:validation:Optional
	WarmupDuration *int64 `json:"warmupDuration,omitempty" tf:"warmup_duration,omitempty"`
}

type ClusterConfigObservation struct {
}

type ClusterConfigParameters struct {

	// +kubebuilder:validation:Optional
	Hadoop []HadoopParameters `json:"hadoop,omitempty" tf:"hadoop,omitempty"`

	// +kubebuilder:validation:Required
	SubclusterSpec []SubclusterSpecParameters `json:"subclusterSpec" tf:"subcluster_spec,omitempty"`

	// +kubebuilder:validation:Optional
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type ClusterObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
}

type ClusterParameters struct {

	// +kubebuilder:validation:Optional
	Bucket *string `json:"bucket,omitempty" tf:"bucket,omitempty"`

	// +kubebuilder:validation:Required
	ClusterConfig []ClusterConfigParameters `json:"clusterConfig" tf:"cluster_config,omitempty"`

	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	HostGroupIds []*string `json:"hostGroupIds,omitempty" tf:"host_group_ids,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupIds []*string `json:"securityGroupIds,omitempty" tf:"security_group_ids,omitempty"`

	// +kubebuilder:validation:Required
	ServiceAccountID *string `json:"serviceAccountId" tf:"service_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	UIProxy *bool `json:"uiProxy,omitempty" tf:"ui_proxy,omitempty"`

	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type HadoopObservation struct {
}

type HadoopParameters struct {

	// +kubebuilder:validation:Optional
	Properties map[string]*string `json:"properties,omitempty" tf:"properties,omitempty"`

	// +kubebuilder:validation:Optional
	SSHPublicKeys []*string `json:"sshPublicKeys,omitempty" tf:"ssh_public_keys,omitempty"`

	// +kubebuilder:validation:Optional
	Services []*string `json:"services,omitempty" tf:"services,omitempty"`
}

type ResourcesObservation struct {
}

type ResourcesParameters struct {

	// +kubebuilder:validation:Required
	DiskSize *int64 `json:"diskSize" tf:"disk_size,omitempty"`

	// +kubebuilder:validation:Optional
	DiskTypeID *string `json:"diskTypeId,omitempty" tf:"disk_type_id,omitempty"`

	// +kubebuilder:validation:Required
	ResourcePresetID *string `json:"resourcePresetId" tf:"resource_preset_id,omitempty"`
}

type SubclusterSpecObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SubclusterSpecParameters struct {

	// +kubebuilder:validation:Optional
	AutoscalingConfig []AutoscalingConfigParameters `json:"autoscalingConfig,omitempty" tf:"autoscaling_config,omitempty"`

	// +kubebuilder:validation:Required
	HostsCount *int64 `json:"hostsCount" tf:"hosts_count,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Resources []ResourcesParameters `json:"resources" tf:"resources,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	SubnetID *string `json:"subnetId" tf:"subnet_id,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}
