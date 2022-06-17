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

	"github.com/crossplane/terrajet/pkg/config"

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
		r.References["settings.mysql_target.connection.on_premise.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.postgres_source.connection.on_premise.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["settings.postgres_target.connection.on_premise.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
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
	})
}