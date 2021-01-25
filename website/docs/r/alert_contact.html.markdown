---
- layout: `uptimerobot`
- page_title: `UptimeRobot: uptimerobot_alert_contact`
- sidebar_current: `docs-uptimerobot-resource-alert-contact`
description: |-
  Set up an alert contact
---

# Resource: uptimerobot_alert_contact

Use this resource to create an alert contact

## Example Usage

```hcl
resource `uptimerobot_alert_contact` `slack` {
  friendly_name = `Slack Alert`
  type          = `slack`
  value   = `https://hooks.slack.com/services/XXXXXXX`
}
```

## Arguments Reference

* `friendly_name` - friendly name of the alert contact (for making it easier to distinguish from others).
* `type` - the type of the alert contact notified (Zapier, HipChat and Slack are not supported in the api yet)

  Possible values are the following:
  - `sms`
  - `email`
  - `twitter-dm`
  - `boxcar`
  - `webhook`
  - `pushbullet`
  - `zapier`
  - `pushover`
  - `hipchat`
  - `slack`
  - `telegram`
  - `hangouts`
* `value` - alert contact's address/phone/url

## Attributes Reference

* `id` - the ID of the alert contact.
* `status` - the status of the alert contact (`not activated`, `paused` or `active`)
