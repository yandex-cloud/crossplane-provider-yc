package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/crossplane-contrib/terrajet/pkg/pipeline"

	"github.com/crossplane-contrib/provider-tf-template/config"
)

func main() {
	rootDir := os.Args[1]
	if rootDir == "" {
		panic("root directory is required to be given as argument")
	}
	absRootDir, err := filepath.Abs(rootDir)
	if err != nil {
		panic(fmt.Sprintf("cannot calculate the absolute path with %s", rootDir))
	}
	pipeline.Run(config.GetProvider(), absRootDir)
}
