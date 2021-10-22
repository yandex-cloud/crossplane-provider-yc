package main

import (
	"github.com/crossplane-contrib/provider-tf-template/config"
	"github.com/crossplane-contrib/terrajet/pkg/pipeline"
)

func main() {
	pipeline.Run(config.GetProvider())
}
