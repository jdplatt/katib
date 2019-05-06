// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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
// Code generated by main. DO NOT EDIT.

package v1alpha2

import (
	trialv1alpha2 "github.com/kubeflow/katib/pkg/api/operators/apis/trial/v1alpha2"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlgorithmSetting) DeepCopyInto(out *AlgorithmSetting) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlgorithmSetting.
func (in *AlgorithmSetting) DeepCopy() *AlgorithmSetting {
	if in == nil {
		return nil
	}
	out := new(AlgorithmSetting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlgorithmSpec) DeepCopyInto(out *AlgorithmSpec) {
	*out = *in
	if in.AlgorithmSettings != nil {
		in, out := &in.AlgorithmSettings, &out.AlgorithmSettings
		*out = make([]AlgorithmSetting, len(*in))
		copy(*out, *in)
	}
	if in.EarlyStopping != nil {
		in, out := &in.EarlyStopping, &out.EarlyStopping
		*out = new(EarlyStoppingSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlgorithmSpec.
func (in *AlgorithmSpec) DeepCopy() *AlgorithmSpec {
	if in == nil {
		return nil
	}
	out := new(AlgorithmSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EarlyStoppingSpec) DeepCopyInto(out *EarlyStoppingSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EarlyStoppingSpec.
func (in *EarlyStoppingSpec) DeepCopy() *EarlyStoppingSpec {
	if in == nil {
		return nil
	}
	out := new(EarlyStoppingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Experiment) DeepCopyInto(out *Experiment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Experiment.
func (in *Experiment) DeepCopy() *Experiment {
	if in == nil {
		return nil
	}
	out := new(Experiment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Experiment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentCondition) DeepCopyInto(out *ExperimentCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentCondition.
func (in *ExperimentCondition) DeepCopy() *ExperimentCondition {
	if in == nil {
		return nil
	}
	out := new(ExperimentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentList) DeepCopyInto(out *ExperimentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Experiment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentList.
func (in *ExperimentList) DeepCopy() *ExperimentList {
	if in == nil {
		return nil
	}
	out := new(ExperimentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ExperimentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentSpec) DeepCopyInto(out *ExperimentSpec) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ParameterSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Objective != nil {
		in, out := &in.Objective, &out.Objective
		*out = new(ObjectiveSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
		*out = new(AlgorithmSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.TrialTemplate != nil {
		in, out := &in.TrialTemplate, &out.TrialTemplate
		*out = new(TrialTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.ParallelTrialCount != nil {
		in, out := &in.ParallelTrialCount, &out.ParallelTrialCount
		*out = new(int)
		**out = **in
	}
	if in.MaxTrialCount != nil {
		in, out := &in.MaxTrialCount, &out.MaxTrialCount
		*out = new(int)
		**out = **in
	}
	if in.MaxFailedTrialCount != nil {
		in, out := &in.MaxFailedTrialCount, &out.MaxFailedTrialCount
		*out = new(int)
		**out = **in
	}
	if in.MetricsCollectorSpec != nil {
		in, out := &in.MetricsCollectorSpec, &out.MetricsCollectorSpec
		*out = new(MetricsCollectorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NasConfig != nil {
		in, out := &in.NasConfig, &out.NasConfig
		*out = new(NasConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentSpec.
func (in *ExperimentSpec) DeepCopy() *ExperimentSpec {
	if in == nil {
		return nil
	}
	out := new(ExperimentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentStatus) DeepCopyInto(out *ExperimentStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		*out = (*in).DeepCopy()
	}
	if in.LastReconcileTime != nil {
		in, out := &in.LastReconcileTime, &out.LastReconcileTime
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]ExperimentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CurrentOptimalTrial.DeepCopyInto(&out.CurrentOptimalTrial)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentStatus.
func (in *ExperimentStatus) DeepCopy() *ExperimentStatus {
	if in == nil {
		return nil
	}
	out := new(ExperimentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeasibleSpace) DeepCopyInto(out *FeasibleSpace) {
	*out = *in
	if in.List != nil {
		in, out := &in.List, &out.List
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeasibleSpace.
func (in *FeasibleSpace) DeepCopy() *FeasibleSpace {
	if in == nil {
		return nil
	}
	out := new(FeasibleSpace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoTemplate) DeepCopyInto(out *GoTemplate) {
	*out = *in
	if in.TemplateSpec != nil {
		in, out := &in.TemplateSpec, &out.TemplateSpec
		*out = new(TemplateSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoTemplate.
func (in *GoTemplate) DeepCopy() *GoTemplate {
	if in == nil {
		return nil
	}
	out := new(GoTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GraphConfig) DeepCopyInto(out *GraphConfig) {
	*out = *in
	if in.NumLayers != nil {
		in, out := &in.NumLayers, &out.NumLayers
		*out = new(int32)
		**out = **in
	}
	if in.InputSizes != nil {
		in, out := &in.InputSizes, &out.InputSizes
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	if in.OutputSizes != nil {
		in, out := &in.OutputSizes, &out.OutputSizes
		*out = make([]int32, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GraphConfig.
func (in *GraphConfig) DeepCopy() *GraphConfig {
	if in == nil {
		return nil
	}
	out := new(GraphConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsCollectorSpec) DeepCopyInto(out *MetricsCollectorSpec) {
	*out = *in
	in.GoTemplate.DeepCopyInto(&out.GoTemplate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsCollectorSpec.
func (in *MetricsCollectorSpec) DeepCopy() *MetricsCollectorSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsCollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NasConfig) DeepCopyInto(out *NasConfig) {
	*out = *in
	in.GraphConfig.DeepCopyInto(&out.GraphConfig)
	if in.Operations != nil {
		in, out := &in.Operations, &out.Operations
		*out = make([]Operation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NasConfig.
func (in *NasConfig) DeepCopy() *NasConfig {
	if in == nil {
		return nil
	}
	out := new(NasConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectiveSpec) DeepCopyInto(out *ObjectiveSpec) {
	*out = *in
	if in.Goal != nil {
		in, out := &in.Goal, &out.Goal
		*out = new(float64)
		**out = **in
	}
	if in.AdditionalMetricsNames != nil {
		in, out := &in.AdditionalMetricsNames, &out.AdditionalMetricsNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectiveSpec.
func (in *ObjectiveSpec) DeepCopy() *ObjectiveSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectiveSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Operation) DeepCopyInto(out *Operation) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make([]ParameterSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Operation.
func (in *Operation) DeepCopy() *Operation {
	if in == nil {
		return nil
	}
	out := new(Operation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptimalTrial) DeepCopyInto(out *OptimalTrial) {
	*out = *in
	if in.ParameterAssignments != nil {
		in, out := &in.ParameterAssignments, &out.ParameterAssignments
		*out = make([]trialv1alpha2.ParameterAssignment, len(*in))
		copy(*out, *in)
	}
	in.Observation.DeepCopyInto(&out.Observation)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptimalTrial.
func (in *OptimalTrial) DeepCopy() *OptimalTrial {
	if in == nil {
		return nil
	}
	out := new(OptimalTrial)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParameterSpec) DeepCopyInto(out *ParameterSpec) {
	*out = *in
	in.FeasibleSpace.DeepCopyInto(&out.FeasibleSpace)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParameterSpec.
func (in *ParameterSpec) DeepCopy() *ParameterSpec {
	if in == nil {
		return nil
	}
	out := new(ParameterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TemplateSpec) DeepCopyInto(out *TemplateSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TemplateSpec.
func (in *TemplateSpec) DeepCopy() *TemplateSpec {
	if in == nil {
		return nil
	}
	out := new(TemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrialTemplate) DeepCopyInto(out *TrialTemplate) {
	*out = *in
	if in.GoTemplate != nil {
		in, out := &in.GoTemplate, &out.GoTemplate
		*out = new(GoTemplate)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrialTemplate.
func (in *TrialTemplate) DeepCopy() *TrialTemplate {
	if in == nil {
		return nil
	}
	out := new(TrialTemplate)
	in.DeepCopyInto(out)
	return out
}
