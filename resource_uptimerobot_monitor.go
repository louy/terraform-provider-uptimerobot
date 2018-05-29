package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

var monitorTypes = map[string]int{
	"http":    1,
	"keyword": 2,
	"ping":    3,
	"port":    4,
}

var monitorSubTypes = map[string]int{
	"http":   1,
	"https":  2,
	"ftp":    3,
	"smtp":   4,
	"pop3":   5,
	"imap":   6,
	"custom": 99,
}

var monitorStatuses = map[string]int{
	"paused":          0,
	"not checked yet": 1,
	"up":              2,
	"seems down":      8,
	"down":            9,
}

var monitorKeywordTypes = map[string]int{
	"exists":     1,
	"not exists": 2,
}

func resourceMonitor() *schema.Resource {
	return &schema.Resource{
		Create: resourceMonitorCreate,
		Read:   resourceMonitorRead,
		Update: resourceMonitorUpdate,
		Delete: resourceMonitorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"friendly_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"url": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice(mapKeys(monitorTypes), false),
			},
			"sub_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(mapKeys(monitorSubTypes), false),
				// required for port monitoring
			},
			"port": {
				Type:     schema.TypeString,
				Optional: true,
				// required for port monitoring
			},
			"keyword_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(mapKeys(monitorKeywordTypes), false),
				// required for keyword monitoring
			},
			"keyword_value": {
				Type:     schema.TypeString,
				Optional: true,
				// required for keyword monitoring
			},
			"interval": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"http_username": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_password": {
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"alert_contact": {
				Type:     schema.TypeList,
				Optional: true,
				// PromoteSingle: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"threshold": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"recurrence": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			// TODO - mwindows
			// TODO - custom_http_headers
			// TODO - ignore_ssl_errors
		},
	}
}

func resourceMonitorCreate(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("friendly_name", d.Get("friendly_name").(string))
	data.Add("url", d.Get("url").(string))
	t := d.Get("type").(string)
	data.Add("type", fmt.Sprintf("%d", monitorTypes[t]))
	switch t {
	case "port":
		data.Add("sub_type", fmt.Sprintf("%d", monitorSubTypes[d.Get("sub_type").(string)]))
		data.Add("port", fmt.Sprintf("%d", d.Get("port").(int)))
		break
	case "keyword":
		data.Add("keyword_type", fmt.Sprintf("%d", monitorKeywordTypes[d.Get("keyword_type").(string)]))
		data.Add("keyword_value", d.Get("keyword_value").(string))

		data.Add("http_username", d.Get("http_username").(string))
		data.Add("http_password", d.Get("http_password").(string))
		break
	case "http":
		data.Add("http_username", d.Get("http_username").(string))
		data.Add("http_password", d.Get("http_password").(string))
		break
	}
	acs := Map(d.Get("alert_contact").([]interface{}), func(ac interface{}) string {
		a := ac.(map[string]interface{})
		return fmt.Sprintf("%d", a["id"].(int))
	})
	data.Add("alert_contacts", strings.Join(acs, "-"))

	body, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"newMonitor",
		data.Encode(),
	)
	if err != nil {
		return err
	}
	monitor := body["monitor"].(map[string]interface{})
	d.SetId(fmt.Sprintf("%d", int(monitor["id"].(float64))))
	d.Set("status", intToString(monitorStatuses, int(monitor["status"].(float64))))
	return nil
}

func resourceMonitorRead(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("monitors", d.Id())

	body, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"getMonitors",
		data.Encode(),
	)
	if err != nil {
		return err
	}

	monitors, ok := body["monitors"].([]interface{})
	if !ok {
		j, _ := json.Marshal(body)
		return errors.New("Unknown response from the server: " + string(j))
	}

	monitor := monitors[0].(map[string]interface{})

	d.Set("friendly_name", monitor["friendly_name"].(string))
	d.Set("url", monitor["url"].(string))
	t := intToString(monitorTypes, int(monitor["type"].(float64)))
	d.Set("type", t)
	d.Set("status", intToString(monitorStatuses, int(monitor["status"].(float64))))
	d.Set("interval", int(monitor["interval"].(float64)))

	switch t {
	case "port":
		d.Set("sub_type", monitor["sub_type"].(string))
		d.Set("port", int(monitor["port"].(float64)))
		break
	case "keyword":
		d.Set("keyword_type", intToString(monitorKeywordTypes, int(monitor["keyword_type"].(float64))))
		d.Set("keyword_value", monitor["keyword_value"].(string))

		d.Set("http_username", monitor["http_username"].(string))
		d.Set("http_password", monitor["http_password"].(string))
		break
	case "http":
		d.Set("http_username", monitor["http_username"].(string))
		d.Set("http_password", monitor["http_password"].(string))
		break
	}

	return nil
}

func resourceMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("id", d.Id())
	data.Add("friendly_name", d.Get("friendly_name").(string))
	data.Add("url", d.Get("url").(string))
	acs := Map(d.Get("alert_contact").([]interface{}), func(ac interface{}) string {
		a := ac.(map[string]interface{})
		return fmt.Sprintf("%d", a["id"].(int))
	})
	data.Add("alert_contacts", strings.Join(acs, "-"))

	_, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"editMonitor",
		data.Encode(),
	)
	if err != nil {
		return err
	}
	return nil
}

func resourceMonitorDelete(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("id", d.Id())

	_, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"deleteMonitor",
		data.Encode(),
	)
	if err != nil {
		return err
	}
	return nil
}
