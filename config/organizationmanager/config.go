package organizationmanager

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

import (
	"github.com/crossplane/upjet/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/yandex-cloud/crossplane-provider-yc/apis/organizationmanager/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/yandex-cloud/crossplane-provider-yc/config/organizationmanager"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_organizationmanager_group_membership", func(r *config.Resource) {
		r.References["group_id"] = config.Reference{
			Type: "Group",
		}
	})
	p.AddResourceConfigurator("yandex_organizationmanager_organization_iam_binding", func(r *config.Resource) {
		r.References["group_id"] = config.Reference{
			Type: "Group",
		}
	})

}
