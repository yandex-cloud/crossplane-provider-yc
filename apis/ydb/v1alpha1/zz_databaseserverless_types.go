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

type DatabaseServerlessInitParameters struct {

	// Inhibits deletion of the database. Can be either true or false
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A description for the Yandex Database serverless cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder that the Yandex Database serverless cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/crossplane-provider-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the Yandex Database serverless cluster.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Location ID for the Yandex Database serverless cluster.
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// Name for the Yandex Database serverless cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ServerlessDatabase []ServerlessDatabaseInitParameters `json:"serverlessDatabase,omitempty" tf:"serverless_database,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type DatabaseServerlessObservation struct {

	// The Yandex Database serverless cluster creation timestamp.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// Full database path of the Yandex Database serverless cluster.
	// Useful for SDK configuration.
	DatabasePath *string `json:"databasePath,omitempty" tf:"database_path,omitempty"`

	// Inhibits deletion of the database. Can be either true or false
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A description for the Yandex Database serverless cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Document API endpoint of the Yandex Database serverless cluster.
	DocumentAPIEndpoint *string `json:"documentApiEndpoint,omitempty" tf:"document_api_endpoint,omitempty"`

	// ID of the folder that the Yandex Database serverless cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// ID of the Yandex Database serverless cluster.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A set of key/value label pairs to assign to the Yandex Database serverless cluster.
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Location ID for the Yandex Database serverless cluster.
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// Name for the Yandex Database serverless cluster.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	ServerlessDatabase []ServerlessDatabaseObservation `json:"serverlessDatabase,omitempty" tf:"serverless_database,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`

	// Status of the Yandex Database serverless cluster.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Whether TLS is enabled for the Yandex Database serverless cluster.
	// Useful for SDK configuration.
	TLSEnabled *bool `json:"tlsEnabled,omitempty" tf:"tls_enabled,omitempty"`

	// API endpoint of the Yandex Database serverless cluster.
	// Useful for SDK configuration.
	YdbAPIEndpoint *string `json:"ydbApiEndpoint,omitempty" tf:"ydb_api_endpoint,omitempty"`

	// Full endpoint of the Yandex Database serverless cluster.
	YdbFullEndpoint *string `json:"ydbFullEndpoint,omitempty" tf:"ydb_full_endpoint,omitempty"`
}

type DatabaseServerlessParameters struct {

	// Inhibits deletion of the database. Can be either true or false
	// +kubebuilder:validation:Optional
	DeletionProtection *bool `json:"deletionProtection,omitempty" tf:"deletion_protection,omitempty"`

	// A description for the Yandex Database serverless cluster.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// ID of the folder that the Yandex Database serverless cluster belongs to.
	// It will be deduced from provider configuration if not set explicitly.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/crossplane-provider-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// A set of key/value label pairs to assign to the Yandex Database serverless cluster.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Location ID for the Yandex Database serverless cluster.
	// +kubebuilder:validation:Optional
	LocationID *string `json:"locationId,omitempty" tf:"location_id,omitempty"`

	// Name for the Yandex Database serverless cluster.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	ServerlessDatabase []ServerlessDatabaseParameters `json:"serverlessDatabase,omitempty" tf:"serverless_database,omitempty"`

	// +kubebuilder:validation:Optional
	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type ServerlessDatabaseInitParameters struct {
	EnableThrottlingRcuLimit *bool `json:"enableThrottlingRcuLimit,omitempty" tf:"enable_throttling_rcu_limit,omitempty"`

	ProvisionedRcuLimit *float64 `json:"provisionedRcuLimit,omitempty" tf:"provisioned_rcu_limit,omitempty"`

	StorageSizeLimit *float64 `json:"storageSizeLimit,omitempty" tf:"storage_size_limit,omitempty"`

	ThrottlingRcuLimit *float64 `json:"throttlingRcuLimit,omitempty" tf:"throttling_rcu_limit,omitempty"`
}

type ServerlessDatabaseObservation struct {
	EnableThrottlingRcuLimit *bool `json:"enableThrottlingRcuLimit,omitempty" tf:"enable_throttling_rcu_limit,omitempty"`

	ProvisionedRcuLimit *float64 `json:"provisionedRcuLimit,omitempty" tf:"provisioned_rcu_limit,omitempty"`

	StorageSizeLimit *float64 `json:"storageSizeLimit,omitempty" tf:"storage_size_limit,omitempty"`

	ThrottlingRcuLimit *float64 `json:"throttlingRcuLimit,omitempty" tf:"throttling_rcu_limit,omitempty"`
}

type ServerlessDatabaseParameters struct {

	// +kubebuilder:validation:Optional
	EnableThrottlingRcuLimit *bool `json:"enableThrottlingRcuLimit,omitempty" tf:"enable_throttling_rcu_limit,omitempty"`

	// +kubebuilder:validation:Optional
	ProvisionedRcuLimit *float64 `json:"provisionedRcuLimit,omitempty" tf:"provisioned_rcu_limit,omitempty"`

	// +kubebuilder:validation:Optional
	StorageSizeLimit *float64 `json:"storageSizeLimit,omitempty" tf:"storage_size_limit,omitempty"`

	// +kubebuilder:validation:Optional
	ThrottlingRcuLimit *float64 `json:"throttlingRcuLimit,omitempty" tf:"throttling_rcu_limit,omitempty"`
}

// DatabaseServerlessSpec defines the desired state of DatabaseServerless
type DatabaseServerlessSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseServerlessParameters `json:"forProvider"`
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
	InitProvider DatabaseServerlessInitParameters `json:"initProvider,omitempty"`
}

// DatabaseServerlessStatus defines the observed state of DatabaseServerless.
type DatabaseServerlessStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseServerlessObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DatabaseServerless is the Schema for the DatabaseServerlesss API. Manages Yandex Database serverless cluster.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type DatabaseServerless struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   DatabaseServerlessSpec   `json:"spec"`
	Status DatabaseServerlessStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseServerlessList contains a list of DatabaseServerlesss
type DatabaseServerlessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DatabaseServerless `json:"items"`
}

// Repository type metadata.
var (
	DatabaseServerless_Kind             = "DatabaseServerless"
	DatabaseServerless_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DatabaseServerless_Kind}.String()
	DatabaseServerless_KindAPIVersion   = DatabaseServerless_Kind + "." + CRDGroupVersion.String()
	DatabaseServerless_GroupVersionKind = CRDGroupVersion.WithKind(DatabaseServerless_Kind)
)

func init() {
	SchemeBuilder.Register(&DatabaseServerless{}, &DatabaseServerlessList{})
}
