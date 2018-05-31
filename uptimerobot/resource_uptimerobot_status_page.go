package uptimerobot

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
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
				Default:      "a-z",
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				// ValidateFunc: validation.StringInSlice(mapKeys(statusPageStatus), false),
				// Optional:     true,
				// Default:      "active",
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
	data := url.Values{}
	data.Add("type", fmt.Sprintf("%d", 1))
	data.Add("friendly_name", d.Get("friendly_name").(string))
	data.Add("custom_domain", d.Get("custom_domain").(string))
	if d.Get("password").(string) != "" {
		data.Add("password", d.Get("password").(string))
	}
	// log.Printf("[DEBUG] [monitors] %+v", d.Get("monitors").([]interface{}))
	if len(d.Get("monitors").([]interface{})) == 0 {
		data.Add("monitors", "0")
	} else {
		// log.Printf("[DEBUG] [monitors type] %s", reflect.TypeOf(d.Get("monitors").([]interface{})[0]))
		var monitors = d.Get("monitors").([]interface{})
		var strMonitors = make([]string, len(monitors))
		for i, v := range monitors {
			strMonitors[i] = strconv.Itoa(v.(int))
		}
		data.Add("monitors", strings.Join(strMonitors, "-"))
	}
	// log.Printf("[DEBUG] Sort: %s %d", d.Get("sort").(string), statusPageSort[d.Get("sort").(string)])
	data.Add("sort", fmt.Sprintf("%d", statusPageSort[d.Get("sort").(string)]))
	// data.Add("status", fmt.Sprintf("%d", statusPageStatus[d.Get("status").(string)]))

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
	d.Set("standard_url", psp["standard_url"].(string))
	if psp["custom_url"] != nil {
		d.Set("custom_url", psp["custom_url"].(string))
	} else {
		d.Set("custom_url", nil)
	}

	d.Set("dns_address", "stats.uptimerobot.com")
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
	if psp["custom_url"] != nil {
		d.Set("custom_url", psp["custom_url"].(string))
	} else {
		d.Set("custom_url", nil)
	}
	d.Set("sort", intToString(statusPageSort, int(psp["sort"].(float64))))
	d.Set("status", intToString(statusPageStatus, int(psp["status"].(float64))))

	d.Set("dns_address", "stats.uptimerobot.com")

	return nil
}

func resourceStatusPageUpdate(d *schema.ResourceData, m interface{}) error {
	data := url.Values{}
	data.Add("id", d.Id())
	data.Add("type", fmt.Sprintf("%d", 1))
	data.Add("friendly_name", d.Get("friendly_name").(string))
	data.Add("custom_domain", d.Get("custom_domain").(string))
	if d.Get("password").(string) != "" {
		data.Add("password", d.Get("password").(string))
	}
	// log.Printf("[DEBUG] [monitors] %+v", d.Get("monitors").([]interface{}))
	if len(d.Get("monitors").([]interface{})) == 0 {
		data.Add("monitors", "0")
	} else {
		var monitors = d.Get("monitors").([]interface{})
		var strMonitors = make([]string, len(monitors))
		for i, v := range monitors {
			strMonitors[i] = strconv.Itoa(v.(int))
		}
		data.Add("monitors", strings.Join(strMonitors, "-"))
	}
	data.Add("sort", fmt.Sprintf("%d", statusPageSort[d.Get("sort").(string)]))
	// data.Add("status", fmt.Sprintf("%d", statusPageStatus[d.Get("status").(string)]))

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
