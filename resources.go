// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rabbitmq

import (
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/tokens"
	"github.com/terraform-providers/terraform-provider-rabbitmq/rabbitmq"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "rabbitmq"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := rabbitmq.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "rabbitmq",
		Description: "A Pulumi package for creating and managing rabbitmq cloud resources.",
		Keywords:    []string{"pulumi", "rabbitmq"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/pulumi/pulumi-rabbitmq",
		Config: map[string]*tfbridge.SchemaInfo{
			"endpoint": {
				Type: "string",
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"RABBITMQ_MANAGEMENT_ENDPOINT"},
				},
			},
			"username": {
				Type: "string",
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"RABBITMQ_USERNAME"},
				},
			},
			"password": {
				Type: "string",
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"RABBITMQ_PASSWORD"},
				},
			},
			"insecure": {
				Type: "boolean",
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"RABBITMQ_INSECURE"},
				},
			},
			"cacert_file": {
				Type: "string",
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"RABBITMQ_CACERT_FILE"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"rabbitmq_binding":     {Tok: makeResource(mainMod, "Binding")},
			"rabbitmq_exchange":    {Tok: makeResource(mainMod, "Exchange")},
			"rabbitmq_permissions": {Tok: makeResource(mainMod, "Permissions")},
			"rabbitmq_policy":      {Tok: makeResource(mainMod, "Policy")},
			"rabbitmq_queue":       {Tok: makeResource(mainMod, "Queue")},
			"rabbitmq_user":        {Tok: makeResource(mainMod, "User")},
			"rabbitmq_vhost":       {Tok: makeResource(mainMod, "VHost")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^1.0.0-rc.1",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=0.17.28,<2.0.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
