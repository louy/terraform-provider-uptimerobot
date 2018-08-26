# Terraform UptimeRobot Provider

## Getting started

```tf

provider "uptimerobot" {
  api_key = "[YOUR MAIN API KEY]"
}

resource "uptimerobot_alert_contact" "slack" {
  friendly_name = "Slack Alert"
  type          = "slack"
  value         = "https://hooks.slack.com/services/XXXXXXX"
}

resource "uptimerobot_monitor" "main" {
  friendly_name = "My Monitor"
  type          = "http"
  url           = "http://example.com"
  # pro allows 60 seconds
  interval      = 300

  alert_contact {
    id = "${resource.uptimerobot_alert_contact.slack.id}"
    # threshold  = 0  # pro only
    # recurrence = 0  # pro only
  }
}

resource "uptimerobot_status_page" "main" {
  friendly_name  = "My Status Page"
  custom_domain  = "status.example.com"
  password       = "WeAreAwsome"
  sort_monitors  = "down-up-paused"
  monitors       = ["${resource.uptimerobot_monitor.main.id}"]
  hide_url_links = false # pro only
}

resource "aws_route53_record" {
  zone_id = "[MY ZONE ID]"
  type    = "CNAME"
  records = ["${resource.uptimerobot_status_page.main.dns_address}"]
}

```
