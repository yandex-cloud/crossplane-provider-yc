package config

import (
	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	tf "github.com/hashicorp/terraform-provider-hashicups/hashicups"
)

const resourcePrefix = "template"

// GetProvider returns provider configuration
func GetProvider() tjconfig.Provider {
	pc := tjconfig.NewProvider(tf.Provider(), resourcePrefix, "github.com/crossplane-contrib/provider-tf-template")

	for _, configure := range []func(provider tjconfig.Provider){
		// add custom config functions
	} {
		configure(pc)
	}

	return pc
}
