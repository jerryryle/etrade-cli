package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func GetQuotes(
	eTradeClient client.ETradeClient, symbols []string, detail constants.QuoteDetailFlag, requireEarningsDate bool,
	skipMiniOptionsCheck bool,
) (
	jsonmap.JsonMap, error,
) {
	response, err := eTradeClient.GetQuotes(symbols, detail, requireEarningsDate, skipMiniOptionsCheck)
	if err != nil {
		return nil, err
	}
	quoteList, err := etradelib.CreateETradeQuoteListFromResponse(response)
	if err != nil {
		return nil, err
	}
	return quoteList.AsJsonMap(), nil
}
