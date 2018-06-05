package uptimerobotapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

var alertContactType = map[string]int{
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
var AlertContactType = mapKeys(alertContactType)

var alertContactStatus = map[string]int{
	"not activated": 0,
	"paused":        1,
	"active":        2,
}

var AlertContactStatus = mapKeys(alertContactStatus)

type AlertContact struct {
	ID           int    `json:"id"`
	FriendlyName string `json:"friendly_name"`
	Value        string `json:"value"`
	Type         string
	Status       string
}

func (client UptimeRobotApiClient) GetAlertContact(id int) (ac AlertContact, err error) {
	ac.ID = id
	data := url.Values{}
	data.Add("alert_contacts", fmt.Sprintf("%d", id))

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

	return client.GetAlertContact(int(alertcontact["id"].(float64)))
}

func (client UptimeRobotApiClient) DeleteAlertContact(id int) (err error) {
	data := url.Values{}
	data.Add("id", fmt.Sprintf("%d", id))

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
	ID           int
	FriendlyName string
	Value        string
}

func (client UptimeRobotApiClient) UpdateAlertContact(req AlertContactUpdateRequest) (err error) {
	data := url.Values{}
	data.Add("id", fmt.Sprintf("%d", req.ID))
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
