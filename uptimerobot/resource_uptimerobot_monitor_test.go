package uptimerobot

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	uptimerobotapi "github.com/louy/terraform-provider-uptimerobot/uptimerobot/api"
)

func TestUptimeRobotDataResourceMonitor_http_monitor(t *testing.T) {
	var FriendlyName = "TF Test: http monitor"
	var Type = "http"
	var URL = "https://google.com"
	var URL2 = "https://yahoo.com"
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
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
				}
				`, FriendlyName, Type, URL2),
				Check: resource.ComposeTestCheckFunc(
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

func TestUptimeRobotDataResourceMonitor_keyword_monitor(t *testing.T) {
	var FriendlyName = "TF Test: keyword"
	var Type = "keyword"
	var URL = "https://google.com"
	var KeywordType = "not exists"
	var KeywordType2 = "exists"
	var KeywordValue = "yahoo"
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
					keyword_type  = "%s"
					keyword_value = "%s"
				}
				`, FriendlyName, Type, URL, KeywordType, KeywordValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "keyword_type", KeywordType),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "keyword_value", KeywordValue),
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
					keyword_type  = "%s"
					keyword_value = "%s"
				}
				`, FriendlyName, Type, URL, KeywordType2, KeywordValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "keyword_type", KeywordType2),
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

func TestUptimeRobotDataResourceMonitor_http_port_monitor(t *testing.T) {
	var FriendlyName = "TF Test: http port monitor"
	var Type = "port"
	var URL = "google.com"
	var URL2 = "yahoo.com"
	var SubType = "http"
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
					sub_type      = "%s"
				}
				`, FriendlyName, Type, URL, SubType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "sub_type", SubType),
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
					sub_type      = "%s"
				}
				`, FriendlyName, Type, URL2, SubType),
				Check: resource.ComposeTestCheckFunc(
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

func TestUptimeRobotDataResourceMonitor_custom_port_monitor(t *testing.T) {
	var FriendlyName = "TF Test: custom port monitor"
	var Type = "port"
	var URL = "google.com"
	var SubType = "custom"
	var Port = 8080
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
					sub_type      = "%s"
					port          = %d
				}
				`, FriendlyName, Type, URL, SubType, Port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "sub_type", SubType),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "port", fmt.Sprintf(`%d`, Port)),
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

func TestUptimeRobotDataResourceMonitor_custom_alert_contact_threshold_and_recurrence(t *testing.T) {
	var FriendlyName = "TF Test: custom alert contact threshold & recurrence"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_alert_contact" "test" {
					friendly_name = "#infrastructure"
					type = "slack"
					value = "https://slack.com/services/xxxx"
				}
				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					alert_contact {
						id         = uptimerobot_alert_contact.test.id
						threshold  = 0
						recurrence = 0
					}
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.#", "1"),
					resource.TestCheckResourceAttrSet("uptimerobot_monitor.test", "alert_contact.0.id"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.0.threshold", "0"),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.0.recurrence", "0"),
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

func TestUptimeRobotDataResourceMonitor_custom_http_headers(t *testing.T) {
	var FriendlyName = "TF Test:  custom http headers"
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
					custom_http_headers = {
						// Accept-Language = "en"
					}
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "custom_http_headers.%", "0"),
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
	var AuthType = "basic"
	var AuthType2 = "digest"
	var URL = fmt.Sprintf("https://httpbin.org/basic-auth/%s/%s", Username, Password)
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name  = "%s"
					type           = "%s"
					url            = "%s"
					http_username  = "%s"
					http_password  = "%s"
					http_auth_type = "%s"
				}
				`, FriendlyName, Type, URL, Username, Password, AuthType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_username", Username),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_password", Password),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_auth_type", AuthType),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_auth_type issue
				// ImportStateVerify: true,
			},
			resource.TestStep{
				Config: fmt.Sprintf(`
				resource "uptimerobot_monitor" "test" {
					friendly_name  = "%s"
					type           = "%s"
					url            = "%s"
					http_username  = "%s"
					http_password  = "%s"
					http_auth_type = "%s"
				}
				`, FriendlyName, Type, URL, Username, Password, AuthType2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_username", Username),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_password", Password),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "http_auth_type", AuthType2),
				),
			},
			resource.TestStep{
				ResourceName: "uptimerobot_monitor.test",
				ImportState:  true,
				// NB: Disabled due to http_auth_type issue
				// ImportStateVerify: true,
			},
		},
	})
}

func TestUptimeRobotDataResourceMonitor_default_alert_contact(t *testing.T) {
	var FriendlyName = "TF Test: using the default alert contact"
	var Type = "http"
	var URL = "https://google.com"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitorDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: fmt.Sprintf(`

				data "uptimerobot_account" "account" {}

				data "uptimerobot_alert_contact" "default" {
				friendly_name = data.uptimerobot_account.account.email
				}

				resource "uptimerobot_monitor" "test" {
					friendly_name = "%s"
					type          = "%s"
					url           = "%s"
					alert_contact {
						id         = data.uptimerobot_alert_contact.default.id
					}
				}
				`, FriendlyName, Type, URL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "friendly_name", FriendlyName),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "type", Type),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "url", URL),
					resource.TestCheckResourceAttr("uptimerobot_monitor.test", "alert_contact.#", "1"),
					resource.TestCheckResourceAttrSet("uptimerobot_monitor.test", "alert_contact.0.id"),
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
