package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"time"
)

func ListOrders(
	eTradeClient client.ETradeClient, accountId string, fromDate *time.Time, toDate *time.Time, symbols []string,
	securityType constants.OrderSecurityType, transactionType constants.OrderTransactionType,
	marketSession constants.MarketSession,
) (jsonmap.JsonMap, error) {
	// This determines how many order items will be retrieved in each request.
	// This should normally be set to the max for efficiency, but can be
	// lowered to test the pagination logic.
	const countPerRequest = constants.OrdersMaxCount

	account, err := GetAccountById(eTradeClient, accountId)
	if err != nil {
		return nil, err
	}
	response, err := eTradeClient.ListOrders(
		account.GetIdKey(), "", countPerRequest, constants.OrderStatusNil, fromDate, toDate, symbols, securityType,
		transactionType, marketSession,
	)
	if err != nil {
		return nil, err
	}
	orderList, err := etradelib.CreateETradeOrderListFromResponse(response)
	if err != nil {
		return nil, err
	}

	for orderList.NextPage() != "" {
		response, err = eTradeClient.ListOrders(
			account.GetIdKey(), orderList.NextPage(), countPerRequest, constants.OrderStatusNil, fromDate, toDate,
			symbols, securityType, transactionType, marketSession,
		)
		if err != nil {
			return nil, err
		}
		err = orderList.AddPageFromResponse(response)
		if err != nil {
			return nil, err
		}
	}
	return orderList.AsJsonMap(), nil
}
