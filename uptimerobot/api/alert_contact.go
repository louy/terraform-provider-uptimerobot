package uptimerobotapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// maximum pagination depth to allow (10*50=500 entries)
const page_limit = 10

var alertContactType = map[string]int{
	"sms":        1,
	"email":      2,
	"twitter-dm": 3,
	"boxcar":     4,
	"webhook":    5,
	"pushbullet": 6,
	"zapier":     7,
	"pushover":   9,
	"hipchat":    10,
	"slack":      11,
	"hangouts":	  21,
}
var AlertContactType = mapKeys(alertContactType)

var alertContactStatus = map[string]int{
	"not activated": 0,
	"paused":        1,
	"active":        2,
}

var AlertContactStatus = mapKeys(alertContactStatus)

type AlertContact struct {
	ID           string `json:"id"`
	FriendlyName string `json:"friendly_name"`
	Value        string `json:"value"`
	Type         string
	Status       string
}

func (client UptimeRobotApiClient) GetAlertContacts() (acs []AlertContact, err error) {
	data := url.Values{}

	var total float64

	for i := 0; i < page_limit; i++ {
		body, err := client.MakeCall(
			"getAlertContacts",
			data.Encode(),
		)
		if err != nil {
			return nil, err
		}

		alertcontacts, ok := body["alert_contacts"].([]interface{})
		if !ok {
			j, _ := json.Marshal(body)
			err = errors.New("Unknown response from the server: " + string(j))
			return nil, err
		}

		for _, i := range alertcontacts {
			alertcontact := i.(map[string]interface{})
			id := alertcontact["id"].(string)
			friendlyName := alertcontact["friendly_name"].(string)
			value := ""
			if alertcontact["value"] != nil {
				value = alertcontact["value"].(string)
			}
			ac := AlertContact{
				id,
				friendlyName,
				value,
				intToString(alertContactType, int(alertcontact["type"].(float64))),
				intToString(alertContactStatus, int(alertcontact["status"].(float64))),
			}
			acs = append(acs, ac)
		}

		total = body["total"].(float64)
		if float64(len(acs)) == total {
			break
		}
	}

	if float64(len(acs)) != total {
		err = errors.New("Hitting pagination limit of: " + string(page_limit))
	}

	return
}

func (client UptimeRobotApiClient) GetAlertContact(id string) (ac AlertContact, err error) {
	ac.ID = id
	data := url.Values{}
	data.Add("alert_contacts", id)

	body, err := client.MakeCall(
		"getAlertContacts",
		data.Encode(),
	)
	if err != nil {
		return
	}

	alertcontacts, ok := body["alert_contacts"].([]interface{})
	if !ok {
		j, _ := json.Marshal(body)
		err = errors.New("Unknown response from the server: " + string(j))
		return
	}

	alertcontact := alertcontacts[0].(map[string]interface{})

	ac.FriendlyName = alertcontact["friendly_name"].(string)
	ac.Value = alertcontact["value"].(string)
	ac.Type = intToString(alertContactType, int(alertcontact["type"].(float64)))
	ac.Status = intToString(alertContactStatus, int(alertcontact["status"].(float64)))

	return
}

type AlertContactCreateRequest struct {
	FriendlyName string
	Type         string
	Value        string
}

func (client UptimeRobotApiClient) CreateAlertContact(req AlertContactCreateRequest) (ac AlertContact, err error) {
	data := url.Values{}
	data.Add("friendly_name", req.FriendlyName)
	data.Add("type", fmt.Sprintf("%d", alertContactType[req.Type]))
	data.Add("value", req.Value)

	body, err := client.MakeCall(
		"newAlertContact",
		data.Encode(),
	)
	if err != nil {
		return
	}

	alertcontact, ok := body["alertcontact"].(map[string]interface{})
	if !ok {
		j, _ := json.Marshal(body)
		err = errors.New("Unknown response from the server: " + string(j))
		return
	}

	// The alert contact ID is a string value according to API docs but is
	// returned as a integer value by the newAlertContact API JSON. In other
	// places the API does correctly handle it as a string value.
	// The difference made by it being a string is that a zero prefix to the ID // number is preserved. A zero prefixed alert contact ID is thus far only
	// been observed on the default alert contact (created at account creation).
	// https://github.com/louy/terraform-provider-uptimerobot/pull/21
	return client.GetAlertContact(fmt.Sprintf("%.0f", alertcontact["id"].(float64)))
}

func (client UptimeRobotApiClient) DeleteAlertContact(id string) (err error) {
	data := url.Values{}
	data.Add("id", id)

	_, err = client.MakeCall(
		"deleteAlertContact",
		data.Encode(),
	)
	if err != nil {
		return
	}
	return
}

type AlertContactUpdateRequest struct {
	ID           string
	FriendlyName string
	Value        string
}

func (client UptimeRobotApiClient) UpdateAlertContact(req AlertContactUpdateRequest) (err error) {
	data := url.Values{}
	data.Add("id", req.ID)
	data.Add("friendly_name", req.FriendlyName)
	data.Add("value", req.Value)

	_, err = client.MakeCall(
		"editAlertContact",
		data.Encode(),
	)
	if err != nil {
		return
	}

	return
}
