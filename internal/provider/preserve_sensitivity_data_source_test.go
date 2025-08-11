// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"
)

func TestAccPreserveSensitivity(t *testing.T) {
	// resource.Test(t, resource.TestCase{
	// 	PreCheck:                 func() { testAccPreCheck(t) },
	// 	ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
	// 	Steps: []resource.TestStep{
	// 		// Read testing
	// 		{
	// 			Config: testAccExampleDataSourceConfig,
	// 			ConfigStateChecks: []statecheck.StateCheck{
	// 				statecheck.ExpectKnownValue(
	// 					"data.scaffolding_example.test",
	// 					tfjsonpath.New("id"),
	// 					knownvalue.StringExact("example-id"),
	// 				),
	// 			},
	// 		},
	// 	},
	// })
}
