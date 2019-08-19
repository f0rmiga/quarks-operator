// +build !ignore_autogenerated

/*

Don't alter this file, it was generated.

*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHDeployment) DeepCopyInto(out *BOSHDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHDeployment.
func (in *BOSHDeployment) DeepCopy() *BOSHDeployment {
	if in == nil {
		return nil
	}
	out := new(BOSHDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BOSHDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHDeploymentList) DeepCopyInto(out *BOSHDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BOSHDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHDeploymentList.
func (in *BOSHDeploymentList) DeepCopy() *BOSHDeploymentList {
	if in == nil {
		return nil
	}
	out := new(BOSHDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BOSHDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHDeploymentSpec) DeepCopyInto(out *BOSHDeploymentSpec) {
	*out = *in
	out.Manifest = in.Manifest
	if in.Ops != nil {
		in, out := &in.Ops, &out.Ops
		*out = make([]ResourceReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHDeploymentSpec.
func (in *BOSHDeploymentSpec) DeepCopy() *BOSHDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(BOSHDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BOSHDeploymentStatus) DeepCopyInto(out *BOSHDeploymentStatus) {
	*out = *in
	if in.LastReconcile != nil {
		in, out := &in.LastReconcile, &out.LastReconcile
		*out = (*in).DeepCopy()
	}
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BOSHDeploymentStatus.
func (in *BOSHDeploymentStatus) DeepCopy() *BOSHDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(BOSHDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceReference) DeepCopyInto(out *ResourceReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceReference.
func (in *ResourceReference) DeepCopy() *ResourceReference {
	if in == nil {
		return nil
	}
	out := new(ResourceReference)
	in.DeepCopyInto(out)
	return out
}
