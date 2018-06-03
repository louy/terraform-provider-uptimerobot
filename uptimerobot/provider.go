package uptimerobot

/*
	Usage:
	```
	provider "uptimerobot" {
	  api_key = "[YOUR MAIN API KEY]"
	}
	```
*/

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

type UptimeRobotConfig struct {
	apiKey string
}

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("UPTIMEROBOT_API_KEY", nil),
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"uptimerobot_account": dataSourceAccount(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"uptimerobot_alert_contact": resourceAlertContact(),
			"uptimerobot_monitor":       resourceMonitor(),
			"uptimerobot_status_page":   resourceStatusPage(),
		},
		ConfigureFunc: func(r *schema.ResourceData) (interface{}, error) {
			config := UptimeRobotConfig{
				apiKey: r.Get("api_key").(string),
			}
			return config, nil
		},
	}
}
