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

	// balancesBalanceResponsePath is the path to the map of balance info
	balancesBalanceResponsePath = ".balanceResponse"
)

func CreateETradeBalancesFromResponse(response []byte) (ETradeBalances, error) {
	responseMap, err := NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	return CreateETradeBalances(responseMap)
}

func CreateETradeBalances(responseMap jsonmap.JsonMap) (ETradeBalances, error) {
	balancesMap, err := responseMap.GetMapAtPath(balancesBalanceResponsePath)
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
