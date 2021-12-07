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

type DhcpOptionsObservation struct {
}

type DhcpOptionsParameters struct {

	// +kubebuilder:validation:Optional
	DomainName *string `json:"domainName,omitempty" tf:"domain_name,omitempty"`

	// +kubebuilder:validation:Optional
	DomainNameServers []*string `json:"domainNameServers,omitempty" tf:"domain_name_servers,omitempty"`

	// +kubebuilder:validation:Optional
	NtpServers []*string `json:"ntpServers,omitempty" tf:"ntp_servers,omitempty"`
}

type SubnetObservation struct {
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	V6CidrBlocks []*string `json:"v6CidrBlocks,omitempty" tf:"v6_cidr_blocks,omitempty"`
}

type SubnetParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DhcpOptions []DhcpOptionsParameters `json:"dhcpOptions,omitempty" tf:"dhcp_options,omitempty"`

	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	NetworkID *string `json:"networkId" tf:"network_id,omitempty"`

	// +kubebuilder:validation:Optional
	RouteTableID *string `json:"routeTableId,omitempty" tf:"route_table_id,omitempty"`

	// +kubebuilder:validation:Required
	V4CidrBlocks []*string `json:"v4CidrBlocks" tf:"v4_cidr_blocks,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// SubnetSpec defines the desired state of Subnet
type SubnetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubnetParameters `json:"forProvider"`
}

// SubnetStatus defines the observed state of Subnet.
type SubnetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubnetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Subnet is the Schema for the Subnets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type Subnet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetSpec   `json:"spec"`
	Status            SubnetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubnetList contains a list of Subnets
type SubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subnet `json:"items"`
}

// Repository type metadata.
var (
	Subnet_Kind             = "Subnet"
	Subnet_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subnet_Kind}.String()
	Subnet_KindAPIVersion   = Subnet_Kind + "." + CRDGroupVersion.String()
	Subnet_GroupVersionKind = CRDGroupVersion.WithKind(Subnet_Kind)
)

func init() {
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
}