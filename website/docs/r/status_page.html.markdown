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
  friendly_name  = "My Status Page"
  custom_domain  = "status.example.com"
  password       = "WeAreAwsome"
  sort_monitors  = "down-up-paused"
  monitors       = [uptimerobot_monitor.main.id]
}
```

## Arguments Reference

* `friendly_name` - friendly name of the monitor (for making it easier to distinguish from others).
* `monitors` - The monitors to be displayed. Use `[0]` for all monitors (default).
* `custom_domain` - (optional) the domain or subdomain that the status page will run on.
* `password` - (optional) the password for the status page.
* `sort` - (optional) the sorting of the status page. Can be one of the following:
  - `a-z`
  - `z-a`
  - `up-down-paused`
  - `down-up-paused`
* `status` - the status of the status page (`paused` or `active`). Defaults to `active`

## Attributes Reference

* `id` - the ID of the status page
* `dns_address` - the dns address that you need to point your custom domain to (`stats.uptimerobot.com`)
* `standard_url` - the full url of the page on uptimerobot.com
* `custom_url` - the full url of the page (only if `custom_domain` is set)
