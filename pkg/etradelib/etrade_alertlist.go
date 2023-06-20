package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeAlertList interface {
	GetAllAlerts() []ETradeAlert
	GetAlertById(alertID int64) ETradeAlert
}

type eTradeAlertList struct {
	alerts []ETradeAlert
}

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

	// alertsSliceResponsePath is the path to a slice of alerts.
	alertsSliceResponsePath = "alertsResponse.alert"
)

func CreateETradeAlertList(alertListResponseMap jsonmap.JsonMap) (ETradeAlertList, error) {
	alertsSlice, err := alertListResponseMap.GetSliceOfMapsAtPath(alertsSliceResponsePath)
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
