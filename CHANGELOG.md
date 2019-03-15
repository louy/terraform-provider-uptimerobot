# v0.3.0
- resource monitor: fix support for `recurrence` and `threshold` in `alert_contact` ([#18](https://github.com/louy/terraform-provider-uptimerobot/pull/18) by [@drubin](https://github.com/drubin))
- make release builds more portable by adding the `CGO_ENABLED=0` flag ([#19](https://github.com/louy/terraform-provider-uptimerobot/pull/19) by [@Novex](https://github.com/Novex))

# v0.2.1
- resource monitor: fix broken `port` monitor type ([#17](https://github.com/louy/terraform-provider-uptimerobot/pull/17) by [@louy](https://github.com/louy))

# v0.2.0
- resource monitor: add support for interval attribute ([#7](https://github.com/louy/terraform-provider-uptimerobot/pull/7) by [@nhamlh](https://github.com/nhamlh))
- resource monitor: remove ForceNew on url ([#12](https://github.com/louy/terraform-provider-uptimerobot/pull/12) by [@drubin](https://github.com/drubin))
- various documentation updates ([#8](https://github.com/louy/terraform-provider-uptimerobot/pull/8) and [#10](https://github.com/louy/terraform-provider-uptimerobot/pull/10) by [@drubin](https://github.com/drubin))

# v0.1.1
- Fix a bug in resourceMonitorUpdate ([#5](https://github.com/louy/terraform-provider-uptimerobot/pull/5) by [@nhamlh](https://github.com/nhamlh))
- Fix a crash in resourceMonitorGet when the monitor doesn't exist ([#6](https://github.com/louy/terraform-provider-uptimerobot/pull/6) by [@louy](https://github.com/louy))

# v0.1.0
- initial release (by [@louy](https://github.com/louy))
