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
	"fmt"

	"github.com/crossplane/upjet/pkg/config"

	"github.com/yandex-cloud/crossplane-provider-yc/config/resourcemanager"
)

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
