//go:build !ignore_autogenerated

/*
Copyright 2024.

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
func (in *AbsolutePowerCapInWattsSpec) DeepCopyInto(out *AbsolutePowerCapInWattsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbsolutePowerCapInWattsSpec.
func (in *AbsolutePowerCapInWattsSpec) DeepCopy() *AbsolutePowerCapInWattsSpec {
	if in == nil {
		return nil
	}
	out := new(AbsolutePowerCapInWattsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AbsoluteTemperatureThresholdInCelsiusSpec) DeepCopyInto(out *AbsoluteTemperatureThresholdInCelsiusSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AbsoluteTemperatureThresholdInCelsiusSpec.
func (in *AbsoluteTemperatureThresholdInCelsiusSpec) DeepCopy() *AbsoluteTemperatureThresholdInCelsiusSpec {
	if in == nil {
		return nil
	}
	out := new(AbsoluteTemperatureThresholdInCelsiusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PowerCappingConfig) DeepCopyInto(out *PowerCappingConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PowerCappingConfig.
func (in *PowerCappingConfig) DeepCopy() *PowerCappingConfig {
	if in == nil {
		return nil
	}
	out := new(PowerCappingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PowerCappingConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PowerCappingConfigList) DeepCopyInto(out *PowerCappingConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PowerCappingConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PowerCappingConfigList.
func (in *PowerCappingConfigList) DeepCopy() *PowerCappingConfigList {
	if in == nil {
		return nil
	}
	out := new(PowerCappingConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PowerCappingConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PowerCappingConfigSpec) DeepCopyInto(out *PowerCappingConfigSpec) {
	*out = *in
	out.PowerCappingSpec = in.PowerCappingSpec
	out.TemperatureThresholdSpec = in.TemperatureThresholdSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PowerCappingConfigSpec.
func (in *PowerCappingConfigSpec) DeepCopy() *PowerCappingConfigSpec {
	if in == nil {
		return nil
	}
	out := new(PowerCappingConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PowerCappingConfigStatus) DeepCopyInto(out *PowerCappingConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PowerCappingConfigStatus.
func (in *PowerCappingConfigStatus) DeepCopy() *PowerCappingConfigStatus {
	if in == nil {
		return nil
	}
	out := new(PowerCappingConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PowerCappingSpec) DeepCopyInto(out *PowerCappingSpec) {
	*out = *in
	out.AbsolutePowerCapInWattsSpec = in.AbsolutePowerCapInWattsSpec
	out.RelativePowerCapInPercentageSpec = in.RelativePowerCapInPercentageSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PowerCappingSpec.
func (in *PowerCappingSpec) DeepCopy() *PowerCappingSpec {
	if in == nil {
		return nil
	}
	out := new(PowerCappingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelativePowerCapInPercentageSpec) DeepCopyInto(out *RelativePowerCapInPercentageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelativePowerCapInPercentageSpec.
func (in *RelativePowerCapInPercentageSpec) DeepCopy() *RelativePowerCapInPercentageSpec {
	if in == nil {
		return nil
	}
	out := new(RelativePowerCapInPercentageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RelativeTemperatureThresholdInPercentageSpec) DeepCopyInto(out *RelativeTemperatureThresholdInPercentageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RelativeTemperatureThresholdInPercentageSpec.
func (in *RelativeTemperatureThresholdInPercentageSpec) DeepCopy() *RelativeTemperatureThresholdInPercentageSpec {
	if in == nil {
		return nil
	}
	out := new(RelativeTemperatureThresholdInPercentageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemperatureThresholdSpec) DeepCopyInto(out *TemperatureThresholdSpec) {
	*out = *in
	out.AbsoluteTemperatureThresholdInCelsiusSpec = in.AbsoluteTemperatureThresholdInCelsiusSpec
	out.RelativeTemperatureThresholdInPercentageSpec = in.RelativeTemperatureThresholdInPercentageSpec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemperatureThresholdSpec.
func (in *TemperatureThresholdSpec) DeepCopy() *TemperatureThresholdSpec {
	if in == nil {
		return nil
	}
	out := new(TemperatureThresholdSpec)
	in.DeepCopyInto(out)
	return out
}
