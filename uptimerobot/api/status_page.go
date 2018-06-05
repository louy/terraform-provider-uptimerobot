package uptimerobotapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

var statusPageStatus = map[string]int{
	"paused": 0,
	"active": 1,
}
var StatusPageStatus = mapKeys(statusPageStatus)

var statusPageSort = map[string]int{
	"a-z":            1,
	"z-a":            2,
	"up-down-paused": 3,
	"down-up-paused": 4,
}
var StatusPageSort = mapKeys(statusPageSort)

type StatusPage struct {
	ID           int
	FriendlyName string `json:"friendly_name"`
	// CustomDomain string `json:"custom_domain"`
	StandardURL string `json:"standard_url"`
	CustomURL   string `json:"custom_url"`
	Sort        string
	Status      string
	DNSAddress  string
	Monitors    []int `json:"monitors"`
}

func (client UptimeRobotApiClient) GetStatusPage(id int) (sp StatusPage, err error) {
	sp.ID = id
	data := url.Values{}
	data.Add("psps", fmt.Sprintf("%d", id))

	body, err := client.MakeCall(
		"getPSPs",
		data.Encode(),
	)
	if err != nil {
		return
	}

	psps, ok := body["psps"].([]interface{})
	if !ok {
		j, _ := json.Marshal(body)
		err = errors.New("Unknown response from the server: " + string(j))
		return
	}

	psp := psps[0].(map[string]interface{})

	sp.FriendlyName = psp["friendly_name"].(string)
	sp.StandardURL = psp["standard_url"].(string)
	if psp["custom_url"] != nil {
		sp.CustomURL = psp["custom_url"].(string)
	}
	sp.Sort = intToString(statusPageSort, int(psp["sort"].(float64)))
	sp.Status = intToString(statusPageStatus, int(psp["status"].(float64)))
	// sp.CustomDomain = psp["custom_domain"].(string)

	monitor, ok := psp["monitors"].(float64)
	if ok {
		sp.Monitors = []int{int(monitor)}
	} else {
		rawMonitors, ok := psp["monitors"].([]interface{})
		if ok {
			monitors := make([]int, len(rawMonitors))
			for k, v := range rawMonitors {
				monitors[k] = int(v.(float64))
			}
			sp.Monitors = monitors
		}
	}

	sp.DNSAddress = "stats.uptimerobot.com"

	return
}

type StatusPageCreateRequest struct {
	FriendlyName string
	CustomDomain string
	Password     string
	Monitors     []int
	Status       string
	Sort         string
}

// CreateStatusPage creates a new status page
func (client UptimeRobotApiClient) CreateStatusPage(req StatusPageCreateRequest) (sp StatusPage, err error) {
	data := url.Values{}
	data.Add("type", fmt.Sprintf("%d", 1))
	data.Add("friendly_name", req.FriendlyName)
	data.Add("custom_domain", req.CustomDomain)
	if req.Password != "" {
		data.Add("password", req.Password)
	}
	if len(req.Monitors) == 0 {
		data.Add("monitors", "0")
	} else {
		var monitors = req.Monitors
		var strMonitors = make([]string, len(monitors))
		for i, v := range monitors {
			strMonitors[i] = strconv.Itoa(v)
		}
		data.Add("monitors", strings.Join(strMonitors, "-"))
	}
	data.Add("sort", fmt.Sprintf("%d", statusPageSort[req.Sort]))
	// FIXME - Got error from UptimeRobot: {"message":"\"status\" is not allowed","parameter_name":"status","passed_value":"1","type":"invalid_parameter"}
	// data.Add("status", fmt.Sprintf("%d", statusPageStatus[req.Status]))

	body, err := client.MakeCall(
		"newPSP",
		data.Encode(),
	)
	if err != nil {
		return
	}
	psp := body["psp"].(map[string]interface{})
	return client.GetStatusPage(int(psp["id"].(float64)))
}

type StatusPageUpdateRequest struct {
	ID           int
	FriendlyName string
	CustomDomain string
	Password     string
	Monitors     []int
	Status       string
	Sort         string
}

// UpdateStatusPage updates an existing status page
func (client UptimeRobotApiClient) UpdateStatusPage(req StatusPageUpdateRequest) (sp StatusPage, err error) {
	data := url.Values{}
	data.Add("id", fmt.Sprintf("%d", req.ID))
	data.Add("type", fmt.Sprintf("%d", 1))
	data.Add("friendly_name", req.FriendlyName)
	data.Add("custom_domain", req.CustomDomain)
	if req.Password != "" {
		data.Add("password", req.Password)
	}
	if len(req.Monitors) == 0 {
		data.Add("monitors", "0")
	} else {
		var monitors = req.Monitors
		var strMonitors = make([]string, len(monitors))
		for i, v := range monitors {
			strMonitors[i] = strconv.Itoa(v)
		}
		data.Add("monitors", strings.Join(strMonitors, "-"))
	}
	data.Add("sort", fmt.Sprintf("%d", statusPageSort[req.Sort]))
	data.Add("status", fmt.Sprintf("%d", statusPageStatus[req.Status]))

	_, err = client.MakeCall(
		"editPSP",
		data.Encode(),
	)
	if err != nil {
		return
	}
	return client.GetStatusPage(req.ID)
}

// DeleteStatusPage updates an existing status page
func (client UptimeRobotApiClient) DeleteStatusPage(id int) (err error) {
	data := url.Values{}
	data.Add("id", fmt.Sprintf("%d", id))

	_, err = client.MakeCall(
		"deletePSP",
		data.Encode(),
	)
	if err != nil {
		return
	}
	return
}
