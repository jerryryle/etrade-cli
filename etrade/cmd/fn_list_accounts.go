package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func ListAccounts(eTradeClient client.ETradeClient) (jsonmap.JsonMap, error) {
	response, err := eTradeClient.ListAccounts()
	if err != nil {
		return nil, err
	}
	accountList, err := etradelib.CreateETradeAccountListFromResponse(response)
	if err != nil {
		return nil, err
	}
	return accountList.AsJsonMap(), nil
}
