// Copyright 2026 EasyRedir Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccHostResource adopts an existing host with its security settings off.
// Because hosts are provisioned via DNS (not the API), the test needs a real
// host to adopt (URLLO_TEST_HOST) and is skipped when that variable is unset.
//
// The config leaves security out: the acceptance-test org's subscription
// rejects any host security field with a 422 ("Your subscription does not
// support host security options"), even when every field is false. The
// https_upgrade check instead verifies the host's existing (off) setting is read
// back correctly.
func TestAccHostResource(t *testing.T) {
	host := testAccHost(t)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccHostConfig(host),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("urllo_host.test", "name", host),
					resource.TestCheckResourceAttrSet("urllo_host.test", "id"),
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

func testAccHostConfig(host string) string {
	return fmt.Sprintf(`
resource "urllo_host" "test" {
  name = %q
}
`, host)
}
