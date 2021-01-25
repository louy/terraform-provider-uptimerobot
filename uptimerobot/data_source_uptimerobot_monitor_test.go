package uptimerobot

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestUptimeRobotDataSourceMonitor_http_monitor(t *testing.T) {
	friendlyName := "TF Test: http monitor"
	monitorType := "http"
	monitorURL := "https://google.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUptimeRobotMonitorHTTPMonitor(friendlyName, monitorType, monitorURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet("data.uptimerobot_monitor.test", "id"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "url", monitorURL),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "type", monitorType),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "status", "not checked yet"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "interval", "300"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "port", "0"),
				),
			},
		},
	})
}

func TestUptimeRobotDataSourceMonitor_keyword_monitor(t *testing.T) {
	friendlyName := "TF Test: keyword"
	monitorType := "keyword"
	monitorURL := "https://google.com"
	keywordType := "not exists"
	keywordValue := "google"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUptimeRobotMonitorKeywordMonitor(friendlyName, monitorType, monitorURL, keywordType, keywordValue),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet("data.uptimerobot_monitor.test", "id"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "url", monitorURL),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "type", monitorType),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "status", "not checked yet"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "interval", "300"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "port", "0"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "keyword_type", keywordType),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "keyword_value", keywordValue),
				),
			},
		},
	})
}

func TestUptimeRobotDataSourceMonitor_http_port_monitor(t *testing.T) {
	friendlyName := "TF Test: http port monitor"
	monitorType := "port"
	monitorURL := "google.com"
	subType := "http"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUptimeRobotMonitorHTTPPortMonitor(friendlyName, monitorType, monitorURL, subType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet("data.uptimerobot_monitor.test", "id"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "url", monitorURL),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "type", monitorType),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "status", "not checked yet"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "interval", "300"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "port", "0"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "sub_type", subType),
				),
			},
		},
	})
}

func TestUptimeRobotDataSourceMonitor_custom_port_monitor(t *testing.T) {
	friendlyName := "TF Test: custom port monitor"
	monitorType := "port"
	monitorURL := "google.com"
	subType := "custom"
	port := 8080

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUptimeRobotMonitorCustomPortMonitor(friendlyName, monitorType, monitorURL, subType, port),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet("data.uptimerobot_monitor.test", "id"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "url", monitorURL),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "type", monitorType),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "status", "not checked yet"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "interval", "300"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "port", fmt.Sprintf("%d", port)),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "sub_type", subType),
				),
			},
		},
	})
}

func TestUptimeRobotDataSourceMonitor_alert_contact(t *testing.T) {
	friendlyName := "TF Test: alert contact"
	monitorType := "http"
	monitorURL := "https://google.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUptimeRobotMonitorAlertContact(friendlyName, monitorType, monitorURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet("data.uptimerobot_monitor.test", "id"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "url", monitorURL),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "type", monitorType),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "status", "not checked yet"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "interval", "300"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "port", "0"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "alert_contact.#", "2"),
					resource.TestCheckResourceAttrSet("data.uptimerobot_monitor.test", "alert_contact.0.id"),
					resource.TestCheckResourceAttrSet("data.uptimerobot_monitor.test", "alert_contact.0.recurrence"),
					resource.TestCheckResourceAttrSet("data.uptimerobot_monitor.test", "alert_contact.0.threshold"),
				),
			},
		},
	})
}

func TestUptimeRobotDataSourceMonitor_custom_http_headers(t *testing.T) {
	friendlyName := "TF Test: custom http headers"
	monitorType := "http"
	monitorURL := "https://google.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUptimeRobotMonitorCustomHTTPHeaders(friendlyName, monitorType, monitorURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet("data.uptimerobot_monitor.test", "id"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "url", monitorURL),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "type", monitorType),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "status", "not checked yet"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "interval", "300"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "port", "0"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "custom_http_headers.%", "0"),
				),
			},
		},
	})
}

func TestUptimeRobotDataSourceMonitor_http_auth_monitor(t *testing.T) {
	friendlyName := "TF Test: http auth monitor"
	monitorType := "http"
	username := "tester"
	password := "secret"
	authType := "basic"
	monitorURL := fmt.Sprintf("https://httpbin.org/basic-auth/%s/%s", username, password)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccUptimeRobotMonitorHTTPAuthMonitor(friendlyName, monitorType, monitorURL, username, password, authType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "friendly_name", friendlyName),
					resource.TestCheckResourceAttrSet("data.uptimerobot_monitor.test", "id"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "url", monitorURL),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "type", monitorType),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "status", "not checked yet"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "interval", "300"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "port", "0"),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "http_username", username),
					resource.TestCheckResourceAttr("data.uptimerobot_monitor.test", "http_password", password),
				),
			},
		},
	})
}

func testAccUptimeRobotMonitorHTTPMonitor(friendlyName, monitorType, monitorURL string) string {
	return fmt.Sprintf(`
resource "uptimerobot_monitor" "monitor" {
	friendly_name = "%s"
	type          = "%s"
	url           = "%s"
}

data "uptimerobot_monitor" "test" {
	friendly_name = uptimerobot_monitor.monitor.friendly_name
}
`, friendlyName, monitorType, monitorURL)
}

func testAccUptimeRobotMonitorKeywordMonitor(friendlyName, monitorType, monitorURL, keywordType, keywordValue string) string {
	return fmt.Sprintf(`
resource "uptimerobot_monitor" "monitor" {
	friendly_name = "%s"
	type          = "%s"
	url           = "%s"
	keyword_type  = "%s"
	keyword_value = "%s"
}

data "uptimerobot_monitor" "test" {
	friendly_name = uptimerobot_monitor.monitor.friendly_name
}
`, friendlyName, monitorType, monitorURL, keywordType, keywordValue)
}

func testAccUptimeRobotMonitorHTTPPortMonitor(friendlyName, monitorType, monitorURL, subType string) string {
	return fmt.Sprintf(`
resource "uptimerobot_monitor" "monitor" {
	friendly_name = "%s"
	type          = "%s"
	url           = "%s"
	sub_type      = "%s"
}

data "uptimerobot_monitor" "test" {
	friendly_name = uptimerobot_monitor.monitor.friendly_name
}
`, friendlyName, monitorType, monitorURL, subType)
}

func testAccUptimeRobotMonitorCustomPortMonitor(friendlyName, monitorType, monitorURL, subType string, port int) string {
	return fmt.Sprintf(`
resource "uptimerobot_monitor" "monitor" {
	friendly_name = "%s"
	type          = "%s"
	url           = "%s"
	sub_type      = "%s"
	port          = %d
}

data "uptimerobot_monitor" "test" {
	friendly_name = uptimerobot_monitor.monitor.friendly_name
}
`, friendlyName, monitorType, monitorURL, subType, port)
}

func testAccUptimeRobotMonitorAlertContact(friendlyName, monitorType, monitorURL string) string {
	return fmt.Sprintf(`
resource "uptimerobot_alert_contact" "slack" {
	friendly_name = "TF Test: Slack"
	type          = "slack"
	value         = "https://slack.com/services/xxxx"
}

resource "uptimerobot_alert_contact" "email" {
	friendly_name = "TF Test: Email"
	type          = "email"
	value         = "tf-test@example.com"
}

resource "uptimerobot_monitor" "monitor" {
	friendly_name = "%s"
	type          = "%s"
	url           = "%s"

	alert_contact {
		id = uptimerobot_alert_contact.slack.id
	}

	alert_contact {
		id = uptimerobot_alert_contact.email.id
	}
}

data "uptimerobot_monitor" "test" {
	friendly_name = uptimerobot_monitor.monitor.friendly_name
}
`, friendlyName, monitorType, monitorURL)
}

func testAccUptimeRobotMonitorCustomHTTPHeaders(friendlyName, monitorType, monitorURL string) string {
	return fmt.Sprintf(`
resource "uptimerobot_monitor" "monitor" {
	friendly_name = "%s"
	type          = "%s"
	url           = "%s"

	custom_http_headers = {
		// Accept-Language = "en"
	}
}

data "uptimerobot_monitor" "test" {
	friendly_name = uptimerobot_monitor.monitor.friendly_name
}
`, friendlyName, monitorType, monitorURL)
}

func testAccUptimeRobotMonitorHTTPAuthMonitor(friendlyName, monitorType, monitorURL, username, password, authType string) string {
	return fmt.Sprintf(`
resource "uptimerobot_monitor" "monitor" {
	friendly_name  = "%s"
	type           = "%s"
	url            = "%s"
	http_username  = "%s"
	http_password  = "%s"
	http_auth_type = "%s"
}

data "uptimerobot_monitor" "test" {
	friendly_name = uptimerobot_monitor.monitor.friendly_name
}
`, friendlyName, monitorType, monitorURL, username, password, authType)
}
