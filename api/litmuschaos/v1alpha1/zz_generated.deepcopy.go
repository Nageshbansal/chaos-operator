//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2019 LitmusChaos Authors

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
	"k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationParams) DeepCopyInto(out *ApplicationParams) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationParams.
func (in *ApplicationParams) DeepCopy() *ApplicationParams {
	if in == nil {
		return nil
	}
	out := new(ApplicationParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosEngine) DeepCopyInto(out *ChaosEngine) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosEngine.
func (in *ChaosEngine) DeepCopy() *ChaosEngine {
	if in == nil {
		return nil
	}
	out := new(ChaosEngine)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosEngine) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosEngineList) DeepCopyInto(out *ChaosEngineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChaosEngine, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosEngineList.
func (in *ChaosEngineList) DeepCopy() *ChaosEngineList {
	if in == nil {
		return nil
	}
	out := new(ChaosEngineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosEngineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosEngineSpec) DeepCopyInto(out *ChaosEngineSpec) {
	*out = *in
	out.Appinfo = in.Appinfo
	in.Components.DeepCopyInto(&out.Components)
	if in.Experiments != nil {
		in, out := &in.Experiments, &out.Experiments
		*out = make([]ExperimentList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = new(Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosEngineSpec.
func (in *ChaosEngineSpec) DeepCopy() *ChaosEngineSpec {
	if in == nil {
		return nil
	}
	out := new(ChaosEngineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosEngineStatus) DeepCopyInto(out *ChaosEngineStatus) {
	*out = *in
	if in.Experiments != nil {
		in, out := &in.Experiments, &out.Experiments
		*out = make([]ExperimentStatuses, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosEngineStatus.
func (in *ChaosEngineStatus) DeepCopy() *ChaosEngineStatus {
	if in == nil {
		return nil
	}
	out := new(ChaosEngineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosExperiment) DeepCopyInto(out *ChaosExperiment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosExperiment.
func (in *ChaosExperiment) DeepCopy() *ChaosExperiment {
	if in == nil {
		return nil
	}
	out := new(ChaosExperiment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosExperiment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosExperimentList) DeepCopyInto(out *ChaosExperimentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChaosExperiment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosExperimentList.
func (in *ChaosExperimentList) DeepCopy() *ChaosExperimentList {
	if in == nil {
		return nil
	}
	out := new(ChaosExperimentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosExperimentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosExperimentSpec) DeepCopyInto(out *ChaosExperimentSpec) {
	*out = *in
	in.Definition.DeepCopyInto(&out.Definition)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosExperimentSpec.
func (in *ChaosExperimentSpec) DeepCopy() *ChaosExperimentSpec {
	if in == nil {
		return nil
	}
	out := new(ChaosExperimentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosExperimentStatus) DeepCopyInto(out *ChaosExperimentStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosExperimentStatus.
func (in *ChaosExperimentStatus) DeepCopy() *ChaosExperimentStatus {
	if in == nil {
		return nil
	}
	out := new(ChaosExperimentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosResult) DeepCopyInto(out *ChaosResult) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosResult.
func (in *ChaosResult) DeepCopy() *ChaosResult {
	if in == nil {
		return nil
	}
	out := new(ChaosResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosResult) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosResultList) DeepCopyInto(out *ChaosResultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChaosResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosResultList.
func (in *ChaosResultList) DeepCopy() *ChaosResultList {
	if in == nil {
		return nil
	}
	out := new(ChaosResultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosResultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosResultSpec) DeepCopyInto(out *ChaosResultSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosResultSpec.
func (in *ChaosResultSpec) DeepCopy() *ChaosResultSpec {
	if in == nil {
		return nil
	}
	out := new(ChaosResultSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosResultStatus) DeepCopyInto(out *ChaosResultStatus) {
	*out = *in
	in.ExperimentStatus.DeepCopyInto(&out.ExperimentStatus)
	if in.ProbeStatuses != nil {
		in, out := &in.ProbeStatuses, &out.ProbeStatuses
		*out = make([]ProbeStatuses, len(*in))
		copy(*out, *in)
	}
	if in.History != nil {
		in, out := &in.History, &out.History
		*out = new(HistoryDetails)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosResultStatus.
func (in *ChaosResultStatus) DeepCopy() *ChaosResultStatus {
	if in == nil {
		return nil
	}
	out := new(ChaosResultStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CmdProbeInputs) DeepCopyInto(out *CmdProbeInputs) {
	*out = *in
	out.Comparator = in.Comparator
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(SourceDetails)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CmdProbeInputs.
func (in *CmdProbeInputs) DeepCopy() *CmdProbeInputs {
	if in == nil {
		return nil
	}
	out := new(CmdProbeInputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComparatorInfo) DeepCopyInto(out *ComparatorInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComparatorInfo.
func (in *ComparatorInfo) DeepCopy() *ComparatorInfo {
	if in == nil {
		return nil
	}
	out := new(ComparatorInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentParams) DeepCopyInto(out *ComponentParams) {
	*out = *in
	in.Runner.DeepCopyInto(&out.Runner)
	if in.Sidecar != nil {
		in, out := &in.Sidecar, &out.Sidecar
		*out = make([]Sidecar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentParams.
func (in *ComponentParams) DeepCopy() *ComponentParams {
	if in == nil {
		return nil
	}
	out := new(ComponentParams)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMap) DeepCopyInto(out *ConfigMap) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMap.
func (in *ConfigMap) DeepCopy() *ConfigMap {
	if in == nil {
		return nil
	}
	out := new(ConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorOutput) DeepCopyInto(out *ErrorOutput) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorOutput.
func (in *ErrorOutput) DeepCopy() *ErrorOutput {
	if in == nil {
		return nil
	}
	out := new(ErrorOutput)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvaluationWindow) DeepCopyInto(out *EvaluationWindow) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvaluationWindow.
func (in *EvaluationWindow) DeepCopy() *EvaluationWindow {
	if in == nil {
		return nil
	}
	out := new(EvaluationWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentAttributes) DeepCopyInto(out *ExperimentAttributes) {
	*out = *in
	in.Components.DeepCopyInto(&out.Components)
	if in.Probe != nil {
		in, out := &in.Probe, &out.Probe
		*out = make([]ProbeAttributes, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentAttributes.
func (in *ExperimentAttributes) DeepCopy() *ExperimentAttributes {
	if in == nil {
		return nil
	}
	out := new(ExperimentAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentComponents) DeepCopyInto(out *ExperimentComponents) {
	*out = *in
	if in.ENV != nil {
		in, out := &in.ENV, &out.ENV
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]ConfigMap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]Secret, len(*in))
		copy(*out, *in)
	}
	if in.ExperimentAnnotations != nil {
		in, out := &in.ExperimentAnnotations, &out.ExperimentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExperimentImagePullSecrets != nil {
		in, out := &in.ExperimentImagePullSecrets, &out.ExperimentImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.StatusCheckTimeouts = in.StatusCheckTimeouts
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentComponents.
func (in *ExperimentComponents) DeepCopy() *ExperimentComponents {
	if in == nil {
		return nil
	}
	out := new(ExperimentComponents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentDef) DeepCopyInto(out *ExperimentDef) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = make([]rbacv1.PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ENVList != nil {
		in, out := &in.ENVList, &out.ENVList
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]ConfigMap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]Secret, len(*in))
		copy(*out, *in)
	}
	if in.HostFileVolumes != nil {
		in, out := &in.HostFileVolumes, &out.HostFileVolumes
		*out = make([]HostFile, len(*in))
		copy(*out, *in)
	}
	if in.ExperimentAnnotations != nil {
		in, out := &in.ExperimentAnnotations, &out.ExperimentAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.SecurityContext.DeepCopyInto(&out.SecurityContext)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentDef.
func (in *ExperimentDef) DeepCopy() *ExperimentDef {
	if in == nil {
		return nil
	}
	out := new(ExperimentDef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentList) DeepCopyInto(out *ExperimentList) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExperimentStatuses) DeepCopyInto(out *ExperimentStatuses) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExperimentStatuses.
func (in *ExperimentStatuses) DeepCopy() *ExperimentStatuses {
	if in == nil {
		return nil
	}
	out := new(ExperimentStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GetMethod) DeepCopyInto(out *GetMethod) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetMethod.
func (in *GetMethod) DeepCopy() *GetMethod {
	if in == nil {
		return nil
	}
	out := new(GetMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPMethod) DeepCopyInto(out *HTTPMethod) {
	*out = *in
	if in.Get != nil {
		in, out := &in.Get, &out.Get
		*out = new(GetMethod)
		**out = **in
	}
	if in.Post != nil {
		in, out := &in.Post, &out.Post
		*out = new(PostMethod)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPMethod.
func (in *HTTPMethod) DeepCopy() *HTTPMethod {
	if in == nil {
		return nil
	}
	out := new(HTTPMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPProbeInputs) DeepCopyInto(out *HTTPProbeInputs) {
	*out = *in
	in.Method.DeepCopyInto(&out.Method)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPProbeInputs.
func (in *HTTPProbeInputs) DeepCopy() *HTTPProbeInputs {
	if in == nil {
		return nil
	}
	out := new(HTTPProbeInputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistoryDetails) DeepCopyInto(out *HistoryDetails) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]TargetDetails, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistoryDetails.
func (in *HistoryDetails) DeepCopy() *HistoryDetails {
	if in == nil {
		return nil
	}
	out := new(HistoryDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostFile) DeepCopyInto(out *HostFile) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostFile.
func (in *HostFile) DeepCopy() *HostFile {
	if in == nil {
		return nil
	}
	out := new(HostFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Identifier) DeepCopyInto(out *Identifier) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Identifier.
func (in *Identifier) DeepCopy() *Identifier {
	if in == nil {
		return nil
	}
	out := new(Identifier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *K8sProbeInputs) DeepCopyInto(out *K8sProbeInputs) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new K8sProbeInputs.
func (in *K8sProbeInputs) DeepCopy() *K8sProbeInputs {
	if in == nil {
		return nil
	}
	out := new(K8sProbeInputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pod) DeepCopyInto(out *Pod) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pod.
func (in *Pod) DeepCopy() *Pod {
	if in == nil {
		return nil
	}
	out := new(Pod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostMethod) DeepCopyInto(out *PostMethod) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostMethod.
func (in *PostMethod) DeepCopy() *PostMethod {
	if in == nil {
		return nil
	}
	out := new(PostMethod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeAttributes) DeepCopyInto(out *ProbeAttributes) {
	*out = *in
	if in.K8sProbeInputs != nil {
		in, out := &in.K8sProbeInputs, &out.K8sProbeInputs
		*out = new(K8sProbeInputs)
		**out = **in
	}
	if in.HTTPProbeInputs != nil {
		in, out := &in.HTTPProbeInputs, &out.HTTPProbeInputs
		*out = new(HTTPProbeInputs)
		(*in).DeepCopyInto(*out)
	}
	if in.CmdProbeInputs != nil {
		in, out := &in.CmdProbeInputs, &out.CmdProbeInputs
		*out = new(CmdProbeInputs)
		(*in).DeepCopyInto(*out)
	}
	if in.PromProbeInputs != nil {
		in, out := &in.PromProbeInputs, &out.PromProbeInputs
		*out = new(PromProbeInputs)
		**out = **in
	}
	if in.SLOProbeInputs != nil {
		in, out := &in.SLOProbeInputs, &out.SLOProbeInputs
		*out = new(SLOProbeInputs)
		(*in).DeepCopyInto(*out)
	}
	out.RunProperties = in.RunProperties
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeAttributes.
func (in *ProbeAttributes) DeepCopy() *ProbeAttributes {
	if in == nil {
		return nil
	}
	out := new(ProbeAttributes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeStatus) DeepCopyInto(out *ProbeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeStatus.
func (in *ProbeStatus) DeepCopy() *ProbeStatus {
	if in == nil {
		return nil
	}
	out := new(ProbeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeStatuses) DeepCopyInto(out *ProbeStatuses) {
	*out = *in
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeStatuses.
func (in *ProbeStatuses) DeepCopy() *ProbeStatuses {
	if in == nil {
		return nil
	}
	out := new(ProbeStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromProbeInputs) DeepCopyInto(out *PromProbeInputs) {
	*out = *in
	out.Comparator = in.Comparator
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromProbeInputs.
func (in *PromProbeInputs) DeepCopy() *PromProbeInputs {
	if in == nil {
		return nil
	}
	out := new(PromProbeInputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunProperty) DeepCopyInto(out *RunProperty) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunProperty.
func (in *RunProperty) DeepCopy() *RunProperty {
	if in == nil {
		return nil
	}
	out := new(RunProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerInfo) DeepCopyInto(out *RunnerInfo) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.RunnerAnnotation != nil {
		in, out := &in.RunnerAnnotation, &out.RunnerAnnotation
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.RunnerLabels != nil {
		in, out := &in.RunnerLabels, &out.RunnerLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ConfigMaps != nil {
		in, out := &in.ConfigMaps, &out.ConfigMaps
		*out = make([]ConfigMap, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]Secret, len(*in))
		copy(*out, *in)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerInfo.
func (in *RunnerInfo) DeepCopy() *RunnerInfo {
	if in == nil {
		return nil
	}
	out := new(RunnerInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLOProbeInputs) DeepCopyInto(out *SLOProbeInputs) {
	*out = *in
	if in.EvaluationWindow != nil {
		in, out := &in.EvaluationWindow, &out.EvaluationWindow
		*out = new(EvaluationWindow)
		**out = **in
	}
	out.SLOSourceMetadata = in.SLOSourceMetadata
	out.Comparator = in.Comparator
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLOProbeInputs.
func (in *SLOProbeInputs) DeepCopy() *SLOProbeInputs {
	if in == nil {
		return nil
	}
	out := new(SLOProbeInputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SLOSourceMetadata) DeepCopyInto(out *SLOSourceMetadata) {
	*out = *in
	out.Scope = in.Scope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SLOSourceMetadata.
func (in *SLOSourceMetadata) DeepCopy() *SLOSourceMetadata {
	if in == nil {
		return nil
	}
	out := new(SLOSourceMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Secret) DeepCopyInto(out *Secret) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Secret.
func (in *Secret) DeepCopy() *Secret {
	if in == nil {
		return nil
	}
	out := new(Secret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecurityContext) DeepCopyInto(out *SecurityContext) {
	*out = *in
	in.PodSecurityContext.DeepCopyInto(&out.PodSecurityContext)
	in.ContainerSecurityContext.DeepCopyInto(&out.ContainerSecurityContext)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecurityContext.
func (in *SecurityContext) DeepCopy() *SecurityContext {
	if in == nil {
		return nil
	}
	out := new(SecurityContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
	if in.Workloads != nil {
		in, out := &in.Workloads, &out.Workloads
		*out = make([]Workload, len(*in))
		copy(*out, *in)
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]Pod, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sidecar) DeepCopyInto(out *Sidecar) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make([]Secret, len(*in))
		copy(*out, *in)
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ENV != nil {
		in, out := &in.ENV, &out.ENV
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sidecar.
func (in *Sidecar) DeepCopy() *Sidecar {
	if in == nil {
		return nil
	}
	out := new(Sidecar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SourceDetails) DeepCopyInto(out *SourceDetails) {
	*out = *in
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ENVList != nil {
		in, out := &in.ENVList, &out.ENVList
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumesMount != nil {
		in, out := &in.VolumesMount, &out.VolumesMount
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SourceDetails.
func (in *SourceDetails) DeepCopy() *SourceDetails {
	if in == nil {
		return nil
	}
	out := new(SourceDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatusCheckTimeout) DeepCopyInto(out *StatusCheckTimeout) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatusCheckTimeout.
func (in *StatusCheckTimeout) DeepCopy() *StatusCheckTimeout {
	if in == nil {
		return nil
	}
	out := new(StatusCheckTimeout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetDetails) DeepCopyInto(out *TargetDetails) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetDetails.
func (in *TargetDetails) DeepCopy() *TargetDetails {
	if in == nil {
		return nil
	}
	out := new(TargetDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TestStatus) DeepCopyInto(out *TestStatus) {
	*out = *in
	if in.ErrorOutput != nil {
		in, out := &in.ErrorOutput, &out.ErrorOutput
		*out = new(ErrorOutput)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TestStatus.
func (in *TestStatus) DeepCopy() *TestStatus {
	if in == nil {
		return nil
	}
	out := new(TestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}
