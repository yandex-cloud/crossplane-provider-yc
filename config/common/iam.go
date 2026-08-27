/*
Copyright 2026 YANDEX LLC

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

package common

import (
	"fmt"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// IAMExternalName reconstructs an IAM resource's external name from its
// identity fields when the Terraform provider does not return an ID in state.
func IAMExternalName(identityFields ...string) config.ExternalName {
	fields := append([]string(nil), identityFields...)

	return config.ExternalName{
		SetIdentifierArgumentFn: config.NopSetIdentifierArgument,
		GetExternalNameFn: func(tfState map[string]any) (string, error) {
			if id, ok := tfState["id"].(string); ok && id != "" {
				return id, nil
			}

			values := make([]string, 0, len(fields))
			for _, field := range fields {
				value, ok := tfState[field].(string)
				if !ok || value == "" {
					return "", fmt.Errorf("cannot reconstruct IAM external name: field %q is missing or empty in Terraform state", field)
				}
				values = append(values, value)
			}

			return strings.Join(values, "/"), nil
		},
		GetIDFn:                config.ExternalNameAsID,
		DisableNameInitializer: true,
		IdentifierFields:       fields,
	}
}
