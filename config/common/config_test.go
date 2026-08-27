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
	"reflect"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestFrameworkComputedIdentifiers(t *testing.T) {
	for resourceName, identifier := range frameworkComputedIdentifierFields {
		t.Run(resourceName, func(t *testing.T) {
			r := &config.Resource{
				Name: resourceName,
				TerraformResource: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			}
			DefaultResourceOverrides("")(r)

			params := map[string]any{}
			r.ExternalName.SetIdentifierArgumentFn(params, "")

			if got := params[identifier]; got != MissingResourceID {
				t.Errorf("identifier placeholder = %v, want %q", got, MissingResourceID)
			}
			if got, want := r.ExternalName.TFPluginFrameworkOptions.ComputedIdentifierAttributes, []string{identifier}; !reflect.DeepEqual(got, want) {
				t.Errorf("ComputedIdentifierAttributes = %v, want %v", got, want)
			}
		})
	}
}

func TestFrameworkComputedIdentifierNotFoundDiagnostic(t *testing.T) {
	isNotFound := frameworkResourceWithComputedIdentifier("resource_id").IsNotFoundDiagnosticFn
	tests := map[string]struct {
		diagnostics []*tfprotov6.Diagnostic
		want        bool
	}{
		"placeholder in error detail": {
			diagnostics: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Detail:   "invalid resource id '" + MissingResourceID + "'",
			}},
			want: true,
		},
		"placeholder in error summary": {
			diagnostics: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Summary:  "resource " + MissingResourceID + " was not found",
			}},
			want: true,
		},
		"unrelated error": {
			diagnostics: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityError,
				Detail:   "permission denied",
			}},
		},
		"warning is not suppressed": {
			diagnostics: []*tfprotov6.Diagnostic{{
				Severity: tfprotov6.DiagnosticSeverityWarning,
				Detail:   "invalid resource id '" + MissingResourceID + "'",
			}},
		},
		"nil diagnostic": {
			diagnostics: []*tfprotov6.Diagnostic{nil},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isNotFound(test.diagnostics); got != test.want {
				t.Errorf("IsNotFoundDiagnosticFn() = %t, want %t", got, test.want)
			}
		})
	}
}
