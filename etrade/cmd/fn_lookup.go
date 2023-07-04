package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func Lookup(eTradeClient client.ETradeClient, search string) (jsonmap.JsonMap, error) {
	response, err := eTradeClient.LookupProduct(search)
	if err != nil {
		return nil, err
	}
	lookupResultsList, err := etradelib.CreateETradeLookupResultListFromResponse(response)
	if err != nil {
		return nil, err
	}
	return lookupResultsList.AsJsonMap(), nil
}
