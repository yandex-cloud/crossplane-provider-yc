package common

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/upbound/upjet/pkg/config"
	registry "github.com/upbound/upjet/pkg/registry"

	"github.com/yandex-cloud/provider-jet-yc/config/resourcemanager"
)

// DefaultResourceFn returns a default resource configuration to be used while
// building resource configurations.
func DefaultResourceFn(name string, terraformResource *schema.Resource, terraformRegistry *registry.Resource, opts ...config.ResourceOption) *config.Resource {
	r := config.DefaultResource(name, terraformResource, terraformRegistry)
	// Add any provider-specific defaulting here. For example:
	r.ExternalName = config.IdentifierFromProvider
	defaultFolderIDFn(r)
	return r
}

// DefaultResourceOverrides returns a default resource configuration to be used while
// building resource configurations.
func DefaultResourceOverrides() config.ResourceOption {
	return func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		defaultFolderIDFn(r)
	}
}

func defaultFolderIDFn(r *config.Resource) {
	if r.ShortGroup != "resourcemanager" ||
		// Fix for group change from "resourcemanager" to "iam"
		r.Name == "yandex_resourcemanager_folder_iam_member" ||
		r.Name == "yandex_resourcemanager_folder_iam_binding" {
		r.References["folder_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", resourcemanager.ApisPackagePath, "Folder"),
		}
	} else {
		r.References["folder_id"] = config.Reference{
			Type: "Folder",
		}
	}
}
