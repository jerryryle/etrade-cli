package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func DeleteAlerts(eTradeClient client.ETradeClient, alertIds []string) (jsonmap.JsonMap, error) {
	_, err := eTradeClient.DeleteAlerts(alertIds)
	if err != nil {
		return nil, fmt.Errorf("requested Alert Id(s) may not exist (%w)", err)
	}
	response, err := eTradeClient.ListAlerts(
		constants.AlertsMaxCount, constants.AlertCategoryNil, constants.AlertStatusNil, constants.SortOrderNil, "",
	)
	if err != nil {
		return nil, err
	}
	alertsList, err := etradelib.CreateETradeAlertListFromResponse(response)
	if err != nil {
		return nil, err
	}
	return alertsList.AsJsonMap(), nil
}
