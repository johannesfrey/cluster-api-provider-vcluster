//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by defaulter-gen. DO NOT EDIT.

package management

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&ProjectClusters{}, func(obj interface{}) { SetObjectDefaults_ProjectClusters(obj.(*ProjectClusters)) })
	scheme.AddTypeDefaultingFunc(&ProjectClustersList{}, func(obj interface{}) { SetObjectDefaults_ProjectClustersList(obj.(*ProjectClustersList)) })
	scheme.AddTypeDefaultingFunc(&Runner{}, func(obj interface{}) { SetObjectDefaults_Runner(obj.(*Runner)) })
	scheme.AddTypeDefaultingFunc(&RunnerList{}, func(obj interface{}) { SetObjectDefaults_RunnerList(obj.(*RunnerList)) })
	return nil
}

func SetObjectDefaults_ProjectClusters(in *ProjectClusters) {
	for i := range in.Runners {
		a := &in.Runners[i]
		SetObjectDefaults_Runner(a)
	}
}

func SetObjectDefaults_ProjectClustersList(in *ProjectClustersList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_ProjectClusters(a)
	}
}

func SetObjectDefaults_Runner(in *Runner) {
	if in.Spec.RunnerSpec.ClusterRef != nil {
		if in.Spec.RunnerSpec.ClusterRef.PodTemplate != nil {
			for i := range in.Spec.RunnerSpec.ClusterRef.PodTemplate.Spec.InitContainers {
				a := &in.Spec.RunnerSpec.ClusterRef.PodTemplate.Spec.InitContainers[i]
				for j := range a.Ports {
					b := &a.Ports[j]
					if b.Protocol == "" {
						b.Protocol = "TCP"
					}
				}
				if a.LivenessProbe != nil {
					if a.LivenessProbe.ProbeHandler.GRPC != nil {
						if a.LivenessProbe.ProbeHandler.GRPC.Service == nil {
							var ptrVar1 string = ""
							a.LivenessProbe.ProbeHandler.GRPC.Service = &ptrVar1
						}
					}
				}
				if a.ReadinessProbe != nil {
					if a.ReadinessProbe.ProbeHandler.GRPC != nil {
						if a.ReadinessProbe.ProbeHandler.GRPC.Service == nil {
							var ptrVar1 string = ""
							a.ReadinessProbe.ProbeHandler.GRPC.Service = &ptrVar1
						}
					}
				}
				if a.StartupProbe != nil {
					if a.StartupProbe.ProbeHandler.GRPC != nil {
						if a.StartupProbe.ProbeHandler.GRPC.Service == nil {
							var ptrVar1 string = ""
							a.StartupProbe.ProbeHandler.GRPC.Service = &ptrVar1
						}
					}
				}
			}
		}
	}
}

func SetObjectDefaults_RunnerList(in *RunnerList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Runner(a)
	}
}
