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

type RouteTableInitParameters struct {

	// An optional description of the route table. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder to which the resource belongs.
	// If omitted, the provider folder is used.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/crossplane-provider-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// Labels to assign to this route table. A list of key/value pairs.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the route table. Provided by the client when the route table is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network this route table belongs to.
	// +crossplane:generate:reference:type=Network
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// A list of static route records for the route table. The structure is documented below.
	StaticRoute []StaticRouteInitParameters `json:"staticRoute,omitempty" tf:"static_route,omitempty"`
}

type RouteTableObservation struct {

	// Creation timestamp of the route table.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// An optional description of the route table. Provide this property when
	// you create the resource.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder to which the resource belongs.
	// If omitted, the provider folder is used.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Labels to assign to this route table. A list of key/value pairs.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the route table. Provided by the client when the route table is created.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network this route table belongs to.
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// A list of static route records for the route table. The structure is documented below.
	StaticRoute []StaticRouteObservation `json:"staticRoute,omitempty" tf:"static_route,omitempty"`
}

type RouteTableParameters struct {

	// An optional description of the route table. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the folder to which the resource belongs.
	// If omitted, the provider folder is used.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/crossplane-provider-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// Labels to assign to this route table. A list of key/value pairs.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Name of the route table. Provided by the client when the route table is created.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the network this route table belongs to.
	// +crossplane:generate:reference:type=Network
	// +kubebuilder:validation:Optional
	NetworkID *string `json:"networkId,omitempty" tf:"network_id,omitempty"`

	// Reference to a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDRef *v1.Reference `json:"networkIdRef,omitempty" tf:"-"`

	// Selector for a Network to populate networkId.
	// +kubebuilder:validation:Optional
	NetworkIDSelector *v1.Selector `json:"networkIdSelector,omitempty" tf:"-"`

	// A list of static route records for the route table. The structure is documented below.
	// +kubebuilder:validation:Optional
	StaticRoute []StaticRouteParameters `json:"staticRoute,omitempty" tf:"static_route,omitempty"`
}

type StaticRouteInitParameters struct {

	// Route prefix in CIDR notation.
	DestinationPrefix *string `json:"destinationPrefix,omitempty" tf:"destination_prefix,omitempty"`

	// ID of the gateway used ad next hop.
	// +crossplane:generate:reference:type=Gateway
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// Reference to a Gateway to populate gatewayId.
	// +kubebuilder:validation:Optional
	GatewayIDRef *v1.Reference `json:"gatewayIdRef,omitempty" tf:"-"`

	// Selector for a Gateway to populate gatewayId.
	// +kubebuilder:validation:Optional
	GatewayIDSelector *v1.Selector `json:"gatewayIdSelector,omitempty" tf:"-"`

	// Address of the next hop.
	NextHopAddress *string `json:"nextHopAddress,omitempty" tf:"next_hop_address,omitempty"`
}

type StaticRouteObservation struct {

	// Route prefix in CIDR notation.
	DestinationPrefix *string `json:"destinationPrefix,omitempty" tf:"destination_prefix,omitempty"`

	// ID of the gateway used ad next hop.
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// Address of the next hop.
	NextHopAddress *string `json:"nextHopAddress,omitempty" tf:"next_hop_address,omitempty"`
}

type StaticRouteParameters struct {

	// Route prefix in CIDR notation.
	// +kubebuilder:validation:Optional
	DestinationPrefix *string `json:"destinationPrefix,omitempty" tf:"destination_prefix,omitempty"`

	// ID of the gateway used ad next hop.
	// +crossplane:generate:reference:type=Gateway
	// +kubebuilder:validation:Optional
	GatewayID *string `json:"gatewayId,omitempty" tf:"gateway_id,omitempty"`

	// Reference to a Gateway to populate gatewayId.
	// +kubebuilder:validation:Optional
	GatewayIDRef *v1.Reference `json:"gatewayIdRef,omitempty" tf:"-"`

	// Selector for a Gateway to populate gatewayId.
	// +kubebuilder:validation:Optional
	GatewayIDSelector *v1.Selector `json:"gatewayIdSelector,omitempty" tf:"-"`

	// Address of the next hop.
	// +kubebuilder:validation:Optional
	NextHopAddress *string `json:"nextHopAddress,omitempty" tf:"next_hop_address,omitempty"`
}

// RouteTableSpec defines the desired state of RouteTable
type RouteTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RouteTableParameters `json:"forProvider"`
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
	InitProvider RouteTableInitParameters `json:"initProvider,omitempty"`
}

// RouteTableStatus defines the observed state of RouteTable.
type RouteTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RouteTable is the Schema for the RouteTables API. A VPC route table is a virtual version of the traditional route table on router device.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RouteTableSpec   `json:"spec"`
	Status            RouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RouteTableList contains a list of RouteTables
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

// Repository type metadata.
var (
	RouteTable_Kind             = "RouteTable"
	RouteTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RouteTable_Kind}.String()
	RouteTable_KindAPIVersion   = RouteTable_Kind + "." + CRDGroupVersion.String()
	RouteTable_GroupVersionKind = CRDGroupVersion.WithKind(RouteTable_Kind)
)

func init() {
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
}
