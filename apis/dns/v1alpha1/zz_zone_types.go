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

type ZoneObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
}

type ZoneParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	PrivateNetworks []*string `json:"privateNetworks,omitempty" tf:"private_networks,omitempty"`

	// +kubebuilder:validation:Optional
	Public *bool `json:"public,omitempty" tf:"public,omitempty"`

	// +kubebuilder:validation:Required
	Zone *string `json:"zone" tf:"zone,omitempty"`
}

// ZoneSpec defines the desired state of Zone
type ZoneSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ZoneParameters `json:"forProvider"`
}

// ZoneStatus defines the observed state of Zone.
type ZoneStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ZoneObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Zone is the Schema for the Zones API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type Zone struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ZoneSpec   `json:"spec"`
	Status            ZoneStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ZoneList contains a list of Zones
type ZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Zone `json:"items"`
}

// Repository type metadata.
var (
	Zone_Kind             = "Zone"
	Zone_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Zone_Kind}.String()
	Zone_KindAPIVersion   = Zone_Kind + "." + CRDGroupVersion.String()
	Zone_GroupVersionKind = CRDGroupVersion.WithKind(Zone_Kind)
)

func init() {
	SchemeBuilder.Register(&Zone{}, &ZoneList{})
}
