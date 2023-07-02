package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func ViewPortfolio(
	eTradeClient client.ETradeClient, accountId string, sortBy constants.PortfolioSortBy, sortOrder constants.SortOrder,
	marketSession constants.MarketSession, totalsRequired bool, portfolioView constants.PortfolioView, withLots bool,
) (jsonmap.JsonMap, error) {
	// This determines how many portfolio items will be retrieved in each
	// request. This should normally be set to the max for efficiency, but can
	// be lowered to test the pagination logic.
	const countPerRequest = constants.PortfolioMaxCount

	account, err := GetAccountById(eTradeClient, accountId)
	if err != nil {
		return nil, err
	}

	response, err := eTradeClient.ViewPortfolio(
		account.GetIdKey(), countPerRequest, sortBy, sortOrder, "", marketSession, totalsRequired, true, portfolioView,
	)
	if err != nil {
		return nil, err
	}

	positionList, err := etradelib.CreateETradePositionListFromResponse(response)
	if err != nil {
		return nil, err
	}

	for positionList.NextPage() != "" {
		response, err = eTradeClient.ViewPortfolio(
			account.GetIdKey(), countPerRequest, sortBy, sortOrder, positionList.NextPage(), marketSession,
			totalsRequired, true, portfolioView,
		)
		if err != nil {
			return nil, err
		}
		err = positionList.AddPageFromResponse(response)
		if err != nil {
			return nil, err
		}
	}

	if withLots {
		for _, position := range positionList.GetAllPositions() {
			response, err = eTradeClient.ListPositionLotsDetails(account.GetIdKey(), position.GetId())
			if err != nil {
				return nil, err
			}
			err = position.AddLotsFromResponse(response)
			if err != nil {
				return nil, err
			}
		}
	}
	return positionList.AsJsonMap(), nil
}
