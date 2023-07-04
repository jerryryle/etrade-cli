package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeAlert interface {
	GetId() int64
	AsJsonMap() jsonmap.JsonMap
}

type eTradeAlert struct {
	id      int64
	jsonMap jsonmap.JsonMap
}

const (
	// The alert response JSON looks like this:
	// {
	//   "id": 1234,
	//   <other alert keys/values>
	// }

	// alertAlertIdResponseKey is the key for the alert ID
	alertAlertIdResponseKey = "id"
)

func CreateETradeAlert(responseMap jsonmap.JsonMap) (ETradeAlert, error) {
	alertId, err := responseMap.GetInt(alertAlertIdResponseKey)
	if err != nil {
		return nil, err
	}

	return &eTradeAlert{
		id:      alertId,
		jsonMap: responseMap,
	}, nil
}

func (e *eTradeAlert) GetId() int64 {
	return e.id
}

func (e *eTradeAlert) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
