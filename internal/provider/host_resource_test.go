// Copyright 2026 EasyRedir Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccHostResource adopts an existing host and toggles its settings. Because
// hosts are provisioned via DNS (not the API), the test needs a real host to
// adopt (URLLO_TEST_HOST) and is skipped when that variable is unset.
func TestAccHostResource(t *testing.T) {
	host := testAccHost(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccHostConfig(host, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("urllo_host.test", "name", host),
					resource.TestCheckResourceAttrSet("urllo_host.test", "id"),
					resource.TestCheckResourceAttr("urllo_host.test", "security.https_upgrade", "true"),
				),
			},
			{
				Config: testAccHostConfig(host, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("urllo_host.test", "security.https_upgrade", "false"),
				),
			},
			{
				ResourceName:      "urllo_host.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccHostConfig(host string, httpsUpgrade bool) string {
	return fmt.Sprintf(`
resource "urllo_host" "test" {
  name = %q
  security = {
    https_upgrade = %t
  }
}
`, host, httpsUpgrade)
}
