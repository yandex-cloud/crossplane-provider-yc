/*
Copyright 2021 The Crossplane Authors.
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

package iam

import (
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/terrajet/pkg/config"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/yandex-cloud/provider-jet-yc/apis/iam/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/yandex-cloud/provider-jet-yc/config/iam"
	// ServiceAccountRefValueFn is the name of resolver.
	ServiceAccountRefValueFn = "ServiceAccountRefValue()"
)

// Configure adds configurations for iam group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_iam_service_account_key", func(r *config.Resource) {
		r.References["service_account_id"] = config.Reference{
			Type: "ServiceAccount",
		}
	})
	p.AddResourceConfigurator("yandex_iam_service_account_static_access_key", func(r *config.Resource) {
		r.References["service_account_id"] = config.Reference{
			Type: "ServiceAccount",
		}
	})
	p.AddResourceConfigurator("yandex_iam_service_account_iam_member", func(r *config.Resource) {
		r.References["service_account_id"] = config.Reference{
			Type: "ServiceAccount",
		}
	})
	p.AddResourceConfigurator("yandex_iam_service_account_iam_member", func(r *config.Resource) {
		r.References["member"] = config.Reference{
			Type:              "ServiceAccount",
			Extractor:         fmt.Sprintf("%s.%s", ConfigPath, ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})
}

// ServiceAccountRefValue returns an extractor that returns templated value with service account id of ServiceAccount.
func ServiceAccountRefValue() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		return func(mg metav1.Object) string {
			externalName := meta.GetExternalName(mg)
			return fmt.Sprintf("serviceAccount:%s", externalName)
		}(mg)
	}
}
