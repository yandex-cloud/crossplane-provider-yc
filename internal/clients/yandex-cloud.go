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

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/terraform"
	sdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/yandex-cloud/crossplane-provider-yc/apis/v1beta1"
)

const (
	folderID              = "folder_id"
	cloudID               = "cloud_id"
	endpoint              = "endpoint"
	yqEndpoint            = "yq_endpoint"
	serviceAccountKeyFile = "service_account_key_file"
	token                 = "token"
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

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		// Validate that only one authentication method is specified
		if pc.Spec.Credentials.ServiceAccountKeyFile != nil && pc.Spec.Credentials.Token != nil {
			return ps, errors.New(errBothTokenAndKeyFile)
		}

		// set provider configuration
		ps.Configuration = map[string]interface{}{}
		ps.Configuration[folderID] = pc.Spec.Credentials.FolderID
		ps.Configuration[cloudID] = pc.Spec.Credentials.CloudID
		ps.Configuration[endpoint] = pc.Spec.Credentials.Endpoint
		ps.Configuration[yqEndpoint] = pc.Spec.Credentials.YQEndpoint

		// Handle authentication based on the specified method
		if pc.Spec.Credentials.Token != nil {
			// Use token authentication - direct specification
			ps.Configuration[token] = *pc.Spec.Credentials.Token
		} else if pc.Spec.Credentials.ServiceAccountKeyFile != nil {
			// Use service account key file authentication - direct specification
			ps.Configuration[serviceAccountKeyFile] = *pc.Spec.Credentials.ServiceAccountKeyFile
		} else {
			// This handles secret references and other credential sources
			data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
			if err != nil {
				return ps, errors.Wrap(err, errExtractCredentials)
			}

			// Try to determine if this is a token or service account key based on content
			dataStr := string(data)
			dataStr = strings.TrimSpace(dataStr)

			if len(dataStr) == 0 {
				return ps, errors.New("credential data is empty")
			}

			// If it looks like JSON (starts with '{'), treat it as service account key
			if dataStr[0] == '{' {
				// Validate it's proper JSON
				creds := map[string]string{}
				if err := json.Unmarshal([]byte(dataStr), &creds); err != nil {
					return ps, errors.Wrap(err, errUnmarshalCredentials)
				}
				ps.Configuration[serviceAccountKeyFile] = dataStr
			} else {
				// Treat as token (plain string)
				// Ensure the token is not empty after trimming
				if len(dataStr) == 0 {
					return ps, errors.New("token is empty")
				}
				ps.Configuration[token] = dataStr
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
