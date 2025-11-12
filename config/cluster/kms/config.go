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

package kms

import (
	"fmt"

	"github.com/crossplane/upjet/v2/pkg/config"

	"github.com/yandex-cloud/crossplane-provider-yc/config/cluster/iam"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/yandex-cloud/crossplane-provider-yc/apis/cluster/kms/v1alpha1"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_kms_symmetric_key_iam_binding", func(r *config.Resource) {
		r.References["symmetric_key_id"] = config.Reference{
			Type: "SymmetricKey",
		}
		r.References["members"] = config.Reference{
			Type:              "ServiceAccount",
			Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})
}
