package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func uptimerobotAPICall(
	apiKey string,
	endpoint string,
	params string,
) (map[string]interface{}, error) {
	url := "https://api.uptimerobot.com/v2/" + endpoint

	payload := strings.NewReader(
		fmt.Sprintf("api_key=%s&format=json&%s", apiKey, params),
	)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(res)
	fmt.Println(string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)

	if result["stat"] != "ok" {
		message, _ := json.Marshal(result["error"])
		return nil, errors.New("Got error from UptimeRobot: " + string(message))
	}

	return result, nil
}
