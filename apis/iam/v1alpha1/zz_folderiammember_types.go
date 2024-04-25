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

type FolderIAMMemberInitParameters struct {

	// ID of the folder to attach a policy to.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// The identity that will be granted the privilege that is specified in the role field.
	// This field can have one of the following values:
	// +crossplane:generate:reference:type=ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/yandex-cloud/provider-jet-yc/config/iam.ServiceAccountRefValue()
	// +crossplane:generate:reference:refFieldName=ServiceAccountRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountSelector
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// The role that should be assigned.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a ServiceAccount to populate member.
	// +kubebuilder:validation:Optional
	ServiceAccountRef *v1.Reference `json:"serviceAccountRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount to populate member.
	// +kubebuilder:validation:Optional
	ServiceAccountSelector *v1.Selector `json:"serviceAccountSelector,omitempty" tf:"-"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type FolderIAMMemberObservation struct {

	// ID of the folder to attach a policy to.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The identity that will be granted the privilege that is specified in the role field.
	// This field can have one of the following values:
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// The role that should be assigned.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type FolderIAMMemberParameters struct {

	// ID of the folder to attach a policy to.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// The identity that will be granted the privilege that is specified in the role field.
	// This field can have one of the following values:
	// +crossplane:generate:reference:type=ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/yandex-cloud/provider-jet-yc/config/iam.ServiceAccountRefValue()
	// +crossplane:generate:reference:refFieldName=ServiceAccountRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountSelector
	// +kubebuilder:validation:Optional
	Member *string `json:"member,omitempty" tf:"member,omitempty"`

	// The role that should be assigned.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a ServiceAccount to populate member.
	// +kubebuilder:validation:Optional
	ServiceAccountRef *v1.Reference `json:"serviceAccountRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount to populate member.
	// +kubebuilder:validation:Optional
	ServiceAccountSelector *v1.Selector `json:"serviceAccountSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// FolderIAMMemberSpec defines the desired state of FolderIAMMember
type FolderIAMMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderIAMMemberParameters `json:"forProvider"`
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
	InitProvider FolderIAMMemberInitParameters `json:"initProvider,omitempty"`
}

// FolderIAMMemberStatus defines the observed state of FolderIAMMember.
type FolderIAMMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderIAMMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FolderIAMMember is the Schema for the FolderIAMMembers API. Allows management of a single member for a single IAM binding for a Yandex Resource Manager folder.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type FolderIAMMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   FolderIAMMemberSpec   `json:"spec"`
	Status FolderIAMMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderIAMMemberList contains a list of FolderIAMMembers
type FolderIAMMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FolderIAMMember `json:"items"`
}

// Repository type metadata.
var (
	FolderIAMMember_Kind             = "FolderIAMMember"
	FolderIAMMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FolderIAMMember_Kind}.String()
	FolderIAMMember_KindAPIVersion   = FolderIAMMember_Kind + "." + CRDGroupVersion.String()
	FolderIAMMember_GroupVersionKind = CRDGroupVersion.WithKind(FolderIAMMember_Kind)
)

func init() {
	SchemeBuilder.Register(&FolderIAMMember{}, &FolderIAMMemberList{})
}
