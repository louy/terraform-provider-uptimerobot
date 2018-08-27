package uptimerobot

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/louy/terraform-provider-uptimerobot/uptimerobot/api"
)

func TestUptimeRobotDataResourceMonitor_http_monitor(t *testing.T) {
	var FriendlyName = "TF Test: http monitor"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
				),
			},
			resource.TestStep{
				ResourceName:      "uptimerobot_monitor.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_change_url(t *testing.T) {
	var FriendlyName = "TF Test: http monitor"
	var Type = "http"
	var URL = "https://google.com"
	var URL2 = "https://google.co.uk"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
				),
			},
			resource.TestStep{
				ResourceName:      "uptimerobot_monitor.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Change url
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
				}
				`, FriendlyName, Type, URL2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL2),
				),
			},
			resource.TestStep{
				ResourceName:      "uptimerobot_monitor.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_ping_monitor(t *testing.T) {
	var FriendlyName = "TF Test: ping monitor"
	var Type = "ping"
	var URL = "1.1.1.1"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
				),
			},
			resource.TestStep{
				ResourceName:      "uptimerobot_monitor.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_custom_interval(t *testing.T) {
	var FriendlyName = "TF Test: ping monitor"
	var Type = "ping"
	var URL = "1.1.1.1"
	var Interval = 300
	var Interval2 = 360
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					interval      = %d
				}
				`, FriendlyName, Type, URL, Interval),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "interval", fmt.Sprintf(`%d`, Interval)),
				),
			},
			resource.TestStep{
				ResourceName:      "uptimerobot_monitor.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					interval      = %d
				}
				`, FriendlyName, Type, URL, Interval2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "interval", fmt.Sprintf(`%d`, Interval2)),
				),
			},
			resource.TestStep{
				ResourceName:      "uptimerobot_monitor.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_http_auth_monitor(t *testing.T) {
	var FriendlyName = "TF Test: http auth monitor"
	var Type = "http"
	var Username = "tester"
	var Password = "secret"
	var URL = fmt.Sprintf("https://httpbin.org/basic-auth/%s/%s", Username, Password)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					http_username = "%s"
					http_password = "%s"
				}
				`, FriendlyName, Type, URL, Username, Password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_username", Username),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_password", Password),
				),
			},
			resource.TestStep{
				ResourceName:      "uptimerobot_monitor.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckMonitorDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(uptimerobotapi.UptimeRobotApiClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "uptimerobot_monitor" {
			continue
		}

		id, err := strconv.Atoi(rs.Primary.ID)
		if err != nil {
			return err
		}

		_, err = client.GetMonitor(id)

		if err == nil {
			return fmt.Errorf("Monitor still exists")
		}

		// Verify the error is what we want
		if strings.Contains(err.Error(), "test") {
			return err
		}
	}

	return nil
}
