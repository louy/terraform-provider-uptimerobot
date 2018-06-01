---
layout: "uptimerobot"
page_title: "UptimeRobot: uptimerobot_monitor"
sidebar_current: "docs-uptimerobot-resource-monitor"
description: |-
  Set up a monitor
---

# Resource: uptimerobot_monitor

Use this resource to create a monitor

## Example Usage

```hcl
resource "uptimerobot_monitor" "my_website" {
  friendly_name = "My Monitor"
  type          = "http"
  url           = "http://example.com"
}
```

## Attributes Reference

TODO
