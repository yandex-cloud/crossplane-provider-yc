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

package nlb

import (
	"fmt"

	xpref "github.com/crossplane/crossplane-runtime/v2/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/resource"

	"github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/compute"
	"github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/vpc"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath       = "github.com/yandex-cloud/crossplane-provider-yc/apis/namespaced/lb/v1alpha1"
	ConfigPath            = "github.com/yandex-cloud/crossplane-provider-yc/config/namespaced/nlb"
	ExtractComputeAddress = ConfigPath + ".ExtractComputeIP()"
)

// Configure adds configurations for vpc group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_lb_network_load_balancer", func(r *config.Resource) {
		r.References["attached_target_group.target_group_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", ApisPackagePath, "TargetGroup"),
		}
	})

	p.AddResourceConfigurator("yandex_lb_target_group", func(r *config.Resource) {
		r.References["target.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["target.address"] = config.Reference{
			Type:      fmt.Sprintf("%s.%s", compute.ApisPackagePath, "Instance"),
			Extractor: ExtractComputeAddress,
		}
	})
}

func ExtractComputeIP() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}

		o, err := tr.GetParameters()
		if err != nil {
			return ""
		}
		networkInterfaces, ok := o["network_interface"].([]any)
		if !ok || len(networkInterfaces) == 0 {
			return ""
		}

		firstNetworkInterface, ok := networkInterfaces[0].(map[string]any)
		if !ok {
			return ""
		}

		ipAddress, ok := firstNetworkInterface["ip_address"].(string)
		if !ok {
			return ""
		}

		return ipAddress
	}
}
