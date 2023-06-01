package etradelib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSandboxUrls(t *testing.T) {
	assert := assert.New(t)
	var urlSet = NewEndpointUrlSet(true)

	assert.Equal(
		"https://api.etrade.com/oauth/request_token",
		urlSet.GetRequestTokenUrl())
	assert.Equal(
		"https://us.etrade.com/e/t/etws/authorize",
		urlSet.AuthorizeApplicationUrl())
	assert.Equal(
		"https://api.etrade.com/oauth/access_token",
		urlSet.GetAccessTokenUrl())
	assert.Equal(
		"https://api.etrade.com/oauth/renew_access_token",
		urlSet.RenewAccessTokenUrl())
	assert.Equal(
		"https://api.etrade.com/oauth/revoke_access_token",
		urlSet.RevokeAccessTokenUrl())
	assert.Equal(
		"https://apisb.etrade.com/v1/accounts/list",
		urlSet.ListAccountsUrl())
	assert.Equal(
		"https://apisb.etrade.com/v1/accounts/1234/balance",
		urlSet.GetAccountBalancesUrl("1234"))
	assert.Equal(
		"https://apisb.etrade.com/v1/accounts/1234/transactions",
		urlSet.ListTransactionsUrl("1234"))
	assert.Equal(
		"https://apisb.etrade.com/v1/accounts/1234/portfolio",
		urlSet.ViewPortfolioUrl("1234"))
	assert.Equal(
		"https://apisb.etrade.com/v1/user/alerts",
		urlSet.ListAlertsUrl())
	assert.Equal(
		"https://apisb.etrade.com/v1/user/alerts/1234",
		urlSet.ListAlertDetailsUrl("1234"))
	assert.Equal(
		"https://apisb.etrade.com/v1/user/alerts/1234,5678",
		urlSet.DeleteAlertUrl("1234,5678"))
	assert.Equal(
		"https://apisb.etrade.com/v1/market/quote/FLIP,FLOP",
		urlSet.GetQuotesUrl("FLIP,FLOP"))
	assert.Equal(
		"https://apisb.etrade.com/v1/market/lookup/FLIP",
		urlSet.LookUpProductUrl("FLIP"))
	assert.Equal(
		"https://apisb.etrade.com/v1/market/optionchains",
		urlSet.GetOptionChainsUrl())
	assert.Equal(
		"https://apisb.etrade.com/v1/market/optionexpiredate",
		urlSet.GetOptionExpireDatesUrl())
	assert.Equal(
		"https://apisb.etrade.com/v1/accounts/1234/orders",
		urlSet.ListOrdersUrl("1234"))
	assert.Equal(
		"https://apisb.etrade.com/v1/accounts/1234/orders/preview",
		urlSet.PreviewOrderUrl("1234"))
	assert.Equal(
		"https://apisb.etrade.com/v1/accounts/1234/orders/place",
		urlSet.PlaceOrderUrl("1234"))
	assert.Equal(
		"https://apisb.etrade.com/v1/accounts/1234/orders/cancel",
		urlSet.CancelOrderUrl("1234"))
	assert.Equal(
		"https://apisb.etrade.com/v1/accounts/1234/orders/5678/change/preview",
		urlSet.ChangePreviewedOrderUrl("1234", "5678"))
	assert.Equal(
		"https://apisb.etrade.com/v1/accounts/1234/orders/5678/change/place",
		urlSet.PlaceChangedOrderUrl("1234", "5678"))
}

func TestProdUrls(t *testing.T) {
	assert := assert.New(t)
	var urlSet = NewEndpointUrlSet(false)

	assert.Equal(
		"https://api.etrade.com/oauth/request_token",
		urlSet.GetRequestTokenUrl())
	assert.Equal(
		"https://us.etrade.com/e/t/etws/authorize",
		urlSet.AuthorizeApplicationUrl())
	assert.Equal(
		"https://api.etrade.com/oauth/access_token",
		urlSet.GetAccessTokenUrl())
	assert.Equal(
		"https://api.etrade.com/oauth/renew_access_token",
		urlSet.RenewAccessTokenUrl())
	assert.Equal(
		"https://api.etrade.com/oauth/revoke_access_token",
		urlSet.RevokeAccessTokenUrl())
	assert.Equal(
		"https://api.etrade.com/v1/accounts/list",
		urlSet.ListAccountsUrl())
	assert.Equal(
		"https://api.etrade.com/v1/accounts/1234/balance",
		urlSet.GetAccountBalancesUrl("1234"))
	assert.Equal(
		"https://api.etrade.com/v1/accounts/1234/transactions",
		urlSet.ListTransactionsUrl("1234"))
	assert.Equal(
		"https://api.etrade.com/v1/accounts/1234/portfolio",
		urlSet.ViewPortfolioUrl("1234"))
	assert.Equal(
		"https://api.etrade.com/v1/user/alerts",
		urlSet.ListAlertsUrl())
	assert.Equal(
		"https://api.etrade.com/v1/user/alerts/1234",
		urlSet.ListAlertDetailsUrl("1234"))
	assert.Equal(
		"https://api.etrade.com/v1/user/alerts/1234,5678",
		urlSet.DeleteAlertUrl("1234,5678"))
	assert.Equal(
		"https://api.etrade.com/v1/market/quote/FLIP,FLOP",
		urlSet.GetQuotesUrl("FLIP,FLOP"))
	assert.Equal(
		"https://api.etrade.com/v1/market/lookup/FLIP",
		urlSet.LookUpProductUrl("FLIP"))
	assert.Equal(
		"https://api.etrade.com/v1/market/optionchains",
		urlSet.GetOptionChainsUrl())
	assert.Equal(
		"https://api.etrade.com/v1/market/optionexpiredate",
		urlSet.GetOptionExpireDatesUrl())
	assert.Equal(
		"https://api.etrade.com/v1/accounts/1234/orders",
		urlSet.ListOrdersUrl("1234"))
	assert.Equal(
		"https://api.etrade.com/v1/accounts/1234/orders/preview",
		urlSet.PreviewOrderUrl("1234"))
	assert.Equal(
		"https://api.etrade.com/v1/accounts/1234/orders/place",
		urlSet.PlaceOrderUrl("1234"))
	assert.Equal(
		"https://api.etrade.com/v1/accounts/1234/orders/cancel",
		urlSet.CancelOrderUrl("1234"))
	assert.Equal(
		"https://api.etrade.com/v1/accounts/1234/orders/5678/change/preview",
		urlSet.ChangePreviewedOrderUrl("1234", "5678"))
	assert.Equal(
		"https://api.etrade.com/v1/accounts/1234/orders/5678/change/place",
		urlSet.PlaceChangedOrderUrl("1234", "5678"))
}
