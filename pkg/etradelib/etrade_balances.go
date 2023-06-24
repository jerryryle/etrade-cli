package etradelib

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

type ETradeBalances interface {
	AsJsonMap() jsonmap.JsonMap
}

type eTradeBalances struct {
	balancesMap jsonmap.JsonMap
}

const (
	// The portfolio (position list) response JSON looks like this:
	// {
	//   "BalanceResponse": {
	//       <balance info>
	//   }
	// }
	//

	// balanceMapResponsePath is the path to the map of balance info
	balanceMapResponsePath = ".balanceResponse"
)

func CreateETradeBalancesFromResponse(response []byte) (ETradeBalances, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeBalances(responseMap)
}

func CreateETradeBalances(positionListResponseMap jsonmap.JsonMap) (ETradeBalances, error) {
	balancesMap, err := positionListResponseMap.GetMapAtPath(balanceMapResponsePath)
	if err != nil {
		return nil, err
	}

	balances := eTradeBalances{
		balancesMap: balancesMap,
	}
	return &balances, nil
}

func (e *eTradeBalances) AsJsonMap() jsonmap.JsonMap {
	return e.balancesMap
}
