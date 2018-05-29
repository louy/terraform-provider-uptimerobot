package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitorCreate,
		Read:   resourceMonitorRead,
		Update: resourceMonitorUpdate,
		Delete: resourceMonitorDelete,

		Schema: map[string]*schema.Schema{
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"sub_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// required for port monitoring
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// required for port monitoring
			},
			"keyword_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// required for keyword monitoring
			},
			"keyword_value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				// required for keyword monitoring
			},
			"interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"alert_contacts": &schema.Schema{
				// FIXME - array of id,threshold,recurrence
				Type:     schema.TypeString,
				Optional: true,
			},
			// TODO - mwindows
			// TODO - custom_http_headers
			// TODO - ignore_ssl_errors
		},
	}
}

func resourceMonitorCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMonitorRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMonitorDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
