package common

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane/terrajet/pkg/config"
	tjconfig "github.com/crossplane/terrajet/pkg/config"

	"github.com/yandex-cloud/provider-jet-yc/config/resourcemanager"
)

// DefaultResourceFn returns a default resource configuration to be used while
// building resource configurations.
func DefaultResourceFn(name string, terraformResource *schema.Resource, opts ...tjconfig.ResourceOption) *tjconfig.Resource {
	r := tjconfig.DefaultResource(name, terraformResource)
	// Add any provider-specific defaulting here. For example:
	r.ExternalName = tjconfig.IdentifierFromProvider
	defaultFolderIDFn(r)
	return r
}

func defaultFolderIDFn(r *tjconfig.Resource) {
	if r.ShortGroup != "resourcemanager" {
		r.References["folder_id"] = config.Reference{
			Type: fmt.Sprintf("%s.%s", resourcemanager.ApisPackagePath, "Folder"),
		}
	} else {
		r.References["folder_id"] = config.Reference{
			Type: "Folder",
		}
	}
}
