package uptimerobot

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestUptimeRobotDataSourceAccount_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testUptimeRobotDataSourceAccount,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.uptimerobot_account.test", "email"),
					resource.TestCheckResourceAttrSet("data.uptimerobot_account.test", "monitor_limit"),
					resource.TestCheckResourceAttrSet("data.uptimerobot_account.test", "monitor_interval"),
					resource.TestCheckResourceAttrSet("data.uptimerobot_account.test", "up_monitors"),
					resource.TestCheckResourceAttrSet("data.uptimerobot_account.test", "down_monitors"),
					resource.TestCheckResourceAttrSet("data.uptimerobot_account.test", "paused_monitors"),
				),
			},
		},
	})
}

var testUptimeRobotDataSourceAccount = `
data "uptimerobot_account" "test" {}
`
