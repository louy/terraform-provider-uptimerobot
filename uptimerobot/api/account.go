package uptimerobotapi

type Account struct {
	Email           string
	MonitorLimit    int
	MonitorInterval int
	UpMonitors      int
	DownMonitors    int
	PausedMonitors  int
}

func (client UptimeRobotApiClient) GetAccountDetails() (acc Account, err error) {
	body, err := client.MakeCall(
		"getAccountDetails",
		"",
	)

	if err != nil {
		return
	}
	account := body["account"].(map[string]interface{})

	acc.Email = account["email"].(string)
	acc.MonitorLimit = int(account["monitor_limit"].(float64))
	acc.MonitorInterval = int(account["monitor_interval"].(float64))
	acc.UpMonitors = int(account["up_monitors"].(float64))
	acc.DownMonitors = int(account["down_monitors"].(float64))
	acc.PausedMonitors = int(account["paused_monitors"].(float64))

	return
}
