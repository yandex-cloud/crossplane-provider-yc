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

type CloudIAMBindingInitParameters struct {

	// ID of the cloud to attach the policy to.
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +crossplane:generate:reference:type=ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/yandex-cloud/crossplane-provider-yc/config/iam.ServiceAccountRefValue()
	// +crossplane:generate:reference:refFieldName=ServiceAccountRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountSelector
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_resourcemanager_cloud_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// References to ServiceAccount to populate members.
	// +kubebuilder:validation:Optional
	ServiceAccountRef []v1.Reference `json:"serviceAccountRef,omitempty" tf:"-"`

	// Selector for a list of ServiceAccount to populate members.
	// +kubebuilder:validation:Optional
	ServiceAccountSelector *v1.Selector `json:"serviceAccountSelector,omitempty" tf:"-"`

	SleepAfter *int64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type CloudIAMBindingObservation struct {

	// ID of the cloud to attach the policy to.
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_resourcemanager_cloud_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *int64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type CloudIAMBindingParameters struct {

	// ID of the cloud to attach the policy to.
	// +kubebuilder:validation:Optional
	CloudID *string `json:"cloudId,omitempty" tf:"cloud_id,omitempty"`

	// An array of identities that will be granted the privilege in the role.
	// Each entry can have one of the following values:
	// +crossplane:generate:reference:type=ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/yandex-cloud/crossplane-provider-yc/config/iam.ServiceAccountRefValue()
	// +crossplane:generate:reference:refFieldName=ServiceAccountRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_resourcemanager_cloud_iam_binding can be used per role.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// References to ServiceAccount to populate members.
	// +kubebuilder:validation:Optional
	ServiceAccountRef []v1.Reference `json:"serviceAccountRef,omitempty" tf:"-"`

	// Selector for a list of ServiceAccount to populate members.
	// +kubebuilder:validation:Optional
	ServiceAccountSelector *v1.Selector `json:"serviceAccountSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SleepAfter *int64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// CloudIAMBindingSpec defines the desired state of CloudIAMBinding
type CloudIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CloudIAMBindingParameters `json:"forProvider"`
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
	InitProvider CloudIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// CloudIAMBindingStatus defines the observed state of CloudIAMBinding.
type CloudIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CloudIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CloudIAMBinding is the Schema for the CloudIAMBindings API. Allows management of a single IAM binding for a Yandex Resource Manager cloud.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type CloudIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.cloudId) || (has(self.initProvider) && has(self.initProvider.cloudId))",message="spec.forProvider.cloudId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   CloudIAMBindingSpec   `json:"spec"`
	Status CloudIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CloudIAMBindingList contains a list of CloudIAMBindings
type CloudIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CloudIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	CloudIAMBinding_Kind             = "CloudIAMBinding"
	CloudIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CloudIAMBinding_Kind}.String()
	CloudIAMBinding_KindAPIVersion   = CloudIAMBinding_Kind + "." + CRDGroupVersion.String()
	CloudIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(CloudIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&CloudIAMBinding{}, &CloudIAMBindingList{})
}
