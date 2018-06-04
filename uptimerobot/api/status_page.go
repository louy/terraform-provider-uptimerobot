package uptimerobotapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
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

type StatusPage struct {
	ID           int
	FriendlyName string `json:"friendly_name"`
	StandardURL  string `json:"standard_url"`
	CustomURL    string `json:"custom_url"`
	Sort         string
	Status       string
	DNSAddress   string
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

	sp.DNSAddress = "stats.uptimerobot.com"

	return
}
