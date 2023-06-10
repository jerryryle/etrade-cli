package responses

type AlertsResponse struct {
	TotalAlerts int64         `xml:"totalAlerts"`
	Alerts      []AlertsAlert `xml:"Alert"`
}

type AlertsAlert struct {
	Id         int64      `xml:"id"`
	CreateTime ETradeTime `xml:"createTime"`
	Subject    string     `xml:"subject"`
	Status     string     `xml:"status"`
}
