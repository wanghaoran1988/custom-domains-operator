// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomain) DeepCopyInto(out *CustomDomain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomain.
func (in *CustomDomain) DeepCopy() *CustomDomain {
	if in == nil {
		return nil
	}
	out := new(CustomDomain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomDomain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainList) DeepCopyInto(out *CustomDomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomDomain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainList.
func (in *CustomDomainList) DeepCopy() *CustomDomainList {
	if in == nil {
		return nil
	}
	out := new(CustomDomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomDomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainSpec) DeepCopyInto(out *CustomDomainSpec) {
	*out = *in
	out.TLSSecret = in.TLSSecret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainSpec.
func (in *CustomDomainSpec) DeepCopy() *CustomDomainSpec {
	if in == nil {
		return nil
	}
	out := new(CustomDomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomDomainStatus) DeepCopyInto(out *CustomDomainStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomDomainStatus.
func (in *CustomDomainStatus) DeepCopy() *CustomDomainStatus {
	if in == nil {
		return nil
	}
	out := new(CustomDomainStatus)
	in.DeepCopyInto(out)
	return out
}
