// +build generate

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

// NOTE: See the below link for details on what is happening here.
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module

// Remove existing CRDs
//go:generate rm -rf ../package/crds

// Remove generated files
//go:generate bash -c "find . -iname 'zz_*' -delete"
//go:generate bash -c "find . -type d -empty -delete"
//go:generate bash -c "find ../internal/controller -iname 'zz_*' -delete"
//go:generate bash -c "find ../internal/controller -type d -empty -delete"

// Run Terrajet generator
//go:generate go run -tags generate ../cmd/generator/main.go ..

// Generate description for resources
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/kubernetes_cluster.html.markdown ./kubernetes/v1alpha1/zz_cluster_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/kubernetes_node_group.html.markdown ./kubernetes/v1alpha1/zz_nodegroup_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/compute_instance.html.markdown ./compute/v1alpha1/zz_instance_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/container_registry.html.markdown ./container/v1alpha1/zz_registry_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/container_repository.html.markdown ./container/v1alpha1/zz_repository_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/dns_zone.html.markdown ./dns/v1alpha1/zz_zone_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/dns_recordset.html.markdown ./dns/v1alpha1/zz_recordset_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/kms_symmetric_key.html.markdown ./kms/v1alpha1/zz_symmetrickey_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/mdb_postgresql_cluster.html.markdown ./mdb/v1alpha1/zz_postgresqlcluster_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/mdb_postgresql_user.html.markdown ./mdb/v1alpha1/zz_postgresqluser_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/mdb_postgresql_database.html.markdown ./mdb/v1alpha1/zz_postgresqldatabase_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/mdb_mysql_cluster.html.markdown ./mdb/v1alpha1/zz_mysqlcluster_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/mdb_mysql_user.html.markdown ./mdb/v1alpha1/zz_mysqluser_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/mdb_mysql_database.html.markdown ./mdb/v1alpha1/zz_mysqldatabase_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/mdb_mongodb_cluster.html.markdown ./mdb/v1alpha1/zz_mongodbcluster_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/mdb_redis_cluster.html.markdown ./mdb/v1alpha1/zz_rediscluster_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/storage_bucket.html.markdown ./storage/v1alpha1/zz_bucket_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/storage_object.html.markdown ./storage/v1alpha1/zz_object_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/vpc_network.html.markdown ./vpc/v1alpha1/zz_network_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/vpc_subnet.html.markdown ./vpc/v1alpha1/zz_subnet_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/vpc_security_group.html.markdown ./vpc/v1alpha1/zz_securitygroup_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/vpc_default_security_group.html.markdown ./vpc/v1alpha1/zz_defaultsecuritygroup_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/vpc_security_group_rule.html.markdown ./vpc/v1alpha1/zz_securitygrouprule_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/resourcemanager_folder.html.markdown ./resourcemanager/v1alpha1/zz_folder_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/resourcemanager_folder.html.markdown ./resourcemanager/v1alpha1/zz_folder_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/resourcemanager_folder_iam_binding.html.markdown ./iam/v1alpha1/zz_folderiambinding_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/resourcemanager_folder_iam_member.html.markdown ./iam/v1alpha1/zz_folderiammember_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/iam_service_account.html.markdown ./iam/v1alpha1/zz_serviceaccount_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/iam_service_account_iam_member.html.markdown ./iam/v1alpha1/zz_serviceaccountiammember_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/iam_service_account_key.html.markdown ./iam/v1alpha1/zz_serviceaccountkey_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/iam_service_account_static_access_key.html.markdown ./iam/v1alpha1/zz_serviceaccountstaticaccesskey_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/datatransfer_endpoint.html.markdown ./datatransfer/v1alpha1/zz_endpoint_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/datatransfer_transfer.html.markdown ./datatransfer/v1alpha1/zz_transfer_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/ydb_database_dedicated.html.markdown ./ydb/v1alpha1/zz_databasededicated_types.go
//go:generate ../terraform-docs-parser.py ../terraform-provider-yandex/website/docs/r/ydb_database_serverless.html.markdown ./ydb/v1alpha1/zz_databaseserverless_types.go

// Generate deepcopy methodsets and CRD manifests
//go:generate go run -tags generate sigs.k8s.io/controller-tools/cmd/controller-gen object:headerFile=../hack/boilerplate.go.txt paths=./... crd:allowDangerousTypes=true,crdVersions=v1 output:artifacts:config=../package/crds

// Generate crossplane-runtime methodsets (resource.Claim, etc)
//go:generate go run -tags generate github.com/crossplane/crossplane-tools/cmd/angryjet generate-methodsets --header-file=../hack/boilerplate.go.txt ./...

// Generate Composite Resources
//go:generate go run -tags generate github.com/benagricola/crossplane-composition-generator
package apis

import (
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen" //nolint:typecheck

	_ "github.com/crossplane/crossplane-tools/cmd/angryjet" //nolint:typecheck

	_ "github.com/benagricola/crossplane-composition-generator" //nolint:typecheck
)
