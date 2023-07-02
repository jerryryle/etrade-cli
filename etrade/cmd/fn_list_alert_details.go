package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func ListAlertDetails(eTradeClient client.ETradeClient, alertId string) (jsonmap.JsonMap, error) {
	response, err := eTradeClient.ListAlertDetails(alertId, false)
	if err != nil {
		return nil, err
	}
	alertDetails, err := etradelib.CreateETradeAlertDetailsFromResponse(response)
	if err != nil {
		return nil, err
	}
	return alertDetails.AsJsonMap(), nil
}
