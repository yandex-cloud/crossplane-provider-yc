/*
Copyright 2022 YANDEX LLC
This is modified version of the software, made by the Crossplane Authors
and available at: https://github.com/crossplane-contrib/provider-jet-template

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

// Package apis contains Kubernetes API groups for the Yandex Cloud provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	clusterapis "github.com/yandex-cloud/crossplane-provider-yc/apis/cluster"
	namespacedapis "github.com/yandex-cloud/crossplane-provider-yc/apis/namespaced"
)

func init() {
	// AddToSchemes may be used to add all resources defined in the project to a Scheme
	AddToSchemes = append(AddToSchemes, clusterapis.AddToScheme, namespacedapis.AddToScheme)
}

// AddToSchemes is a list of functions to add all resources to a scheme.
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all resources to the scheme.
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
