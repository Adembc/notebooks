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

package v1beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityProbe) DeepCopyInto(out *ActivityProbe) {
	*out = *in
	if in.Exec != nil {
		in, out := &in.Exec, &out.Exec
		*out = new(ActivityProbeExec)
		(*in).DeepCopyInto(*out)
	}
	if in.Jupyter != nil {
		in, out := &in.Jupyter, &out.Jupyter
		*out = new(ActivityProbeJupyter)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityProbe.
func (in *ActivityProbe) DeepCopy() *ActivityProbe {
	if in == nil {
		return nil
	}
	out := new(ActivityProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityProbeExec) DeepCopyInto(out *ActivityProbeExec) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityProbeExec.
func (in *ActivityProbeExec) DeepCopy() *ActivityProbeExec {
	if in == nil {
		return nil
	}
	out := new(ActivityProbeExec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActivityProbeJupyter) DeepCopyInto(out *ActivityProbeJupyter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActivityProbeJupyter.
func (in *ActivityProbeJupyter) DeepCopy() *ActivityProbeJupyter {
	if in == nil {
		return nil
	}
	out := new(ActivityProbeJupyter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPProxy) DeepCopyInto(out *HTTPProxy) {
	*out = *in
	if in.RemovePathPrefix != nil {
		in, out := &in.RemovePathPrefix, &out.RemovePathPrefix
		*out = new(bool)
		**out = **in
	}
	if in.RequestHeaders != nil {
		in, out := &in.RequestHeaders, &out.RequestHeaders
		*out = new(IstioHeaderOperations)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPProxy.
func (in *HTTPProxy) DeepCopy() *HTTPProxy {
	if in == nil {
		return nil
	}
	out := new(HTTPProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfig) DeepCopyInto(out *ImageConfig) {
	*out = *in
	out.Spawner = in.Spawner
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]ImageConfigValue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfig.
func (in *ImageConfig) DeepCopy() *ImageConfig {
	if in == nil {
		return nil
	}
	out := new(ImageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfigSpec) DeepCopyInto(out *ImageConfigSpec) {
	*out = *in
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]ImagePort, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfigSpec.
func (in *ImageConfigSpec) DeepCopy() *ImageConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ImageConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfigValue) DeepCopyInto(out *ImageConfigValue) {
	*out = *in
	in.Spawner.DeepCopyInto(&out.Spawner)
	if in.Redirect != nil {
		in, out := &in.Redirect, &out.Redirect
		*out = new(OptionRedirect)
		(*in).DeepCopyInto(*out)
	}
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfigValue.
func (in *ImageConfigValue) DeepCopy() *ImageConfigValue {
	if in == nil {
		return nil
	}
	out := new(ImageConfigValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagePort) DeepCopyInto(out *ImagePort) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagePort.
func (in *ImagePort) DeepCopy() *ImagePort {
	if in == nil {
		return nil
	}
	out := new(ImagePort)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IstioHeaderOperations) DeepCopyInto(out *IstioHeaderOperations) {
	*out = *in
	if in.Set != nil {
		in, out := &in.Set, &out.Set
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Add != nil {
		in, out := &in.Add, &out.Add
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Remove != nil {
		in, out := &in.Remove, &out.Remove
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IstioHeaderOperations.
func (in *IstioHeaderOperations) DeepCopy() *IstioHeaderOperations {
	if in == nil {
		return nil
	}
	out := new(IstioHeaderOperations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionMetric) DeepCopyInto(out *OptionMetric) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionMetric.
func (in *OptionMetric) DeepCopy() *OptionMetric {
	if in == nil {
		return nil
	}
	out := new(OptionMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionRedirect) DeepCopyInto(out *OptionRedirect) {
	*out = *in
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(RedirectMessage)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionRedirect.
func (in *OptionRedirect) DeepCopy() *OptionRedirect {
	if in == nil {
		return nil
	}
	out := new(OptionRedirect)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionSpawnerInfo) DeepCopyInto(out *OptionSpawnerInfo) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]OptionSpawnerLabel, len(*in))
		copy(*out, *in)
	}
	if in.Hidden != nil {
		in, out := &in.Hidden, &out.Hidden
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionSpawnerInfo.
func (in *OptionSpawnerInfo) DeepCopy() *OptionSpawnerInfo {
	if in == nil {
		return nil
	}
	out := new(OptionSpawnerInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionSpawnerLabel) DeepCopyInto(out *OptionSpawnerLabel) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionSpawnerLabel.
func (in *OptionSpawnerLabel) DeepCopy() *OptionSpawnerLabel {
	if in == nil {
		return nil
	}
	out := new(OptionSpawnerLabel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsSpawnerConfig) DeepCopyInto(out *OptionsSpawnerConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsSpawnerConfig.
func (in *OptionsSpawnerConfig) DeepCopy() *OptionsSpawnerConfig {
	if in == nil {
		return nil
	}
	out := new(OptionsSpawnerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodConfig) DeepCopyInto(out *PodConfig) {
	*out = *in
	out.Spawner = in.Spawner
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]PodConfigValue, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodConfig.
func (in *PodConfig) DeepCopy() *PodConfig {
	if in == nil {
		return nil
	}
	out := new(PodConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodConfigSpec) DeepCopyInto(out *PodConfigSpec) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodConfigSpec.
func (in *PodConfigSpec) DeepCopy() *PodConfigSpec {
	if in == nil {
		return nil
	}
	out := new(PodConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodConfigValue) DeepCopyInto(out *PodConfigValue) {
	*out = *in
	in.Spawner.DeepCopyInto(&out.Spawner)
	if in.Redirect != nil {
		in, out := &in.Redirect, &out.Redirect
		*out = new(OptionRedirect)
		(*in).DeepCopyInto(*out)
	}
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodConfigValue.
func (in *PodConfigValue) DeepCopy() *PodConfigValue {
	if in == nil {
		return nil
	}
	out := new(PodConfigValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTemplateOptionsMetrics) DeepCopyInto(out *PodTemplateOptionsMetrics) {
	*out = *in
	if in.ImageConfig != nil {
		in, out := &in.ImageConfig, &out.ImageConfig
		*out = make([]OptionMetric, len(*in))
		copy(*out, *in)
	}
	if in.PodConfig != nil {
		in, out := &in.PodConfig, &out.PodConfig
		*out = make([]OptionMetric, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTemplateOptionsMetrics.
func (in *PodTemplateOptionsMetrics) DeepCopy() *PodTemplateOptionsMetrics {
	if in == nil {
		return nil
	}
	out := new(PodTemplateOptionsMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodVolumeMount) DeepCopyInto(out *PodVolumeMount) {
	*out = *in
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodVolumeMount.
func (in *PodVolumeMount) DeepCopy() *PodVolumeMount {
	if in == nil {
		return nil
	}
	out := new(PodVolumeMount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedirectMessage) DeepCopyInto(out *RedirectMessage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedirectMessage.
func (in *RedirectMessage) DeepCopy() *RedirectMessage {
	if in == nil {
		return nil
	}
	out := new(RedirectMessage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workspace) DeepCopyInto(out *Workspace) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workspace.
func (in *Workspace) DeepCopy() *Workspace {
	if in == nil {
		return nil
	}
	out := new(Workspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workspace) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceActivity) DeepCopyInto(out *WorkspaceActivity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceActivity.
func (in *WorkspaceActivity) DeepCopy() *WorkspaceActivity {
	if in == nil {
		return nil
	}
	out := new(WorkspaceActivity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKind) DeepCopyInto(out *WorkspaceKind) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKind.
func (in *WorkspaceKind) DeepCopy() *WorkspaceKind {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceKind) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindConfigMap) DeepCopyInto(out *WorkspaceKindConfigMap) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindConfigMap.
func (in *WorkspaceKindConfigMap) DeepCopy() *WorkspaceKindConfigMap {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindConfigMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindCullingConfig) DeepCopyInto(out *WorkspaceKindCullingConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.MaxInactiveSeconds != nil {
		in, out := &in.MaxInactiveSeconds, &out.MaxInactiveSeconds
		*out = new(int32)
		**out = **in
	}
	if in.MinimumProbeIntervalSeconds != nil {
		in, out := &in.MinimumProbeIntervalSeconds, &out.MinimumProbeIntervalSeconds
		*out = new(int32)
		**out = **in
	}
	in.ActivityProbe.DeepCopyInto(&out.ActivityProbe)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindCullingConfig.
func (in *WorkspaceKindCullingConfig) DeepCopy() *WorkspaceKindCullingConfig {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindCullingConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindIcon) DeepCopyInto(out *WorkspaceKindIcon) {
	*out = *in
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = new(string)
		**out = **in
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(WorkspaceKindConfigMap)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindIcon.
func (in *WorkspaceKindIcon) DeepCopy() *WorkspaceKindIcon {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindIcon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindList) DeepCopyInto(out *WorkspaceKindList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkspaceKind, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindList.
func (in *WorkspaceKindList) DeepCopy() *WorkspaceKindList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceKindList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindPodMetadata) DeepCopyInto(out *WorkspaceKindPodMetadata) {
	*out = *in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindPodMetadata.
func (in *WorkspaceKindPodMetadata) DeepCopy() *WorkspaceKindPodMetadata {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindPodMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindPodOptions) DeepCopyInto(out *WorkspaceKindPodOptions) {
	*out = *in
	in.ImageConfig.DeepCopyInto(&out.ImageConfig)
	in.PodConfig.DeepCopyInto(&out.PodConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindPodOptions.
func (in *WorkspaceKindPodOptions) DeepCopy() *WorkspaceKindPodOptions {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindPodOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindPodTemplate) DeepCopyInto(out *WorkspaceKindPodTemplate) {
	*out = *in
	if in.PodMetadata != nil {
		in, out := &in.PodMetadata, &out.PodMetadata
		*out = new(WorkspaceKindPodMetadata)
		(*in).DeepCopyInto(*out)
	}
	out.ServiceAccount = in.ServiceAccount
	if in.Culling != nil {
		in, out := &in.Culling, &out.Culling
		*out = new(WorkspaceKindCullingConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Probes != nil {
		in, out := &in.Probes, &out.Probes
		*out = new(WorkspaceKindProbes)
		(*in).DeepCopyInto(*out)
	}
	out.VolumeMounts = in.VolumeMounts
	if in.HTTPProxy != nil {
		in, out := &in.HTTPProxy, &out.HTTPProxy
		*out = new(HTTPProxy)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraEnv != nil {
		in, out := &in.ExtraEnv, &out.ExtraEnv
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumeMounts != nil {
		in, out := &in.ExtraVolumeMounts, &out.ExtraVolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumes != nil {
		in, out := &in.ExtraVolumes, &out.ExtraVolumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ContainerSecurityContext != nil {
		in, out := &in.ContainerSecurityContext, &out.ContainerSecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	in.Options.DeepCopyInto(&out.Options)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindPodTemplate.
func (in *WorkspaceKindPodTemplate) DeepCopy() *WorkspaceKindPodTemplate {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindPodTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindProbes) DeepCopyInto(out *WorkspaceKindProbes) {
	*out = *in
	if in.StartupProbe != nil {
		in, out := &in.StartupProbe, &out.StartupProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(v1.Probe)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindProbes.
func (in *WorkspaceKindProbes) DeepCopy() *WorkspaceKindProbes {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindProbes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindServiceAccount) DeepCopyInto(out *WorkspaceKindServiceAccount) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindServiceAccount.
func (in *WorkspaceKindServiceAccount) DeepCopy() *WorkspaceKindServiceAccount {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindSpawner) DeepCopyInto(out *WorkspaceKindSpawner) {
	*out = *in
	if in.Hidden != nil {
		in, out := &in.Hidden, &out.Hidden
		*out = new(bool)
		**out = **in
	}
	if in.Deprecated != nil {
		in, out := &in.Deprecated, &out.Deprecated
		*out = new(bool)
		**out = **in
	}
	if in.DeprecationMessage != nil {
		in, out := &in.DeprecationMessage, &out.DeprecationMessage
		*out = new(string)
		**out = **in
	}
	in.Icon.DeepCopyInto(&out.Icon)
	in.Logo.DeepCopyInto(&out.Logo)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindSpawner.
func (in *WorkspaceKindSpawner) DeepCopy() *WorkspaceKindSpawner {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindSpawner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindSpec) DeepCopyInto(out *WorkspaceKindSpec) {
	*out = *in
	in.Spawner.DeepCopyInto(&out.Spawner)
	in.PodTemplate.DeepCopyInto(&out.PodTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindSpec.
func (in *WorkspaceKindSpec) DeepCopy() *WorkspaceKindSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindStatus) DeepCopyInto(out *WorkspaceKindStatus) {
	*out = *in
	in.PodTemplateOptions.DeepCopyInto(&out.PodTemplateOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindStatus.
func (in *WorkspaceKindStatus) DeepCopy() *WorkspaceKindStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceKindVolumeMounts) DeepCopyInto(out *WorkspaceKindVolumeMounts) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceKindVolumeMounts.
func (in *WorkspaceKindVolumeMounts) DeepCopy() *WorkspaceKindVolumeMounts {
	if in == nil {
		return nil
	}
	out := new(WorkspaceKindVolumeMounts)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceList) DeepCopyInto(out *WorkspaceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workspace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceList.
func (in *WorkspaceList) DeepCopy() *WorkspaceList {
	if in == nil {
		return nil
	}
	out := new(WorkspaceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkspaceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacePodMetadata) DeepCopyInto(out *WorkspacePodMetadata) {
	*out = *in
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
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacePodMetadata.
func (in *WorkspacePodMetadata) DeepCopy() *WorkspacePodMetadata {
	if in == nil {
		return nil
	}
	out := new(WorkspacePodMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacePodOptionInfo) DeepCopyInto(out *WorkspacePodOptionInfo) {
	*out = *in
	if in.RedirectChain != nil {
		in, out := &in.RedirectChain, &out.RedirectChain
		*out = make([]WorkspacePodOptionRedirectStep, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacePodOptionInfo.
func (in *WorkspacePodOptionInfo) DeepCopy() *WorkspacePodOptionInfo {
	if in == nil {
		return nil
	}
	out := new(WorkspacePodOptionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacePodOptionRedirectStep) DeepCopyInto(out *WorkspacePodOptionRedirectStep) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacePodOptionRedirectStep.
func (in *WorkspacePodOptionRedirectStep) DeepCopy() *WorkspacePodOptionRedirectStep {
	if in == nil {
		return nil
	}
	out := new(WorkspacePodOptionRedirectStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacePodOptions) DeepCopyInto(out *WorkspacePodOptions) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacePodOptions.
func (in *WorkspacePodOptions) DeepCopy() *WorkspacePodOptions {
	if in == nil {
		return nil
	}
	out := new(WorkspacePodOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacePodOptionsStatus) DeepCopyInto(out *WorkspacePodOptionsStatus) {
	*out = *in
	in.ImageConfig.DeepCopyInto(&out.ImageConfig)
	in.PodConfig.DeepCopyInto(&out.PodConfig)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacePodOptionsStatus.
func (in *WorkspacePodOptionsStatus) DeepCopy() *WorkspacePodOptionsStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspacePodOptionsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacePodTemplate) DeepCopyInto(out *WorkspacePodTemplate) {
	*out = *in
	if in.PodMetadata != nil {
		in, out := &in.PodMetadata, &out.PodMetadata
		*out = new(WorkspacePodMetadata)
		(*in).DeepCopyInto(*out)
	}
	in.Volumes.DeepCopyInto(&out.Volumes)
	out.Options = in.Options
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacePodTemplate.
func (in *WorkspacePodTemplate) DeepCopy() *WorkspacePodTemplate {
	if in == nil {
		return nil
	}
	out := new(WorkspacePodTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspacePodVolumes) DeepCopyInto(out *WorkspacePodVolumes) {
	*out = *in
	if in.Home != nil {
		in, out := &in.Home, &out.Home
		*out = new(string)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]PodVolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspacePodVolumes.
func (in *WorkspacePodVolumes) DeepCopy() *WorkspacePodVolumes {
	if in == nil {
		return nil
	}
	out := new(WorkspacePodVolumes)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceSpec) DeepCopyInto(out *WorkspaceSpec) {
	*out = *in
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.DisableCulling != nil {
		in, out := &in.DisableCulling, &out.DisableCulling
		*out = new(bool)
		**out = **in
	}
	if in.DeferUpdates != nil {
		in, out := &in.DeferUpdates, &out.DeferUpdates
		*out = new(bool)
		**out = **in
	}
	in.PodTemplate.DeepCopyInto(&out.PodTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceSpec.
func (in *WorkspaceSpec) DeepCopy() *WorkspaceSpec {
	if in == nil {
		return nil
	}
	out := new(WorkspaceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceStatus) DeepCopyInto(out *WorkspaceStatus) {
	*out = *in
	out.Activity = in.Activity
	in.PodTemplateOptions.DeepCopyInto(&out.PodTemplateOptions)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceStatus.
func (in *WorkspaceStatus) DeepCopy() *WorkspaceStatus {
	if in == nil {
		return nil
	}
	out := new(WorkspaceStatus)
	in.DeepCopyInto(out)
	return out
}
