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

package compute

import (
	"fmt"

	"github.com/crossplane/upjet/pkg/config"

	"github.com/yandex-cloud/crossplane-provider-yc/config/iam"
	"github.com/yandex-cloud/crossplane-provider-yc/config/vpc"
)

// Configure adds configurations for compute group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_compute_instance", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["service_account_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccount"),
		}
		r.References["network_interface.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["network_interface.security_group_ids"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.UseAsync = true
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			data := make(map[string][]byte)
			if v, ok := attr["fqdn"].(string); ok && v != "" {
				data["fqdn"] = []byte(v)
			}
			if networkInterfaces, ok := attr["network_interface"].([]interface{}); ok {
				if len(networkInterfaces) > 0 {
					if networkInterface, ok := networkInterfaces[0].(map[string]interface{}); ok {
						if v, ok := networkInterface["ip_address"].(string); ok && v != "" {
							data["internal_ip"] = []byte(v)
						}
						if v, ok := networkInterface["nat_ip_address"].(string); ok && v != "" {
							data["external_ip"] = []byte(v)
						}
					}
				}
			}
			return data, nil
		}
	})
}
