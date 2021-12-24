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

	"github.com/yandex-cloud/provider-jet-yc/config/vpc"
)

func usernames(attr map[string]interface{}) map[string]string {
	usersInterface, ok := attr["user"]
	if !ok {
		return nil
	}
	users, ok := usersInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make(map[string]string)
	for i, userInterface := range users {
		user, ok := userInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if username, ok := user["name"].(string); ok {
			result[fmt.Sprintf("attribute.user.%d.name", i)] = username
		}
	}
	return result
}

func databases(attr map[string]interface{}) map[string]string {
	databasesInterface, ok := attr["database"]
	if !ok {
		return nil
	}
	databases, ok := databasesInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make(map[string]string)
	for i, databaseInterface := range databases {
		database, ok := databaseInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if databaseName, ok := database["name"].(string); ok {
			result[fmt.Sprintf("attribute.database.%d.name", i)] = databaseName
		}
	}
	return result
}

func fqdns(attr map[string]interface{}) map[string]string {
	hostsInterface, ok := attr["host"]
	if !ok {
		return nil
	}
	hosts, ok := hostsInterface.([]interface{})
	if !ok {
		return nil
	}
	result := make(map[string]string)
	for i, hostInterface := range hosts {
		host, ok := hostInterface.(map[string]interface{})
		if !ok {
			continue
		}
		if fqdn, ok := host["fqdn"].(string); ok {
			result[fmt.Sprintf("attribute.host.%d.fqdn", i)] = fqdn
		}
	}
	return result
}

func connDetails(attr map[string]interface{}) (map[string][]byte, error) {
	conn := make(map[string][]byte)
	for k, v := range fqdns(attr) {
		conn[k] = []byte(v)
	}
	for k, v := range usernames(attr) {
		conn[k] = []byte(v)
	}
	for k, v := range databases(attr) {
		conn[k] = []byte(v)
	}
	return conn, nil
}

// Configure adds configurations for mdb group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_mdb_postgresql_cluster", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = connDetails
	})
	p.AddResourceConfigurator("yandex_mdb_redis_cluster", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = connDetails
	})
	p.AddResourceConfigurator("yandex_mdb_mongodb_cluster", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["host.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = connDetails
	})
}
