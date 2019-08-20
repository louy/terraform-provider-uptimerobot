# Terraform UptimeRobot Provider
[![All Contributors](https://img.shields.io/badge/all_contributors-8-orange.svg?style=flat-square)](#contributors)
[![CircleCI](https://circleci.com/gh/louy/terraform-provider-uptimerobot.svg?style=svg)](https://circleci.com/gh/louy/terraform-provider-uptimerobot)

## Getting started

```tf

provider "uptimerobot" {
  api_key = "[YOUR MAIN API KEY]" # or pass via environment variable UPTIMEROBOT_API_KEY
}

data "uptimerobot_account" "account" {}

data "uptimerobot_alert_contact" "default_alert_contact" {
  friendly_name = "${data.uptimerobot_account.account.email}"
}

resource "uptimerobot_alert_contact" "slack" {
  friendly_name = "Slack Alert"
  type          = "slack"
  value         = "https://hooks.slack.com/services/XXXXXXX"
}

resource "uptimerobot_monitor" "main" {
  friendly_name = "My Monitor"
  type          = "http"
  url           = "http://example.com"
  # pro allows 60 seconds
  interval      = 300

  alert_contact {
    id = "${uptimerobot_alert_contact.slack.id}"
    # threshold  = 0  # pro only
    # recurrence = 0  # pro only
  }

  alert_contact {
    id = "${data.uptimerobot_alert_contact.default_alert_contact.id}"
  }
}

resource "uptimerobot_monitor" "custom_port" {
  url           = "doe.john.me"
  type          = "port"
  sub_type      = "custom"
  port          = 5678
  friendly_name = "Custom port"
}

resource "uptimerobot_status_page" "main" {
  friendly_name  = "My Status Page"
  custom_domain  = "status.example.com"
  password       = "WeAreAwsome"
  sort_monitors  = "down-up-paused"
  monitors       = ["${uptimerobot_monitor.main.id}"]
  hide_url_links = false # pro only
}

resource "aws_route53_record" {
  zone_id = "[MY ZONE ID]"
  type    = "CNAME"
  records = ["${uptimerobot_status_page.main.dns_address}"]
}

```

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
<table>
  <tr>
    <td align="center"><a href="https://nhamlh.space"><img src="https://avatars3.githubusercontent.com/u/11173217?v=4" width="100px;" alt="Nham Le"/><br /><sub><b>Nham Le</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=nhamlh" title="Code">💻</a></td>
    <td align="center"><a href="http://louy.alakkad.me"><img src="https://avatars3.githubusercontent.com/u/349850?v=4" width="100px;" alt="Louay Alakkad"/><br /><sub><b>Louay Alakkad</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=louy" title="Code">💻</a> <a href="#maintenance-louy" title="Maintenance">🚧</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=louy" title="Tests">⚠️</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=louy" title="Documentation">📖</a> <a href="#tool-louy" title="Tools">🔧</a></td>
    <td align="center"><a href="http://blog.smartcube.co.za"><img src="https://avatars0.githubusercontent.com/u/237513?v=4" width="100px;" alt="David Rubin"/><br /><sub><b>David Rubin</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=drubin" title="Code">💻</a> <a href="#maintenance-drubin" title="Maintenance">🚧</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=drubin" title="Tests">⚠️</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=drubin" title="Documentation">📖</a></td>
    <td align="center"><a href="https://ijohan.nl"><img src="https://avatars2.githubusercontent.com/u/365827?v=4" width="100px;" alt="Johan Bloemberg"/><br /><sub><b>Johan Bloemberg</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=aequitas" title="Code">💻</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=aequitas" title="Tests">⚠️</a> <a href="#ideas-aequitas" title="Ideas, Planning, & Feedback">🤔</a></td>
    <td align="center"><a href="https://twitch.tv/sebbity"><img src="https://avatars1.githubusercontent.com/u/564860?v=4" width="100px;" alt="Seb Patane"/><br /><sub><b>Seb Patane</b></sub></a><br /><a href="#platform-Novex" title="Packaging/porting to new platform">📦</a></td>
    <td align="center"><a href="https://github.com/leeif"><img src="https://avatars1.githubusercontent.com/u/15794005?v=4" width="100px;" alt="YIFAN LI"/><br /><sub><b>YIFAN LI</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=leeif" title="Code">💻</a> <a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=leeif" title="Tests">⚠️</a></td>
    <td align="center"><a href="https://nicolas.lamirault.xyz"><img src="https://avatars0.githubusercontent.com/u/29233?v=4" width="100px;" alt="Nicolas Lamirault"/><br /><sub><b>Nicolas Lamirault</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=nlamirault" title="Documentation">📖</a></td>
  </tr>
  <tr>
    <td align="center"><a href="https://blog.apatin.ru"><img src="https://avatars2.githubusercontent.com/u/35623?v=4" width="100px;" alt="Daniel Apatin"/><br /><sub><b>Daniel Apatin</b></sub></a><br /><a href="https://github.com/louy/terraform-provider-uptimerobot/commits?author=ad" title="Documentation">📖</a></td>
  </tr>
</table>

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!