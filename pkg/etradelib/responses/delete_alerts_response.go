package responses

type DeleteAlertsResponse struct {
	Result       string                     `xml:"result"`
	FailedAlerts []DeleteAlertsFailedAlerts `xml:"failedAlerts"`
}

type DeleteAlertsFailedAlerts struct {
	AlertIds []int64 `xml:"alertId"`
}
