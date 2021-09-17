package uptimerobotapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

func New(apiKey string) UptimeRobotApiClient {
	return UptimeRobotApiClient{apiKey}
}

type UptimeRobotApiClient struct {
	apiKey string
}

func (client UptimeRobotApiClient) MakeCall(
	endpoint string,
	params string,
) (map[string]interface{}, error) {
	log.Printf("[DEBUG] Making request to: %#v", endpoint)

	url := "https://api.uptimerobot.com/v2/" + endpoint

	payload := strings.NewReader(
		fmt.Sprintf("api_key=%s&format=json&%s", client.apiKey, params),
	)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10
	standardClient := retryClient.StandardClient()

	res, err := standardClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Got %d response from UptimeRobot: %s", res.StatusCode, body)
	}

	log.Printf("[DEBUG] Got response: %#v", res)
	log.Printf("[DEBUG] Got body: %#v", string(body))

	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil, fmt.Errorf("Got decoding json from UptimeRobot: %s. Response body: %s", err.Error(), body)
	}

	if result["stat"] != "ok" {
		message, _ := json.Marshal(result["error"])
		return nil, errors.New("Got error from UptimeRobot: " + string(message))
	}

	return result, nil
}
