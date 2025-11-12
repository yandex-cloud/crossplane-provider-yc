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

	tjconfig "github.com/crossplane/upjet/v2/pkg/config"
	sdk "github.com/yandex-cloud/terraform-provider-yandex/yandex"
	framework "github.com/yandex-cloud/terraform-provider-yandex/yandex-framework/provider"

	"github.com/yandex-cloud/crossplane-provider-yc/config/common"

	// Cluster-scoped configs
	albCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/alb"
	cdnCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/cdn"
	computeCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/compute"
	datatransferCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/datatransfer"
	dnsCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/dns"
	iamCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/iam"
	kubernetesCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/kubernetes"
	mdbCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/mdb"
	ymqCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/message"
	nlbCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/nlb"
	organizationmanagerCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/organizationmanager"
	resourcemanagerCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/resourcemanager"
	storageCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/storage"
	vpcCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/vpc"
	ydbCluster "github.com/yandex-cloud/crossplane-provider-yc/config/cluster/ydb"

	// Namespaced configs
	albNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/alb"
	cdnNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/cdn"
	computeNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/compute"
	datatransferNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/datatransfer"
	dnsNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/dns"
	iamNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/iam"
	kubernetesNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/kubernetes"
	mdbNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/mdb"
	ymqNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/message"
	nlbNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/nlb"
	organizationmanagerNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/organizationmanager"
	resourcemanagerNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/resourcemanager"
	storageNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/storage"
	vpcNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/vpc"
	ydbNamespaced "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/ydb"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

const (
	resourcePrefix = "yandex-cloud"
	modulePath     = "github.com/yandex-cloud/crossplane-provider-yc"
)

var (
	tfFrameworkIncludeList = []string{
		"yandex_mdb_mongodb_database$",
		"yandex_mdb_mongodb_user$",
		// "yandex_mdb_opensearch_cluster$",

		"yandex_vpc_security_group_rule$",
	}

	tfSDKIncludeList = []string{
		"yandex_mdb_mongodb_cluster$",
		"yandex_mdb_redis_cluster$",
		"yandex_mdb_kafka_cluster$",
		"yandex_mdb_kafka_connector$",

		"yandex_storage_object$",
	}

	tfForkingIncludeList = []string{
		"yandex_alb_backend_group$",
		"yandex_alb_http_router$",
		"yandex_alb_load_balancer$",
		"yandex_alb_target_group$",
		"yandex_alb_virtual_host$",

		"yandex_cdn_origin_group$",
		// CDN Resource disabled pending the fix of CDN API
		// "yandex_cdn_resource$",

		"yandex_compute_disk$",
		"yandex_compute_disk_placement_group$",
		"yandex_compute_filesystem$",
		"yandex_compute_gpu_cluster$",
		"yandex_compute_image$",
		"yandex_compute_instance$",
		"yandex_compute_instance_group$",
		"yandex_compute_placement_group$",
		"yandex_compute_snapshot$",
		"yandex_compute_snapshot_schedule$",

		"yandex_container_repository$",
		"yandex_container_registry$",

		"yandex_datatransfer_endpoint$",
		"yandex_datatransfer_transfer$",

		"yandex_dns_zone$",
		"yandex_dns_recordset$",

		"yandex_iam_service_account$",
		"yandex_iam_service_account_api_key$",
		"yandex_iam_service_account_iam_binding$",
		"yandex_iam_service_account_iam_member$",
		"yandex_iam_service_account_key$",
		"yandex_iam_service_account_static_access_key$",

		"yandex_kms_symmetric_key$",
		"yandex_kms_symmetric_key_iam_binding$",

		"yandex_kubernetes_cluster$",
		"yandex_kubernetes_node_group$",

		"yandex_lb_network_load_balancer$",
		"yandex_lb_target_group$",

		"yandex_mdb_clickhouse_cluster$",
		"yandex_mdb_kafka_topic$",
		"yandex_mdb_kafka_user$",
		"yandex_mdb_elasticsearch_cluster$",
		"yandex_mdb_mysql_cluster$",
		"yandex_mdb_mysql_database$",
		"yandex_mdb_mysql_user$",
		"yandex_mdb_postgresql_cluster$",
		"yandex_mdb_postgresql_database$",
		"yandex_mdb_postgresql_user$",

		"yandex_message_queue$",

		"yandex_organizationmanager_saml_federation_user_account$",
		"yandex_organizationmanager_saml_federation$",
		"yandex_organizationmanager_organization_iam_binding$",
		"yandex_organizationmanager_group_iam_member$",
		"yandex_organizationmanager_group$",

		"yandex_resourcemanager_cloud_iam_binding$",
		"yandex_resourcemanager_cloud_iam_member$",
		"yandex_resourcemanager_cloud$",
		"yandex_resourcemanager_folder$",
		"yandex_resourcemanager_folder_iam_member$",
		"yandex_resourcemanager_folder_iam_binding$",

		"yandex_vpc_address$",
		"yandex_vpc_default_security_group$",
		"yandex_vpc_gateway$",
		"yandex_vpc_network$",
		"yandex_vpc_route_table$",
		"yandex_vpc_security_group$",
		"yandex_vpc_subnet$",

		"yandex_storage_bucket$",

		"yandex_ydb_database_dedicated$",
		"yandex_ydb_database_serverless$",
	}
)

// GetProvider returns provider configuration.
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		tjconfig.WithShortName("yandex-cloud"),
		tjconfig.WithRootGroup("yandex-cloud.jet.crossplane.io"),
		tjconfig.WithFeaturesPackage("internal/features"),
		tjconfig.WithDefaultResourceOptions(common.DefaultResourceOverrides(resourcemanagerCluster.FolderType)),
		tjconfig.WithTerraformProvider(sdk.NewSDKProvider()),
		tjconfig.WithTerraformPluginSDKIncludeList(tfSDKIncludeList),
		tjconfig.WithTerraformPluginFrameworkProvider(framework.NewFrameworkProvider()),
		tjconfig.WithTerraformPluginFrameworkIncludeList(tfFrameworkIncludeList),
		tjconfig.WithIncludeList(tfForkingIncludeList),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		albCluster.Configure,
		computeCluster.Configure,
		datatransferCluster.Configure,
		dnsCluster.Configure,
		iamCluster.Configure,
		kubernetesCluster.Configure,
		mdbCluster.Configure,
		storageCluster.Configure,
		vpcCluster.Configure,
		resourcemanagerCluster.Configure,
		ydbCluster.Configure,
		ymqCluster.Configure,
		cdnCluster.Configure,
		organizationmanagerCluster.Configure,
		nlbCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns provider configuration for namespaced resources.
func GetProviderNamespaced() *tjconfig.Provider {
	pc := tjconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		tjconfig.WithShortName("yandex-cloud"),
		tjconfig.WithRootGroup("yandex-cloud.m.jet.crossplane.io"),
		tjconfig.WithFeaturesPackage("internal/features"),
		tjconfig.WithDefaultResourceOptions(common.DefaultResourceOverrides(resourcemanagerNamespaced.FolderType)),
		tjconfig.WithTerraformProvider(sdk.NewSDKProvider()),
		tjconfig.WithTerraformPluginSDKIncludeList(tfSDKIncludeList),
		tjconfig.WithTerraformPluginFrameworkProvider(framework.NewFrameworkProvider()),
		tjconfig.WithTerraformPluginFrameworkIncludeList(tfFrameworkIncludeList),
		tjconfig.WithIncludeList(tfForkingIncludeList),
	)

	for _, configure := range []func(provider *tjconfig.Provider){
		albNamespaced.Configure,
		computeNamespaced.Configure,
		datatransferNamespaced.Configure,
		dnsNamespaced.Configure,
		iamNamespaced.Configure,
		kubernetesNamespaced.Configure,
		mdbNamespaced.Configure,
		storageNamespaced.Configure,
		vpcNamespaced.Configure,
		resourcemanagerNamespaced.Configure,
		ydbNamespaced.Configure,
		ymqNamespaced.Configure,
		cdnNamespaced.Configure,
		organizationmanagerNamespaced.Configure,
		nlbNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
