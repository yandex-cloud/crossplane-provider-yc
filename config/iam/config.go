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

package iam

import (
	"encoding/json"
	"fmt"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reference"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/upbound/upjet/pkg/config"
)

const (
	// ApisPackagePath is the golang path for this package.
	ApisPackagePath = "github.com/yandex-cloud/provider-jet-yc/apis/iam/v1alpha1"
	// ConfigPath is the golang path for this package.
	ConfigPath = "github.com/yandex-cloud/provider-jet-yc/config/iam"
	// ServiceAccountRefValueFn is the name of resolver.
	ServiceAccountRefValueFn = "ServiceAccountRefValue()"
)

func serviceAccountKey(attr map[string]interface{}) ([]byte, error) {
	required := []string{
		"id",
		"service_account_id",
		"created_at",
		"key_algorithm",
		"public_key",
		"private_key",
	}

	for _, field := range required {
		if _, ok := attr[field]; !ok {
			return nil, fmt.Errorf("required field %s is missing", field)
		}
	}
	result := map[string]string{
		"id":                 attr["id"].(string),
		"service_account_id": attr["service_account_id"].(string),
		"created_at":         attr["created_at"].(string),
		"key_algorithm":      attr["key_algorithm"].(string),
		"public_key":         attr["public_key"].(string),
		"private_key":        attr["private_key"].(string),
	}
	encoded, err := json.Marshal(result)
	return encoded, err
}

func serviceAccountStaticKey(attr map[string]interface{}) (map[string][]byte, error) {
	if _, ok := attr["access_key"]; !ok {
		return nil, nil
	}
	if _, ok := attr["id"]; !ok {
		return nil, nil
	}
	return map[string][]byte{
		"attribute.access_key": []byte(attr["access_key"].(string)),
		"attribute.id":         []byte(attr["id"].(string)),
	}, nil
}

// Configure adds configurations for iam group.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("yandex_iam_service_account_key", func(r *config.Resource) {
		r.References["service_account_id"] = config.Reference{
			Type: "ServiceAccount",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			bb, err := serviceAccountKey(attr)
			return map[string][]byte{
				"service_account_key": bb,
			}, err
		}
	})
	p.AddResourceConfigurator("yandex_iam_service_account_static_access_key", func(r *config.Resource) {
		r.References["service_account_id"] = config.Reference{
			Type: "ServiceAccount",
		}
		r.Sensitive.AdditionalConnectionDetailsFn = serviceAccountStaticKey
	})
	p.AddResourceConfigurator("yandex_iam_service_account_iam_member", func(r *config.Resource) {
		r.References["service_account_id"] = config.Reference{
			Type: "ServiceAccount",
		}
	})
	p.AddResourceConfigurator("yandex_iam_service_account_iam_member", func(r *config.Resource) {
		r.References["member"] = config.Reference{
			Type:              "ServiceAccount",
			Extractor:         fmt.Sprintf("%s.%s", ConfigPath, ServiceAccountRefValueFn),
			RefFieldName:      "ServiceAccountRef",
			SelectorFieldName: "ServiceAccountSelector",
		}
	})
}

// ServiceAccountRefValue returns an extractor that returns templated value with service account id of ServiceAccount.
func ServiceAccountRefValue() reference.ExtractValueFn {
	return func(mg resource.Managed) string {
		if externalName := meta.GetExternalName(mg); externalName != "" {
			return fmt.Sprintf("serviceAccount:%s", externalName)
		}
		return ""
	}
}
