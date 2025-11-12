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
	"context"
	"testing"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/yandex-cloud/crossplane-provider-yc/apis/cluster/v1beta1"
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

func TestSeparateSecrets(t *testing.T) {
	type args struct {
		pc      *v1beta1.ProviderConfig
		secrets []client.Object
	}
	type want struct {
		hasServiceAccountKey bool
		hasStorageAccessKey  bool
		hasStorageSecretKey  bool
		errContains          string
	}

	cases := map[string]struct {
		args args
		want want
	}{
		"ServiceAccountKeyInSeparateSecret": {
			args: args{
				pc: &v1beta1.ProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: "test-pc"},
					Spec: v1beta1.ProviderConfigSpec{
						Credentials: v1beta1.ProviderCredentials{
							Source: xpv1.CredentialsSourceSecret,
							ServiceAccountKeySecretRef: &xpv1.SecretKeySelector{
								SecretReference: xpv1.SecretReference{
									Name:      "sa-secret",
									Namespace: "default",
								},
								Key: "sa-key",
							},
							FolderID: "test-folder",
							CloudID:  "test-cloud",
						},
					},
				},
				secrets: []client.Object{
					&corev1.Secret{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "sa-secret",
							Namespace: "default",
						},
						Data: map[string][]byte{
							"sa-key": []byte(`{"id":"test-id","service_account_id":"test-sa"}`),
						},
					},
				},
			},
			want: want{
				hasServiceAccountKey: true,
				hasStorageAccessKey:  false,
				hasStorageSecretKey:  false,
			},
		},
		"StorageCredentialsInSeparateSecrets": {
			args: args{
				pc: &v1beta1.ProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: "test-pc"},
					Spec: v1beta1.ProviderConfigSpec{
						Credentials: v1beta1.ProviderCredentials{
							Source: xpv1.CredentialsSourceSecret,
							ServiceAccountKeySecretRef: &xpv1.SecretKeySelector{
								SecretReference: xpv1.SecretReference{
									Name:      "sa-secret",
									Namespace: "default",
								},
								Key: "sa-key",
							},
							StorageAccessKeySecretRef: &xpv1.SecretKeySelector{
								SecretReference: xpv1.SecretReference{
									Name:      "storage-access-secret",
									Namespace: "default",
								},
								Key: "access-key",
							},
							StorageSecretKeySecretRef: &xpv1.SecretKeySelector{
								SecretReference: xpv1.SecretReference{
									Name:      "storage-secret-secret",
									Namespace: "default",
								},
								Key: "secret-key",
							},
							FolderID: "test-folder",
							CloudID:  "test-cloud",
						},
					},
				},
				secrets: []client.Object{
					&corev1.Secret{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "sa-secret",
							Namespace: "default",
						},
						Data: map[string][]byte{
							"sa-key": []byte(`{"id":"test-id","service_account_id":"test-sa"}`),
						},
					},
					&corev1.Secret{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "storage-access-secret",
							Namespace: "default",
						},
						Data: map[string][]byte{
							"access-key": []byte("test-access-key"),
						},
					},
					&corev1.Secret{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "storage-secret-secret",
							Namespace: "default",
						},
						Data: map[string][]byte{
							"secret-key": []byte("test-secret-key"),
						},
					},
				},
			},
			want: want{
				hasServiceAccountKey: true,
				hasStorageAccessKey:  true,
				hasStorageSecretKey:  true,
			},
		},
		"MixedDirectAndSecretRef": {
			args: args{
				pc: &v1beta1.ProviderConfig{
					ObjectMeta: metav1.ObjectMeta{Name: "test-pc"},
					Spec: v1beta1.ProviderConfigSpec{
						Credentials: v1beta1.ProviderCredentials{
							Source: xpv1.CredentialsSourceSecret,
							ServiceAccountKeySecretRef: &xpv1.SecretKeySelector{
								SecretReference: xpv1.SecretReference{
									Name:      "sa-secret",
									Namespace: "default",
								},
								Key: "sa-key",
							},
							StorageAccessKey: stringPtr("direct-access-key"),
							StorageSecretKey: stringPtr("direct-secret-key"),
							FolderID:         "test-folder",
							CloudID:          "test-cloud",
						},
					},
				},
				secrets: []client.Object{
					&corev1.Secret{
						ObjectMeta: metav1.ObjectMeta{
							Name:      "sa-secret",
							Namespace: "default",
						},
						Data: map[string][]byte{
							"sa-key": []byte(`{"id":"test-id","service_account_id":"test-sa"}`),
						},
					},
				},
			},
			want: want{
				hasServiceAccountKey: true,
				hasStorageAccessKey:  true,
				hasStorageSecretKey:  true,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			scheme := runtime.NewScheme()
			_ = corev1.AddToScheme(scheme)
			_ = v1beta1.SchemeBuilder.AddToScheme(scheme)

			fakeClient := fake.NewClientBuilder().
				WithScheme(scheme).
				WithObjects(tc.args.secrets...).
				Build()

			// We can't fully test TerraformSetupBuilder without a real managed resource
			// and provider setup, but we can verify the credential extraction logic
			ctx := context.Background()

			// Test service account key extraction
			if tc.args.pc.Spec.Credentials.ServiceAccountKeySecretRef != nil {
				data, err := resource.CommonCredentialExtractor(ctx, xpv1.CredentialsSourceSecret, fakeClient, xpv1.CommonCredentialSelectors{
					SecretRef: &xpv1.SecretKeySelector{
						SecretReference: tc.args.pc.Spec.Credentials.ServiceAccountKeySecretRef.SecretReference,
						Key:             tc.args.pc.Spec.Credentials.ServiceAccountKeySecretRef.Key,
					},
				})

				if tc.want.errContains != "" {
					if err == nil {
						t.Errorf("expected error containing %q, got nil", tc.want.errContains)
					}
					return
				}

				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if tc.want.hasServiceAccountKey && len(data) == 0 {
					t.Error("expected service account key data, got empty")
				}
			}

			// Test storage access key extraction
			if tc.args.pc.Spec.Credentials.StorageAccessKeySecretRef != nil {
				data, err := resource.CommonCredentialExtractor(ctx, xpv1.CredentialsSourceSecret, fakeClient, xpv1.CommonCredentialSelectors{
					SecretRef: &xpv1.SecretKeySelector{
						SecretReference: tc.args.pc.Spec.Credentials.StorageAccessKeySecretRef.SecretReference,
						Key:             tc.args.pc.Spec.Credentials.StorageAccessKeySecretRef.Key,
					},
				})

				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if tc.want.hasStorageAccessKey && len(data) == 0 {
					t.Error("expected storage access key data, got empty")
				}
			}

			// Test storage secret key extraction
			if tc.args.pc.Spec.Credentials.StorageSecretKeySecretRef != nil {
				data, err := resource.CommonCredentialExtractor(ctx, xpv1.CredentialsSourceSecret, fakeClient, xpv1.CommonCredentialSelectors{
					SecretRef: &xpv1.SecretKeySelector{
						SecretReference: tc.args.pc.Spec.Credentials.StorageSecretKeySecretRef.SecretReference,
						Key:             tc.args.pc.Spec.Credentials.StorageSecretKeySecretRef.Key,
					},
				})

				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if tc.want.hasStorageSecretKey && len(data) == 0 {
					t.Error("expected storage secret key data, got empty")
				}
			}
		})
	}
}

func stringPtr(s string) *string {
	return &s
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
