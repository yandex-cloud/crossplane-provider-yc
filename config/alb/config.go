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

package alb

import (
	"fmt"

	"github.com/crossplane/upjet/pkg/config"

	"github.com/yandex-cloud/crossplane-provider-yc/config/vpc"
)

// Configure adds configurations for alb group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_alb_load_balancer", func(r *config.Resource) {
		r.References["network_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Network"),
		}
		r.References["security_group_ids"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "SecurityGroup"),
		}
		r.References["allocation_policy.location.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["listener.endpoint.address.internal_ipv4_address.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}
		r.References["listener.http.handler.http_router_id"] = config.Reference{
			Type: "HTTPRouter",
		}
		r.References["sni_handler.handler.stream_handler.backend_group_id"] = config.Reference{
			Type: "BackendGroup",
		}
		r.References["sni_handler.handler.http_handler.http_router_id"] = config.Reference{
			Type: "HTTPRouter",
		}
		r.References["default_handler.handler.stream_handler.backend_group_id"] = config.Reference{
			Type: "BackendGroup",
		}
		r.References["default_handler.handler.http_handler.http_router_id"] = config.Reference{
			Type: "HTTPRouter",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_alb_backend_group", func(r *config.Resource) {
		r.References["http_backend.target_group_ids"] = config.Reference{
			Type: "TargetGroup",
		}
		r.References["stream_backend.target_group_ids"] = config.Reference{
			Type: "TargetGroup",
		}
		r.References["grpc_backend.target_group_ids"] = config.Reference{
			Type: "TargetGroup",
		}
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_alb_http_router", func(r *config.Resource) {
		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_alb_target_group", func(r *config.Resource) {
		r.References["target.subnet_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", vpc.ApisPackagePath, "Subnet"),
		}

		r.UseAsync = true
	})
	p.AddResourceConfigurator("yandex_alb_virtual_host", func(r *config.Resource) {
		r.References["http_router_id"] = config.Reference{
			Type: "HTTPRouter",
		}
		r.References["route.http_route.http_route_action.backend_group_id"] = config.Reference{
			Type: "BackendGroup",
		}
		r.References["route.grpc_route.grpc_route_action.backend_group_id"] = config.Reference{
			Type: "BackendGroup",
		}
		r.UseAsync = true
	})
}
