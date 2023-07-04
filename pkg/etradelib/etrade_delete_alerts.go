package etradelib

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeDeleteAlerts interface {
	IsSuccess() bool
	GetFailedAlerts() []int64
	AsJsonMap() jsonmap.JsonMap
}

type eTradeDeleteAlerts struct {
	isSuccess    bool
	failedAlerts []int64
}

const (
	// The AsJsonMap() map looks like this:
	// {
	//   "status": "error",
	//   "error": "some alerts could not be deleted",
	//   "failedAlerts": [1234]
	// }

	// DeleteAlertsStatusKey is the key for the status
	DeleteAlertsStatusKey = "status"

	// DeleteAlertsErrorKey is the key for the error message
	DeleteAlertsErrorKey = "error"

	// DeleteAlertsFailedAlertsKey is the key for a slice of failed
	// alert IDs
	DeleteAlertsFailedAlertsKey = "failedAlerts"
)

const (
	// The delete alert response JSON looks like this:
	// {
	//   "AlertsResponse": {
	//     "failedAlerts": {
	//       "alertId": [
	//         1234
	//       ]
	//     },
	//     "result": "ERROR"
	//   }
	// }

	// deleteAlertsResultResponsePath is the path to the result
	deleteAlertsResultResponsePath = ".alertsResponse.result"

	// deleteAlertsFailedAlertsResponsePath is the path to a slice of failed
	// alert IDs
	deleteAlertsFailedAlertsResponsePath = ".alertsResponse.failedAlerts.alertId"
)

func CreateETradeDeleteAlertsFromResponse(response []byte) (ETradeDeleteAlerts, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeDeleteAlerts(responseMap)
}

func CreateETradeDeleteAlerts(responseMap jsonmap.JsonMap) (ETradeDeleteAlerts, error) {
	result, err := responseMap.GetStringAtPath(deleteAlertsResultResponsePath)
	if err != nil {
		return nil, err
	}

	isSuccess := false
	switch result {
	case "SUCCESS":
		isSuccess = true
	case "ERROR":
		isSuccess = false
	default:
		return nil, fmt.Errorf("unexpected result '%s'", result)
	}

	failedAlerts, err := responseMap.GetSliceOfIntsAtPathWithDefault(deleteAlertsFailedAlertsResponsePath, []int64{})
	if err != nil {
		return nil, err
	}

	return &eTradeDeleteAlerts{
		isSuccess:    isSuccess,
		failedAlerts: failedAlerts,
	}, nil
}

func (e *eTradeDeleteAlerts) IsSuccess() bool {
	return e.isSuccess
}

func (e *eTradeDeleteAlerts) GetFailedAlerts() []int64 {
	return e.failedAlerts
}

func (e *eTradeDeleteAlerts) AsJsonMap() jsonmap.JsonMap {
	if e.isSuccess {
		return jsonmap.JsonMap{
			DeleteAlertsStatusKey: "success",
		}
	} else {
		return jsonmap.JsonMap{
			DeleteAlertsStatusKey:       "error",
			DeleteAlertsErrorKey:        "some alerts could not be deleted",
			DeleteAlertsFailedAlertsKey: jsonmap.NewJsonSliceFromSlice(e.failedAlerts),
		}
	}
}
