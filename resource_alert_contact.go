package main

import (
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform/helper/schema"
)

var alertContactTypes = map[string]int{
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

func resourceAlertContact() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlertContactCreate,
		Read:   resourceAlertContactRead,
		Update: resourceAlertContactUpdate,
		Delete: resourceAlertContactDelete,

		Schema: map[string]*schema.Schema{
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: func(i interface{}, k string) (s []string, es []error) {
					v, ok := i.(string)
					if !ok {
						es = append(es, fmt.Errorf("expected type of %s to be string", k))
						return
					}

					if _, ok := alertContactTypes[v]; !ok {
						es = append(es, fmt.Errorf("unknown type for %s", v))
						return
					}

					return
				},
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
	return nil
}

func resourceAlertContactRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAlertContactUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAlertContactDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
