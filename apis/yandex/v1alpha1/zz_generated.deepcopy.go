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
func (in *ContentObservation) DeepCopyInto(out *ContentObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentObservation.
func (in *ContentObservation) DeepCopy() *ContentObservation {
	if in == nil {
		return nil
	}
	out := new(ContentObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentParameters) DeepCopyInto(out *ContentParameters) {
	*out = *in
	if in.ZipFilename != nil {
		in, out := &in.ZipFilename, &out.ZipFilename
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentParameters.
func (in *ContentParameters) DeepCopy() *ContentParameters {
	if in == nil {
		return nil
	}
	out := new(ContentParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Function) DeepCopyInto(out *Function) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Function.
func (in *Function) DeepCopy() *Function {
	if in == nil {
		return nil
	}
	out := new(Function)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Function) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionList) DeepCopyInto(out *FunctionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Function, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionList.
func (in *FunctionList) DeepCopy() *FunctionList {
	if in == nil {
		return nil
	}
	out := new(FunctionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FunctionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionObservation) DeepCopyInto(out *FunctionObservation) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = new(string)
		**out = **in
	}
	if in.ImageSize != nil {
		in, out := &in.ImageSize, &out.ImageSize
		*out = new(int64)
		**out = **in
	}
	if in.LoggroupID != nil {
		in, out := &in.LoggroupID, &out.LoggroupID
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
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
	if in.Content != nil {
		in, out := &in.Content, &out.Content
		*out = make([]ContentParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Entrypoint != nil {
		in, out := &in.Entrypoint, &out.Entrypoint
		*out = new(string)
		**out = **in
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
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
	if in.ExecutionTimeout != nil {
		in, out := &in.ExecutionTimeout, &out.ExecutionTimeout
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
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(int64)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Package != nil {
		in, out := &in.Package, &out.Package
		*out = make([]PackageParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(string)
		**out = **in
	}
	if in.ServiceAccountID != nil {
		in, out := &in.ServiceAccountID, &out.ServiceAccountID
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UserHash != nil {
		in, out := &in.UserHash, &out.UserHash
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
func (in *FunctionSpec) DeepCopyInto(out *FunctionSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionSpec.
func (in *FunctionSpec) DeepCopy() *FunctionSpec {
	if in == nil {
		return nil
	}
	out := new(FunctionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FunctionStatus) DeepCopyInto(out *FunctionStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FunctionStatus.
func (in *FunctionStatus) DeepCopy() *FunctionStatus {
	if in == nil {
		return nil
	}
	out := new(FunctionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageObservation) DeepCopyInto(out *PackageObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageObservation.
func (in *PackageObservation) DeepCopy() *PackageObservation {
	if in == nil {
		return nil
	}
	out := new(PackageObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageParameters) DeepCopyInto(out *PackageParameters) {
	*out = *in
	if in.BucketName != nil {
		in, out := &in.BucketName, &out.BucketName
		*out = new(string)
		**out = **in
	}
	if in.ObjectName != nil {
		in, out := &in.ObjectName, &out.ObjectName
		*out = new(string)
		**out = **in
	}
	if in.Sha256 != nil {
		in, out := &in.Sha256, &out.Sha256
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageParameters.
func (in *PackageParameters) DeepCopy() *PackageParameters {
	if in == nil {
		return nil
	}
	out := new(PackageParameters)
	in.DeepCopyInto(out)
	return out
}
