package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeAlertList interface {
	GetAllAlerts() []ETradeAlert
	GetAlertById(alertID int64) ETradeAlert
	AsJsonMap() jsonmap.JsonMap
}

type eTradeAlertList struct {
	alerts []ETradeAlert
}

const (
	// The AsJsonMap() map looks like this:
	// "alerts": [
	//   {
	//     <alert info>
	//   }
	// ]

	// AlertListAlertsPath is the path to a slice of alerts.
	AlertListAlertsPath = ".alerts"
)

const (
	// The alert list response JSON looks like this:
	// {
	//   "AlertsResponse": {
	//     "Alert": [
	//       {
	//         <alert info>
	//       }
	//     ]
	//   }
	// }

	// alertListAlertsResponsePath is the path to a slice of alerts.
	alertListAlertsResponsePath = ".alertsResponse.alert"
)

func CreateETradeAlertListFromResponse(response []byte) (ETradeAlertList, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeAlertList(responseMap)
}

func CreateETradeAlertList(alertListResponseMap jsonmap.JsonMap) (ETradeAlertList, error) {
	alertsSlice, err := alertListResponseMap.GetSliceOfMapsAtPath(alertListAlertsResponsePath)
	if err != nil {
		return nil, err
	}
	allAlerts := make([]ETradeAlert, 0, len(alertsSlice))
	for _, alertJsonMap := range alertsSlice {
		alert, err := CreateETradeAlert(alertJsonMap)
		if err != nil {
			return nil, err
		}
		allAlerts = append(allAlerts, alert)
	}
	return &eTradeAlertList{alerts: allAlerts}, nil
}

func (e *eTradeAlertList) GetAllAlerts() []ETradeAlert {
	return e.alerts
}

func (e *eTradeAlertList) GetAlertById(alertID int64) ETradeAlert {
	for _, alert := range e.alerts {
		if alert.GetId() == alertID {
			return alert
		}
	}
	return nil
}

func (e *eTradeAlertList) AsJsonMap() jsonmap.JsonMap {
	alertsSlice := make(jsonmap.JsonSlice, 0, len(e.alerts))
	for _, alert := range e.alerts {
		alertsSlice = append(alertsSlice, alert.AsJsonMap())
	}
	var alertListMap = jsonmap.JsonMap{}
	err := alertListMap.SetSliceAtPath(AlertListAlertsPath, alertsSlice)
	if err != nil {
		panic(err)
	}
	return alertListMap
}
