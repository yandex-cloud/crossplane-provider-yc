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

package mdb

import (
	"fmt"

	"github.com/crossplane-contrib/terrajet/pkg/config"

	"bb.yandex-team.ru/crossplane/provider-jet-yc/config/resourcemanager"
	"bb.yandex-team.ru/crossplane/provider-jet-yc/config/vpc"
)

// Configure adds configurations for vpc group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_mdb_postgresql_cluster", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["folder_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", resourcemanager.ApisPackagePath, "Folder"),
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_mdb_redis_cluster", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["folder_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", resourcemanager.ApisPackagePath, "Folder"),
		}
		r.UseAsync = true
	})
}
