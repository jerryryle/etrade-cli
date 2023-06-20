package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradePositionList interface {
	GetAllPositions() []ETradePosition
	GetPositionById(positionID int64) ETradePosition
	NextPage() string
	AddPage(positionListResponseMap jsonmap.JsonMap) error
}

type eTradePositionList struct {
	positions     []ETradePosition
	totalsJsonMap jsonmap.JsonMap
	nextPage      string
}

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

	// positionListTotalsMapResponsePath is the path to a map of totals
	positionListTotalsMapResponsePath = "portfolioResponse.totals"

	// positionListPositionsSliceResponsePath is the path to a slice of positions.
	positionListPositionsSliceResponsePath = "portfolioResponse.accountPortfolio[0].position"

	// positionListNextPageStringPath is the path to the next page number string
	positionListNextPageStringPath = "portfolioResponse.accountPortfolio[0].nextPageNo"
)

func CreateETradePositionList(positionListResponseMap jsonmap.JsonMap) (ETradePositionList, error) {
	// the totals are optional, so ignore any error and accept a possibly-nil map.
	totalsMap, _ := positionListResponseMap.GetMapAtPath(positionListTotalsMapResponsePath)

	// Create a new positionList with the totals and everything else
	// initialized to its zero value.
	positionList := eTradePositionList{
		positions:     []ETradePosition{},
		totalsJsonMap: totalsMap,
		nextPage:      "",
	}
	err := positionList.AddPage(positionListResponseMap)
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

func (e *eTradePositionList) AddPage(positionListResponseMap jsonmap.JsonMap) error {
	positionsSlice, err := positionListResponseMap.GetSliceOfMapsAtPath(positionListPositionsSliceResponsePath)
	if err != nil {
		return err
	}

	// the nextPage key only appears if there are more pages, so ignore any
	// error and accept a possibly-zero int.
	nextPage, _ := positionListResponseMap.GetStringAtPath(positionListNextPageStringPath)

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
