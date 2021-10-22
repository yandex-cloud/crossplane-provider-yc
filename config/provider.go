package config

import (
	// TODO(turkenh): try to find another package name for configs here or in terrajet
	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	"github.com/hashicorp/terraform-provider-hashicups/hashicups"
)

func GetProvider() tjconfig.Provider {
	pc := tjconfig.Provider{
		Schema:         hashicups.Provider(),
		GroupSuffix:    ".hashicups.tf.crossplane.io",
		ResourcePrefix: "hashicups_",
		ShortName:      "tfhashicups",
		ModulePath:     "github.com/crossplane-contrib/provider-tf-template",
		IncludeList: []string{
			"hashicups_.+",
		},

		Resource: map[string]tjconfig.Resource{},
	}

	for _, setup := range []func(provider tjconfig.Provider){

	} {
		setup(pc)
	}

	return pc
}
