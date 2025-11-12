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

package clients

import (
	"context"
	"encoding/json"
	"strings"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/crossplane/upjet/v2/pkg/terraform"
	sdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	clusterv1beta1 "github.com/yandex-cloud/crossplane-provider-yc/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/yandex-cloud/crossplane-provider-yc/apis/namespaced/v1beta1"
)

const (
	folderID              = "folder_id"
	cloudID               = "cloud_id"
	endpoint              = "endpoint"
	yqEndpoint            = "yq_endpoint"
	serviceAccountKeyFile = "service_account_key_file"
	token                 = "token"
	storageAccessKey      = "storage_access_key"
	storageSecretKey      = "storage_secret_key"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal template credentials as JSON"
	errBothTokenAndKeyFile  = "both token and serviceAccountKeyFile are specified, only one should be provided"
	errNoAuthMethod         = "neither token nor serviceAccountKeyFile is specified, one must be provided"
	errUnknownPCType        = "unknown provider config type"
	errUnknownGVK           = "unknown GVK for ProviderConfig"
	errNotAnObject          = "provider config is not an Object"
	errNotManagedResource   = "resource is not a managed resource"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string, ujprovider *config.Provider) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		// Resolve the appropriate ProviderConfig based on MR type (cluster vs namespaced)
		pcSpec, err := resolveProviderConfig(ctx, client, mg)
		if err != nil {
			return terraform.Setup{}, errors.Wrap(err, "cannot resolve provider config")
		}

		// Validate that only one authentication method is specified
		if pcSpec.Credentials.ServiceAccountKeyFile != nil && pcSpec.Credentials.Token != nil {
			return ps, errors.New(errBothTokenAndKeyFile)
		}

		// set provider configuration
		ps.Configuration = map[string]interface{}{}
		ps.Configuration[folderID] = pcSpec.Credentials.FolderID
		ps.Configuration[cloudID] = pcSpec.Credentials.CloudID
		ps.Configuration[endpoint] = pcSpec.Credentials.Endpoint
		ps.Configuration[yqEndpoint] = pcSpec.Credentials.YQEndpoint

		// Handle storage credentials - direct specification
		if pcSpec.Credentials.StorageAccessKey != nil {
			ps.Configuration[storageAccessKey] = *pcSpec.Credentials.StorageAccessKey
		}
		if pcSpec.Credentials.StorageSecretKey != nil {
			ps.Configuration[storageSecretKey] = *pcSpec.Credentials.StorageSecretKey
		}

		// Handle storage credentials from separate secrets
		if pcSpec.Credentials.StorageAccessKeySecretRef != nil {
			data, err := resource.CommonCredentialExtractor(ctx, xpv1.CredentialsSourceSecret, client, xpv1.CommonCredentialSelectors{
				SecretRef: &xpv1.SecretKeySelector{
					SecretReference: pcSpec.Credentials.StorageAccessKeySecretRef.SecretReference,
					Key:             pcSpec.Credentials.StorageAccessKeySecretRef.Key,
				},
			})
			if err != nil {
				return ps, errors.Wrap(err, "cannot extract storage access key from secret")
			}
			ps.Configuration[storageAccessKey] = string(data)
		}

		if pcSpec.Credentials.StorageSecretKeySecretRef != nil {
			data, err := resource.CommonCredentialExtractor(ctx, xpv1.CredentialsSourceSecret, client, xpv1.CommonCredentialSelectors{
				SecretRef: &xpv1.SecretKeySelector{
					SecretReference: pcSpec.Credentials.StorageSecretKeySecretRef.SecretReference,
					Key:             pcSpec.Credentials.StorageSecretKeySecretRef.Key,
				},
			})
			if err != nil {
				return ps, errors.Wrap(err, "cannot extract storage secret key from secret")
			}
			ps.Configuration[storageSecretKey] = string(data)
		}

		// Handle authentication based on the specified method
		if pcSpec.Credentials.Token != nil {
			// Use token authentication - direct specification
			ps.Configuration[token] = *pcSpec.Credentials.Token
		} else if pcSpec.Credentials.ServiceAccountKeyFile != nil {
			// Use service account key file authentication - direct specification
			ps.Configuration[serviceAccountKeyFile] = *pcSpec.Credentials.ServiceAccountKeyFile
		} else if pcSpec.Credentials.ServiceAccountKeySecretRef != nil {
			// Use service account key from separate secret
			data, err := resource.CommonCredentialExtractor(ctx, xpv1.CredentialsSourceSecret, client, xpv1.CommonCredentialSelectors{
				SecretRef: &xpv1.SecretKeySelector{
					SecretReference: pcSpec.Credentials.ServiceAccountKeySecretRef.SecretReference,
					Key:             pcSpec.Credentials.ServiceAccountKeySecretRef.Key,
				},
			})
			if err != nil {
				return ps, errors.Wrap(err, "cannot extract service account key from secret")
			}
			ps.Configuration[serviceAccountKeyFile] = string(data)
		} else {
			// This handles secret references and other credential sources (backward compatibility)
			data, err := resource.CommonCredentialExtractor(ctx, pcSpec.Credentials.Source, client, pcSpec.Credentials.CommonCredentialSelectors)
			if err != nil {
				return ps, errors.Wrap(err, errExtractCredentials)
			}

			// Handle different types of credentials from secrets
			if err := handleCredentialsFromSecret(data, ps.Configuration); err != nil {
				return ps, err
			}
		}

		// Ensure we have at least one main authentication method
		if _, hasToken := ps.Configuration[token]; !hasToken {
			if _, hasKeyFile := ps.Configuration[serviceAccountKeyFile]; !hasKeyFile {
				return ps, errors.New(errNoAuthMethod)
			}
		}

		diag := ujprovider.TerraformProvider.Configure(ctx, &sdk.ResourceConfig{
			Config: ps.Configuration,
		})
		if diag != nil && diag.HasError() {
			return ps, errors.Errorf("failed to configure the SDK provider: %v", diag)
		}
		ps.Meta = ujprovider.TerraformProvider.Meta()
		// Configure the framework provider - this is done by upjet's framework external client
		// We just need to set the FrameworkProvider and ensure Configuration is available
		ps.FrameworkProvider = ujprovider.TerraformPluginFrameworkProvider

		return ps, nil
	}
}

// toSharedPCSpec converts cluster-scoped ProviderConfig to the shared namespaced spec type
func toSharedPCSpec(pc *clusterv1beta1.ProviderConfig) (*namespacedv1beta1.ProviderConfigSpec, error) {
	if pc == nil {
		return nil, nil
	}
	data, err := json.Marshal(pc.Spec)
	if err != nil {
		return nil, err
	}

	var mSpec namespacedv1beta1.ProviderConfigSpec
	err = json.Unmarshal(data, &mSpec)
	return &mSpec, err
}

// resolveProviderConfig resolves the appropriate ProviderConfig based on the MR type
func resolveProviderConfig(ctx context.Context, crClient client.Client, mg resource.Managed) (*namespacedv1beta1.ProviderConfigSpec, error) {
	switch managed := mg.(type) {
	case resource.LegacyManaged:
		return resolveLegacy(ctx, crClient, managed)
	case resource.ModernManaged:
		return resolveModern(ctx, crClient, managed)
	default:
		return nil, errors.New(errNotManagedResource)
	}
}

// resolveLegacy resolves cluster-scoped ProviderConfig for legacy cluster-scoped MRs
func resolveLegacy(ctx context.Context, client client.Client, mg resource.LegacyManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}
	pc := &clusterv1beta1.ProviderConfig{}
	if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	t := resource.NewLegacyProviderConfigUsageTracker(client, &clusterv1beta1.ProviderConfigUsage{})
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}

	return toSharedPCSpec(pc)
}

// resolveModern resolves ProviderConfig for modern namespaced MRs
// It handles both namespaced ProviderConfig and cluster-scoped ClusterProviderConfig
func resolveModern(ctx context.Context, crClient client.Client, mg resource.ModernManaged) (*namespacedv1beta1.ProviderConfigSpec, error) {
	configRef := mg.GetProviderConfigReference()
	if configRef == nil {
		return nil, errors.New(errNoProviderConfig)
	}

	pcRuntimeObj, err := crClient.Scheme().New(namespacedv1beta1.SchemeGroupVersion.WithKind(configRef.Kind))
	if err != nil {
		return nil, errors.Wrap(err, errUnknownGVK)
	}
	pcObj, ok := pcRuntimeObj.(client.Object)
	if !ok {
		// This indicates a programming error, types are not properly generated
		return nil, errors.New(errNotAnObject)
	}

	// Namespace will be ignored if the PC is a cluster-scoped type (ClusterProviderConfig)
	if err := crClient.Get(ctx, types.NamespacedName{Name: configRef.Name, Namespace: mg.GetNamespace()}, pcObj); err != nil {
		return nil, errors.Wrap(err, errGetProviderConfig)
	}

	var pcSpec namespacedv1beta1.ProviderConfigSpec
	pcu := &namespacedv1beta1.ProviderConfigUsage{}
	switch pc := pcObj.(type) {
	case *namespacedv1beta1.ProviderConfig:
		pcSpec = pc.Spec
		// For namespaced ProviderConfig, override secret namespace to MR namespace
		if pcSpec.Credentials.SecretRef != nil {
			pcSpec.Credentials.SecretRef.Namespace = mg.GetNamespace()
		}
		if pcSpec.Credentials.ServiceAccountKeySecretRef != nil {
			pcSpec.Credentials.ServiceAccountKeySecretRef.Namespace = mg.GetNamespace()
		}
		if pcSpec.Credentials.StorageAccessKeySecretRef != nil {
			pcSpec.Credentials.StorageAccessKeySecretRef.Namespace = mg.GetNamespace()
		}
		if pcSpec.Credentials.StorageSecretKeySecretRef != nil {
			pcSpec.Credentials.StorageSecretKeySecretRef.Namespace = mg.GetNamespace()
		}
	case *namespacedv1beta1.ClusterProviderConfig:
		pcSpec = pc.Spec
		// ClusterProviderConfig secrets are already cluster-scoped, no namespace override needed
	default:
		return nil, errors.New(errUnknownPCType)
	}
	t := resource.NewProviderConfigUsageTracker(crClient, pcu)
	if err := t.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackUsage)
	}
	return &pcSpec, nil
}

// handleCredentialsFromSecret processes credential data from secrets and populates the configuration
func handleCredentialsFromSecret(data []byte, config map[string]interface{}) error {
	dataStr := string(data)
	dataStr = strings.TrimSpace(dataStr)

	if len(dataStr) == 0 {
		return errors.New("credential data is empty")
	}

	// Try to parse as JSON first - could be service account key or structured credentials
	var creds map[string]interface{}
	if err := json.Unmarshal([]byte(dataStr), &creds); err == nil {
		// Successfully parsed as JSON

		// Always check for and extract storage credentials first
		if accessKey, ok := creds[storageAccessKey].(string); ok && accessKey != "" {
			config[storageAccessKey] = accessKey
		}
		if secretKey, ok := creds[storageSecretKey].(string); ok && secretKey != "" {
			config[storageSecretKey] = secretKey
		}

		// Check if it contains a service_account_key_file field (separate from service account key JSON)
		if saKeyFile, ok := creds[serviceAccountKeyFile].(string); ok && saKeyFile != "" {
			// Validate that the service account key file is valid JSON
			var saKeyData map[string]interface{}
			if err := json.Unmarshal([]byte(saKeyFile), &saKeyData); err != nil {
				return errors.Wrap(err, "service_account_key_file contains invalid JSON")
			}
			config[serviceAccountKeyFile] = saKeyFile
			return nil
		}

		// Check if it contains token
		if tokenVal, ok := creds[token].(string); ok && tokenVal != "" {
			config[token] = tokenVal
			return nil
		}

		// Check if it's a service account key JSON (has specific fields)
		if _, hasID := creds["id"]; hasID {
			if _, hasServiceAccountID := creds["service_account_id"]; hasServiceAccountID {
				// This looks like a service account key JSON
				config[serviceAccountKeyFile] = dataStr
				return nil
			}
		}

		// If no main authentication method found in JSON, treat the whole thing as service account key
		config[serviceAccountKeyFile] = dataStr
		return nil
	}

	// Not valid JSON, treat as plain token
	if len(dataStr) == 0 {
		return errors.New("token is empty")
	}
	config[token] = dataStr
	return nil
}
