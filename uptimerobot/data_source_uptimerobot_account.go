package uptimerobot

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/louy/terraform-provider-uptimerobot/uptimerobot/api"
)

func dataSourceAccount() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceAccountRead,
		Schema: map[string]*schema.Schema{
			"email":            {Computed: true, Type: schema.TypeString},
			"monitor_limit":    {Computed: true, Type: schema.TypeInt},
			"monitor_interval": {Computed: true, Type: schema.TypeInt},
			"up_monitors":      {Computed: true, Type: schema.TypeInt},
			"down_monitors":    {Computed: true, Type: schema.TypeInt},
			"paused_monitors":  {Computed: true, Type: schema.TypeInt},
		},
	}
}

func dataSourceAccountRead(d *schema.ResourceData, m interface{}) error {
	account, err := m.(uptimerobotapi.UptimeRobotApiClient).GetAccountDetails()
	if err != nil {
		return err
	}

	d.SetId(time.Now().UTC().String())

	d.Set("email", account.Email)
	d.Set("monitor_limit", account.MonitorLimit)
	d.Set("monitor_interval", account.MonitorInterval)
	d.Set("up_monitors", account.UpMonitors)
	d.Set("down_monitors", account.DownMonitors)
	d.Set("paused_monitors", account.PausedMonitors)

	return nil
}
