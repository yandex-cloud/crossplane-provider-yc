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

type MySQLDatabaseInitParameters struct {

	// +crossplane:generate:reference:type=MySQLCluster
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a MySQLCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a MySQLCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// The name of the database.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type MySQLDatabaseObservation struct {
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the database.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type MySQLDatabaseParameters struct {

	// +crossplane:generate:reference:type=MySQLCluster
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a MySQLCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a MySQLCluster to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// The name of the database.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// MySQLDatabaseSpec defines the desired state of MySQLDatabase
type MySQLDatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MySQLDatabaseParameters `json:"forProvider"`
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
	InitProvider MySQLDatabaseInitParameters `json:"initProvider,omitempty"`
}

// MySQLDatabaseStatus defines the observed state of MySQLDatabase.
type MySQLDatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MySQLDatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MySQLDatabase is the Schema for the MySQLDatabases API. Manages a MySQL database within Yandex.Cloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type MySQLDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   MySQLDatabaseSpec   `json:"spec"`
	Status MySQLDatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLDatabaseList contains a list of MySQLDatabases
type MySQLDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLDatabase `json:"items"`
}

// Repository type metadata.
var (
	MySQLDatabase_Kind             = "MySQLDatabase"
	MySQLDatabase_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MySQLDatabase_Kind}.String()
	MySQLDatabase_KindAPIVersion   = MySQLDatabase_Kind + "." + CRDGroupVersion.String()
	MySQLDatabase_GroupVersionKind = CRDGroupVersion.WithKind(MySQLDatabase_Kind)
)

func init() {
	SchemeBuilder.Register(&MySQLDatabase{}, &MySQLDatabaseList{})
}
