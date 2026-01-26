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

	"github.com/ghodss/yaml"
)

// kubernetesClusterConnectionDetails extracts connection details from a Kubernetes cluster.
func KubernetesClusterConnectionDetails(attr map[string]any) map[string][]byte {
	conn := make(map[string][]byte)

	// Extract master information
	if masterList, ok := attr["master"].([]any); ok && len(masterList) > 0 {
		if master, ok := masterList[0].(map[string]any); ok {
			// Extract CA certificate
			if caCert, ok := master["cluster_ca_certificate"].(string); ok && caCert != "" {
				conn["cluster_ca_certificate"] = []byte(caCert)
			}

			// Extract endpoints
			if endpoint, ok := master["external_v4_endpoint"].(string); ok && endpoint != "" {
				conn["endpoint"] = []byte(endpoint)
			}
			if endpoint, ok := master["internal_v4_endpoint"].(string); ok && endpoint != "" {
				conn["internal_endpoint"] = []byte(endpoint)
			}

			// Generate kubeconfig if we have the necessary data
			if caCert, hasCert := master["cluster_ca_certificate"].(string); hasCert && caCert != "" {
				if endpoint, hasEndpoint := master["external_v4_endpoint"].(string); hasEndpoint && endpoint != "" {
					if clusterName, hasName := attr["name"].(string); hasName {
						kubeconfig := generateKubeconfig(clusterName, endpoint, caCert)
						conn["kubeconfig"] = []byte(kubeconfig)
					}
				}
			}
		}
	}

	return conn
}

// generateKubeconfig creates a kubeconfig YAML for the cluster.
func generateKubeconfig(clusterName, endpoint, caCert string) string {
	kubeconfig := map[string]any{
		"apiVersion": "v1",
		"kind":       "Config",
		"clusters": []map[string]any{
			{
				"name": clusterName,
				"cluster": map[string]any{
					"server":                     endpoint,
					"certificate-authority-data": base64.StdEncoding.EncodeToString([]byte(caCert)),
				},
			},
		},
		"contexts": []map[string]any{
			{
				"name": clusterName,
				"context": map[string]any{
					"cluster": clusterName,
					"user":    clusterName,
				},
			},
		},
		"current-context": clusterName,
		"users": []map[string]any{
			{
				"name": clusterName,
				"user": map[string]any{
					"exec": map[string]any{
						"apiVersion": "client.authentication.k8s.io/v1beta1",
						"command":    "yc",
						"args": []string{
							"managed-kubernetes",
							"create-token",
							"--cluster-name=" + clusterName,
						},
					},
				},
			},
		},
	}

	data, err := yaml.Marshal(kubeconfig)
	if err != nil {
		return ""
	}
	return string(data)
}
