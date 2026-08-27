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

package common

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/crossplane/crossplane-runtime/v2/pkg/fieldpath"
	"github.com/crossplane/crossplane-runtime/v2/pkg/reconciler/managed"
	xpresource "github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Provider version. Will be re-defined upon build.
var Version = "0.0.0-dev"

// MissingResourceID is a syntactically valid Yandex Cloud resource ID used
// only for the initial read of Plugin Framework resources whose identifiers
// are computed by the provider.
const MissingResourceID = "aaaaaaaaaaaaaaaaaaaa"

var frameworkComputedIdentifierFields = map[string]string{
	"yandex_compute_disk_placement_group": "disk_placement_group_id",
	"yandex_compute_filesystem":           "filesystem_id",
	"yandex_compute_gpu_cluster":          "gpu_cluster_id",
	"yandex_container_registry":           "registry_id",
	"yandex_container_repository":         "repository_id",
	"yandex_datatransfer_endpoint":        "endpoint_id",
	"yandex_datatransfer_transfer":        "transfer_id",
	"yandex_iam_service_account":          "service_account_id",
	"yandex_kms_symmetric_key":            "symmetric_key_id",
	"yandex_lb_target_group":              "target_group_id",
	"yandex_organizationmanager_group":    "group_id",
	"yandex_resourcemanager_cloud":        "cloud_id",
	"yandex_resourcemanager_folder":       "folder_id",
}

// DefaultResourceOverrides returns a default resource configuration to be used while
// building resource configurations.
// folderAPIPath should be the full API path to the Folder type, e.g.:
// - "github.com/yandex-cloud/crossplane-provider-yc/apis/cluster/resourcemanager/v1alpha1.Folder"
// - "github.com/yandex-cloud/crossplane-provider-yc/apis/namespaced/resourcemanager/v1alpha1.Folder"
func DefaultResourceOverrides(folderAPIPath string) config.ResourceOption {
	return func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		if identifier, ok := frameworkComputedIdentifierFields[r.Name]; ok {
			r.ExternalName = frameworkResourceWithComputedIdentifier(identifier)
		}
		defaultFolderIDFn(r, folderAPIPath)

		if s, ok := r.TerraformResource.Schema["labels"]; ok && s.Type == schema.TypeMap {
			r.InitializerFns = append(r.InitializerFns, func(client client.Client) managed.Initializer {
				return NewLabeller(client, "labels")
			})
		}
	}
}

func frameworkResourceWithComputedIdentifier(identifier string) config.ExternalName {
	externalName := config.FrameworkResourceWithComputedIdentifier(identifier, MissingResourceID)
	externalName.IsNotFoundDiagnosticFn = func(diagnostics []*tfprotov6.Diagnostic) bool {
		for _, diagnostic := range diagnostics {
			if diagnostic == nil || diagnostic.Severity != tfprotov6.DiagnosticSeverityError {
				continue
			}
			if strings.Contains(diagnostic.Summary, MissingResourceID) || strings.Contains(diagnostic.Detail, MissingResourceID) {
				return true
			}
		}
		return false
	}
	return externalName
}

func defaultFolderIDFn(r *config.Resource, folderAPIPath string) {
	// Skip adding folder_id reference for resourcemanager resources themselves
	// except for IAM binding resources that were moved from resourcemanager group
	if r.ShortGroup == "resourcemanager" &&
		r.Name != "yandex_resourcemanager_folder_iam_member" &&
		r.Name != "yandex_resourcemanager_folder_iam_binding" {
		// For resourcemanager group resources, use local Folder type
		r.References["folder_id"] = config.Reference{
			Type: "Folder",
		}
	} else if folderAPIPath != "" {
		// For all other resources, use the provided folder API path
		r.References["folder_id"] = config.Reference{
			Type: folderAPIPath,
		}
	}
}

// Labeller implements the Initialize function to set YC labels
type Labeller struct {
	kube      client.Client
	fieldName string
}

// NewLabeller returns a Labeller object.
func NewLabeller(kube client.Client, fieldName string) *Labeller {
	return &Labeller{kube: kube, fieldName: fieldName}
}

// Initialize is a custom initializer for setting YC labels.
func (t *Labeller) Initialize(ctx context.Context, mg xpresource.Managed) error {
	paved, err := fieldpath.PaveObject(mg)
	if err != nil {
		return fmt.Errorf("failed to pave Managed resource: %w", err)
	}
	pavedByte, err := setYCLabelsWithPaved(paved, t.fieldName)
	if err != nil {
		return fmt.Errorf("failed to set YC labels in paved resource: %w", err)
	}
	if err := json.Unmarshal(pavedByte, mg); err != nil {
		return fmt.Errorf("failed to unmarshal paved resource: %w", err)
	}
	if err := t.kube.Update(ctx, mg); err != nil {
		return fmt.Errorf("failed to update Managed resource using k8s client: %w", err)
	}
	return nil
}

func setYCLabelsWithPaved(paved *fieldpath.Paved, fieldName string) ([]byte, error) {
	// our version is usually semVer, but some resources validate labels as "[-_0-9a-z]*"
	ver := strings.ReplaceAll(Version, ".", "-")
	ver = strings.ReplaceAll(ver, "+", "_")

	tags := map[string]*string{
		"managed-by":                  ptr.To("crossplane-provider-yc"),
		"crossplane-provider-version": ptr.To(ver),
	}

	if err := paved.SetValue(fmt.Sprintf("spec.forProvider.%s", fieldName), tags); err != nil {
		return nil, err
	}
	pavedByte, err := paved.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return pavedByte, nil
}

// SingletonMergeStrategy is useful to avoid passing an empty object during
// server-side apply.
var SingletonMergeStrategy = config.MergeStrategy{
	ListMergeStrategy: config.ListMergeStrategy{
		ListMapKeys: config.ListMapKeys{
			InjectedKey: config.InjectedKey{
				Key:          "index",
				DefaultValue: "default",
			},
		},
		MergeStrategy: config.ListTypeMap,
	},
}
