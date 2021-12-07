// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DlqObservation) DeepCopyInto(out *DlqObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DlqObservation.
func (in *DlqObservation) DeepCopy() *DlqObservation {
	if in == nil {
		return nil
	}
	out := new(DlqObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DlqParameters) DeepCopyInto(out *DlqParameters) {
	*out = *in
	if in.QueueID != nil {
		in, out := &in.QueueID, &out.QueueID
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DlqParameters.
func (in *DlqParameters) DeepCopy() *DlqParameters {
	if in == nil {
		return nil
	}
	out := new(DlqParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionObservation) DeepCopyInto(out *FunctionObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionObservation.
func (in *FunctionObservation) DeepCopy() *FunctionObservation {
	if in == nil {
		return nil
	}
	out := new(FunctionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionParameters) DeepCopyInto(out *FunctionParameters) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.RetryAttempts != nil {
		in, out := &in.RetryAttempts, &out.RetryAttempts
		*out = new(string)
		**out = **in
	}
	if in.RetryInterval != nil {
		in, out := &in.RetryInterval, &out.RetryInterval
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionParameters.
func (in *FunctionParameters) DeepCopy() *FunctionParameters {
	if in == nil {
		return nil
	}
	out := new(FunctionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBinding) DeepCopyInto(out *IamBinding) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBinding.
func (in *IamBinding) DeepCopy() *IamBinding {
	if in == nil {
		return nil
	}
	out := new(IamBinding)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamBinding) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingList) DeepCopyInto(out *IamBindingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IamBinding, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingList.
func (in *IamBindingList) DeepCopy() *IamBindingList {
	if in == nil {
		return nil
	}
	out := new(IamBindingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IamBindingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingObservation) DeepCopyInto(out *IamBindingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingObservation.
func (in *IamBindingObservation) DeepCopy() *IamBindingObservation {
	if in == nil {
		return nil
	}
	out := new(IamBindingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingParameters) DeepCopyInto(out *IamBindingParameters) {
	*out = *in
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.Members != nil {
		in, out := &in.Members, &out.Members
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Role != nil {
		in, out := &in.Role, &out.Role
		*out = new(string)
		**out = **in
	}
	if in.SleepAfter != nil {
		in, out := &in.SleepAfter, &out.SleepAfter
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingParameters.
func (in *IamBindingParameters) DeepCopy() *IamBindingParameters {
	if in == nil {
		return nil
	}
	out := new(IamBindingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingSpec) DeepCopyInto(out *IamBindingSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingSpec.
func (in *IamBindingSpec) DeepCopy() *IamBindingSpec {
	if in == nil {
		return nil
	}
	out := new(IamBindingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamBindingStatus) DeepCopyInto(out *IamBindingStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamBindingStatus.
func (in *IamBindingStatus) DeepCopy() *IamBindingStatus {
	if in == nil {
		return nil
	}
	out := new(IamBindingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IotObservation) DeepCopyInto(out *IotObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IotObservation.
func (in *IotObservation) DeepCopy() *IotObservation {
	if in == nil {
		return nil
	}
	out := new(IotObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IotParameters) DeepCopyInto(out *IotParameters) {
	*out = *in
	if in.DeviceID != nil {
		in, out := &in.DeviceID, &out.DeviceID
		*out = new(string)
		**out = **in
	}
	if in.RegistryID != nil {
		in, out := &in.RegistryID, &out.RegistryID
		*out = new(string)
		**out = **in
	}
	if in.Topic != nil {
		in, out := &in.Topic, &out.Topic
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IotParameters.
func (in *IotParameters) DeepCopy() *IotParameters {
	if in == nil {
		return nil
	}
	out := new(IotParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogGroupObservation) DeepCopyInto(out *LogGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogGroupObservation.
func (in *LogGroupObservation) DeepCopy() *LogGroupObservation {
	if in == nil {
		return nil
	}
	out := new(LogGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogGroupParameters) DeepCopyInto(out *LogGroupParameters) {
	*out = *in
	if in.BatchCutoff != nil {
		in, out := &in.BatchCutoff, &out.BatchCutoff
		*out = new(string)
		**out = **in
	}
	if in.BatchSize != nil {
		in, out := &in.BatchSize, &out.BatchSize
		*out = new(string)
		**out = **in
	}
	if in.LogGroupIds != nil {
		in, out := &in.LogGroupIds, &out.LogGroupIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogGroupParameters.
func (in *LogGroupParameters) DeepCopy() *LogGroupParameters {
	if in == nil {
		return nil
	}
	out := new(LogGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingObservation) DeepCopyInto(out *LoggingObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingObservation.
func (in *LoggingObservation) DeepCopy() *LoggingObservation {
	if in == nil {
		return nil
	}
	out := new(LoggingObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingParameters) DeepCopyInto(out *LoggingParameters) {
	*out = *in
	if in.BatchCutoff != nil {
		in, out := &in.BatchCutoff, &out.BatchCutoff
		*out = new(string)
		**out = **in
	}
	if in.BatchSize != nil {
		in, out := &in.BatchSize, &out.BatchSize
		*out = new(string)
		**out = **in
	}
	if in.GroupID != nil {
		in, out := &in.GroupID, &out.GroupID
		*out = new(string)
		**out = **in
	}
	if in.Levels != nil {
		in, out := &in.Levels, &out.Levels
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ResourceIds != nil {
		in, out := &in.ResourceIds, &out.ResourceIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ResourceTypes != nil {
		in, out := &in.ResourceTypes, &out.ResourceTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingParameters.
func (in *LoggingParameters) DeepCopy() *LoggingParameters {
	if in == nil {
		return nil
	}
	out := new(LoggingParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessageQueueObservation) DeepCopyInto(out *MessageQueueObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessageQueueObservation.
func (in *MessageQueueObservation) DeepCopy() *MessageQueueObservation {
	if in == nil {
		return nil
	}
	out := new(MessageQueueObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MessageQueueParameters) DeepCopyInto(out *MessageQueueParameters) {
	*out = *in
	if in.BatchCutoff != nil {
		in, out := &in.BatchCutoff, &out.BatchCutoff
		*out = new(string)
		**out = **in
	}
	if in.BatchSize != nil {
		in, out := &in.BatchSize, &out.BatchSize
		*out = new(string)
		**out = **in
	}
	if in.QueueID != nil {
		in, out := &in.QueueID, &out.QueueID
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
	if in.VisibilityTimeout != nil {
		in, out := &in.VisibilityTimeout, &out.VisibilityTimeout
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MessageQueueParameters.
func (in *MessageQueueParameters) DeepCopy() *MessageQueueParameters {
	if in == nil {
		return nil
	}
	out := new(MessageQueueParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageObservation) DeepCopyInto(out *ObjectStorageObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageObservation.
func (in *ObjectStorageObservation) DeepCopy() *ObjectStorageObservation {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStorageParameters) DeepCopyInto(out *ObjectStorageParameters) {
	*out = *in
	if in.BucketID != nil {
		in, out := &in.BucketID, &out.BucketID
		*out = new(string)
		**out = **in
	}
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = new(bool)
		**out = **in
	}
	if in.Delete != nil {
		in, out := &in.Delete, &out.Delete
		*out = new(bool)
		**out = **in
	}
	if in.Prefix != nil {
		in, out := &in.Prefix, &out.Prefix
		*out = new(string)
		**out = **in
	}
	if in.Suffix != nil {
		in, out := &in.Suffix, &out.Suffix
		*out = new(string)
		**out = **in
	}
	if in.Update != nil {
		in, out := &in.Update, &out.Update
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStorageParameters.
func (in *ObjectStorageParameters) DeepCopy() *ObjectStorageParameters {
	if in == nil {
		return nil
	}
	out := new(ObjectStorageParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyObservation) DeepCopyInto(out *PolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyObservation.
func (in *PolicyObservation) DeepCopy() *PolicyObservation {
	if in == nil {
		return nil
	}
	out := new(PolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PolicyParameters) DeepCopyInto(out *PolicyParameters) {
	*out = *in
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
	if in.ZoneInstancesLimit != nil {
		in, out := &in.ZoneInstancesLimit, &out.ZoneInstancesLimit
		*out = new(int64)
		**out = **in
	}
	if in.ZoneRequestsLimit != nil {
		in, out := &in.ZoneRequestsLimit, &out.ZoneRequestsLimit
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PolicyParameters.
func (in *PolicyParameters) DeepCopy() *PolicyParameters {
	if in == nil {
		return nil
	}
	out := new(PolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicy) DeepCopyInto(out *ScalingPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicy.
func (in *ScalingPolicy) DeepCopy() *ScalingPolicy {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyList) DeepCopyInto(out *ScalingPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScalingPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyList.
func (in *ScalingPolicyList) DeepCopy() *ScalingPolicyList {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyObservation) DeepCopyInto(out *ScalingPolicyObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyObservation.
func (in *ScalingPolicyObservation) DeepCopy() *ScalingPolicyObservation {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyParameters) DeepCopyInto(out *ScalingPolicyParameters) {
	*out = *in
	if in.FunctionID != nil {
		in, out := &in.FunctionID, &out.FunctionID
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = make([]PolicyParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyParameters.
func (in *ScalingPolicyParameters) DeepCopy() *ScalingPolicyParameters {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicySpec) DeepCopyInto(out *ScalingPolicySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicySpec.
func (in *ScalingPolicySpec) DeepCopy() *ScalingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyStatus) DeepCopyInto(out *ScalingPolicyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyStatus.
func (in *ScalingPolicyStatus) DeepCopy() *ScalingPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimerObservation) DeepCopyInto(out *TimerObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimerObservation.
func (in *TimerObservation) DeepCopy() *TimerObservation {
	if in == nil {
		return nil
	}
	out := new(TimerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimerParameters) DeepCopyInto(out *TimerParameters) {
	*out = *in
	if in.CronExpression != nil {
		in, out := &in.CronExpression, &out.CronExpression
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimerParameters.
func (in *TimerParameters) DeepCopy() *TimerParameters {
	if in == nil {
		return nil
	}
	out := new(TimerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerList) DeepCopyInto(out *TriggerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Trigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerList.
func (in *TriggerList) DeepCopy() *TriggerList {
	if in == nil {
		return nil
	}
	out := new(TriggerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TriggerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerObservation) DeepCopyInto(out *TriggerObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerObservation.
func (in *TriggerObservation) DeepCopy() *TriggerObservation {
	if in == nil {
		return nil
	}
	out := new(TriggerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerParameters) DeepCopyInto(out *TriggerParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Dlq != nil {
		in, out := &in.Dlq, &out.Dlq
		*out = make([]DlqParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
	}
	if in.Function != nil {
		in, out := &in.Function, &out.Function
		*out = make([]FunctionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Iot != nil {
		in, out := &in.Iot, &out.Iot
		*out = make([]IotParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.LogGroup != nil {
		in, out := &in.LogGroup, &out.LogGroup
		*out = make([]LogGroupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = make([]LoggingParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MessageQueue != nil {
		in, out := &in.MessageQueue, &out.MessageQueue
		*out = make([]MessageQueueParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.ObjectStorage != nil {
		in, out := &in.ObjectStorage, &out.ObjectStorage
		*out = make([]ObjectStorageParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Timer != nil {
		in, out := &in.Timer, &out.Timer
		*out = make([]TimerParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerParameters.
func (in *TriggerParameters) DeepCopy() *TriggerParameters {
	if in == nil {
		return nil
	}
	out := new(TriggerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpec) DeepCopyInto(out *TriggerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpec.
func (in *TriggerSpec) DeepCopy() *TriggerSpec {
	if in == nil {
		return nil
	}
	out := new(TriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerStatus) DeepCopyInto(out *TriggerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerStatus.
func (in *TriggerStatus) DeepCopy() *TriggerStatus {
	if in == nil {
		return nil
	}
	out := new(TriggerStatus)
	in.DeepCopyInto(out)
	return out
}
