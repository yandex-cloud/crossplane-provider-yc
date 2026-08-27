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
	"context"
	"reflect"
	"strings"
	"testing"
)

func TestIAMExternalName(t *testing.T) {
	tests := map[string]struct {
		state   map[string]any
		fields  []string
		want    string
		wantErr string
	}{
		"provider ID is preferred": {
			state:  map[string]any{"id": "provider/resource/id"},
			fields: []string{"resource_id", "role"},
			want:   "provider/resource/id",
		},
		"binding ID is reconstructed": {
			state: map[string]any{
				"resource_id": "resource-id",
				"role":        "viewer",
			},
			fields: []string{"resource_id", "role"},
			want:   "resource-id/viewer",
		},
		"member ID is reconstructed": {
			state: map[string]any{
				"resource_id": "resource-id",
				"role":        "viewer",
				"member":      "serviceAccount:service-account-id",
			},
			fields: []string{"resource_id", "role", "member"},
			want:   "resource-id/viewer/serviceAccount:service-account-id",
		},
		"missing field is rejected": {
			state: map[string]any{
				"resource_id": "resource-id",
			},
			fields:  []string{"resource_id", "role"},
			wantErr: `field "role" is missing or empty`,
		},
		"non-string field is rejected": {
			state: map[string]any{
				"resource_id": "resource-id",
				"role":        42,
			},
			fields:  []string{"resource_id", "role"},
			wantErr: `field "role" is missing or empty`,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			externalName := IAMExternalName(test.fields...)
			got, err := externalName.GetExternalNameFn(test.state)
			if test.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), test.wantErr) {
					t.Fatalf("GetExternalNameFn() error = %v, want an error containing %q", err, test.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("GetExternalNameFn() unexpected error: %v", err)
			}
			if got != test.want {
				t.Errorf("GetExternalNameFn() = %q, want %q", got, test.want)
			}
			if !reflect.DeepEqual(externalName.IdentifierFields, test.fields) {
				t.Errorf("IdentifierFields = %v, want %v", externalName.IdentifierFields, test.fields)
			}

			id, err := externalName.GetIDFn(context.Background(), got, nil, nil)
			if err != nil {
				t.Fatalf("GetIDFn() unexpected error: %v", err)
			}
			if id != got {
				t.Errorf("GetIDFn() = %q, want %q", id, got)
			}
		})
	}
}
