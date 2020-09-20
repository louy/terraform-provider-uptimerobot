package uptimerobot

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/louy/terraform-provider-uptimerobot/uptimerobot/api"
)

func TestUptimeRobotDataResourceStatusPage_basic(t *testing.T) {
	var friendlyName = "TF Test: Basic"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStatusPageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_status_page" "test" {
					friendly_name = "%s"
				}
				`, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_status_page.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttr("uptimerobot_status_page.test", "monitors.#", "1"),
					resource.TestCheckResourceAttr("uptimerobot_status_page.test", "monitors.0", "0"),
				),
			},
			resource.TestStep{
				ResourceName:      "uptimerobot_status_page.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceStatusPage_update_name(t *testing.T) {
	var friendlyName = "TF Test: Update name"
	var friendlyName2 = "TF Test: Update name 2"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStatusPageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_status_page" "test" {
					friendly_name = "%s"
				}
				`, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_status_page.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttr("uptimerobot_status_page.test", "monitors.#", "1"),
					resource.TestCheckResourceAttr("uptimerobot_status_page.test", "monitors.0", "0"),
				),
			},
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_status_page" "test" {
					friendly_name = "%s"
				}
				`, friendlyName2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_status_page.test", "friendly_name", friendlyName2),
				),
			},
			resource.TestStep{
				ResourceName:      "uptimerobot_status_page.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceStatusPage_custom_monitors(t *testing.T) {
	var friendlyName = "TF Test: custom monitors"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckStatusPageDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "TF Test: status page custom monitors"
					type          = "http"
					url           = "https://google.com"
				}
				resource "uptimerobot_status_page" "test" {
					friendly_name = "%s"
					monitors      = [uptimerobot_monitor.test.id]
				}
				`, friendlyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_status_page.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttr("uptimerobot_status_page.test", "monitors.#", "1"),
					resource.TestCheckResourceAttrSet("uptimerobot_status_page.test", "monitors.0"),
				),
			},
			resource.TestStep{
				ResourceName:      "uptimerobot_status_page.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckStatusPageDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(uptimerobotapi.UptimeRobotApiClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "uptimerobot_status_page" {
			continue
		}

		id, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = client.GetStatusPage(id)

		if err == nil {
			return fmt.Errorf("Status page still exists")
		}

		// Verify the error is what we want
		if strings.Contains(err.Error(), "test") {
			return err
		}
	}

	return nil
}
