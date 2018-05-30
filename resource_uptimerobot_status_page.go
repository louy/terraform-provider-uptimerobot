package uptimerobot

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

var statusPageStatus = map[string]int{
	"paused": 0,
	"active": 1,
}
var statusPageSort = map[string]int{
	"a-z":            1,
	"z-a":            2,
	"up-down-paused": 3,
	"down-up-paused": 4,
}

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
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(mapKeys(statusPageSort), false),
				Default:      mapKeys(monitorTypes)[0],
			},
			"status": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(mapKeys(statusPageStatus), false),
				Optional:     true,
				Default:      "active",
			},
			"monitors": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				PromoteSingle: true,
			},
			"dns_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				DefaultFunc: func() (interface{}, error) {
					return "stats.uptimerobot.com", nil
				},
			},
		},
	}
}

func resourceStatusPageCreate(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("type", fmt.Sprintf("%d", 1))
	data.Add("friendly_name", d.Get("friendly_name").(string))
	data.Add("custom_domain", d.Get("custom_domain").(string))
	data.Add("password", d.Get("password").(string))
	if len(d.Get("monitors").([]string)) == 0 {
		data.Add("monitors", "0")
	} else {
		data.Add("monitors", strings.Join(d.Get("monitors").([]string), "-"))
	}
	data.Add("sort", fmt.Sprintf("%d", statusPageSort[d.Get("sort").(string)]))
	data.Add("status", fmt.Sprintf("%d", statusPageStatus[d.Get("status").(string)]))

	body, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"newPSP",
		data.Encode(),
	)
	if err != nil {
		return err
	}
	psp := body["psp"].(map[string]interface{})
	d.SetId(fmt.Sprintf("%d", int(psp["id"].(float64))))
	d.Set("status", "active")
	return nil
}

func resourceStatusPageRead(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("psps", d.Id())

	body, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"getPSPs",
		data.Encode(),
	)
	if err != nil {
		return err
	}

	psps, ok := body["psps"].([]interface{})
	if !ok {
		j, _ := json.Marshal(body)
		return errors.New("Unknown response from the server: " + string(j))
	}

	psp := psps[0].(map[string]interface{})

	d.Set("friendly_name", psp["friendly_name"].(string))
	d.Set("standard_url", psp["standard_url"].(string))
	d.Set("custom_url", psp["custom_url"].(string))
	d.Set("sort", intToString(statusPageSort, psp["sort"].(int)))
	d.Set("status", intToString(statusPageStatus, psp["status"].(int)))

	return nil
}

func resourceStatusPageUpdate(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("id", d.Id())
	data.Add("type", fmt.Sprintf("%d", 1))
	data.Add("friendly_name", d.Get("friendly_name").(string))
	data.Add("custom_domain", d.Get("custom_domain").(string))
	data.Add("password", d.Get("password").(string))
	if len(d.Get("monitors").([]string)) == 0 {
		data.Add("monitors", "0")
	} else {
		data.Add("monitors", strings.Join(d.Get("monitors").([]string), "-"))
	}
	data.Add("sort", fmt.Sprintf("%d", statusPageSort[d.Get("sort").(string)]))
	data.Add("status", fmt.Sprintf("%d", statusPageStatus[d.Get("status").(string)]))

	_, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"editPSP",
		data.Encode(),
	)
	if err != nil {
		return err
	}

	return nil
}

func resourceStatusPageDelete(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("id", d.Id())

	_, err := uptimerobotAPICall(
		m.(UptimeRobotConfig).apiKey,
		"deletePSP",
		data.Encode(),
	)
	if err != nil {
		return err
	}

	return nil
}
