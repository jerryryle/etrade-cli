package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func GetAccountBalances(eTradeClient client.ETradeClient, accountId string, realTimeBalance bool) (
	jsonmap.JsonMap, error,
) {
	account, err := GetAccountById(eTradeClient, accountId)
	if err != nil {
		return nil, err
	}
	response, err := eTradeClient.GetAccountBalances(account.GetIdKey(), realTimeBalance)
	if err != nil {
		return nil, err
	}
	balances, err := etradelib.CreateETradeBalancesFromResponse(response)
	if err != nil {
		return nil, err
	}
	return balances.AsJsonMap(), nil
}
