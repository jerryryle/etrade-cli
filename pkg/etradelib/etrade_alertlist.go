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
	//     "totalAlerts": 1,
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
	accountsSlice, err := alertListResponseMap.GetMapSliceAtPath(alertsSliceResponsePath)
	if err != nil {
		return nil, err
	}
	allAlerts := make([]ETradeAlert, 0, len(accountsSlice))
	for _, alertInfoMap := range accountsSlice {
		account, err := CreateETradeAlert(alertInfoMap)
		if err != nil {
			return nil, err
		}
		allAlerts = append(allAlerts, account)
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
