package uptimerobot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/realiotech/terraform-provider-uptimerobot/uptimerobot/api"
)

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
				ValidateFunc: validation.StringInSlice(uptimerobotapi.AlertContactType, false),
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
	ac, err := m.(uptimerobotapi.UptimeRobotApiClient).CreateAlertContact(
		uptimerobotapi.AlertContactCreateRequest{
			FriendlyName: d.Get("friendly_name").(string),
			Type:         d.Get("type").(string),
			Value:        d.Get("value").(string),
		})
	if err != nil {
		return err
	}

	d.SetId(ac.ID)
	updateAlertContactResource(d, ac)

	return nil
}

func resourceAlertContactRead(d *schema.ResourceData, m interface{}) error {
	id := d.Id()

	ac, err := m.(uptimerobotapi.UptimeRobotApiClient).GetAlertContact(id)
	if err != nil {
		return err
	}

	updateAlertContactResource(d, ac)

	return nil
}

func resourceAlertContactUpdate(d *schema.ResourceData, m interface{}) error {
	id := d.Id()

	err := m.(uptimerobotapi.UptimeRobotApiClient).UpdateAlertContact(
		uptimerobotapi.AlertContactUpdateRequest{
			ID:           id,
			FriendlyName: d.Get("friendly_name").(string),
			Value:        d.Get("value").(string),
		})
	if err != nil {
		return err
	}

	return nil
}

func resourceAlertContactDelete(d *schema.ResourceData, m interface{}) error {
	id := d.Id()

	err := m.(uptimerobotapi.UptimeRobotApiClient).DeleteAlertContact(id)
	if err != nil {
		return err
	}

	return nil
}

func updateAlertContactResource(d *schema.ResourceData, ac uptimerobotapi.AlertContact) {
	d.Set("friendly_name", ac.FriendlyName)
	d.Set("value", ac.Value)
	d.Set("type", ac.Type)
	d.Set("status", ac.Status)
}
