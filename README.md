# Terraform UptimeRobot Provider

## Getting started

```tf

provider "uptime-robot" {
  api_key = "[YOUR MAIN API KEY]"
}

resource "uptime_robot_monitor" "main" {
  title     = "My Monitor"
  type      = "http"
  endpoint  = "http://example.com"
  frequency = 300
}

resource "uptime_robot_status_page" "main" {
  title         = "My Status Page"
  custom_domain = "status.example.com"
  password      = "WeAreAwsome"
  sort_monitors = "down-up-paused"
  monitors      = ["${resource.uptime_robot_monitor.main.id}"]
  hide_logo     = false # pro only
}

resource "aws_route53_record" {
  zone_id = "[MY ZONE ID]"
  type    = "CNAME"
  records = ["${resource.uptime_robot_status_page.main.id}"]
}

```
