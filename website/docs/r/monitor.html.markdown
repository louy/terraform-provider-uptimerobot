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

## Arguments Reference

TODO

## Attributes Reference

* `id` - the ID of the monitor (can be used for monitor-specific requests)
* `status` - the status of the monitor (`paused`, `not checked yet`, `up`, `seems down`, or `down`)

TODO
