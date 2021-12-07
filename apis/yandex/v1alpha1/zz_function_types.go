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

type ContentObservation struct {
}

type ContentParameters struct {

	// +kubebuilder:validation:Required
	ZipFilename *string `json:"zipFilename" tf:"zip_filename,omitempty"`
}

type FunctionObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	ImageSize *int64 `json:"imageSize,omitempty" tf:"image_size,omitempty"`

	LoggroupID *string `json:"loggroupId,omitempty" tf:"loggroup_id,omitempty"`

	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type FunctionParameters struct {

	// +kubebuilder:validation:Optional
	Content []ContentParameters `json:"content,omitempty" tf:"content,omitempty"`

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Entrypoint *string `json:"entrypoint" tf:"entrypoint,omitempty"`

	// +kubebuilder:validation:Optional
	Environment map[string]*string `json:"environment,omitempty" tf:"environment,omitempty"`

	// +kubebuilder:validation:Optional
	ExecutionTimeout *string `json:"executionTimeout,omitempty" tf:"execution_timeout,omitempty"`

	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Required
	Memory *int64 `json:"memory" tf:"memory,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Package []PackageParameters `json:"package,omitempty" tf:"package,omitempty"`

	// +kubebuilder:validation:Required
	Runtime *string `json:"runtime" tf:"runtime,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Required
	UserHash *string `json:"userHash" tf:"user_hash,omitempty"`
}

type PackageObservation struct {
}

type PackageParameters struct {

	// +kubebuilder:validation:Required
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`

	// +kubebuilder:validation:Required
	ObjectName *string `json:"objectName" tf:"object_name,omitempty"`

	// +kubebuilder:validation:Optional
	Sha256 *string `json:"sha256,omitempty" tf:"sha_256,omitempty"`
}

// FunctionSpec defines the desired state of Function
type FunctionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FunctionParameters `json:"forProvider"`
}

// FunctionStatus defines the observed state of Function.
type FunctionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FunctionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Function is the Schema for the Functions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type Function struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FunctionSpec   `json:"spec"`
	Status            FunctionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FunctionList contains a list of Functions
type FunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Function `json:"items"`
}

// Repository type metadata.
var (
	Function_Kind             = "Function"
	Function_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Function_Kind}.String()
	Function_KindAPIVersion   = Function_Kind + "." + CRDGroupVersion.String()
	Function_GroupVersionKind = CRDGroupVersion.WithKind(Function_Kind)
)

func init() {
	SchemeBuilder.Register(&Function{}, &FunctionList{})
}