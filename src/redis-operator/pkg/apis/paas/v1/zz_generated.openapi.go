// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"redis-operator/pkg/apis/paas/v1.RedisCluster":       schema_pkg_apis_paas_v1_RedisCluster(ref),
		"redis-operator/pkg/apis/paas/v1.RedisClusterSpec":   schema_pkg_apis_paas_v1_RedisClusterSpec(ref),
		"redis-operator/pkg/apis/paas/v1.RedisClusterStatus": schema_pkg_apis_paas_v1_RedisClusterStatus(ref),
	}
}

func schema_pkg_apis_paas_v1_RedisCluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RedisCluster is the Schema for the redisclusters API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("redis-operator/pkg/apis/paas/v1.RedisClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("redis-operator/pkg/apis/paas/v1.RedisClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta", "redis-operator/pkg/apis/paas/v1.RedisClusterSpec", "redis-operator/pkg/apis/paas/v1.RedisClusterStatus"},
	}
}

func schema_pkg_apis_paas_v1_RedisClusterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RedisClusterSpec defines the desired state of RedisCluster",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_paas_v1_RedisClusterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "RedisClusterStatus defines the observed state of RedisCluster",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
