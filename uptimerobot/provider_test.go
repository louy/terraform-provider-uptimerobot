package uptimerobot

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"uptimerobot": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}
func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("UPTIMEROBOT_API_KEY"); v == "" {
		t.Fatal("UPTIMEROBOT_API_KEY must be set for acceptance tests")
	}
}

func testProAccPreCheck(t *testing.T) {
	if v := os.Getenv("UPTIMEROBOT_PRO"); v == "" {
		t.Fatal("UPTIMEROBOT_PRO must be set to true for pro features acceptance tests")
	}
}
