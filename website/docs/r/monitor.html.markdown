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

resource "uptimerobot_monitor" "my_cron" {
  friendly_name = "My Cronjob"
  type          = "heartbeat"
  interval      = 300
}
```

## Arguments Reference

* `api_key` - (Required) uptimerobot API key
* `friendly_name` - (Required) friendly name of the monitor (for making it easier to distinguish from others).
* `url` - (Required for all types EXCEPT for the heartbeat one) the URL/IP of the monitor.
* `type` - (Required) the type of the monitor. Can be one of the following:
  - *`http`*
  - *`heartbeat`*
  - *`keyword`* - will also enable the following options:
    - `keyword_type` - if the monitor will be flagged as down when the keyword exists or not exists. Can be one of the following:
      - `exists` -  (required for keyword monitoring)
      - `not exists` -  (required for keyword monitoring)  
    - `keyword_value` -  (required for keyword monitoring) the value of the keyword.
  - *`ping`*
  - *`port`* - will also enable the following options:
    - `sub_type` - (Required for port monitoring and custom) which pre-defined port/service is monitored or if a custom port is monitored. Can be one of the following:
      - `http`
      - `https`
      - `ftp`
      - `smtp`
      - `pop3`
      - `imap`
      - `custom`
    - `port` - the port monitored (only if subtype is `custom` or `port`)
* `http_username` - (Optional) used for password-protected web pages (HTTP basic or digest). Available for HTTP and keyword monitoring.
* `http_password` - (Optional) used for password-protected web pages (HTTP basic or digest). Available for HTTP and keyword monitoring.
* `http_auth_type` - (Optional) used for password-protected web pages (HTTP basic or digest). Available for HTTP and keyword monitoring. Can be one of the following:
  - `basic`
  - `digest`
* `interval` - (Optional) the interval for the monitoring check (300 seconds by default).
* `alert_contact` - (Optional) the alert contacts to be notified when the monitor goes up/down.Multiple

## Attributes Reference

* `id` - the ID of the monitor (can be used for monitor-specific requests)
* `status` - the status of the monitor (`paused`, `not checked yet`, `up`, `seems down`, or `down`)
