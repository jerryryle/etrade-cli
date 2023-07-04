package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func ListTransactionDetails(eTradeClient client.ETradeClient, accountId string, transactionId string) (
	jsonmap.JsonMap, error,
) {
	account, err := GetAccountById(eTradeClient, accountId)
	if err != nil {
		return nil, err
	}
	response, err := eTradeClient.ListTransactionDetails(account.GetIdKey(), transactionId)
	if err != nil {
		return nil, err
	}
	transactionDetails, err := etradelib.CreateETradeTransactionDetailsFromResponse(response)
	if err != nil {
		return nil, err
	}
	return transactionDetails.AsJsonMap(), nil
}
