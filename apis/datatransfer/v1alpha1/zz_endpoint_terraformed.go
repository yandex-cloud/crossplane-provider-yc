/*
Copyright 2022 YANDEX LLC
This is modified version of the software, made by the Crossplane Authors
and available at: https://github.com/crossplane-contrib/provider-jet-template

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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	"dario.cat/mergo"
	"github.com/pkg/errors"

	"github.com/crossplane/upjet/pkg/resource"
	"github.com/crossplane/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this Endpoint
func (mg *Endpoint) GetTerraformResourceType() string {
	return "yandex_datatransfer_endpoint"
}

// GetConnectionDetailsMapping for this Endpoint
func (tr *Endpoint) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"settings[*].clickhouse_source[*].connection[*].connection_options[*].password[*].raw": "spec.forProvider.settings[*].clickhouseSource[*].connection[*].connectionOptions[*].password[*].rawSecretRef", "settings[*].clickhouse_target[*].connection[*].connection_options[*].password[*].raw": "spec.forProvider.settings[*].clickhouseTarget[*].connection[*].connectionOptions[*].password[*].rawSecretRef", "settings[*].kafka_source[*].auth[*].sasl[*].password[*].raw": "spec.forProvider.settings[*].kafkaSource[*].auth[*].sasl[*].password[*].rawSecretRef", "settings[*].kafka_target[*].auth[*].sasl[*].password[*].raw": "spec.forProvider.settings[*].kafkaTarget[*].auth[*].sasl[*].password[*].rawSecretRef", "settings[*].mongo_source[*].connection[*].connection_options[*].password[*].raw": "spec.forProvider.settings[*].mongoSource[*].connection[*].connectionOptions[*].password[*].rawSecretRef", "settings[*].mongo_target[*].connection[*].connection_options[*].password[*].raw": "spec.forProvider.settings[*].mongoTarget[*].connection[*].connectionOptions[*].password[*].rawSecretRef", "settings[*].mysql_source[*].password[*].raw": "spec.forProvider.settings[*].mysqlSource[*].password[*].rawSecretRef", "settings[*].mysql_target[*].password[*].raw": "spec.forProvider.settings[*].mysqlTarget[*].password[*].rawSecretRef", "settings[*].postgres_source[*].password[*].raw": "spec.forProvider.settings[*].postgresSource[*].password[*].rawSecretRef", "settings[*].postgres_target[*].password[*].raw": "spec.forProvider.settings[*].postgresTarget[*].password[*].rawSecretRef", "settings[*].ydb_source[*].sa_key_content": "spec.forProvider.settings[*].ydbSource[*].saKeyContentSecretRef", "settings[*].ydb_target[*].sa_key_content": "spec.forProvider.settings[*].ydbTarget[*].saKeyContentSecretRef"}
}

// GetObservation of this Endpoint
func (tr *Endpoint) GetObservation() (map[string]any, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Endpoint
func (tr *Endpoint) SetObservation(obs map[string]any) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Endpoint
func (tr *Endpoint) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Endpoint
func (tr *Endpoint) GetParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Endpoint
func (tr *Endpoint) SetParameters(params map[string]any) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// GetInitParameters of this Endpoint
func (tr *Endpoint) GetInitParameters() (map[string]any, error) {
	p, err := json.TFParser.Marshal(tr.Spec.InitProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]any{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// GetInitParameters of this Endpoint
func (tr *Endpoint) GetMergedParameters(shouldMergeInitProvider bool) (map[string]any, error) {
	params, err := tr.GetParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get parameters for resource '%q'", tr.GetName())
	}
	if !shouldMergeInitProvider {
		return params, nil
	}

	initParams, err := tr.GetInitParameters()
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get init parameters for resource '%q'", tr.GetName())
	}

	// Note(lsviben): mergo.WithSliceDeepCopy is needed to merge the
	// slices from the initProvider to forProvider. As it also sets
	// overwrite to true, we need to set it back to false, we don't
	// want to overwrite the forProvider fields with the initProvider
	// fields.
	err = mergo.Merge(&params, initParams, mergo.WithSliceDeepCopy, func(c *mergo.Config) {
		c.Overwrite = false
	})
	if err != nil {
		return nil, errors.Wrapf(err, "cannot merge spec.initProvider and spec.forProvider parameters for resource '%q'", tr.GetName())
	}

	return params, nil
}

// LateInitialize this Endpoint using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Endpoint) LateInitialize(attrs []byte) (bool, error) {
	params := &EndpointParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Endpoint) GetTerraformSchemaVersion() int {
	return 1
}