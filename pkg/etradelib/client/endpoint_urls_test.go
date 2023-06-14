package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSandboxUrls(t *testing.T) {
	var urls = GetEndpointUrls(false)

	assert.Equal(t,
		"https://apisb.etrade.com/oauth/request_token",
		urls.GetRequestTokenUrl())
	assert.Equal(t,
		"https://us.etrade.com/e/t/etws/authorize",
		urls.AuthorizeApplicationUrl())
	assert.Equal(t,
		"https://apisb.etrade.com/oauth/access_token",
		urls.GetAccessTokenUrl())
	assert.Equal(t,
		"https://apisb.etrade.com/oauth/renew_access_token",
		urls.RenewAccessTokenUrl())
	assert.Equal(t,
		"https://apisb.etrade.com/oauth/revoke_access_token",
		urls.RevokeAccessTokenUrl())
	assert.Equal(t,
		"https://apisb.etrade.com/v1/accounts/list",
		urls.ListAccountsUrl())
	assert.Equal(t,
		"https://apisb.etrade.com/v1/accounts/1234/balance",
		urls.GetAccountBalancesUrl("1234"))
	assert.Equal(t,
		"https://apisb.etrade.com/v1/accounts/1234/transactions",
		urls.ListTransactionsUrl("1234"))
	assert.Equal(t,
		"https://apisb.etrade.com/v1/accounts/1234/portfolio",
		urls.ViewPortfolioUrl("1234"))
	assert.Equal(t,
		"https://apisb.etrade.com/v1/user/alerts",
		urls.ListAlertsUrl())
	assert.Equal(t,
		"https://apisb.etrade.com/v1/user/alerts/1234",
		urls.ListAlertDetailsUrl("1234"))
	assert.Equal(t,
		"https://apisb.etrade.com/v1/user/alerts/1234,5678",
		urls.DeleteAlertUrl("1234,5678"))
	assert.Equal(t,
		"https://apisb.etrade.com/v1/market/quote/FLIP,FLOP",
		urls.GetQuotesUrl("FLIP,FLOP"))
	assert.Equal(t,
		"https://apisb.etrade.com/v1/market/lookup/FLIP",
		urls.LookUpProductUrl("FLIP"))
	assert.Equal(t,
		"https://apisb.etrade.com/v1/market/optionchains",
		urls.GetOptionChainsUrl())
	assert.Equal(t,
		"https://apisb.etrade.com/v1/market/optionexpiredate",
		urls.GetOptionExpireDatesUrl())
	assert.Equal(t,
		"https://apisb.etrade.com/v1/accounts/1234/orders",
		urls.ListOrdersUrl("1234"))
	assert.Equal(t,
		"https://apisb.etrade.com/v1/accounts/1234/orders/preview",
		urls.PreviewOrderUrl("1234"))
	assert.Equal(t,
		"https://apisb.etrade.com/v1/accounts/1234/orders/place",
		urls.PlaceOrderUrl("1234"))
	assert.Equal(t,
		"https://apisb.etrade.com/v1/accounts/1234/orders/cancel",
		urls.CancelOrderUrl("1234"))
	assert.Equal(t,
		"https://apisb.etrade.com/v1/accounts/1234/orders/5678/change/preview",
		urls.ChangePreviewedOrderUrl("1234", "5678"))
	assert.Equal(t,
		"https://apisb.etrade.com/v1/accounts/1234/orders/5678/change/place",
		urls.PlaceChangedOrderUrl("1234", "5678"))
}

func TestProductionUrls(t *testing.T) {
	var urls = GetEndpointUrls(true)

	assert.Equal(t,
		"https://api.etrade.com/oauth/request_token",
		urls.GetRequestTokenUrl())
	assert.Equal(t,
		"https://us.etrade.com/e/t/etws/authorize",
		urls.AuthorizeApplicationUrl())
	assert.Equal(t,
		"https://api.etrade.com/oauth/access_token",
		urls.GetAccessTokenUrl())
	assert.Equal(t,
		"https://api.etrade.com/oauth/renew_access_token",
		urls.RenewAccessTokenUrl())
	assert.Equal(t,
		"https://api.etrade.com/oauth/revoke_access_token",
		urls.RevokeAccessTokenUrl())
	assert.Equal(t,
		"https://api.etrade.com/v1/accounts/list",
		urls.ListAccountsUrl())
	assert.Equal(t,
		"https://api.etrade.com/v1/accounts/1234/balance",
		urls.GetAccountBalancesUrl("1234"))
	assert.Equal(t,
		"https://api.etrade.com/v1/accounts/1234/transactions",
		urls.ListTransactionsUrl("1234"))
	assert.Equal(t,
		"https://api.etrade.com/v1/accounts/1234/portfolio",
		urls.ViewPortfolioUrl("1234"))
	assert.Equal(t,
		"https://api.etrade.com/v1/user/alerts",
		urls.ListAlertsUrl())
	assert.Equal(t,
		"https://api.etrade.com/v1/user/alerts/1234",
		urls.ListAlertDetailsUrl("1234"))
	assert.Equal(t,
		"https://api.etrade.com/v1/user/alerts/1234,5678",
		urls.DeleteAlertUrl("1234,5678"))
	assert.Equal(t,
		"https://api.etrade.com/v1/market/quote/FLIP,FLOP",
		urls.GetQuotesUrl("FLIP,FLOP"))
	assert.Equal(t,
		"https://api.etrade.com/v1/market/lookup/FLIP",
		urls.LookUpProductUrl("FLIP"))
	assert.Equal(t,
		"https://api.etrade.com/v1/market/optionchains",
		urls.GetOptionChainsUrl())
	assert.Equal(t,
		"https://api.etrade.com/v1/market/optionexpiredate",
		urls.GetOptionExpireDatesUrl())
	assert.Equal(t,
		"https://api.etrade.com/v1/accounts/1234/orders",
		urls.ListOrdersUrl("1234"))
	assert.Equal(t,
		"https://api.etrade.com/v1/accounts/1234/orders/preview",
		urls.PreviewOrderUrl("1234"))
	assert.Equal(t,
		"https://api.etrade.com/v1/accounts/1234/orders/place",
		urls.PlaceOrderUrl("1234"))
	assert.Equal(t,
		"https://api.etrade.com/v1/accounts/1234/orders/cancel",
		urls.CancelOrderUrl("1234"))
	assert.Equal(t,
		"https://api.etrade.com/v1/accounts/1234/orders/5678/change/preview",
		urls.ChangePreviewedOrderUrl("1234", "5678"))
	assert.Equal(t,
		"https://api.etrade.com/v1/accounts/1234/orders/5678/change/place",
		urls.PlaceChangedOrderUrl("1234", "5678"))
}