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

package datatransfer

import (
	"fmt"

	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/crossplane/terrajet/pkg/config"
	"github.com/crossplane/terrajet/pkg/resource"

	"github.com/yandex-cloud/provider-jet-yc/config/mdb"
	"github.com/yandex-cloud/provider-jet-yc/config/vpc"
)

// Configure adds configurations for datatransfer group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_datatransfer_endpoint", func(r *config.Resource) {
		r.References["settings.postgres_target.connection.mdb_cluster_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "PostgresqlCluster"),
		}
		r.References["settings.mysql_target.connection.mdb_cluster_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MySQLCluster"),
		}
		r.References["settings.mongo_target.connection.connection_options.mdb_cluster_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MongodbCluster"),
		}
		r.References["settings.postgres_source.connection.mdb_cluster_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "PostgresqlCluster"),
		}
		r.References["settings.mysql_source.connection.mdb_cluster_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MySQLCluster"),
		}
		r.References["settings.mongo_source.connection.connection_options.mdb_cluster_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MongodbCluster"),
		}
		r.References["settings.mysql_target.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.postgres_target.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.mongo_target.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.clickhouse_target.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.mysql_source.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.postgres_source.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.mongo_source.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.clickhouse_source.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.mysql_source.connection.on_premise.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.mysql_source.user"] = config.Reference{
			Type:      fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MySQLUser"),
			Extractor: ExtractUsernameFunc,
		}
		r.References["settings.mysql_source.database"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MySQLDatabase"),
		}
		r.References["settings.mysql_source.security_groups"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["settings.mysql_target.connection.on_premise.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.mysql_target.user"] = config.Reference{
			Type:      fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MySQLUser"),
			Extractor: ExtractUsernameFunc,
		}
		r.References["settings.mysql_target.database"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "MySQLDatabase"),
		}
		r.References["settings.mysql_target.security_groups"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["settings.postgres_source.connection.on_premise.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.postgres_source.user"] = config.Reference{
			Type:      fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "PostgresqlUser"),
			Extractor: ExtractUsernameFunc,
		}
		r.References["settings.postgres_source.database"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "PostgresqlDatabase"),
		}
		r.References["settings.postgres_source.security_groups"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["settings.postgres_target.connection.on_premise.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.postgres_target.user"] = config.Reference{
			Type:      fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "PostgresqlUser"),
			Extractor: ExtractUsernameFunc,
		}
		r.References["settings.postgres_target.database"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", mdb.ApisPackagePath, "PostgresqlDatabase"),
		}
		r.References["settings.postgres_target.security_groups"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["settings.mongo_source.connection.connection_options.on_premise.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.clickhouse_source.connection.connection_options.on_premise.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
	})
	p.AddResourceConfigurator("yandex_datatransfer_transfer", func(r *config.Resource) {
		r.References["source_id"] = config.Reference{
			Type: "Endpoint",
		}
		r.References["target_id"] = config.Reference{
			Type: "Endpoint",
		}
		r.UseAsync = true
	})
}

const (
	// APISPackagePath is the package path for generated APIs root package
	APISPackagePath = "github.com/yandex-cloud/provider-jet-yc/config/datatransfer"

	// ExtractUsernameFunc extracts username from MySQLUser or PotgresqlUser resource
	ExtractUsernameFunc = APISPackagePath + ".ExtractUsername()"
)

// ExtractUsername extracts the value of `spec.forProvider.name`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractUsername() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetParameters()
		if err != nil {
			return ""
		}
		if k := o["name"]; k != nil {
			return k.(string)
		}
		return ""
	}
}
