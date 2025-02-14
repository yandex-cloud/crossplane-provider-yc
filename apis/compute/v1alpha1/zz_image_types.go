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

type ImageInitParameters struct {

	// An optional description of the image. Provide this property when
	// you create a resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the image family to which this image belongs.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/crossplane-provider-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the image.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Minimum size in GB of the disk that will be created from this image.
	MinDiskSize *float64 `json:"minDiskSize,omitempty" tf:"min_disk_size,omitempty"`

	// Name of the disk.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Operating system type that is contained in the image. Possible values: "LINUX", "WINDOWS".
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// Optimize the image to create a disk.
	Pooled *bool `json:"pooled,omitempty" tf:"pooled,omitempty"`

	// License IDs that indicate which licenses are
	// attached to this image.
	// +listType=set
	ProductIds []*string `json:"productIds,omitempty" tf:"product_ids,omitempty"`

	// The ID of a disk to use as the source of the
	// image. Changing this ID forces a new resource to be created.
	SourceDisk *string `json:"sourceDisk,omitempty" tf:"source_disk,omitempty"`

	// The name of the family to use as the source of the new image.
	// The ID of the latest image is taken from the "standard-images" folder. Changing the family forces
	// a new resource to be created.
	SourceFamily *string `json:"sourceFamily,omitempty" tf:"source_family,omitempty"`

	// The ID of an existing image to use as the source of the
	// image. Changing this ID forces a new resource to be created.
	SourceImage *string `json:"sourceImage,omitempty" tf:"source_image,omitempty"`

	// The ID of a snapshot to use as the source of the
	// image. Changing this ID forces a new resource to be created.
	SourceSnapshot *string `json:"sourceSnapshot,omitempty" tf:"source_snapshot,omitempty"`

	// The URL to use as the source of the
	// image. Changing this URL forces a new resource to be created.
	SourceURL *string `json:"sourceUrl,omitempty" tf:"source_url,omitempty"`
}

type ImageObservation struct {

	// Creation timestamp of the image.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// An optional description of the image. Provide this property when
	// you create a resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the image family to which this image belongs.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the image.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Minimum size in GB of the disk that will be created from this image.
	MinDiskSize *float64 `json:"minDiskSize,omitempty" tf:"min_disk_size,omitempty"`

	// Name of the disk.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Operating system type that is contained in the image. Possible values: "LINUX", "WINDOWS".
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// Optimize the image to create a disk.
	Pooled *bool `json:"pooled,omitempty" tf:"pooled,omitempty"`

	// License IDs that indicate which licenses are
	// attached to this image.
	// +listType=set
	ProductIds []*string `json:"productIds,omitempty" tf:"product_ids,omitempty"`

	// The size of the image, specified in GB.
	Size *float64 `json:"size,omitempty" tf:"size,omitempty"`

	// The ID of a disk to use as the source of the
	// image. Changing this ID forces a new resource to be created.
	SourceDisk *string `json:"sourceDisk,omitempty" tf:"source_disk,omitempty"`

	// The name of the family to use as the source of the new image.
	// The ID of the latest image is taken from the "standard-images" folder. Changing the family forces
	// a new resource to be created.
	SourceFamily *string `json:"sourceFamily,omitempty" tf:"source_family,omitempty"`

	// The ID of an existing image to use as the source of the
	// image. Changing this ID forces a new resource to be created.
	SourceImage *string `json:"sourceImage,omitempty" tf:"source_image,omitempty"`

	// The ID of a snapshot to use as the source of the
	// image. Changing this ID forces a new resource to be created.
	SourceSnapshot *string `json:"sourceSnapshot,omitempty" tf:"source_snapshot,omitempty"`

	// The URL to use as the source of the
	// image. Changing this URL forces a new resource to be created.
	SourceURL *string `json:"sourceUrl,omitempty" tf:"source_url,omitempty"`

	// The status of the image.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type ImageParameters struct {

	// An optional description of the image. Provide this property when
	// you create a resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the image family to which this image belongs.
	// +kubebuilder:validation:Optional
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// The ID of the folder that the resource belongs to. If it
	// is not provided, the default provider folder is used.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/crossplane-provider-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the image.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Minimum size in GB of the disk that will be created from this image.
	// +kubebuilder:validation:Optional
	MinDiskSize *float64 `json:"minDiskSize,omitempty" tf:"min_disk_size,omitempty"`

	// Name of the disk.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Operating system type that is contained in the image. Possible values: "LINUX", "WINDOWS".
	// +kubebuilder:validation:Optional
	OsType *string `json:"osType,omitempty" tf:"os_type,omitempty"`

	// Optimize the image to create a disk.
	// +kubebuilder:validation:Optional
	Pooled *bool `json:"pooled,omitempty" tf:"pooled,omitempty"`

	// License IDs that indicate which licenses are
	// attached to this image.
	// +kubebuilder:validation:Optional
	// +listType=set
	ProductIds []*string `json:"productIds,omitempty" tf:"product_ids,omitempty"`

	// The ID of a disk to use as the source of the
	// image. Changing this ID forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SourceDisk *string `json:"sourceDisk,omitempty" tf:"source_disk,omitempty"`

	// The name of the family to use as the source of the new image.
	// The ID of the latest image is taken from the "standard-images" folder. Changing the family forces
	// a new resource to be created.
	// +kubebuilder:validation:Optional
	SourceFamily *string `json:"sourceFamily,omitempty" tf:"source_family,omitempty"`

	// The ID of an existing image to use as the source of the
	// image. Changing this ID forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SourceImage *string `json:"sourceImage,omitempty" tf:"source_image,omitempty"`

	// The ID of a snapshot to use as the source of the
	// image. Changing this ID forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SourceSnapshot *string `json:"sourceSnapshot,omitempty" tf:"source_snapshot,omitempty"`

	// The URL to use as the source of the
	// image. Changing this URL forces a new resource to be created.
	// +kubebuilder:validation:Optional
	SourceURL *string `json:"sourceUrl,omitempty" tf:"source_url,omitempty"`
}

// ImageSpec defines the desired state of Image
type ImageSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ImageParameters `json:"forProvider"`
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
	InitProvider ImageInitParameters `json:"initProvider,omitempty"`
}

// ImageStatus defines the observed state of Image.
type ImageStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ImageObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Image is the Schema for the Images API. Creates a VM image for the Yandex Compute service from an existing tarball.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ImageSpec   `json:"spec"`
	Status            ImageStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ImageList contains a list of Images
type ImageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Image `json:"items"`
}

// Repository type metadata.
var (
	Image_Kind             = "Image"
	Image_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Image_Kind}.String()
	Image_KindAPIVersion   = Image_Kind + "." + CRDGroupVersion.String()
	Image_GroupVersionKind = CRDGroupVersion.WithKind(Image_Kind)
)

func init() {
	SchemeBuilder.Register(&Image{}, &ImageList{})
}
