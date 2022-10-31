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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SecurityGroupEgressObservation struct {
	// Id of the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecurityGroupEgressParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// +kubebuilder:validation:Optional
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// +kubebuilder:validation:Optional
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type SecurityGroupIngressObservation struct {
	// Id of the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SecurityGroupIngressParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	FromPort *float64 `json:"fromPort,omitempty" tf:"from_port,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PredefinedTarget *string `json:"predefinedTarget,omitempty" tf:"predefined_target,omitempty"`

	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// +kubebuilder:validation:Optional
	SecurityGroupID *string `json:"securityGroupId,omitempty" tf:"security_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	ToPort *float64 `json:"toPort,omitempty" tf:"to_port,omitempty"`

	// +kubebuilder:validation:Optional
	V4CidrBlocks []*string `json:"v4CidrBlocks,omitempty" tf:"v4_cidr_blocks,omitempty"`

	// +kubebuilder:validation:Optional
	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type SecurityGroupObservation struct {
	// Creation timestamp of this security group.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// +kubebuilder:validation:Optional
	Egress []SecurityGroupEgressObservation `json:"egress,omitempty" tf:"egress,omitempty"`

	// Id of the rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Ingress []SecurityGroupIngressObservation `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// Status of this security group.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`
}

type SecurityGroupParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	Egress []SecurityGroupEgressParameters `json:"egress,omitempty" tf:"egress,omitempty"`

	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	Ingress []SecurityGroupIngressParameters `json:"ingress,omitempty" tf:"ingress,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`
}

// SecurityGroupSpec defines the desired state of SecurityGroup
type SecurityGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityGroupParameters `json:"forProvider"`
}

// SecurityGroupStatus defines the observed state of SecurityGroup.
type SecurityGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroup is the Schema for the SecurityGroups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type SecurityGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SecurityGroupSpec   `json:"spec"`
	Status            SecurityGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityGroupList contains a list of SecurityGroups
type SecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityGroup `json:"items"`
}

// Repository type metadata.
var (
	SecurityGroup_Kind             = "SecurityGroup"
	SecurityGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityGroup_Kind}.String()
	SecurityGroup_KindAPIVersion   = SecurityGroup_Kind + "." + CRDGroupVersion.String()
	SecurityGroup_GroupVersionKind = CRDGroupVersion.WithKind(SecurityGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityGroup{}, &SecurityGroupList{})
}
