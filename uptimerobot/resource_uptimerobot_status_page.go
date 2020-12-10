package uptimerobot

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/louy/terraform-provider-uptimerobot/uptimerobot/api"
)

func resourceStatusPage() *schema.Resource {
	return &schema.Resource{
		Create: resourceStatusPageCreate,
		Read:   resourceStatusPageRead,
		Update: resourceStatusPageUpdate,
		Delete: resourceStatusPageDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"friendly_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"custom_domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"sort": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(uptimerobotapi.StatusPageSort, false),
				Default:      "a-z",
			},
			"status": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(uptimerobotapi.StatusPageStatus, false),
				Default:      "active",
			},
			"monitors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				PromoteSingle: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if k == "monitors.#" && old == "1" && new == "0" && d.Get("monitors.0").(int) == 0 {
						return true
					}
					return false
				},
			},
			"dns_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"standard_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"custom_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceStatusPageCreate(d *schema.ResourceData, m interface{}) error {
	rawMonitors := d.Get("monitors").([]interface{})
	monitors := make([]int, len(rawMonitors))
	for i := range rawMonitors {
		monitors[i] = rawMonitors[i].(int)
	}

	sp, err := m.(uptimerobotapi.UptimeRobotApiClient).CreateStatusPage(uptimerobotapi.StatusPageCreateRequest{
		FriendlyName: d.Get("friendly_name").(string),
		CustomDomain: d.Get("custom_domain").(string),
		Password:     d.Get("password").(string),
		Monitors:     monitors,
		Sort:         d.Get("sort").(string),
		Status:       d.Get("status").(string),
	})
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%d", sp.ID))
	updateStatusPageResource(d, sp)

	return nil
}

func resourceStatusPageRead(d *schema.ResourceData, m interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	sp, err := m.(uptimerobotapi.UptimeRobotApiClient).GetStatusPage(id)
	if err != nil {
		return err
	}

	updateStatusPageResource(d, sp)

	return nil
}

func resourceStatusPageUpdate(d *schema.ResourceData, m interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	rawMonitors := d.Get("monitors").([]interface{})
	monitors := make([]int, len(rawMonitors))
	for i := range rawMonitors {
		monitors[i] = rawMonitors[i].(int)
	}

	sp, err := m.(uptimerobotapi.UptimeRobotApiClient).UpdateStatusPage(uptimerobotapi.StatusPageUpdateRequest{
		ID:           id,
		FriendlyName: d.Get("friendly_name").(string),
		CustomDomain: d.Get("custom_domain").(string),
		Password:     d.Get("password").(string),
		Monitors:     monitors,
		Sort:         d.Get("sort").(string),
		Status:       d.Get("status").(string),
	})
	if err != nil {
		return err
	}

	updateStatusPageResource(d, sp)

	return nil
}

func resourceStatusPageDelete(d *schema.ResourceData, m interface{}) error {
	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return err
	}

	err = m.(uptimerobotapi.UptimeRobotApiClient).DeleteStatusPage(id)
	if err != nil {
		return err
	}

	return nil
}
func updateStatusPageResource(d *schema.ResourceData, sp uptimerobotapi.StatusPage) {
	d.Set("friendly_name", sp.FriendlyName)
	d.Set("standard_url", sp.StandardURL)
	d.Set("custom_url", sp.CustomURL)
	d.Set("sort", sp.Sort)
	d.Set("status", sp.Status)
	d.Set("dns_address", sp.DNSAddress)
	d.Set("monitors", sp.Monitors)
}
