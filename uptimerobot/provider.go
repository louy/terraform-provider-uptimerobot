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
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/realiotech/terraform-provider-uptimerobot/uptimerobot/api"
)

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
			"uptimerobot_account":       dataSourceAccount(),
			"uptimerobot_alert_contact": dataSourceAlertContact(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"uptimerobot_alert_contact": resourceAlertContact(),
			"uptimerobot_monitor":       resourceMonitor(),
			"uptimerobot_status_page":   resourceStatusPage(),
		},
		ConfigureFunc: func(r *schema.ResourceData) (interface{}, error) {
			config := uptimerobotapi.New(r.Get("api_key").(string))
			return config, nil
		},
	}
}
