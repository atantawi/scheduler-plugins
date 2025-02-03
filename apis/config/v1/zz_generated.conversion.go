//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1 "k8s.io/kube-scheduler/config/v1"
	apisconfig "k8s.io/kubernetes/pkg/scheduler/apis/config"
	config "sigs.k8s.io/scheduler-plugins/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ClusterTopologyPlacementGroupArgs)(nil), (*config.ClusterTopologyPlacementGroupArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ClusterTopologyPlacementGroupArgs_To_config_ClusterTopologyPlacementGroupArgs(a.(*ClusterTopologyPlacementGroupArgs), b.(*config.ClusterTopologyPlacementGroupArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ClusterTopologyPlacementGroupArgs)(nil), (*ClusterTopologyPlacementGroupArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ClusterTopologyPlacementGroupArgs_To_v1_ClusterTopologyPlacementGroupArgs(a.(*config.ClusterTopologyPlacementGroupArgs), b.(*ClusterTopologyPlacementGroupArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CoschedulingArgs)(nil), (*config.CoschedulingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_CoschedulingArgs_To_config_CoschedulingArgs(a.(*CoschedulingArgs), b.(*config.CoschedulingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.CoschedulingArgs)(nil), (*CoschedulingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_CoschedulingArgs_To_v1_CoschedulingArgs(a.(*config.CoschedulingArgs), b.(*CoschedulingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LoadVariationRiskBalancingArgs)(nil), (*config.LoadVariationRiskBalancingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LoadVariationRiskBalancingArgs_To_config_LoadVariationRiskBalancingArgs(a.(*LoadVariationRiskBalancingArgs), b.(*config.LoadVariationRiskBalancingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LoadVariationRiskBalancingArgs)(nil), (*LoadVariationRiskBalancingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LoadVariationRiskBalancingArgs_To_v1_LoadVariationRiskBalancingArgs(a.(*config.LoadVariationRiskBalancingArgs), b.(*LoadVariationRiskBalancingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LowRiskOverCommitmentArgs)(nil), (*config.LowRiskOverCommitmentArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_LowRiskOverCommitmentArgs_To_config_LowRiskOverCommitmentArgs(a.(*LowRiskOverCommitmentArgs), b.(*config.LowRiskOverCommitmentArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LowRiskOverCommitmentArgs)(nil), (*LowRiskOverCommitmentArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LowRiskOverCommitmentArgs_To_v1_LowRiskOverCommitmentArgs(a.(*config.LowRiskOverCommitmentArgs), b.(*LowRiskOverCommitmentArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetricProviderSpec)(nil), (*config.MetricProviderSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_MetricProviderSpec_To_config_MetricProviderSpec(a.(*MetricProviderSpec), b.(*config.MetricProviderSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MetricProviderSpec)(nil), (*MetricProviderSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MetricProviderSpec_To_v1_MetricProviderSpec(a.(*config.MetricProviderSpec), b.(*MetricProviderSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkOverheadArgs)(nil), (*config.NetworkOverheadArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NetworkOverheadArgs_To_config_NetworkOverheadArgs(a.(*NetworkOverheadArgs), b.(*config.NetworkOverheadArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NetworkOverheadArgs)(nil), (*NetworkOverheadArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NetworkOverheadArgs_To_v1_NetworkOverheadArgs(a.(*config.NetworkOverheadArgs), b.(*NetworkOverheadArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeResourceTopologyCache)(nil), (*config.NodeResourceTopologyCache)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NodeResourceTopologyCache_To_config_NodeResourceTopologyCache(a.(*NodeResourceTopologyCache), b.(*config.NodeResourceTopologyCache), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeResourceTopologyCache)(nil), (*NodeResourceTopologyCache)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeResourceTopologyCache_To_v1_NodeResourceTopologyCache(a.(*config.NodeResourceTopologyCache), b.(*NodeResourceTopologyCache), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeResourcesAllocatableArgs)(nil), (*config.NodeResourcesAllocatableArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs(a.(*NodeResourcesAllocatableArgs), b.(*config.NodeResourcesAllocatableArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeResourcesAllocatableArgs)(nil), (*NodeResourcesAllocatableArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeResourcesAllocatableArgs_To_v1_NodeResourcesAllocatableArgs(a.(*config.NodeResourcesAllocatableArgs), b.(*NodeResourcesAllocatableArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PeaksArgs)(nil), (*config.PeaksArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PeaksArgs_To_config_PeaksArgs(a.(*PeaksArgs), b.(*config.PeaksArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PeaksArgs)(nil), (*PeaksArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PeaksArgs_To_v1_PeaksArgs(a.(*config.PeaksArgs), b.(*PeaksArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PowerModel)(nil), (*config.PowerModel)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PowerModel_To_config_PowerModel(a.(*PowerModel), b.(*config.PowerModel), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PowerModel)(nil), (*PowerModel)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PowerModel_To_v1_PowerModel(a.(*config.PowerModel), b.(*PowerModel), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PreemptionTolerationArgs)(nil), (*config.PreemptionTolerationArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_PreemptionTolerationArgs_To_config_PreemptionTolerationArgs(a.(*PreemptionTolerationArgs), b.(*config.PreemptionTolerationArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.PreemptionTolerationArgs)(nil), (*PreemptionTolerationArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_PreemptionTolerationArgs_To_v1_PreemptionTolerationArgs(a.(*config.PreemptionTolerationArgs), b.(*PreemptionTolerationArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ScoringStrategy)(nil), (*config.ScoringStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ScoringStrategy_To_config_ScoringStrategy(a.(*ScoringStrategy), b.(*config.ScoringStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ScoringStrategy)(nil), (*ScoringStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ScoringStrategy_To_v1_ScoringStrategy(a.(*config.ScoringStrategy), b.(*ScoringStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SySchedArgs)(nil), (*config.SySchedArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_SySchedArgs_To_config_SySchedArgs(a.(*SySchedArgs), b.(*config.SySchedArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.SySchedArgs)(nil), (*SySchedArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_SySchedArgs_To_v1_SySchedArgs(a.(*config.SySchedArgs), b.(*SySchedArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TargetLoadPackingArgs)(nil), (*config.TargetLoadPackingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TargetLoadPackingArgs_To_config_TargetLoadPackingArgs(a.(*TargetLoadPackingArgs), b.(*config.TargetLoadPackingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TargetLoadPackingArgs)(nil), (*TargetLoadPackingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TargetLoadPackingArgs_To_v1_TargetLoadPackingArgs(a.(*config.TargetLoadPackingArgs), b.(*TargetLoadPackingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TopologicalSortArgs)(nil), (*config.TopologicalSortArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TopologicalSortArgs_To_config_TopologicalSortArgs(a.(*TopologicalSortArgs), b.(*config.TopologicalSortArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TopologicalSortArgs)(nil), (*TopologicalSortArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TopologicalSortArgs_To_v1_TopologicalSortArgs(a.(*config.TopologicalSortArgs), b.(*TopologicalSortArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TrimaranSpec)(nil), (*config.TrimaranSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_TrimaranSpec_To_config_TrimaranSpec(a.(*TrimaranSpec), b.(*config.TrimaranSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TrimaranSpec)(nil), (*TrimaranSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TrimaranSpec_To_v1_TrimaranSpec(a.(*config.TrimaranSpec), b.(*TrimaranSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.NodeResourceTopologyMatchArgs)(nil), (*NodeResourceTopologyMatchArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeResourceTopologyMatchArgs_To_v1_NodeResourceTopologyMatchArgs(a.(*config.NodeResourceTopologyMatchArgs), b.(*NodeResourceTopologyMatchArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*NodeResourceTopologyMatchArgs)(nil), (*config.NodeResourceTopologyMatchArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NodeResourceTopologyMatchArgs_To_config_NodeResourceTopologyMatchArgs(a.(*NodeResourceTopologyMatchArgs), b.(*config.NodeResourceTopologyMatchArgs), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_ClusterTopologyPlacementGroupArgs_To_config_ClusterTopologyPlacementGroupArgs(in *ClusterTopologyPlacementGroupArgs, out *config.ClusterTopologyPlacementGroupArgs, s conversion.Scope) error {
	if err := metav1.Convert_Pointer_string_To_string(&in.TopologyConfigMapNameSpace, &out.TopologyConfigMapNameSpace, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ClusterTopologyPlacementGroupArgs_To_config_ClusterTopologyPlacementGroupArgs is an autogenerated conversion function.
func Convert_v1_ClusterTopologyPlacementGroupArgs_To_config_ClusterTopologyPlacementGroupArgs(in *ClusterTopologyPlacementGroupArgs, out *config.ClusterTopologyPlacementGroupArgs, s conversion.Scope) error {
	return autoConvert_v1_ClusterTopologyPlacementGroupArgs_To_config_ClusterTopologyPlacementGroupArgs(in, out, s)
}

func autoConvert_config_ClusterTopologyPlacementGroupArgs_To_v1_ClusterTopologyPlacementGroupArgs(in *config.ClusterTopologyPlacementGroupArgs, out *ClusterTopologyPlacementGroupArgs, s conversion.Scope) error {
	if err := metav1.Convert_string_To_Pointer_string(&in.TopologyConfigMapNameSpace, &out.TopologyConfigMapNameSpace, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_ClusterTopologyPlacementGroupArgs_To_v1_ClusterTopologyPlacementGroupArgs is an autogenerated conversion function.
func Convert_config_ClusterTopologyPlacementGroupArgs_To_v1_ClusterTopologyPlacementGroupArgs(in *config.ClusterTopologyPlacementGroupArgs, out *ClusterTopologyPlacementGroupArgs, s conversion.Scope) error {
	return autoConvert_config_ClusterTopologyPlacementGroupArgs_To_v1_ClusterTopologyPlacementGroupArgs(in, out, s)
}

func autoConvert_v1_CoschedulingArgs_To_config_CoschedulingArgs(in *CoschedulingArgs, out *config.CoschedulingArgs, s conversion.Scope) error {
	if err := metav1.Convert_Pointer_int64_To_int64(&in.PermitWaitingTimeSeconds, &out.PermitWaitingTimeSeconds, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_int64_To_int64(&in.PodGroupBackoffSeconds, &out.PodGroupBackoffSeconds, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_CoschedulingArgs_To_config_CoschedulingArgs is an autogenerated conversion function.
func Convert_v1_CoschedulingArgs_To_config_CoschedulingArgs(in *CoschedulingArgs, out *config.CoschedulingArgs, s conversion.Scope) error {
	return autoConvert_v1_CoschedulingArgs_To_config_CoschedulingArgs(in, out, s)
}

func autoConvert_config_CoschedulingArgs_To_v1_CoschedulingArgs(in *config.CoschedulingArgs, out *CoschedulingArgs, s conversion.Scope) error {
	if err := metav1.Convert_int64_To_Pointer_int64(&in.PermitWaitingTimeSeconds, &out.PermitWaitingTimeSeconds, s); err != nil {
		return err
	}
	if err := metav1.Convert_int64_To_Pointer_int64(&in.PodGroupBackoffSeconds, &out.PodGroupBackoffSeconds, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_CoschedulingArgs_To_v1_CoschedulingArgs is an autogenerated conversion function.
func Convert_config_CoschedulingArgs_To_v1_CoschedulingArgs(in *config.CoschedulingArgs, out *CoschedulingArgs, s conversion.Scope) error {
	return autoConvert_config_CoschedulingArgs_To_v1_CoschedulingArgs(in, out, s)
}

func autoConvert_v1_LoadVariationRiskBalancingArgs_To_config_LoadVariationRiskBalancingArgs(in *LoadVariationRiskBalancingArgs, out *config.LoadVariationRiskBalancingArgs, s conversion.Scope) error {
	if err := Convert_v1_TrimaranSpec_To_config_TrimaranSpec(&in.TrimaranSpec, &out.TrimaranSpec, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_float64_To_float64(&in.SafeVarianceMargin, &out.SafeVarianceMargin, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_float64_To_float64(&in.SafeVarianceSensitivity, &out.SafeVarianceSensitivity, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_LoadVariationRiskBalancingArgs_To_config_LoadVariationRiskBalancingArgs is an autogenerated conversion function.
func Convert_v1_LoadVariationRiskBalancingArgs_To_config_LoadVariationRiskBalancingArgs(in *LoadVariationRiskBalancingArgs, out *config.LoadVariationRiskBalancingArgs, s conversion.Scope) error {
	return autoConvert_v1_LoadVariationRiskBalancingArgs_To_config_LoadVariationRiskBalancingArgs(in, out, s)
}

func autoConvert_config_LoadVariationRiskBalancingArgs_To_v1_LoadVariationRiskBalancingArgs(in *config.LoadVariationRiskBalancingArgs, out *LoadVariationRiskBalancingArgs, s conversion.Scope) error {
	if err := Convert_config_TrimaranSpec_To_v1_TrimaranSpec(&in.TrimaranSpec, &out.TrimaranSpec, s); err != nil {
		return err
	}
	if err := metav1.Convert_float64_To_Pointer_float64(&in.SafeVarianceMargin, &out.SafeVarianceMargin, s); err != nil {
		return err
	}
	if err := metav1.Convert_float64_To_Pointer_float64(&in.SafeVarianceSensitivity, &out.SafeVarianceSensitivity, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_LoadVariationRiskBalancingArgs_To_v1_LoadVariationRiskBalancingArgs is an autogenerated conversion function.
func Convert_config_LoadVariationRiskBalancingArgs_To_v1_LoadVariationRiskBalancingArgs(in *config.LoadVariationRiskBalancingArgs, out *LoadVariationRiskBalancingArgs, s conversion.Scope) error {
	return autoConvert_config_LoadVariationRiskBalancingArgs_To_v1_LoadVariationRiskBalancingArgs(in, out, s)
}

func autoConvert_v1_LowRiskOverCommitmentArgs_To_config_LowRiskOverCommitmentArgs(in *LowRiskOverCommitmentArgs, out *config.LowRiskOverCommitmentArgs, s conversion.Scope) error {
	if err := Convert_v1_TrimaranSpec_To_config_TrimaranSpec(&in.TrimaranSpec, &out.TrimaranSpec, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_int64_To_int64(&in.SmoothingWindowSize, &out.SmoothingWindowSize, s); err != nil {
		return err
	}
	out.RiskLimitWeights = *(*map[corev1.ResourceName]float64)(unsafe.Pointer(&in.RiskLimitWeights))
	return nil
}

// Convert_v1_LowRiskOverCommitmentArgs_To_config_LowRiskOverCommitmentArgs is an autogenerated conversion function.
func Convert_v1_LowRiskOverCommitmentArgs_To_config_LowRiskOverCommitmentArgs(in *LowRiskOverCommitmentArgs, out *config.LowRiskOverCommitmentArgs, s conversion.Scope) error {
	return autoConvert_v1_LowRiskOverCommitmentArgs_To_config_LowRiskOverCommitmentArgs(in, out, s)
}

func autoConvert_config_LowRiskOverCommitmentArgs_To_v1_LowRiskOverCommitmentArgs(in *config.LowRiskOverCommitmentArgs, out *LowRiskOverCommitmentArgs, s conversion.Scope) error {
	if err := Convert_config_TrimaranSpec_To_v1_TrimaranSpec(&in.TrimaranSpec, &out.TrimaranSpec, s); err != nil {
		return err
	}
	if err := metav1.Convert_int64_To_Pointer_int64(&in.SmoothingWindowSize, &out.SmoothingWindowSize, s); err != nil {
		return err
	}
	out.RiskLimitWeights = *(*map[corev1.ResourceName]float64)(unsafe.Pointer(&in.RiskLimitWeights))
	return nil
}

// Convert_config_LowRiskOverCommitmentArgs_To_v1_LowRiskOverCommitmentArgs is an autogenerated conversion function.
func Convert_config_LowRiskOverCommitmentArgs_To_v1_LowRiskOverCommitmentArgs(in *config.LowRiskOverCommitmentArgs, out *LowRiskOverCommitmentArgs, s conversion.Scope) error {
	return autoConvert_config_LowRiskOverCommitmentArgs_To_v1_LowRiskOverCommitmentArgs(in, out, s)
}

func autoConvert_v1_MetricProviderSpec_To_config_MetricProviderSpec(in *MetricProviderSpec, out *config.MetricProviderSpec, s conversion.Scope) error {
	out.Type = config.MetricProviderType(in.Type)
	if err := metav1.Convert_Pointer_string_To_string(&in.Address, &out.Address, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_string_To_string(&in.Token, &out.Token, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_bool_To_bool(&in.InsecureSkipVerify, &out.InsecureSkipVerify, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_MetricProviderSpec_To_config_MetricProviderSpec is an autogenerated conversion function.
func Convert_v1_MetricProviderSpec_To_config_MetricProviderSpec(in *MetricProviderSpec, out *config.MetricProviderSpec, s conversion.Scope) error {
	return autoConvert_v1_MetricProviderSpec_To_config_MetricProviderSpec(in, out, s)
}

func autoConvert_config_MetricProviderSpec_To_v1_MetricProviderSpec(in *config.MetricProviderSpec, out *MetricProviderSpec, s conversion.Scope) error {
	out.Type = MetricProviderType(in.Type)
	if err := metav1.Convert_string_To_Pointer_string(&in.Address, &out.Address, s); err != nil {
		return err
	}
	if err := metav1.Convert_string_To_Pointer_string(&in.Token, &out.Token, s); err != nil {
		return err
	}
	if err := metav1.Convert_bool_To_Pointer_bool(&in.InsecureSkipVerify, &out.InsecureSkipVerify, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_MetricProviderSpec_To_v1_MetricProviderSpec is an autogenerated conversion function.
func Convert_config_MetricProviderSpec_To_v1_MetricProviderSpec(in *config.MetricProviderSpec, out *MetricProviderSpec, s conversion.Scope) error {
	return autoConvert_config_MetricProviderSpec_To_v1_MetricProviderSpec(in, out, s)
}

func autoConvert_v1_NetworkOverheadArgs_To_config_NetworkOverheadArgs(in *NetworkOverheadArgs, out *config.NetworkOverheadArgs, s conversion.Scope) error {
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	if err := metav1.Convert_Pointer_string_To_string(&in.WeightsName, &out.WeightsName, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_string_To_string(&in.NetworkTopologyName, &out.NetworkTopologyName, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_NetworkOverheadArgs_To_config_NetworkOverheadArgs is an autogenerated conversion function.
func Convert_v1_NetworkOverheadArgs_To_config_NetworkOverheadArgs(in *NetworkOverheadArgs, out *config.NetworkOverheadArgs, s conversion.Scope) error {
	return autoConvert_v1_NetworkOverheadArgs_To_config_NetworkOverheadArgs(in, out, s)
}

func autoConvert_config_NetworkOverheadArgs_To_v1_NetworkOverheadArgs(in *config.NetworkOverheadArgs, out *NetworkOverheadArgs, s conversion.Scope) error {
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	if err := metav1.Convert_string_To_Pointer_string(&in.WeightsName, &out.WeightsName, s); err != nil {
		return err
	}
	if err := metav1.Convert_string_To_Pointer_string(&in.NetworkTopologyName, &out.NetworkTopologyName, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_NetworkOverheadArgs_To_v1_NetworkOverheadArgs is an autogenerated conversion function.
func Convert_config_NetworkOverheadArgs_To_v1_NetworkOverheadArgs(in *config.NetworkOverheadArgs, out *NetworkOverheadArgs, s conversion.Scope) error {
	return autoConvert_config_NetworkOverheadArgs_To_v1_NetworkOverheadArgs(in, out, s)
}

func autoConvert_v1_NodeResourceTopologyCache_To_config_NodeResourceTopologyCache(in *NodeResourceTopologyCache, out *config.NodeResourceTopologyCache, s conversion.Scope) error {
	out.ForeignPodsDetect = (*config.ForeignPodsDetectMode)(unsafe.Pointer(in.ForeignPodsDetect))
	out.ResyncMethod = (*config.CacheResyncMethod)(unsafe.Pointer(in.ResyncMethod))
	out.InformerMode = (*config.CacheInformerMode)(unsafe.Pointer(in.InformerMode))
	out.ResyncScope = (*config.CacheResyncScope)(unsafe.Pointer(in.ResyncScope))
	return nil
}

// Convert_v1_NodeResourceTopologyCache_To_config_NodeResourceTopologyCache is an autogenerated conversion function.
func Convert_v1_NodeResourceTopologyCache_To_config_NodeResourceTopologyCache(in *NodeResourceTopologyCache, out *config.NodeResourceTopologyCache, s conversion.Scope) error {
	return autoConvert_v1_NodeResourceTopologyCache_To_config_NodeResourceTopologyCache(in, out, s)
}

func autoConvert_config_NodeResourceTopologyCache_To_v1_NodeResourceTopologyCache(in *config.NodeResourceTopologyCache, out *NodeResourceTopologyCache, s conversion.Scope) error {
	out.ForeignPodsDetect = (*ForeignPodsDetectMode)(unsafe.Pointer(in.ForeignPodsDetect))
	out.ResyncMethod = (*CacheResyncMethod)(unsafe.Pointer(in.ResyncMethod))
	out.InformerMode = (*CacheInformerMode)(unsafe.Pointer(in.InformerMode))
	out.ResyncScope = (*CacheResyncScope)(unsafe.Pointer(in.ResyncScope))
	return nil
}

// Convert_config_NodeResourceTopologyCache_To_v1_NodeResourceTopologyCache is an autogenerated conversion function.
func Convert_config_NodeResourceTopologyCache_To_v1_NodeResourceTopologyCache(in *config.NodeResourceTopologyCache, out *NodeResourceTopologyCache, s conversion.Scope) error {
	return autoConvert_config_NodeResourceTopologyCache_To_v1_NodeResourceTopologyCache(in, out, s)
}

func autoConvert_v1_NodeResourceTopologyMatchArgs_To_config_NodeResourceTopologyMatchArgs(in *NodeResourceTopologyMatchArgs, out *config.NodeResourceTopologyMatchArgs, s conversion.Scope) error {
	// WARNING: in.ScoringStrategy requires manual conversion: inconvertible types (*sigs.k8s.io/scheduler-plugins/apis/config/v1.ScoringStrategy vs sigs.k8s.io/scheduler-plugins/apis/config.ScoringStrategy)
	if err := metav1.Convert_Pointer_int64_To_int64(&in.CacheResyncPeriodSeconds, &out.CacheResyncPeriodSeconds, s); err != nil {
		return err
	}
	out.DiscardReservedNodes = in.DiscardReservedNodes
	out.Cache = (*config.NodeResourceTopologyCache)(unsafe.Pointer(in.Cache))
	return nil
}

func autoConvert_config_NodeResourceTopologyMatchArgs_To_v1_NodeResourceTopologyMatchArgs(in *config.NodeResourceTopologyMatchArgs, out *NodeResourceTopologyMatchArgs, s conversion.Scope) error {
	// WARNING: in.ScoringStrategy requires manual conversion: inconvertible types (sigs.k8s.io/scheduler-plugins/apis/config.ScoringStrategy vs *sigs.k8s.io/scheduler-plugins/apis/config/v1.ScoringStrategy)
	if err := metav1.Convert_int64_To_Pointer_int64(&in.CacheResyncPeriodSeconds, &out.CacheResyncPeriodSeconds, s); err != nil {
		return err
	}
	out.DiscardReservedNodes = in.DiscardReservedNodes
	out.Cache = (*NodeResourceTopologyCache)(unsafe.Pointer(in.Cache))
	return nil
}

func autoConvert_v1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs(in *NodeResourcesAllocatableArgs, out *config.NodeResourcesAllocatableArgs, s conversion.Scope) error {
	out.Resources = *(*[]apisconfig.ResourceSpec)(unsafe.Pointer(&in.Resources))
	out.Mode = config.ModeType(in.Mode)
	return nil
}

// Convert_v1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs is an autogenerated conversion function.
func Convert_v1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs(in *NodeResourcesAllocatableArgs, out *config.NodeResourcesAllocatableArgs, s conversion.Scope) error {
	return autoConvert_v1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs(in, out, s)
}

func autoConvert_config_NodeResourcesAllocatableArgs_To_v1_NodeResourcesAllocatableArgs(in *config.NodeResourcesAllocatableArgs, out *NodeResourcesAllocatableArgs, s conversion.Scope) error {
	out.Resources = *(*[]configv1.ResourceSpec)(unsafe.Pointer(&in.Resources))
	out.Mode = ModeType(in.Mode)
	return nil
}

// Convert_config_NodeResourcesAllocatableArgs_To_v1_NodeResourcesAllocatableArgs is an autogenerated conversion function.
func Convert_config_NodeResourcesAllocatableArgs_To_v1_NodeResourcesAllocatableArgs(in *config.NodeResourcesAllocatableArgs, out *NodeResourcesAllocatableArgs, s conversion.Scope) error {
	return autoConvert_config_NodeResourcesAllocatableArgs_To_v1_NodeResourcesAllocatableArgs(in, out, s)
}

func autoConvert_v1_PeaksArgs_To_config_PeaksArgs(in *PeaksArgs, out *config.PeaksArgs, s conversion.Scope) error {
	out.WatcherAddress = in.WatcherAddress
	out.NodePowerModel = *(*map[string]config.PowerModel)(unsafe.Pointer(&in.NodePowerModel))
	return nil
}

// Convert_v1_PeaksArgs_To_config_PeaksArgs is an autogenerated conversion function.
func Convert_v1_PeaksArgs_To_config_PeaksArgs(in *PeaksArgs, out *config.PeaksArgs, s conversion.Scope) error {
	return autoConvert_v1_PeaksArgs_To_config_PeaksArgs(in, out, s)
}

func autoConvert_config_PeaksArgs_To_v1_PeaksArgs(in *config.PeaksArgs, out *PeaksArgs, s conversion.Scope) error {
	out.WatcherAddress = in.WatcherAddress
	out.NodePowerModel = *(*map[string]PowerModel)(unsafe.Pointer(&in.NodePowerModel))
	return nil
}

// Convert_config_PeaksArgs_To_v1_PeaksArgs is an autogenerated conversion function.
func Convert_config_PeaksArgs_To_v1_PeaksArgs(in *config.PeaksArgs, out *PeaksArgs, s conversion.Scope) error {
	return autoConvert_config_PeaksArgs_To_v1_PeaksArgs(in, out, s)
}

func autoConvert_v1_PowerModel_To_config_PowerModel(in *PowerModel, out *config.PowerModel, s conversion.Scope) error {
	out.K0 = in.K0
	out.K1 = in.K1
	out.K2 = in.K2
	return nil
}

// Convert_v1_PowerModel_To_config_PowerModel is an autogenerated conversion function.
func Convert_v1_PowerModel_To_config_PowerModel(in *PowerModel, out *config.PowerModel, s conversion.Scope) error {
	return autoConvert_v1_PowerModel_To_config_PowerModel(in, out, s)
}

func autoConvert_config_PowerModel_To_v1_PowerModel(in *config.PowerModel, out *PowerModel, s conversion.Scope) error {
	out.K0 = in.K0
	out.K1 = in.K1
	out.K2 = in.K2
	return nil
}

// Convert_config_PowerModel_To_v1_PowerModel is an autogenerated conversion function.
func Convert_config_PowerModel_To_v1_PowerModel(in *config.PowerModel, out *PowerModel, s conversion.Scope) error {
	return autoConvert_config_PowerModel_To_v1_PowerModel(in, out, s)
}

func autoConvert_v1_PreemptionTolerationArgs_To_config_PreemptionTolerationArgs(in *PreemptionTolerationArgs, out *config.PreemptionTolerationArgs, s conversion.Scope) error {
	if err := metav1.Convert_Pointer_int32_To_int32(&in.MinCandidateNodesPercentage, &out.MinCandidateNodesPercentage, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_int32_To_int32(&in.MinCandidateNodesAbsolute, &out.MinCandidateNodesAbsolute, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_PreemptionTolerationArgs_To_config_PreemptionTolerationArgs is an autogenerated conversion function.
func Convert_v1_PreemptionTolerationArgs_To_config_PreemptionTolerationArgs(in *PreemptionTolerationArgs, out *config.PreemptionTolerationArgs, s conversion.Scope) error {
	return autoConvert_v1_PreemptionTolerationArgs_To_config_PreemptionTolerationArgs(in, out, s)
}

func autoConvert_config_PreemptionTolerationArgs_To_v1_PreemptionTolerationArgs(in *config.PreemptionTolerationArgs, out *PreemptionTolerationArgs, s conversion.Scope) error {
	if err := metav1.Convert_int32_To_Pointer_int32(&in.MinCandidateNodesPercentage, &out.MinCandidateNodesPercentage, s); err != nil {
		return err
	}
	if err := metav1.Convert_int32_To_Pointer_int32(&in.MinCandidateNodesAbsolute, &out.MinCandidateNodesAbsolute, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_PreemptionTolerationArgs_To_v1_PreemptionTolerationArgs is an autogenerated conversion function.
func Convert_config_PreemptionTolerationArgs_To_v1_PreemptionTolerationArgs(in *config.PreemptionTolerationArgs, out *PreemptionTolerationArgs, s conversion.Scope) error {
	return autoConvert_config_PreemptionTolerationArgs_To_v1_PreemptionTolerationArgs(in, out, s)
}

func autoConvert_v1_ScoringStrategy_To_config_ScoringStrategy(in *ScoringStrategy, out *config.ScoringStrategy, s conversion.Scope) error {
	out.Type = config.ScoringStrategyType(in.Type)
	out.Resources = *(*[]apisconfig.ResourceSpec)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1_ScoringStrategy_To_config_ScoringStrategy is an autogenerated conversion function.
func Convert_v1_ScoringStrategy_To_config_ScoringStrategy(in *ScoringStrategy, out *config.ScoringStrategy, s conversion.Scope) error {
	return autoConvert_v1_ScoringStrategy_To_config_ScoringStrategy(in, out, s)
}

func autoConvert_config_ScoringStrategy_To_v1_ScoringStrategy(in *config.ScoringStrategy, out *ScoringStrategy, s conversion.Scope) error {
	out.Type = ScoringStrategyType(in.Type)
	out.Resources = *(*[]configv1.ResourceSpec)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_config_ScoringStrategy_To_v1_ScoringStrategy is an autogenerated conversion function.
func Convert_config_ScoringStrategy_To_v1_ScoringStrategy(in *config.ScoringStrategy, out *ScoringStrategy, s conversion.Scope) error {
	return autoConvert_config_ScoringStrategy_To_v1_ScoringStrategy(in, out, s)
}

func autoConvert_v1_SySchedArgs_To_config_SySchedArgs(in *SySchedArgs, out *config.SySchedArgs, s conversion.Scope) error {
	if err := metav1.Convert_Pointer_string_To_string(&in.DefaultProfileNamespace, &out.DefaultProfileNamespace, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_string_To_string(&in.DefaultProfileName, &out.DefaultProfileName, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_SySchedArgs_To_config_SySchedArgs is an autogenerated conversion function.
func Convert_v1_SySchedArgs_To_config_SySchedArgs(in *SySchedArgs, out *config.SySchedArgs, s conversion.Scope) error {
	return autoConvert_v1_SySchedArgs_To_config_SySchedArgs(in, out, s)
}

func autoConvert_config_SySchedArgs_To_v1_SySchedArgs(in *config.SySchedArgs, out *SySchedArgs, s conversion.Scope) error {
	if err := metav1.Convert_string_To_Pointer_string(&in.DefaultProfileNamespace, &out.DefaultProfileNamespace, s); err != nil {
		return err
	}
	if err := metav1.Convert_string_To_Pointer_string(&in.DefaultProfileName, &out.DefaultProfileName, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_SySchedArgs_To_v1_SySchedArgs is an autogenerated conversion function.
func Convert_config_SySchedArgs_To_v1_SySchedArgs(in *config.SySchedArgs, out *SySchedArgs, s conversion.Scope) error {
	return autoConvert_config_SySchedArgs_To_v1_SySchedArgs(in, out, s)
}

func autoConvert_v1_TargetLoadPackingArgs_To_config_TargetLoadPackingArgs(in *TargetLoadPackingArgs, out *config.TargetLoadPackingArgs, s conversion.Scope) error {
	if err := Convert_v1_TrimaranSpec_To_config_TrimaranSpec(&in.TrimaranSpec, &out.TrimaranSpec, s); err != nil {
		return err
	}
	out.DefaultRequests = *(*corev1.ResourceList)(unsafe.Pointer(&in.DefaultRequests))
	if err := metav1.Convert_Pointer_string_To_string(&in.DefaultRequestsMultiplier, &out.DefaultRequestsMultiplier, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_int64_To_int64(&in.TargetUtilization, &out.TargetUtilization, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_TargetLoadPackingArgs_To_config_TargetLoadPackingArgs is an autogenerated conversion function.
func Convert_v1_TargetLoadPackingArgs_To_config_TargetLoadPackingArgs(in *TargetLoadPackingArgs, out *config.TargetLoadPackingArgs, s conversion.Scope) error {
	return autoConvert_v1_TargetLoadPackingArgs_To_config_TargetLoadPackingArgs(in, out, s)
}

func autoConvert_config_TargetLoadPackingArgs_To_v1_TargetLoadPackingArgs(in *config.TargetLoadPackingArgs, out *TargetLoadPackingArgs, s conversion.Scope) error {
	if err := Convert_config_TrimaranSpec_To_v1_TrimaranSpec(&in.TrimaranSpec, &out.TrimaranSpec, s); err != nil {
		return err
	}
	out.DefaultRequests = *(*corev1.ResourceList)(unsafe.Pointer(&in.DefaultRequests))
	if err := metav1.Convert_string_To_Pointer_string(&in.DefaultRequestsMultiplier, &out.DefaultRequestsMultiplier, s); err != nil {
		return err
	}
	if err := metav1.Convert_int64_To_Pointer_int64(&in.TargetUtilization, &out.TargetUtilization, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_TargetLoadPackingArgs_To_v1_TargetLoadPackingArgs is an autogenerated conversion function.
func Convert_config_TargetLoadPackingArgs_To_v1_TargetLoadPackingArgs(in *config.TargetLoadPackingArgs, out *TargetLoadPackingArgs, s conversion.Scope) error {
	return autoConvert_config_TargetLoadPackingArgs_To_v1_TargetLoadPackingArgs(in, out, s)
}

func autoConvert_v1_TopologicalSortArgs_To_config_TopologicalSortArgs(in *TopologicalSortArgs, out *config.TopologicalSortArgs, s conversion.Scope) error {
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	return nil
}

// Convert_v1_TopologicalSortArgs_To_config_TopologicalSortArgs is an autogenerated conversion function.
func Convert_v1_TopologicalSortArgs_To_config_TopologicalSortArgs(in *TopologicalSortArgs, out *config.TopologicalSortArgs, s conversion.Scope) error {
	return autoConvert_v1_TopologicalSortArgs_To_config_TopologicalSortArgs(in, out, s)
}

func autoConvert_config_TopologicalSortArgs_To_v1_TopologicalSortArgs(in *config.TopologicalSortArgs, out *TopologicalSortArgs, s conversion.Scope) error {
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	return nil
}

// Convert_config_TopologicalSortArgs_To_v1_TopologicalSortArgs is an autogenerated conversion function.
func Convert_config_TopologicalSortArgs_To_v1_TopologicalSortArgs(in *config.TopologicalSortArgs, out *TopologicalSortArgs, s conversion.Scope) error {
	return autoConvert_config_TopologicalSortArgs_To_v1_TopologicalSortArgs(in, out, s)
}

func autoConvert_v1_TrimaranSpec_To_config_TrimaranSpec(in *TrimaranSpec, out *config.TrimaranSpec, s conversion.Scope) error {
	if err := Convert_v1_MetricProviderSpec_To_config_MetricProviderSpec(&in.MetricProvider, &out.MetricProvider, s); err != nil {
		return err
	}
	if err := metav1.Convert_Pointer_string_To_string(&in.WatcherAddress, &out.WatcherAddress, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_TrimaranSpec_To_config_TrimaranSpec is an autogenerated conversion function.
func Convert_v1_TrimaranSpec_To_config_TrimaranSpec(in *TrimaranSpec, out *config.TrimaranSpec, s conversion.Scope) error {
	return autoConvert_v1_TrimaranSpec_To_config_TrimaranSpec(in, out, s)
}

func autoConvert_config_TrimaranSpec_To_v1_TrimaranSpec(in *config.TrimaranSpec, out *TrimaranSpec, s conversion.Scope) error {
	if err := Convert_config_MetricProviderSpec_To_v1_MetricProviderSpec(&in.MetricProvider, &out.MetricProvider, s); err != nil {
		return err
	}
	if err := metav1.Convert_string_To_Pointer_string(&in.WatcherAddress, &out.WatcherAddress, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_TrimaranSpec_To_v1_TrimaranSpec is an autogenerated conversion function.
func Convert_config_TrimaranSpec_To_v1_TrimaranSpec(in *config.TrimaranSpec, out *TrimaranSpec, s conversion.Scope) error {
	return autoConvert_config_TrimaranSpec_To_v1_TrimaranSpec(in, out, s)
}
