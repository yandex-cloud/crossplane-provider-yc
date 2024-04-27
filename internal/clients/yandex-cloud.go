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

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/terraform"
	fwk "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	sdk "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/yandex-cloud/provider-jet-yc/apis/v1beta1"
)

const (
	folderID              = "folder_id"
	cloudID               = "cloud_id"
	endpoint              = "endpoint"
	serviceAccountKeyFile = "service_account_key_file"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal template credentials as JSON"
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

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// set provider configuration
		ps.Configuration = map[string]interface{}{}
		ps.Configuration[serviceAccountKeyFile] = string(data)
		ps.Configuration[folderID] = pc.Spec.Credentials.FolderID
		ps.Configuration[cloudID] = pc.Spec.Credentials.CloudID
		ps.Configuration[endpoint] = pc.Spec.Credentials.Endpoint
		if ujprovider.TerraformProvider != nil {
			diag := ujprovider.TerraformProvider.Configure(ctx, &sdk.ResourceConfig{
				Config: ps.Configuration,
			})
			if diag != nil && diag.HasError() {
				return ps, errors.Errorf("failed to configure the SDK provider: %v", diag)
			}
			ps.Meta = ujprovider.TerraformProvider.Meta()
		} else {
			return ps, errors.Wrap(err, "unable to set terraform SDK v2 provider")
		}
		if ujprovider.TerraformPluginFrameworkProvider != nil {
			sch := fwk.SchemaResponse{}
			ujprovider.TerraformPluginFrameworkProvider.Schema(ctx, fwk.SchemaRequest{}, &sch)
			resp := fwk.ConfigureResponse{}
			ujprovider.TerraformPluginFrameworkProvider.Configure(ctx, fwk.ConfigureRequest{
				Config: tfsdk.Config{
					Raw: tftypes.NewValue(tftypes.Object{AttributeTypes: map[string]tftypes.Type{
						serviceAccountKeyFile:     tftypes.String,
						folderID:                  tftypes.String,
						cloudID:                   tftypes.String,
						endpoint:                  tftypes.String,
						"organization_id":         tftypes.String,
						"zone":                    tftypes.String,
						"token":                   tftypes.String,
						"plaintext":               tftypes.Bool,
						"insecure":                tftypes.Bool,
						"max_retries":             tftypes.Number,
						"ymq_endpoint":            tftypes.String,
						"storage_secret_key":      tftypes.String,
						"ymq_secret_key":          tftypes.String,
						"storage_endpoint":        tftypes.String,
						"storage_access_key":      tftypes.String,
						"ymq_access_key":          tftypes.String,
						"shared_credentials_file": tftypes.String,
						"profile":                 tftypes.String,
						"region_id":               tftypes.String,
					},
					}, map[string]tftypes.Value{
						serviceAccountKeyFile:     tftypes.NewValue(tftypes.String, ps.Configuration[serviceAccountKeyFile]),
						folderID:                  tftypes.NewValue(tftypes.String, ps.Configuration[folderID]),
						cloudID:                   tftypes.NewValue(tftypes.String, ps.Configuration[cloudID]),
						endpoint:                  tftypes.NewValue(tftypes.String, ps.Configuration[endpoint]),
						"organization_id":         tftypes.NewValue(tftypes.String, ""),
						"zone":                    tftypes.NewValue(tftypes.String, ""),
						"token":                   tftypes.NewValue(tftypes.String, ""),
						"plaintext":               tftypes.NewValue(tftypes.Bool, false),
						"insecure":                tftypes.NewValue(tftypes.Bool, false),
						"max_retries":             tftypes.NewValue(tftypes.Number, 0),
						"ymq_endpoint":            tftypes.NewValue(tftypes.String, ""),
						"storage_secret_key":      tftypes.NewValue(tftypes.String, ""),
						"ymq_secret_key":          tftypes.NewValue(tftypes.String, ""),
						"storage_endpoint":        tftypes.NewValue(tftypes.String, ""),
						"ymq_access_key":          tftypes.NewValue(tftypes.String, ""),
						"storage_access_key":      tftypes.NewValue(tftypes.String, ""),
						"shared_credentials_file": tftypes.NewValue(tftypes.String, ""),
						"profile":                 tftypes.NewValue(tftypes.String, ""),
						"region_id":               tftypes.NewValue(tftypes.String, ""),
					}),
					Schema: sch.Schema,
				},
			}, &resp)
			if resp.Diagnostics != nil && resp.Diagnostics.HasError() {
				return ps, errors.Errorf("failed to configure the Framework provider: %v", resp.Diagnostics)
			}
			ps.FrameworkProvider = ujprovider.TerraformPluginFrameworkProvider
		} else {
			return ps, errors.Wrap(err, "unable to set terraform Plugin Framework provider")
		}

		return ps, nil
	}
}
