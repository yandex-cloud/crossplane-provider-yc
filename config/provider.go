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

package config

import (
	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"bb.yandex-team.ru/crossplane/provider-jet-yc/config/compute"
	"bb.yandex-team.ru/crossplane/provider-jet-yc/config/container"
	"bb.yandex-team.ru/crossplane/provider-jet-yc/config/dns"
	"bb.yandex-team.ru/crossplane/provider-jet-yc/config/iam"
	"bb.yandex-team.ru/crossplane/provider-jet-yc/config/kms"
	"bb.yandex-team.ru/crossplane/provider-jet-yc/config/kubernetes"
	"bb.yandex-team.ru/crossplane/provider-jet-yc/config/mdb"
	"bb.yandex-team.ru/crossplane/provider-jet-yc/config/storage"
	"bb.yandex-team.ru/crossplane/provider-jet-yc/config/vpc"
)

const (
	resourcePrefix = "yandex-cloud"
	modulePath     = "bb.yandex-team.ru/crossplane/provider-jet-yc"
)

// GetProvider returns provider configuration
func GetProvider(tf *schema.Provider) *tjconfig.Provider {
	// Comment out the line below instead of the above, if your Terraform
	// provider uses an old version (<v2) of github.com/hashicorp/terraform-plugin-sdk.
	// resourceMap := conversion.GetV2ResourceMap(tf.Provider())

	defaultResourceFn := func(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
		r := tjconfig.DefaultResource(name, terraformResource)
		// Add any provider-specific defaulting here. For example:
		r.ExternalName = tjconfig.IdentifierFromProvider
		return r
	}

	pc := tjconfig.NewProvider(tf.ResourcesMap, resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(defaultResourceFn),
		//
		// Use this config to generate all terraform provider resources
		//
		// tjconfig.WithSkipList([]string{
		//   "yandex_iot_core_device$",
		// 	 "yandex_iot_core_registry",
		// }),
		tjconfig.WithIncludeList([]string{
			"yandex_vpc_network$",
			"yandex_vpc_subnet$",
			"yandex_compute_instance$",
			"yandex_container_registry$",
			"yandex_container_repository$",
			"yandex_dns_zone$",
			"yandex_dns_recordset$",
			"yandex_iam_service_account$",
			"yandex_iam_service_account_key$",
			"yandex_iam_service_account_static_access_key$",
			"yandex_iam_service_account_iam_member$",
			"yandex_kubernetes_cluster$",
			"yandex_kubernetes_node_group$",
			"yandex_mdb_mongodb_cluster$",
			"yandex_mdb_postgresql_cluster$",
			"yandex_mdb_redis_cluster$",
			"yandex_resourcemanager_folder$",
			"yandex_storage_bucket$",
			"yandex_storage_object$",
			"yandex_kms_symmetric_key$",
		}),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		compute.Configure,
		container.Configure,
		dns.Configure,
		iam.Configure,
		kubernetes.Configure,
		mdb.Configure,
		storage.Configure,
		vpc.Configure,
		kms.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
