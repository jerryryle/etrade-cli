package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
)

func GetAccountById(client client.ETradeClient, accountId string) (etradelib.ETradeAccount, error) {
	response, err := client.ListAccounts()
	if err != nil {
		return nil, err
	}
	responseMap, err := etradelib.NewNormalizedJsonMap(response)
	if err != nil {
		return nil, err
	}
	accountList, err := etradelib.CreateETradeAccountList(responseMap)
	if err != nil {
		return nil, err
	}
	account := accountList.GetAccountById(accountId)
	if account == nil {
		return nil, fmt.Errorf("account with id %s not found", accountId)
	}
	return account, nil
}
