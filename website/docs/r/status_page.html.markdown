---
layout: "uptimerobot"
page_title: "UptimeRobot: uptimerobot_status_page"
sidebar_current: "docs-uptimerobot-resource-status-page"
description: |-
  Set up a status page
---

# Resource: uptimerobot_status_page

Use this resource to create a status page

## Example Usage

```hcl
resource "uptimerobot_status_page" "my_status_page" {
  friendly_name = "My Status Page"
  custom_domain = "status.example.com"
  password      = "WeAreAwsome"
  sort_monitors = "down-up-paused"
  monitors      = ["${resource.uptimerobot_monitor.main.id}"]
  hide_logo     = false # pro only
}
```

## Attributes Reference

TODO
