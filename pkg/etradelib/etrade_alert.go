package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

type ETradeAlert interface {
	GetId() int64
	GetInfoMap() jsonmap.JsonMap
}

type eTradeAlert struct {
	id      int64
	infoMap jsonmap.JsonMap
}

const (
	// The alert response JSON looks like this:
	// {
	//   "id": 1234,
	//   <other alert keys/values>
	// }

	// alertIdResponseKey is the key for the alert ID
	alertIdResponseKey = "id"
)

func CreateETradeAlert(alertResponseMap jsonmap.JsonMap) (ETradeAlert, error) {
	alertId, err := alertResponseMap.GetInt(alertIdResponseKey)
	if err != nil {
		return nil, err
	}

	return &eTradeAlert{
		id:      alertId,
		infoMap: alertResponseMap,
	}, nil
}

func (e *eTradeAlert) GetId() int64 {
	return e.id
}

func (e *eTradeAlert) GetInfoMap() jsonmap.JsonMap {
	return e.infoMap
}
