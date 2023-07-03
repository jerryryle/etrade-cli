package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func DeleteAlerts(eTradeClient client.ETradeClient, alertIds []string) (jsonmap.JsonMap, error) {
	response, err := eTradeClient.DeleteAlerts(alertIds)
	if err != nil {
		return nil, err
	}
	deleteAlerts, err := etradelib.CreateETradeDeleteAlertsFromResponse(response)
	if err != nil {
		return nil, err
	}
	return deleteAlerts.AsJsonMap(), nil
}
