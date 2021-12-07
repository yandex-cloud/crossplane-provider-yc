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
func (in *AttachedTargetGroupObservation) DeepCopyInto(out *AttachedTargetGroupObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachedTargetGroupObservation.
func (in *AttachedTargetGroupObservation) DeepCopy() *AttachedTargetGroupObservation {
	if in == nil {
		return nil
	}
	out := new(AttachedTargetGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AttachedTargetGroupParameters) DeepCopyInto(out *AttachedTargetGroupParameters) {
	*out = *in
	if in.Healthcheck != nil {
		in, out := &in.Healthcheck, &out.Healthcheck
		*out = make([]HealthcheckParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TargetGroupID != nil {
		in, out := &in.TargetGroupID, &out.TargetGroupID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AttachedTargetGroupParameters.
func (in *AttachedTargetGroupParameters) DeepCopy() *AttachedTargetGroupParameters {
	if in == nil {
		return nil
	}
	out := new(AttachedTargetGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAddressSpecObservation) DeepCopyInto(out *ExternalAddressSpecObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAddressSpecObservation.
func (in *ExternalAddressSpecObservation) DeepCopy() *ExternalAddressSpecObservation {
	if in == nil {
		return nil
	}
	out := new(ExternalAddressSpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalAddressSpecParameters) DeepCopyInto(out *ExternalAddressSpecParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalAddressSpecParameters.
func (in *ExternalAddressSpecParameters) DeepCopy() *ExternalAddressSpecParameters {
	if in == nil {
		return nil
	}
	out := new(ExternalAddressSpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPOptionsObservation) DeepCopyInto(out *HTTPOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPOptionsObservation.
func (in *HTTPOptionsObservation) DeepCopy() *HTTPOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(HTTPOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPOptionsParameters) DeepCopyInto(out *HTTPOptionsParameters) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPOptionsParameters.
func (in *HTTPOptionsParameters) DeepCopy() *HTTPOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(HTTPOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcheckObservation) DeepCopyInto(out *HealthcheckObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcheckObservation.
func (in *HealthcheckObservation) DeepCopy() *HealthcheckObservation {
	if in == nil {
		return nil
	}
	out := new(HealthcheckObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HealthcheckParameters) DeepCopyInto(out *HealthcheckParameters) {
	*out = *in
	if in.HTTPOptions != nil {
		in, out := &in.HTTPOptions, &out.HTTPOptions
		*out = make([]HTTPOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HealthyThreshold != nil {
		in, out := &in.HealthyThreshold, &out.HealthyThreshold
		*out = new(int64)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.TCPOptions != nil {
		in, out := &in.TCPOptions, &out.TCPOptions
		*out = make([]TCPOptionsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		*out = new(int64)
		**out = **in
	}
	if in.UnhealthyThreshold != nil {
		in, out := &in.UnhealthyThreshold, &out.UnhealthyThreshold
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HealthcheckParameters.
func (in *HealthcheckParameters) DeepCopy() *HealthcheckParameters {
	if in == nil {
		return nil
	}
	out := new(HealthcheckParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalAddressSpecObservation) DeepCopyInto(out *InternalAddressSpecObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalAddressSpecObservation.
func (in *InternalAddressSpecObservation) DeepCopy() *InternalAddressSpecObservation {
	if in == nil {
		return nil
	}
	out := new(InternalAddressSpecObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InternalAddressSpecParameters) DeepCopyInto(out *InternalAddressSpecParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.IPVersion != nil {
		in, out := &in.IPVersion, &out.IPVersion
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InternalAddressSpecParameters.
func (in *InternalAddressSpecParameters) DeepCopy() *InternalAddressSpecParameters {
	if in == nil {
		return nil
	}
	out := new(InternalAddressSpecParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerObservation) DeepCopyInto(out *ListenerObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerObservation.
func (in *ListenerObservation) DeepCopy() *ListenerObservation {
	if in == nil {
		return nil
	}
	out := new(ListenerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ListenerParameters) DeepCopyInto(out *ListenerParameters) {
	*out = *in
	if in.ExternalAddressSpec != nil {
		in, out := &in.ExternalAddressSpec, &out.ExternalAddressSpec
		*out = make([]ExternalAddressSpecParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InternalAddressSpec != nil {
		in, out := &in.InternalAddressSpec, &out.InternalAddressSpec
		*out = make([]InternalAddressSpecParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(string)
		**out = **in
	}
	if in.TargetPort != nil {
		in, out := &in.TargetPort, &out.TargetPort
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ListenerParameters.
func (in *ListenerParameters) DeepCopy() *ListenerParameters {
	if in == nil {
		return nil
	}
	out := new(ListenerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkLoadBalancer) DeepCopyInto(out *NetworkLoadBalancer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkLoadBalancer.
func (in *NetworkLoadBalancer) DeepCopy() *NetworkLoadBalancer {
	if in == nil {
		return nil
	}
	out := new(NetworkLoadBalancer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkLoadBalancer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkLoadBalancerList) DeepCopyInto(out *NetworkLoadBalancerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NetworkLoadBalancer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkLoadBalancerList.
func (in *NetworkLoadBalancerList) DeepCopy() *NetworkLoadBalancerList {
	if in == nil {
		return nil
	}
	out := new(NetworkLoadBalancerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NetworkLoadBalancerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkLoadBalancerObservation) DeepCopyInto(out *NetworkLoadBalancerObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkLoadBalancerObservation.
func (in *NetworkLoadBalancerObservation) DeepCopy() *NetworkLoadBalancerObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkLoadBalancerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkLoadBalancerParameters) DeepCopyInto(out *NetworkLoadBalancerParameters) {
	*out = *in
	if in.AttachedTargetGroup != nil {
		in, out := &in.AttachedTargetGroup, &out.AttachedTargetGroup
		*out = make([]AttachedTargetGroupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
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
	if in.Listener != nil {
		in, out := &in.Listener, &out.Listener
		*out = make([]ListenerParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RegionID != nil {
		in, out := &in.RegionID, &out.RegionID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkLoadBalancerParameters.
func (in *NetworkLoadBalancerParameters) DeepCopy() *NetworkLoadBalancerParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkLoadBalancerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkLoadBalancerSpec) DeepCopyInto(out *NetworkLoadBalancerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkLoadBalancerSpec.
func (in *NetworkLoadBalancerSpec) DeepCopy() *NetworkLoadBalancerSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkLoadBalancerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkLoadBalancerStatus) DeepCopyInto(out *NetworkLoadBalancerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkLoadBalancerStatus.
func (in *NetworkLoadBalancerStatus) DeepCopy() *NetworkLoadBalancerStatus {
	if in == nil {
		return nil
	}
	out := new(NetworkLoadBalancerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPOptionsObservation) DeepCopyInto(out *TCPOptionsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPOptionsObservation.
func (in *TCPOptionsObservation) DeepCopy() *TCPOptionsObservation {
	if in == nil {
		return nil
	}
	out := new(TCPOptionsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPOptionsParameters) DeepCopyInto(out *TCPOptionsParameters) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPOptionsParameters.
func (in *TCPOptionsParameters) DeepCopy() *TCPOptionsParameters {
	if in == nil {
		return nil
	}
	out := new(TCPOptionsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroup) DeepCopyInto(out *TargetGroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroup.
func (in *TargetGroup) DeepCopy() *TargetGroup {
	if in == nil {
		return nil
	}
	out := new(TargetGroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TargetGroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroupList) DeepCopyInto(out *TargetGroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TargetGroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroupList.
func (in *TargetGroupList) DeepCopy() *TargetGroupList {
	if in == nil {
		return nil
	}
	out := new(TargetGroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TargetGroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroupObservation) DeepCopyInto(out *TargetGroupObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroupObservation.
func (in *TargetGroupObservation) DeepCopy() *TargetGroupObservation {
	if in == nil {
		return nil
	}
	out := new(TargetGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroupParameters) DeepCopyInto(out *TargetGroupParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FolderID != nil {
		in, out := &in.FolderID, &out.FolderID
		*out = new(string)
		**out = **in
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
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.RegionID != nil {
		in, out := &in.RegionID, &out.RegionID
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = make([]TargetParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroupParameters.
func (in *TargetGroupParameters) DeepCopy() *TargetGroupParameters {
	if in == nil {
		return nil
	}
	out := new(TargetGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroupSpec) DeepCopyInto(out *TargetGroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroupSpec.
func (in *TargetGroupSpec) DeepCopy() *TargetGroupSpec {
	if in == nil {
		return nil
	}
	out := new(TargetGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetGroupStatus) DeepCopyInto(out *TargetGroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetGroupStatus.
func (in *TargetGroupStatus) DeepCopy() *TargetGroupStatus {
	if in == nil {
		return nil
	}
	out := new(TargetGroupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetObservation) DeepCopyInto(out *TargetObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetObservation.
func (in *TargetObservation) DeepCopy() *TargetObservation {
	if in == nil {
		return nil
	}
	out := new(TargetObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetParameters) DeepCopyInto(out *TargetParameters) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		*out = new(string)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetParameters.
func (in *TargetParameters) DeepCopy() *TargetParameters {
	if in == nil {
		return nil
	}
	out := new(TargetParameters)
	in.DeepCopyInto(out)
	return out
}
