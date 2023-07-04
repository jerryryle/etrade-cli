package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"time"
)

func ListTransactions(
	eTradeClient client.ETradeClient, accountId string, startDate *time.Time, endDate *time.Time,
	sortOrder constants.SortOrder,
) (jsonmap.JsonMap, error) {
	// This determines how many transaction items will be retrieved in each
	// request. This should normally be set to the max for efficiency, but can
	// be lowered to test the pagination logic.
	const countPerRequest = constants.TransactionsMaxCount

	account, err := GetAccountById(eTradeClient, accountId)
	if err != nil {
		return nil, err
	}
	response, err := eTradeClient.ListTransactions(
		account.GetIdKey(),
		startDate, endDate, sortOrder, "", countPerRequest,
	)
	if err != nil {
		return nil, err
	}
	transactionList, err := etradelib.CreateETradeTransactionListFromResponse(response)
	if err != nil {
		return nil, err
	}

	for transactionList.NextPage() != "" {
		response, err = eTradeClient.ListTransactions(
			account.GetIdKey(),
			startDate, endDate, sortOrder, transactionList.NextPage(), countPerRequest,
		)
		if err != nil {
			return nil, err
		}
		err = transactionList.AddPageFromResponse(response)
		if err != nil {
			return nil, err
		}
	}

	return transactionList.AsJsonMap(), nil
}
