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

type RegistryIamBindingObservation struct {
}

type RegistryIamBindingParameters struct {

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	RegistryID *string `json:"registryId" tf:"registry_id,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	SleepAfter *int64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// RegistryIamBindingSpec defines the desired state of RegistryIamBinding
type RegistryIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegistryIamBindingParameters `json:"forProvider"`
}

// RegistryIamBindingStatus defines the observed state of RegistryIamBinding.
type RegistryIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegistryIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryIamBinding is the Schema for the RegistryIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type RegistryIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegistryIamBindingSpec   `json:"spec"`
	Status            RegistryIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegistryIamBindingList contains a list of RegistryIamBindings
type RegistryIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegistryIamBinding `json:"items"`
}

// Repository type metadata.
var (
	RegistryIamBinding_Kind             = "RegistryIamBinding"
	RegistryIamBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegistryIamBinding_Kind}.String()
	RegistryIamBinding_KindAPIVersion   = RegistryIamBinding_Kind + "." + CRDGroupVersion.String()
	RegistryIamBinding_GroupVersionKind = CRDGroupVersion.WithKind(RegistryIamBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&RegistryIamBinding{}, &RegistryIamBindingList{})
}