---
layout: "uptimerobot"
page_title: "Provider: UptimeRobot"
sidebar_current: "docs-uptimerobot-index"
description: |-
  The UptimeRobot provider is used to interact with the resources supported by UptimeRobot. The provider needs to be configured with the proper credentials before it can be used.
---

# UptimeRobot Provider
The UptimeRobot provider is used to interact with the resources supported by UptimeRobot. The provider needs to be configured with the proper credentials before it can be used.
Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the UptimeRobot Provider
provider "uptimerobot" {
  api_key = "${var.uptimerobot_api_key}"
}

# Create a monitor
resource "uptimerobot_monitor" "web" {
  friendly_name = "My Monitor"
  type          = "http"
  url           = "http://example.com"
}
```

## Authentication
The UptimeRobot provider needs an account-specific (main) api key to work. You can find that key for your account in the [My Settings](https://uptimerobot.com/dashboard#mySettings) page on UptimeRobot's website.

## Argument Reference

The following arguments are supported in the `provider` block:

* `api_key` - (Required) UptimeRobot's account-specific api key.
