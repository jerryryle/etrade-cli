package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradePositionList interface {
	GetAllPositions() []ETradePosition
	GetPositionById(positionID int64) ETradePosition
	NextPage() string
	AddPage(responseMap jsonmap.JsonMap) error
	AddPageFromResponse(response []byte) error
	AsJsonMap() jsonmap.JsonMap
}

type eTradePositionList struct {
	positions []ETradePosition
	totalsMap jsonmap.JsonMap
	nextPage  string
}

const (
	// The AsJsonMap() map looks like this:
	// {
	//   "positions": [
	//     {
	//       <position info>
	//     }
	//   ]
	//   "totals": {
	//     <totals info>
	//   }
	// }
	//

	// PositionsListPositionsPath is the path to a slice of positions.
	PositionsListPositionsPath = ".positions"

	// PositionsListTotalsPath is the path to a map of totals info
	PositionsListTotalsPath = ".totals"
)

const (
	// The portfolio (position list) response JSON looks like this:
	// {
	//   "PortfolioResponse": {
	//     "Totals": {
	//       <totals info>
	//     },
	//     "AccountPortfolio": [
	//       {
	//         "nextPageNo": "2",
	//         "Position": [
	//           {
	//             <position info>
	//           }
	//         ]
	//       }
	//     ]
	//   }
	// }
	//
	// The "Totals" key is optional and only appears if explicitly requested.
	// The "nextPageNo" key only appears if there are more pages to fetch.

	// positionListTotalsResponsePath is the path to a map of totals
	positionListTotalsResponsePath = ".portfolioResponse.totals"

	// positionListPositionsResponsePath is the path to a slice of positions.
	positionListPositionsResponsePath = ".portfolioResponse.accountPortfolio[0].position"

	// positionListNextPageResponsePath is the path to the next page number string
	positionListNextPageResponsePath = ".portfolioResponse.accountPortfolio[0].nextPageNo"
)

func CreateETradePositionListFromResponse(response []byte) (ETradePositionList, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradePositionList(responseMap)
}

func CreateETradePositionList(responseMap jsonmap.JsonMap) (ETradePositionList, error) {
	// the totals are optional, so ignore any error and accept a possibly-nil map.
	totalsMap, _ := responseMap.GetMapAtPath(positionListTotalsResponsePath)

	// Create a new positionList with the totals and everything else
	// initialized to its zero value.
	positionList := eTradePositionList{
		positions: []ETradePosition{},
		totalsMap: totalsMap,
		nextPage:  "",
	}
	err := positionList.AddPage(responseMap)
	if err != nil {
		return nil, err
	}
	return &positionList, nil
}

func (e *eTradePositionList) GetAllPositions() []ETradePosition {
	return e.positions
}

func (e *eTradePositionList) GetPositionById(positionID int64) ETradePosition {
	for _, position := range e.positions {
		if position.GetId() == positionID {
			return position
		}
	}
	return nil
}

func (e *eTradePositionList) NextPage() string {
	return e.nextPage
}

func (e *eTradePositionList) AddPageFromResponse(response []byte) error {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return err
	}
	return e.AddPage(responseMap)
}

func (e *eTradePositionList) AddPage(responseMap jsonmap.JsonMap) error {
	positionsSlice, err := responseMap.GetSliceOfMapsAtPath(positionListPositionsResponsePath)
	if err != nil {
		return err
	}

	// the nextPage key only appears if there are more pages, so ignore any
	// error and accept a possibly-zero int.
	nextPage, _ := responseMap.GetStringAtPath(positionListNextPageResponsePath)

	allPositions := make([]ETradePosition, 0, len(positionsSlice))
	for _, positionJsonMap := range positionsSlice {
		position, err := CreateETradePosition(positionJsonMap)
		if err != nil {
			return err
		}
		allPositions = append(allPositions, position)
	}
	e.positions = append(e.positions, allPositions...)
	e.nextPage = nextPage
	return nil
}

func (e *eTradePositionList) AsJsonMap() jsonmap.JsonMap {
	positionSlice := make(jsonmap.JsonSlice, 0, len(e.positions))
	for _, position := range e.positions {
		positionSlice = append(positionSlice, position.AsJsonMap())
	}
	var positionListMap = jsonmap.JsonMap{}
	err := positionListMap.SetSliceAtPath(PositionsListPositionsPath, positionSlice)
	if err != nil {
		panic(err)
	}

	if e.totalsMap != nil {
		err := positionListMap.SetMapAtPath(PositionsListTotalsPath, e.totalsMap)
		if err != nil {
			panic(err)
		}
	}

	return positionListMap
}
