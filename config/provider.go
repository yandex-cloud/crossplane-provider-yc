package config

import (
	tjconfig "github.com/crossplane-contrib/terrajet/pkg/config"
	tf "github.com/hashicorp/terraform-provider-hashicups/hashicups"
)

const resourcePrefix = "template"

// GetProvider returns provider configuration
func GetProvider() tjconfig.Provider {
	pc := tjconfig.NewProvider(tf.Provider().ResourcesMap, resourcePrefix, "github.com/crossplane-contrib/provider-tf-template")
	// Comment out the line below instead of the above, if your Terraform
	// provider uses an old version (<v2) of github.com/hashicorp/terraform-plugin-sdk.
	// pc := tjconfig.NewProvider(tjconfig.GetV2ResourceMap(tf.Provider()), resourcePrefix, "github.com/crossplane-contrib/provider-tf-template")

	for _, configure := range []func(provider tjconfig.Provider){
		// add custom config functions
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
