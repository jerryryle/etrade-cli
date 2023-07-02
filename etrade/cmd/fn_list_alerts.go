package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func ListAlerts(
	eTradeClient client.ETradeClient, count int, category constants.AlertCategory, status constants.AlertStatus,
	sortOrder constants.SortOrder, search string,
) (jsonmap.JsonMap, error) {
	response, err := eTradeClient.ListAlerts(count, category, status, sortOrder, search)
	if err != nil {
		return nil, err
	}
	alertsList, err := etradelib.CreateETradeAlertListFromResponse(response)
	if err != nil {
		return nil, err
	}
	return alertsList.AsJsonMap(), nil
}
