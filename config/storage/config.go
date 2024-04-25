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

package storage

import (
	"fmt"

	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/resource"

	"github.com/yandex-cloud/provider-jet-yc/config/iam"
)

// Configure adds configurations for storage group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_storage_bucket", func(r *config.Resource) {
		r.References["access_key"] = config.Reference{
			Type:      fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccountStaticAccessKey"),
			Extractor: ExtractPublicKeyFuncPath,
		}
	})
	p.AddResourceConfigurator("yandex_storage_object", func(r *config.Resource) {
		r.References["access_key"] = config.Reference{
			Type:      fmt.Sprintf("%s.%s", iam.ApisPackagePath, "ServiceAccountStaticAccessKey"),
			Extractor: ExtractPublicKeyFuncPath,
		}
		r.References["bucket"] = config.Reference{
			Type: "Bucket",
		}
	})
}

const (
	// APISPackagePath is the package path for generated APIs root package
	APISPackagePath = "github.com/yandex-cloud/provider-jet-yc/config/storage"
	// ExtractPublicKeyFuncPath holds the Azure resource ID extractor func name
	ExtractPublicKeyFuncPath = APISPackagePath + ".ExtractAccessKey()"
)

// ExtractAccessKey extracts the value of `spec.atProvider.accessKey`
// from a Terraformed resource. If mr is not a Terraformed
// resource, returns an empty string.
func ExtractAccessKey() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}
		o, err := tr.GetObservation()
		if err != nil {
			return ""
		}
		if k := o["access_key"]; k != nil {
			return k.(string)
		}
		return ""
	}
}
