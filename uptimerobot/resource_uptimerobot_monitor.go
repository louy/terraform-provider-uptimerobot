package uptimerobot

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	uptimerobotapi "github.com/louy/terraform-provider-uptimerobot/uptimerobot/api"
)

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
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice(uptimerobotapi.MonitorType, false),
			},
			"sub_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(uptimerobotapi.MonitorSubType, false),
				// required for port monitoring
			},
			"port": {
				Type:     schema.TypeInt,
				Optional: true,
				// required for port monitoring
			},
			"keyword_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(uptimerobotapi.MonitorKeywordType, false),
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
			"http_auth_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(uptimerobotapi.MonitorHTTPAuthType, false),
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ignore_ssl_errors": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"alert_contact": {
				Type:     schema.TypeList,
				Optional: true,
				// PromoteSingle: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
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
			"custom_http_headers": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			// TODO - mwindows
		},
	}
}

func resourceMonitorCreate(d *schema.ResourceData, m interface{}) error {
	req := uptimerobotapi.MonitorCreateRequest{
		FriendlyName: d.Get("friendly_name").(string),
		URL:          d.Get("url").(string),
		Type:         d.Get("type").(string),
	}

	switch req.Type {
	case "port":
		req.SubType = d.Get("sub_type").(string)
		req.Port = d.Get("port").(int)
		break
	case "keyword":
		req.KeywordType = d.Get("keyword_type").(string)
		req.KeywordValue = d.Get("keyword_value").(string)

		req.HTTPUsername = d.Get("http_username").(string)
		req.HTTPPassword = d.Get("http_password").(string)
		req.HTTPAuthType = d.Get("http_auth_type").(string)
		break
	case "http":
		req.HTTPUsername = d.Get("http_username").(string)
		req.HTTPPassword = d.Get("http_password").(string)
		req.HTTPAuthType = d.Get("http_auth_type").(string)
		break
	}

	// Add optional attributes
	req.Interval = d.Get("interval").(int)

	req.IgnoreSSLErrors = d.Get("ignore_ssl_errors").(bool)

	req.AlertContacts = make([]uptimerobotapi.MonitorRequestAlertContact, len(d.Get("alert_contact").([]interface{})))
	for k, v := range d.Get("alert_contact").([]interface{}) {
		req.AlertContacts[k] = uptimerobotapi.MonitorRequestAlertContact{
			ID:         v.(map[string]interface{})["id"].(string),
			Threshold:  v.(map[string]interface{})["threshold"].(int),
			Recurrence: v.(map[string]interface{})["recurrence"].(int),
		}
	}

	// custom_http_headers
	httpHeaderMap := d.Get("custom_http_headers").(map[string]interface{})
	req.CustomHTTPHeaders = make(map[string]string, len(httpHeaderMap))
	for k, v := range httpHeaderMap {
		req.CustomHTTPHeaders[k] = v.(string)
	}

	monitor, err := m.(uptimerobotapi.UptimeRobotApiClient).CreateMonitor(req)
	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%d", monitor.ID))
	if err := updateMonitorResource(d, monitor); err != nil {
		return err
	}
	return nil
}

func resourceMonitorRead(d *schema.ResourceData, m interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	monitor, err := m.(uptimerobotapi.UptimeRobotApiClient).GetMonitor(id)
	if err != nil {
		return err
	}
	if err := updateMonitorResource(d, monitor); err != nil {
		return err
	}
	return nil
}

func resourceMonitorUpdate(d *schema.ResourceData, m interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	req := uptimerobotapi.MonitorUpdateRequest{
		ID:           id,
		FriendlyName: d.Get("friendly_name").(string),
		URL:          d.Get("url").(string),
		Type:         d.Get("type").(string),
	}

	switch req.Type {
	case "port":
		req.SubType = d.Get("sub_type").(string)
		req.Port = d.Get("port").(int)
		break
	case "keyword":
		req.KeywordType = d.Get("keyword_type").(string)
		req.KeywordValue = d.Get("keyword_value").(string)

		req.HTTPUsername = d.Get("http_username").(string)
		req.HTTPPassword = d.Get("http_password").(string)
		req.HTTPAuthType = d.Get("http_auth_type").(string)
		break
	case "http":
		req.HTTPUsername = d.Get("http_username").(string)
		req.HTTPPassword = d.Get("http_password").(string)
		req.HTTPAuthType = d.Get("http_auth_type").(string)
		break
	}

	// Add optional attributes
	req.Interval = d.Get("interval").(int)
	req.IgnoreSSLErrors = d.Get("ignore_ssl_errors").(bool)

	req.AlertContacts = make([]uptimerobotapi.MonitorRequestAlertContact, len(d.Get("alert_contact").([]interface{})))
	for k, v := range d.Get("alert_contact").([]interface{}) {
		req.AlertContacts[k] = uptimerobotapi.MonitorRequestAlertContact{
			ID:         v.(map[string]interface{})["id"].(string),
			Threshold:  v.(map[string]interface{})["threshold"].(int),
			Recurrence: v.(map[string]interface{})["recurrence"].(int),
		}
	}
	sort.Slice(req.AlertContacts, func(i, j int) bool {
		return req.AlertContacts[i].ID < req.AlertContacts[j].ID
	})

	// custom_http_headers
	httpHeaderMap := d.Get("custom_http_headers").(map[string]interface{})
	req.CustomHTTPHeaders = make(map[string]string, len(httpHeaderMap))
	for k, v := range httpHeaderMap {
		req.CustomHTTPHeaders[k] = v.(string)
	}

	monitor, err := m.(uptimerobotapi.UptimeRobotApiClient).UpdateMonitor(req)
	if err != nil {
		return err
	}
	if err := updateMonitorResource(d, monitor); err != nil {
		return err
	}
	return nil
}

func resourceMonitorDelete(d *schema.ResourceData, m interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	err = m.(uptimerobotapi.UptimeRobotApiClient).DeleteMonitor(id)
	if err != nil {
		return err
	}
	return nil
}

func updateMonitorResource(d *schema.ResourceData, m uptimerobotapi.Monitor) error {
	d.Set("friendly_name", m.FriendlyName)
	d.Set("url", m.URL)
	d.Set("type", m.Type)
	d.Set("status", m.Status)
	d.Set("interval", m.Interval)

	d.Set("sub_type", m.SubType)
	d.Set("port", m.Port)

	d.Set("keyword_type", m.KeywordType)
	d.Set("keyword_value", m.KeywordValue)

	d.Set("http_username", m.HTTPUsername)
	d.Set("http_password", m.HTTPPassword)
	// PS: There seems to be a bug in the UR api as it never returns this value
	// d.Set("http_auth_type", m.HTTPAuthType)

	d.Set("ignore_ssl_errors", m.IgnoreSSLErrors)

	if err := d.Set("custom_http_headers", m.CustomHTTPHeaders); err != nil {
		return fmt.Errorf("error setting custom_http_headers for resource %s: %s", d.Id(), err)
	}

	rawContacts := make([]map[string]interface{}, len(m.AlertContacts))
	for k, v := range m.AlertContacts {
		rawContacts[k] = map[string]interface{}{
			"id":         v.ID,
			"recurrence": v.Recurrence,
			"threshold":  v.Threshold,
		}
	}
	if err := d.Set("alert_contact", rawContacts); err != nil {
		return fmt.Errorf("error setting alert_contact for resource %s: %s", d.Id(), err)
	}

	return nil
}
