package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func GetOptionChains(
	eTradeClient client.ETradeClient, symbol string, expiryYear int, expiryMonth int, expiryDay int,
	strikePriceNear int, noOfStrikes int, includeWeekly bool, skipAdjusted bool,
	optionCategory constants.OptionCategory, chainType constants.OptionChainType, priceType constants.OptionPriceType,
) (jsonmap.JsonMap, error) {
	response, err := eTradeClient.GetOptionChains(
		symbol, expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes, includeWeekly, skipAdjusted,
		optionCategory, chainType, priceType,
	)
	if err != nil {
		return nil, err
	}
	optionChains, err := etradelib.CreateETradeOptionChainPairListFromResponse(response)
	if err != nil {
		return nil, err
	}
	return optionChains.AsJsonMap(), nil
}
