# v0.4.1
- Switch to [terraform plugin sdk](https://www.terraform.io/docs/extend/plugin-sdk.html) ([#60](https://github.com/louy/terraform-provider-uptimerobot/pull/60) by [@nlamirault](https://github.com/nlamirault))

# v0.4.0
- BREAKING: Terraform 0.12 support ([#58](https://github.com/louy/terraform-provider-uptimerobot/pull/58) by [@caarlos0](https://github.com/caarlos0), [#59](https://github.com/louy/terraform-provider-uptimerobot/pull/59) by [@aequitas](https://github.com/aequitas))

# v0.3.2
- data source alert_contact: Allow default alert contact to be configured ([#21](https://github.com/louy/terraform-provider-uptimerobot/pull/21) by [@aequitas](https://github.com/aequitas))
- resource monitor: Handle nil value in alert contact ([#28](https://github.com/louy/terraform-provider-uptimerobot/pull/28) by [@louy](https://github.com/louy))

# v0.3.1
- resource monitor: add support for `custom_http_headers` ([#20](https://github.com/louy/terraform-provider-uptimerobot/pull/20) by [@leeif](https://github.com/leeif))

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
