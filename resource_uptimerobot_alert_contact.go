package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

type AlertContactType int

var alertContactTypes = map[string]AlertContactType{
	"sms":        1,
	"email":      2,
	"twitter-dm": 3,
	"boxcar":     4,
	"webhook":    5,
	"pushbullet": 6,
	"zapier":     7,
	"pushover":   8,
	"hipchat":    10,
	"slack":      11,
}

func getContactTypeName(value int) string {
	for k, v := range alertContactTypes {
		if int(v) == value {
			return k
		}
	}
	return ""
}

type AlertContactStatus int

var alertContactStatuses = map[string]AlertContactStatus{
	"not activated": 0,
	"paused":        1,
	"active":        2,
}

func getContactStatusName(value int) string {
	for k, v := range alertContactStatuses {
		if int(v) == value {
			return k
		}
	}
	return ""
}

func resourceAlertContact() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertContactCreate,
		Read:   resourceAlertContactRead,
		Update: resourceAlertContactUpdate,
		Delete: resourceAlertContactDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice(mapKeys(alertContactTypes), false),
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAlertContactCreate(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("friendly_name", d.Get("friendly_name").(string))
	data.Add("type", fmt.Sprintf("%d", alertContactTypes[d.Get("type").(string)]))
	data.Add("value", d.Get("value").(string))

	body, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"newAlertContact",
		data.Encode(),
	)
	if err != nil {
		return err
	}
	alertcontact := body["alertcontact"].(map[string]interface{})
	d.SetId(fmt.Sprintf("%d", int(alertcontact["id"].(float64))))
	d.Set("status", getContactStatusName(0))
	return nil
}

func resourceAlertContactRead(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("alert_contacts", d.Id())

	body, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"getAlertContacts",
		data.Encode(),
	)
	if err != nil {
		return err
	}

	alertcontacts, ok := body["alert_contacts"].([]interface{})
	if !ok {
		j, _ := json.Marshal(body)
		return errors.New("Unknown response from the server: " + string(j))
	}

	alertcontact := alertcontacts[0].(map[string]interface{})

	d.Set("friendly_name", alertcontact["friendly_name"].(string))
	d.Set("value", alertcontact["value"].(string))
	d.Set("type", getContactTypeName(int(alertcontact["type"].(float64))))
	d.Set("status", getContactStatusName(int(alertcontact["status"].(float64))))

	return nil
}

func resourceAlertContactUpdate(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("id", d.Id())
	data.Add("friendly_name", d.Get("friendly_name").(string))
	data.Add("value", d.Get("value").(string))

	_, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"editAlertContact",
		data.Encode(),
	)
	if err != nil {
		return err
	}

	return nil
}

func resourceAlertContactDelete(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("id", d.Id())

	_, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"deleteAlertContact",
		data.Encode(),
	)
	if err != nil {
		return err
	}

	return nil
}
