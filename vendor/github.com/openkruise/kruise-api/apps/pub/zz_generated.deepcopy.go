//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Kruise Authors.

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

package pub

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InPlaceUpdateContainerBatch) DeepCopyInto(out *InPlaceUpdateContainerBatch) {
	*out = *in
	in.Timestamp.DeepCopyInto(&out.Timestamp)
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InPlaceUpdateContainerBatch.
func (in *InPlaceUpdateContainerBatch) DeepCopy() *InPlaceUpdateContainerBatch {
	if in == nil {
		return nil
	}
	out := new(InPlaceUpdateContainerBatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InPlaceUpdateContainerStatus) DeepCopyInto(out *InPlaceUpdateContainerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InPlaceUpdateContainerStatus.
func (in *InPlaceUpdateContainerStatus) DeepCopy() *InPlaceUpdateContainerStatus {
	if in == nil {
		return nil
	}
	out := new(InPlaceUpdateContainerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InPlaceUpdatePreCheckBeforeNext) DeepCopyInto(out *InPlaceUpdatePreCheckBeforeNext) {
	*out = *in
	if in.ContainersRequiredReady != nil {
		in, out := &in.ContainersRequiredReady, &out.ContainersRequiredReady
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InPlaceUpdatePreCheckBeforeNext.
func (in *InPlaceUpdatePreCheckBeforeNext) DeepCopy() *InPlaceUpdatePreCheckBeforeNext {
	if in == nil {
		return nil
	}
	out := new(InPlaceUpdatePreCheckBeforeNext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InPlaceUpdateState) DeepCopyInto(out *InPlaceUpdateState) {
	*out = *in
	in.UpdateTimestamp.DeepCopyInto(&out.UpdateTimestamp)
	if in.LastContainerStatuses != nil {
		in, out := &in.LastContainerStatuses, &out.LastContainerStatuses
		*out = make(map[string]InPlaceUpdateContainerStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NextContainerImages != nil {
		in, out := &in.NextContainerImages, &out.NextContainerImages
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NextContainerRefMetadata != nil {
		in, out := &in.NextContainerRefMetadata, &out.NextContainerRefMetadata
		*out = make(map[string]v1.ObjectMeta, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.PreCheckBeforeNext != nil {
		in, out := &in.PreCheckBeforeNext, &out.PreCheckBeforeNext
		*out = new(InPlaceUpdatePreCheckBeforeNext)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerBatchesRecord != nil {
		in, out := &in.ContainerBatchesRecord, &out.ContainerBatchesRecord
		*out = make([]InPlaceUpdateContainerBatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InPlaceUpdateState.
func (in *InPlaceUpdateState) DeepCopy() *InPlaceUpdateState {
	if in == nil {
		return nil
	}
	out := new(InPlaceUpdateState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InPlaceUpdateStrategy) DeepCopyInto(out *InPlaceUpdateStrategy) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InPlaceUpdateStrategy.
func (in *InPlaceUpdateStrategy) DeepCopy() *InPlaceUpdateStrategy {
	if in == nil {
		return nil
	}
	out := new(InPlaceUpdateStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Lifecycle) DeepCopyInto(out *Lifecycle) {
	*out = *in
	if in.PreDelete != nil {
		in, out := &in.PreDelete, &out.PreDelete
		*out = new(LifecycleHook)
		(*in).DeepCopyInto(*out)
	}
	if in.InPlaceUpdate != nil {
		in, out := &in.InPlaceUpdate, &out.InPlaceUpdate
		*out = new(LifecycleHook)
		(*in).DeepCopyInto(*out)
	}
	if in.PreNormal != nil {
		in, out := &in.PreNormal, &out.PreNormal
		*out = new(LifecycleHook)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Lifecycle.
func (in *Lifecycle) DeepCopy() *Lifecycle {
	if in == nil {
		return nil
	}
	out := new(Lifecycle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LifecycleHook) DeepCopyInto(out *LifecycleHook) {
	*out = *in
	if in.LabelsHandler != nil {
		in, out := &in.LabelsHandler, &out.LabelsHandler
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.FinalizersHandler != nil {
		in, out := &in.FinalizersHandler, &out.FinalizersHandler
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LifecycleHook.
func (in *LifecycleHook) DeepCopy() *LifecycleHook {
	if in == nil {
		return nil
	}
	out := new(LifecycleHook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeContainerHashes) DeepCopyInto(out *RuntimeContainerHashes) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeContainerHashes.
func (in *RuntimeContainerHashes) DeepCopy() *RuntimeContainerHashes {
	if in == nil {
		return nil
	}
	out := new(RuntimeContainerHashes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeContainerMeta) DeepCopyInto(out *RuntimeContainerMeta) {
	*out = *in
	out.Hashes = in.Hashes
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeContainerMeta.
func (in *RuntimeContainerMeta) DeepCopy() *RuntimeContainerMeta {
	if in == nil {
		return nil
	}
	out := new(RuntimeContainerMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RuntimeContainerMetaSet) DeepCopyInto(out *RuntimeContainerMetaSet) {
	*out = *in
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]RuntimeContainerMeta, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RuntimeContainerMetaSet.
func (in *RuntimeContainerMetaSet) DeepCopy() *RuntimeContainerMetaSet {
	if in == nil {
		return nil
	}
	out := new(RuntimeContainerMetaSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdatePriorityOrderTerm) DeepCopyInto(out *UpdatePriorityOrderTerm) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdatePriorityOrderTerm.
func (in *UpdatePriorityOrderTerm) DeepCopy() *UpdatePriorityOrderTerm {
	if in == nil {
		return nil
	}
	out := new(UpdatePriorityOrderTerm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdatePriorityStrategy) DeepCopyInto(out *UpdatePriorityStrategy) {
	*out = *in
	if in.OrderPriority != nil {
		in, out := &in.OrderPriority, &out.OrderPriority
		*out = make([]UpdatePriorityOrderTerm, len(*in))
		copy(*out, *in)
	}
	if in.WeightPriority != nil {
		in, out := &in.WeightPriority, &out.WeightPriority
		*out = make([]UpdatePriorityWeightTerm, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdatePriorityStrategy.
func (in *UpdatePriorityStrategy) DeepCopy() *UpdatePriorityStrategy {
	if in == nil {
		return nil
	}
	out := new(UpdatePriorityStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpdatePriorityWeightTerm) DeepCopyInto(out *UpdatePriorityWeightTerm) {
	*out = *in
	in.MatchSelector.DeepCopyInto(&out.MatchSelector)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpdatePriorityWeightTerm.
func (in *UpdatePriorityWeightTerm) DeepCopy() *UpdatePriorityWeightTerm {
	if in == nil {
		return nil
	}
	out := new(UpdatePriorityWeightTerm)
	in.DeepCopyInto(out)
	return out
}