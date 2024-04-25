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

package resourcemanager

import (
	"fmt"

	"github.com/crossplane/upjet/pkg/config"

	"github.com/yandex-cloud/provider-jet-yc/config/iam"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/yandex-cloud/provider-jet-yc/apis/resourcemanager/v1alpha1"
)

// Configure adds configurations for resourcemanager group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_resourcemanager_folder_iam_member", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.References["member"] = config.Reference{
			Type:              "ServiceAccount",
			Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})
	p.AddResourceConfigurator("yandex_resourcemanager_folder_iam_binding", func(r *config.Resource) {
		r.ShortGroup = "iam"
		r.References["members"] = config.Reference{
			Type:              "ServiceAccount",
			Extractor:         fmt.Sprintf("%s.%s", iam.ConfigPath, iam.ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountsRef",
			SelectorFieldName: "ServiceAccountsSelector",
		}
	})
}
