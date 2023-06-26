package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradePosition interface {
	GetId() int64
	AddLots(lotsResponseMap jsonmap.JsonMap) error
	AddLotsFromResponse(response []byte) error
	AsJsonMap() jsonmap.JsonMap
}

type eTradePosition struct {
	id      int64
	jsonMap jsonmap.JsonMap
}

const (
	// The AsJsonMap() map looks like this:
	// {
	//   <position info>
	//   lots: [
	//      { lot info },
	//   ],
	// }
	//

	// PositionLotsPath is the path to a slice of lots for the position.
	PositionLotsPath = ".lots"
)

const (
	// The position response JSON looks like this:
	// {
	//   "positionId": 1234,
	//   <other position info>
	// }

	// positionPositionIdResponseKey is the key for the position ID
	positionPositionIdResponseKey = "positionId"
)

const (
	// The position lots response JSON looks like this:
	// {
	//   "PositionLotsResponse": {
	//     "PositionLot": [
	//       {
	//         <lot info>
	//       }
	//     ]
	//   }
	// }

	// positionLotsPositionLotResponsePath is the path to the slice of lots for
	// the position.
	positionLotsPositionLotResponsePath = "positionLotsResponse.positionLot"
)

func CreateETradePosition(positionJsonMap jsonmap.JsonMap) (ETradePosition, error) {
	positionId, err := positionJsonMap.GetInt(positionPositionIdResponseKey)
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

func (e *eTradePosition) AddLots(lotsResponseMap jsonmap.JsonMap) error {
	lotsSlice, err := lotsResponseMap.GetSliceAtPath(positionLotsPositionLotResponsePath)
	if err != nil {
		return err
	}
	return e.jsonMap.SetSliceAtPath(PositionLotsPath, lotsSlice)
}

func (e *eTradePosition) AddLotsFromResponse(response []byte) error {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return err
	}
	return e.AddLots(responseMap)
}
