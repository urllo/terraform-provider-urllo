// Copyright 2026 EasyRedir Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
)

// envTestHost points live acceptance tests at an existing host in the urllo
// account (e.g. "your.host.here").
// The host resource test adopts it directly. Rule tests create redirects on randomized, unique paths under it.
const envTestHost = "URLLO_TEST_HOST"

// randInt returns a random integer for building unique test fixtures.
func randInt() int {
	return acctest.RandInt()
}

// testAccHost returns the existing host for live tests, skipping the test when
// it is not configured.
func testAccHost(t *testing.T) string {
	t.Helper()
	host := os.Getenv(envTestHost)
	if host == "" {
		t.Skipf("%s must be set to an existing host in your urllo account", envTestHost)
	}
	return host
}

// testAccSource returns a unique source URL: a random path on the test host.
func testAccSource(t *testing.T, prefix string) string {
	return fmt.Sprintf("%s/%s-%d", testAccHost(t), prefix, randInt())
}
