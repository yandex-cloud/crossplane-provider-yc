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

type QueueObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`
}

type QueueParameters struct {

	// +kubebuilder:validation:Optional
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// +kubebuilder:validation:Optional
	ContentBasedDeduplication *bool `json:"contentBasedDeduplication,omitempty" tf:"content_based_deduplication,omitempty"`

	// +kubebuilder:validation:Optional
	DelaySeconds *int64 `json:"delaySeconds,omitempty" tf:"delay_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	FifoQueue *bool `json:"fifoQueue,omitempty" tf:"fifo_queue,omitempty"`

	// +kubebuilder:validation:Optional
	MaxMessageSize *int64 `json:"maxMessageSize,omitempty" tf:"max_message_size,omitempty"`

	// +kubebuilder:validation:Optional
	MessageRetentionSeconds *int64 `json:"messageRetentionSeconds,omitempty" tf:"message_retention_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	NamePrefix *string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`

	// +kubebuilder:validation:Optional
	ReceiveWaitTimeSeconds *int64 `json:"receiveWaitTimeSeconds,omitempty" tf:"receive_wait_time_seconds,omitempty"`

	// +kubebuilder:validation:Optional
	RedrivePolicy *string `json:"redrivePolicy,omitempty" tf:"redrive_policy,omitempty"`

	// +kubebuilder:validation:Optional
	SecretKeySecretRef *v1.SecretKeySelector `json:"secretKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VisibilityTimeoutSeconds *int64 `json:"visibilityTimeoutSeconds,omitempty" tf:"visibility_timeout_seconds,omitempty"`
}

// QueueSpec defines the desired state of Queue
type QueueSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     QueueParameters `json:"forProvider"`
}

// QueueStatus defines the observed state of Queue.
type QueueStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        QueueObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Queue is the Schema for the Queues API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloudjet}
type Queue struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              QueueSpec   `json:"spec"`
	Status            QueueStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// QueueList contains a list of Queues
type QueueList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Queue `json:"items"`
}

// Repository type metadata.
var (
	Queue_Kind             = "Queue"
	Queue_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Queue_Kind}.String()
	Queue_KindAPIVersion   = Queue_Kind + "." + CRDGroupVersion.String()
	Queue_GroupVersionKind = CRDGroupVersion.WithKind(Queue_Kind)
)

func init() {
	SchemeBuilder.Register(&Queue{}, &QueueList{})
}
