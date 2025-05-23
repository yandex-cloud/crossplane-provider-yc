/*
Copyright 2022 YANDEX LLC
This is modified version of the software, made by the Crossplane Authors
and available at: https://github.com/crossplane-contrib/provider-jet-template

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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type GpuClusterInitParameters struct {

	// Description of the GPU cluster. Provide this property when you create a resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder that the GPU cluster belongs to. If it is not provided, the default
	// provider folder is used.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/crossplane-provider-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// Type of interconnect between nodes to use in GPU cluster. Type infiniband is set by default,
	// and it is the only one available at the moment.
	InterconnectType *string `json:"interconnectType,omitempty" tf:"interconnect_type,omitempty"`

	// Labels to assign to this GPU cluster. A list of key/value pairs. For details about the concept,
	// see documentation.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the GPU cluster. Provide this property when you create a resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Availability zone where the GPU cluster will reside.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type GpuClusterObservation struct {

	// Creation timestamp of the GPU cluster.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Description of the GPU cluster. Provide this property when you create a resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder that the GPU cluster belongs to. If it is not provided, the default
	// provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Type of interconnect between nodes to use in GPU cluster. Type infiniband is set by default,
	// and it is the only one available at the moment.
	InterconnectType *string `json:"interconnectType,omitempty" tf:"interconnect_type,omitempty"`

	// Labels to assign to this GPU cluster. A list of key/value pairs. For details about the concept,
	// see documentation.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the GPU cluster. Provide this property when you create a resource.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The status of the GPU cluster.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Availability zone where the GPU cluster will reside.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type GpuClusterParameters struct {

	// Description of the GPU cluster. Provide this property when you create a resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder that the GPU cluster belongs to. If it is not provided, the default
	// provider folder is used.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/crossplane-provider-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// Type of interconnect between nodes to use in GPU cluster. Type infiniband is set by default,
	// and it is the only one available at the moment.
	// +kubebuilder:validation:Optional
	InterconnectType *string `json:"interconnectType,omitempty" tf:"interconnect_type,omitempty"`

	// Labels to assign to this GPU cluster. A list of key/value pairs. For details about the concept,
	// see documentation.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the GPU cluster. Provide this property when you create a resource.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Availability zone where the GPU cluster will reside.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// GpuClusterSpec defines the desired state of GpuCluster
type GpuClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GpuClusterParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider GpuClusterInitParameters `json:"initProvider,omitempty"`
}

// GpuClusterStatus defines the observed state of GpuCluster.
type GpuClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GpuClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// GpuCluster is the Schema for the GpuClusters API. GPU Cluster connects multiple Compute GPU Instances in the same availability zone with high-speed low-latency network.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type GpuCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GpuClusterSpec   `json:"spec"`
	Status            GpuClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GpuClusterList contains a list of GpuClusters
type GpuClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GpuCluster `json:"items"`
}

// Repository type metadata.
var (
	GpuCluster_Kind             = "GpuCluster"
	GpuCluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: GpuCluster_Kind}.String()
	GpuCluster_KindAPIVersion   = GpuCluster_Kind + "." + CRDGroupVersion.String()
	GpuCluster_GroupVersionKind = CRDGroupVersion.WithKind(GpuCluster_Kind)
)

func init() {
	SchemeBuilder.Register(&GpuCluster{}, &GpuClusterList{})
}
