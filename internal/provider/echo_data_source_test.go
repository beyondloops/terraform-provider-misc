// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"
)

func TestAccEcho(t *testing.T) {
	// resource.Test(t, resource.TestCase{
	// 	PreCheck:                 func() { testAccPreCheck(t) },
	// 	ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
	// 	Steps: []resource.TestStep{
	// 		// Read testing
	// 		{
	// 			Config: testAccExampleDataSourceConfig,
	// 			ConfigStateChecks: []statecheck.StateCheck{
	// 				statecheck.ExpectKnownValue(
	// 					"data.misc_echo.test",
	// 					tfjsonpath.New("input"),
	// 					knownvalue.StringExact("example-input"),
	// 				),
	// 			},
	// 		},
	// 	},
	// })
}
