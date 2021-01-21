---
layout: "uptimerobot"
page_title: "UptimeRobot: uptimerobot_account"
sidebar_current: "docs-uptimerobot-datasource-account"
description: |-
  Get information about your account
---

# Data Source: uptimerobot_account

Use this data source to get information about the current UptimeRobot account.

## Example Usage

```hcl
data "uptimerobot_account" "main" {}
```

## Attributes Reference

 * `email` - the account e-mail
 * `monitor_limit` - the max number of monitors that can be created for the account
 * `monitor_interval` - the min monitoring interval (in seconds) supported by the account
 * `up_monitors` - the number of "up" monitors
 * `down_monitors` - the number of "down" monitors
 * `paused_monitors` - the number of "paused" monitors
