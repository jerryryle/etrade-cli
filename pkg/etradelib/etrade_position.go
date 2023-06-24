package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradePosition interface {
	GetId() int64
	AsJsonMap() jsonmap.JsonMap
}

type eTradePosition struct {
	id      int64
	jsonMap jsonmap.JsonMap
}

const (
	// The position response JSON looks like this:
	// {
	//   "positionId": 1234,
	//   <other alert keys/values>
	// }

	// positionIdResponseKey is the key for the position ID
	positionIdResponseKey = "positionId"
)

func CreateETradePosition(positionJsonMap jsonmap.JsonMap) (ETradePosition, error) {
	positionId, err := positionJsonMap.GetInt(positionIdResponseKey)
	if err != nil {
		return nil, err
	}

	return &eTradePosition{
		id:      positionId,
		jsonMap: positionJsonMap,
	}, nil
}

func (e *eTradePosition) GetId() int64 {
	return e.id
}

func (e *eTradePosition) AsJsonMap() jsonmap.JsonMap {
	return e.jsonMap
}
