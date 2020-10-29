package uptimerobot

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	uptimerobotapi "github.com/louy/terraform-provider-uptimerobot/uptimerobot/api"
)

func dataSourceMonitor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceMonitorRead,
		Schema: map[string]*schema.Schema{
			"id":                  {Type: schema.TypeInt, Computed: true},
			"friendly_name":       {Type: schema.TypeString, Required: true},
			"url":                 {Type: schema.TypeString, Computed: true},
			"type":                {Type: schema.TypeString, Computed: true},
			"status":              {Type: schema.TypeString, Computed: true},
			"interval":            {Type: schema.TypeInt, Computed: true},
			"sub_type":            {Type: schema.TypeString, Computed: true},
			"port":                {Type: schema.TypeInt, Computed: true},
			"keyword_type":        {Type: schema.TypeString, Computed: true},
			"keyword_value":       {Type: schema.TypeString, Computed: true},
			"http_username":       {Type: schema.TypeString, Computed: true},
			"http_password":       {Type: schema.TypeString, Computed: true},
			"http_auth_type":      {Type: schema.TypeString, Computed: true},
			"custom_http_headers": {Type: schema.TypeMap, Computed: true},
			"alert_contact": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id":         {Type: schema.TypeString, Computed: true},
						"threshold":  {Type: schema.TypeInt, Computed: true},
						"recurrence": {Type: schema.TypeInt, Computed: true},
					},
				},
			},
		},
	}
}

func dataSourceMonitorRead(d *schema.ResourceData, m interface{}) error {
	friendlyName := d.Get("friendly_name").(string)

	ids, err := m.(uptimerobotapi.UptimeRobotApiClient).GetMonitorIDs(friendlyName)
	if err != nil {
		return err
	}

	if len(ids) < 1 {
		return fmt.Errorf("Failed to find monitor by name %s", friendlyName)
	}

	if len(ids) > 1 {
		return fmt.Errorf("More than one monitor with name %s exists", friendlyName)
	}

	monitor, err := m.(uptimerobotapi.UptimeRobotApiClient).GetMonitor(ids[0])
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%d", monitor.ID))
	if err := updateMonitorResource(d, monitor); err != nil {
		return err
	}

	return nil
}
