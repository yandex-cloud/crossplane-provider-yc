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

package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"
	"github.com/yandex-cloud/crossplane-provider-yc/config/nlb"
	"github.com/yandex-cloud/crossplane-provider-yc/config/organizationmanager"

	tjconfig "github.com/crossplane/upjet/pkg/config"
	sdk "github.com/yandex-cloud/terraform-provider-yandex/yandex"
	framework "github.com/yandex-cloud/terraform-provider-yandex/yandex-framework/provider"

	"github.com/yandex-cloud/crossplane-provider-yc/config/alb"
	"github.com/yandex-cloud/crossplane-provider-yc/config/cdn"
	"github.com/yandex-cloud/crossplane-provider-yc/config/common"
	"github.com/yandex-cloud/crossplane-provider-yc/config/compute"
	"github.com/yandex-cloud/crossplane-provider-yc/config/datatransfer"
	"github.com/yandex-cloud/crossplane-provider-yc/config/dns"
	"github.com/yandex-cloud/crossplane-provider-yc/config/iam"
	"github.com/yandex-cloud/crossplane-provider-yc/config/kubernetes"
	"github.com/yandex-cloud/crossplane-provider-yc/config/mdb"
	ymq "github.com/yandex-cloud/crossplane-provider-yc/config/message"
	"github.com/yandex-cloud/crossplane-provider-yc/config/resourcemanager"
	"github.com/yandex-cloud/crossplane-provider-yc/config/storage"
	"github.com/yandex-cloud/crossplane-provider-yc/config/vpc"
	"github.com/yandex-cloud/crossplane-provider-yc/config/ydb"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

const (
	resourcePrefix = "yandex-cloud"
	modulePath     = "github.com/yandex-cloud/crossplane-provider-yc"
)

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		tjconfig.WithShortName("yandex-cloud"),
		tjconfig.WithRootGroup("yandex-cloud.jet.crossplane.io"),
		tjconfig.WithFeaturesPackage("internal/features"),
		tjconfig.WithDefaultResourceOptions(common.DefaultResourceOverrides()),
		tjconfig.WithTerraformProvider(sdk.NewSDKProvider()),
		tjconfig.WithTerraformPluginSDKIncludeList([]string{
			"yandex_mdb_mongodb_cluster$",
			"yandex_mdb_redis_cluster$",
		}),
		tjconfig.WithTerraformPluginFrameworkProvider(framework.NewFrameworkProvider()),
		tjconfig.WithTerraformPluginFrameworkIncludeList([]string{
			"yandex_mdb_mongodb_database$",
			"yandex_mdb_mongodb_user$",
		}),
		tjconfig.WithIncludeList([]string{
			"yandex_vpc_gateway$",
			"yandex_vpc_route_table$",

			"yandex_lb_network_load_balancer$",
			"yandex_lb_target_group$",

			"yandex_resourcemanager_cloud_iam_binding$",
			"yandex_resourcemanager_cloud_iam_member$",
			"yandex_resourcemanager_cloud$",

			"yandex_kms_symmetric_key_iam_binding$",

			"yandex_organizationmanager_saml_federation_user_account$",
			"yandex_organizationmanager_saml_federation$",
			"yandex_organizationmanager_organization_iam_binding$",
			"yandex_organizationmanager_group_iam_member$",
			"yandex_organizationmanager_group$",

			"yandex_iam_service_account_api_key$",
			"yandex_iam_service_account_iam_binding$",

			"yandex_mdb_opensearch_cluster$",
			"yandex_cdn_resource$",
			"yandex_cdn_origin_group$",
			"yandex_mdb_postgresql_cluster$",
			"yandex_mdb_postgresql_database$",
			"yandex_mdb_postgresql_user$",
			"yandex_mdb_elasticsearch_cluster$",
			"yandex_alb_backend_group$",
			"yandex_alb_http_router$",
			"yandex_alb_load_balancer$",
			"yandex_alb_target_group$",
			"yandex_alb_virtual_host$",
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
			"yandex_mdb_kafka_cluster$",
			"yandex_mdb_kafka_connector$",
			"yandex_mdb_kafka_topic$",
			"yandex_mdb_kafka_user$",
			"yandex_mdb_mysql_cluster$",
			"yandex_mdb_mysql_database$",
			"yandex_mdb_mysql_user$",
			"yandex_resourcemanager_folder$",
			"yandex_resourcemanager_folder_iam_member$",
			"yandex_resourcemanager_folder_iam_binding$",
			"yandex_vpc_default_security_group$",
			"yandex_vpc_network$",
			"yandex_vpc_subnet$",
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
			"yandex_message_queue$",
		}),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		alb.Configure,
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
		ymq.Configure,
		cdn.Configure,
		organizationmanager.Configure,
		nlb.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
