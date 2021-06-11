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

* `friendly_name` - friendly name of the monitor (for making it easier to distinguish from others).
* `url` - the URL/IP of the monitor.
* `type` - the type of the monitor. Can be one of the following:
  - *`http`*
  - *`keyword`* - will also enable the following options:
    - `keyword_type` - if the monitor will be flagged as down when the keyword exists or not exists. Can be one of the following:
      - `exists`
      - `not exists`
    - `keyword_value` - the value of the keyword.
  - *`ping`*
  - *`port`* - will also enable the following options:
    - `sub_type` - which pre-defined port/service is monitored or if a custom port is monitored. Can be one of the following:
      - `http`
      - `https`
      - `ftp`
      - `smtp`
      - `pop3`
      - `imap`
      - `custom`
    - `port` - the port monitored (only if subtype is `custom`)
* `http_method` - the HTTP method to be used. Available for HTTP and keyword monitoring.
* `http_username` - used for password-protected web pages (HTTP basic or digest). Available for HTTP and keyword monitoring.
* `http_password` - used for password-protected web pages (HTTP basic or digest). Available for HTTP and keyword monitoring.
* `http_auth_type` - Used for password-protected web pages (HTTP basic or digest). Available for HTTP and keyword monitoring. Can be one of the following:
  - `basic`
  - `digest`
* `interval` - the interval for the monitoring check (300 seconds by default).

## Attributes Reference

* `id` - the ID of the monitor (can be used for monitor-specific requests)
* `status` - the status of the monitor (`paused`, `not checked yet`, `up`, `seems down`, or `down`)
