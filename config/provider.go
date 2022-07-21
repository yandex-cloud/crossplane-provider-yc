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
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"

	"github.com/yandex-cloud/provider-jet-yc/config/common"
	"github.com/yandex-cloud/provider-jet-yc/config/compute"
	"github.com/yandex-cloud/provider-jet-yc/config/datatransfer"
	"github.com/yandex-cloud/provider-jet-yc/config/dns"
	"github.com/yandex-cloud/provider-jet-yc/config/iam"
	"github.com/yandex-cloud/provider-jet-yc/config/kubernetes"
	"github.com/yandex-cloud/provider-jet-yc/config/mdb"
	"github.com/yandex-cloud/provider-jet-yc/config/resourcemanager"
	"github.com/yandex-cloud/provider-jet-yc/config/storage"
	"github.com/yandex-cloud/provider-jet-yc/config/vpc"
	"github.com/yandex-cloud/provider-jet-yc/config/ydb"
)

//go:embed schema.json
var providerSchema string

const (
	resourcePrefix = "yandex-cloud"
	modulePath     = "github.com/yandex-cloud/provider-jet-yc"
)

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithShortName("yandex-cloud"),
		tjconfig.WithRootGroup("yandex-cloud.jet.crossplane.io"),
		tjconfig.WithDefaultResourceFn(common.DefaultResourceFn),
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
			"yandex_mdb_postgresql_database$",
			"yandex_mdb_postgresql_user$",
			"yandex_mdb_mysql_cluster$",
			"yandex_mdb_mysql_database$",
			"yandex_mdb_mysql_user$",
			"yandex_mdb_redis_cluster$",
			"yandex_mdb_elasticsearch_cluster$",
			"yandex_resourcemanager_folder$",
			"yandex_resourcemanager_folder_iam_member$",
			"yandex_resourcemanager_folder_iam_binding$",
			"yandex_vpc_default_security_group$",
			"yandex_vpc_security_group$",
			"yandex_vpc_security_group_rule$",
			"yandex_vpc_address$",
			"yandex_storage_bucket$",
			"yandex_storage_object$",
			"yandex_kms_symmetric_key$",
			"yandex_datatransfer_endpoint$",
			"yandex_datatransfer_transfer$",
			"yandex_ydb_database_dedicated$",
			"yandex_ydb_database_serverless$",
		}),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		compute.Configure,
		datatransfer.Configure,
		dns.Configure,
		iam.Configure,
		kubernetes.Configure,
		mdb.Configure,
		storage.Configure,
		vpc.Configure,
		resourcemanager.Configure,
		ydb.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
