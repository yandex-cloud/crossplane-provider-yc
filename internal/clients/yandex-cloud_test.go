/*
Copyright 2022 YANDEX LLC

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

package clients

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestHandleCredentialsFromSecret(t *testing.T) {
	tests := []struct {
		name           string
		data           []byte
		expectedConfig map[string]interface{}
		expectedError  string
	}{
		{
			name:          "empty data",
			data:          []byte(""),
			expectedError: "credential data is empty",
		},
		{
			name: "plain token",
			data: []byte("test-token-123"),
			expectedConfig: map[string]interface{}{
				token: "test-token-123",
			},
		},
		{
			name: "service account key JSON",
			data: []byte(`{
				"id": "key-id",
				"service_account_id": "sa-id",
				"created_at": "2023-01-01T00:00:00Z",
				"key_algorithm": "RSA_2048",
				"public_key": "public-key-data",
				"private_key": "private-key-data"
			}`),
			expectedConfig: map[string]interface{}{
				serviceAccountKeyFile: `{
				"id": "key-id",
				"service_account_id": "sa-id",
				"created_at": "2023-01-01T00:00:00Z",
				"key_algorithm": "RSA_2048",
				"public_key": "public-key-data",
				"private_key": "private-key-data"
			}`,
			},
		},
		{
			name: "JSON with token and storage credentials",
			data: []byte(`{
				"token": "test-token-456",
				"storage_access_key": "access-key-123",
				"storage_secret_key": "secret-key-456"
			}`),
			expectedConfig: map[string]interface{}{
				token:            "test-token-456",
				storageAccessKey: "access-key-123",
				storageSecretKey: "secret-key-456",
			},
		},
		{
			name: "JSON with only storage credentials",
			data: []byte(`{
				"storage_access_key": "access-key-123",
				"storage_secret_key": "secret-key-456"
			}`),
			expectedConfig: map[string]interface{}{
				serviceAccountKeyFile: `{
				"storage_access_key": "access-key-123",
				"storage_secret_key": "secret-key-456"
			}`,
				storageAccessKey: "access-key-123",
				storageSecretKey: "secret-key-456",
			},
		},
		{
			name: "JSON with separate service account key file and storage credentials",
			data: []byte(`{
				"service_account_key_file": "{\"id\":\"key-id\",\"service_account_id\":\"sa-id\"}",
				"storage_access_key": "access-key-123",
				"storage_secret_key": "secret-key-456"
			}`),
			expectedConfig: map[string]interface{}{
				serviceAccountKeyFile: `{"id":"key-id","service_account_id":"sa-id"}`,
				storageAccessKey:      "access-key-123",
				storageSecretKey:      "secret-key-456",
			},
		},
		{
			name: "JSON with invalid service_account_key_file and storage credentials",
			data: []byte(`{
				"service_account_key_file": "invalid-json-content",
				"storage_access_key": "access-key-123",
				"storage_secret_key": "secret-key-456"
			}`),
			expectedError: "service_account_key_file contains invalid JSON: invalid character 'i' looking for beginning of value",
		},
		{
			name: "invalid JSON",
			data: []byte(`{"invalid": json}`),
			expectedConfig: map[string]interface{}{
				token: `{"invalid": json}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := make(map[string]interface{})
			err := handleCredentialsFromSecret(tt.data, config)

			if tt.expectedError != "" {
				if err == nil {
					t.Errorf("expected error %q, got nil", tt.expectedError)
					return
				}
				if err.Error() != tt.expectedError {
					t.Errorf("expected error %q, got %q", tt.expectedError, err.Error())
				}
				return
			}

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if diff := cmp.Diff(tt.expectedConfig, config); diff != "" {
				t.Errorf("config mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestHandleCredentialsFromSecret_StorageCredentialsExtraction(t *testing.T) {
	tests := []struct {
		name           string
		data           []byte
		expectedConfig map[string]interface{}
	}{
		{
			name: "JSON with partial storage credentials - only access key",
			data: []byte(`{
				"token": "test-token",
				"storage_access_key": "access-key-123"
			}`),
			expectedConfig: map[string]interface{}{
				token:            "test-token",
				storageAccessKey: "access-key-123",
			},
		},
		{
			name: "JSON with partial storage credentials - only secret key",
			data: []byte(`{
				"token": "test-token",
				"storage_secret_key": "secret-key-456"
			}`),
			expectedConfig: map[string]interface{}{
				token:            "test-token",
				storageSecretKey: "secret-key-456",
			},
		},
		{
			name: "JSON with empty storage credentials",
			data: []byte(`{
				"token": "test-token",
				"storage_access_key": "",
				"storage_secret_key": ""
			}`),
			expectedConfig: map[string]interface{}{
				token: "test-token",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config := make(map[string]interface{})
			err := handleCredentialsFromSecret(tt.data, config)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if diff := cmp.Diff(tt.expectedConfig, config); diff != "" {
				t.Errorf("config mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
