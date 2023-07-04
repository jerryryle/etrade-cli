package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func GetOptionExpireDates(
	eTradeClient client.ETradeClient, symbol string, expiryType constants.OptionExpiryType,
) (jsonmap.JsonMap, error) {
	response, err := eTradeClient.GetOptionExpireDates(symbol, expiryType)
	if err != nil {
		return nil, err
	}
	optionExpireDates, err := etradelib.CreateETradeOptionExpireDateListFromResponse(response)
	if err != nil {
		return nil, err
	}
	return optionExpireDates.AsJsonMap(), nil
}
