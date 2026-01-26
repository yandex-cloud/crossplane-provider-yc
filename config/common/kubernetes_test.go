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
	"encoding/base64"
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

const clusterAttrsJSON = `{
    "id": "cathn0s6n567890abcde",
    "name": "test-k8s-cluster",
    "description": "Test Kubernetes cluster",
    "folder_id": "b1gj2tg21doe4mcdr530",
    "created_at": "2023-12-24T08:51:11Z",
    "status": "RUNNING",
    "health": "HEALTHY",
    "network_id": "enp42t1n32u1r35t1qm3",
    "master": [{
        "version": "1.28",
        "cluster_ca_certificate": "-----BEGIN CERTIFICATE-----\nMIICyDCCAbCgAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl\ncm5ldGVzMB4XDTIzMTIyNDA4NTExMVoXDTMzMTIyMTA4NTExMVowFTETMBEGA1UE\nAxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL\n-----END CERTIFICATE-----",
        "external_v4_endpoint": "https://84.201.123.45",
        "internal_v4_endpoint": "https://10.0.0.10",
        "external_v4_address": "84.201.123.45",
        "internal_v4_address": "10.0.0.10",
        "public_ip": true,
        "version_info": [{
            "current_version": "1.28",
            "new_revision_available": false
        }]
    }],
    "service_account_id": "aje9k8luj4qf890abcde",
    "node_service_account_id": "aje9k8luj4qf890abcde",
    "release_channel": "REGULAR"
}`

func Test_kubernetesClusterConnectionDetails(t *testing.T) {
	var attrs map[string]any
	_ = json.Unmarshal([]byte(clusterAttrsJSON), &attrs)

	tests := []struct {
		name string
		attr map[string]any
		want map[string][]byte
	}{
		{
			name: "full cluster with all connection details",
			attr: attrs,
			want: map[string][]byte{
				"cluster_ca_certificate": []byte("-----BEGIN CERTIFICATE-----\nMIICyDCCAbCgAwIBAgIBADANBgkqhkiG9w0BAQsFADAVMRMwEQYDVQQDEwprdWJl\ncm5ldGVzMB4XDTIzMTIyNDA4NTExMVoXDTMzMTIyMTA4NTExMVowFTETMBEGA1UE\nAxMKa3ViZXJuZXRlczCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAL\n-----END CERTIFICATE-----"),
				"endpoint":               []byte("https://84.201.123.45"),
				"internal_endpoint":      []byte("https://10.0.0.10"),
			},
		},
		{
			name: "cluster without master data",
			attr: map[string]any{
				"id":   "test-id",
				"name": "test-cluster",
			},
			want: map[string][]byte{},
		},
		{
			name: "cluster with empty master array",
			attr: map[string]any{
				"id":     "test-id",
				"name":   "test-cluster",
				"master": []any{},
			},
			want: map[string][]byte{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KubernetesClusterConnectionDetails(tt.attr)

			// Check that all expected keys are present
			for key, wantValue := range tt.want {
				gotValue, exists := got[key]
				if !exists {
					t.Errorf("kubernetesClusterConnectionDetails() missing key %s", key)
					continue
				}
				if !reflect.DeepEqual(gotValue, wantValue) {
					t.Errorf("kubernetesClusterConnectionDetails() key %s = %s, want %s", key, gotValue, wantValue)
				}
			}

			// For the full test case, also check kubeconfig
			if tt.name == "full cluster with all connection details" {
				kubeconfigBytes, exists := got["kubeconfig"]
				if !exists {
					t.Error("kubernetesClusterConnectionDetails() missing kubeconfig")
				} else {
					kubeconfig := string(kubeconfigBytes)
					// Verify kubeconfig contains expected elements
					if !strings.Contains(kubeconfig, "apiVersion: v1") {
						t.Error("kubeconfig missing apiVersion")
					}
					if !strings.Contains(kubeconfig, "kind: Config") {
						t.Error("kubeconfig missing kind")
					}
					if !strings.Contains(kubeconfig, "test-k8s-cluster") {
						t.Error("kubeconfig missing cluster name")
					}
					if !strings.Contains(kubeconfig, "https://84.201.123.45") {
						t.Error("kubeconfig missing server endpoint")
					}
					if !strings.Contains(kubeconfig, "certificate-authority-data") {
						t.Error("kubeconfig missing certificate-authority-data")
					}
				}
			}
		})
	}
}

func Test_generateKubeconfig(t *testing.T) {
	tests := []struct {
		name        string
		clusterName string
		endpoint    string
		caCert      string
		wantContain []string
	}{
		{
			name:        "valid kubeconfig generation",
			clusterName: "test-cluster",
			endpoint:    "https://84.201.123.45",
			caCert:      "-----BEGIN CERTIFICATE-----\ntest\n-----END CERTIFICATE-----",
			wantContain: []string{
				"apiVersion: v1",
				"kind: Config",
				"test-cluster",
				"https://84.201.123.45",
				"certificate-authority-data:",
				"yc",
				"managed-kubernetes",
				"create-token",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateKubeconfig(tt.clusterName, tt.endpoint, tt.caCert)

			if got == "" {
				t.Error("generateKubeconfig() returned empty string")
				return
			}

			for _, want := range tt.wantContain {
				if !strings.Contains(got, want) {
					t.Errorf("generateKubeconfig() missing expected content: %s", want)
				}
			}

			// Verify the CA cert is base64 encoded in the output
			expectedEncoded := base64.StdEncoding.EncodeToString([]byte(tt.caCert))
			if !strings.Contains(got, expectedEncoded) {
				t.Error("generateKubeconfig() CA certificate not properly base64 encoded")
			}
		})
	}
}

func Test_kubernetesClusterConnectionDetails_EdgeCases(t *testing.T) {
	tests := []struct {
		name string
		attr map[string]any
	}{
		{
			name: "nil attributes",
			attr: nil,
		},
		{
			name: "master with missing fields",
			attr: map[string]any{
				"name": "test",
				"master": []any{
					map[string]any{
						"version": "1.28",
					},
				},
			},
		},
		{
			name: "master with empty strings",
			attr: map[string]any{
				"name": "test",
				"master": []any{
					map[string]any{
						"cluster_ca_certificate": "",
						"external_v4_endpoint":   "",
						"internal_v4_endpoint":   "",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Should not panic
			got := KubernetesClusterConnectionDetails(tt.attr)
			if got == nil {
				t.Error("kubernetesClusterConnectionDetails() returned nil, expected empty map")
			}
		})
	}
}
