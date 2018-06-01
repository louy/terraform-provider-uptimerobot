---
layout: "uptimerobot"
page_title: "UptimeRobot: uptimerobot_alert_contact"
sidebar_current: "docs-uptimerobot-resource-alert_contact"
description: |-
  Set up an alert contact
---

# Resource: uptimerobot_alert_contact

Use this resource to create an alert contact

## Example Usage

```hcl
resource "uptimerobot_alert_contact" "slack" {
  friendly_name = "Slack Alert"
  type          = "slack"
  webhook_url   = "https://hooks.slack.com/services/XXXXXXX"
}
```

## Attributes Reference

TODO
