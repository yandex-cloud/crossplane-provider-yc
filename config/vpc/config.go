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

package vpc

import (
	"github.com/crossplane/terrajet/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/yandex-cloud/provider-jet-yc/apis/vpc/v1alpha1"
)

// Configure adds configurations for vpc group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_vpc_subnet", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: "Network",
		}
	})
	p.AddResourceConfigurator("yandex_vpc_default_security_group", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: "Network",
		}
	})
	p.AddResourceConfigurator("yandex_vpc_security_group", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: "Network",
		}
	})
	p.AddResourceConfigurator("yandex_vpc_security_group_rule", func(r *config.Resource) {
		r.References["security_group_id"] = config.Reference{
			Type: "SecurityGroup",
		}
	})
}
