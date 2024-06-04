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

package kubernetes

import (
	"fmt"

	"github.com/crossplane/upjet/pkg/config"

	"github.com/yandex-cloud/crossplane-provider-yc/config/iam"
	"github.com/yandex-cloud/crossplane-provider-yc/config/kms"
	"github.com/yandex-cloud/crossplane-provider-yc/config/vpc"
)

// Configure adds configurations for kubernetes group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_kubernetes_cluster", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["service_account_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
		}
		r.References["node_service_account_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
		}
		r.References["master.regional.location.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["master.zonal.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["master.security_group_ids"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["kms_provider.key_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", kms.ApisPackagePath, "SymmetricKey"),
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_kubernetes_node_group", func(r *config.Resource) {
		r.References["cluster_id"] = config.Reference{
			Type: "Cluster",
		}
		r.References["allocation_policy.location.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["instance_template.network_interface.subnet_ids"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["instance_template.network_interface.security_group_ids"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
	})
}
