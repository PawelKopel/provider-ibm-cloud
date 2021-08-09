// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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
func (in *CloudantDatabase) DeepCopyInto(out *CloudantDatabase) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudantDatabase.
func (in *CloudantDatabase) DeepCopy() *CloudantDatabase {
	if in == nil {
		return nil
	}
	out := new(CloudantDatabase)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudantDatabase) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudantDatabaseList) DeepCopyInto(out *CloudantDatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudantDatabase, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudantDatabaseList.
func (in *CloudantDatabaseList) DeepCopy() *CloudantDatabaseList {
	if in == nil {
		return nil
	}
	out := new(CloudantDatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudantDatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudantDatabaseObservation) DeepCopyInto(out *CloudantDatabaseObservation) {
	*out = *in
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(DatabaseInformationCluster)
		**out = **in
	}
	if in.Sizes != nil {
		in, out := &in.Sizes, &out.Sizes
		*out = new(ContentInformationSizes)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudantDatabaseObservation.
func (in *CloudantDatabaseObservation) DeepCopy() *CloudantDatabaseObservation {
	if in == nil {
		return nil
	}
	out := new(CloudantDatabaseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudantDatabaseParameters) DeepCopyInto(out *CloudantDatabaseParameters) {
	*out = *in
	if in.Partitioned != nil {
		in, out := &in.Partitioned, &out.Partitioned
		*out = new(bool)
		**out = **in
	}
	if in.Q != nil {
		in, out := &in.Q, &out.Q
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudantDatabaseParameters.
func (in *CloudantDatabaseParameters) DeepCopy() *CloudantDatabaseParameters {
	if in == nil {
		return nil
	}
	out := new(CloudantDatabaseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudantDatabaseSpec) DeepCopyInto(out *CloudantDatabaseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudantDatabaseSpec.
func (in *CloudantDatabaseSpec) DeepCopy() *CloudantDatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(CloudantDatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudantDatabaseStatus) DeepCopyInto(out *CloudantDatabaseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudantDatabaseStatus.
func (in *CloudantDatabaseStatus) DeepCopy() *CloudantDatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(CloudantDatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentInformationSizes) DeepCopyInto(out *ContentInformationSizes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentInformationSizes.
func (in *ContentInformationSizes) DeepCopy() *ContentInformationSizes {
	if in == nil {
		return nil
	}
	out := new(ContentInformationSizes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseInformationCluster) DeepCopyInto(out *DatabaseInformationCluster) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseInformationCluster.
func (in *DatabaseInformationCluster) DeepCopy() *DatabaseInformationCluster {
	if in == nil {
		return nil
	}
	out := new(DatabaseInformationCluster)
	in.DeepCopyInto(out)
	return out
}
