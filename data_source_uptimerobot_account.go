package main

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
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
	body, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"getAccountDetails",
		"",
	)

	if err != nil {
		return err
	}
	account := body["account"].(map[string]interface{})

	d.SetId(time.Now().UTC().String())

	d.Set("email", account["email"].(string))
	d.Set("monitor_limit", int(account["monitor_limit"].(float64)))
	d.Set("monitor_interval", int(account["monitor_interval"].(float64)))
	d.Set("up_monitors", int(account["up_monitors"].(float64)))
	d.Set("down_monitors", int(account["down_monitors"].(float64)))
	d.Set("paused_monitors", int(account["paused_monitors"].(float64)))

	return nil
}
