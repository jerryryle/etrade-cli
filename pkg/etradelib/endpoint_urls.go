package etradelib

import (
	"fmt"
	"strings"
)

type EndpointUrls interface {
	GetRequestTokenUrl() string
	AuthorizeApplicationUrl() string
	GetAccessTokenUrl() string
	RenewAccessTokenUrl() string
	RevokeAccessTokenUrl() string
	ListAccountsUrl() string
	GetAccountBalancesUrl(accountIdKey string) string
	ListTransactionsUrl(accountIdKey string) string
	ListTransactionDetailsUrl(accountIdKey string, transactionId string) string
	ViewPortfolioUrl(accountIdKey string) string
	ListAlertsUrl() string
	ListAlertDetailsUrl(alertId string) string
	DeleteAlertUrl(alertIdList string) string
	GetQuotesUrl(symbols string) string
	LookUpProductUrl(search string) string
	GetOptionChainsUrl() string
	GetOptionExpireDatesUrl() string
	ListOrdersUrl(accountIdKey string) string
	PreviewOrderUrl(accountIdKey string) string
	PlaceOrderUrl(accountIdKey string) string
	CancelOrderUrl(accountIdKey string) string
	ChangePreviewedOrderUrl(accountIdKey string, orderId string) string
	PlaceChangedOrderUrl(accountIdKey string, orderId string) string
}

type endpointUrls struct {
	getRequestTokenUrl        string
	authorizeApplicationUrl   string
	getAccessTokenUrl         string
	renewAccessTokenUrl       string
	revokeAccessTokenUrl      string
	listAccountsUrl           string
	getAccountBalancesUrl     string
	listTransactionsUrl       string
	listTransactionDetailsUrl string
	viewPortfolioUrl          string
	listAlertsUrl             string
	listAlertDetailsUrl       string
	deleteAlertUrl            string
	getQuotesUrl              string
	lookUpProductUrl          string
	getOptionChainsUrl        string
	getOptionExpireDatesUrl   string
	listOrdersUrl             string
	previewOrderUrl           string
	placeOrderUrl             string
	cancelOrderUrl            string
	changePreviewedOrderUrl   string
	placeChangedOrderUrl      string
}

func (s *endpointUrls) GetRequestTokenUrl() string {
	return s.getRequestTokenUrl
}

func (s *endpointUrls) AuthorizeApplicationUrl() string {
	return s.authorizeApplicationUrl
}

func (s *endpointUrls) GetAccessTokenUrl() string {
	return s.getAccessTokenUrl
}

func (s *endpointUrls) RenewAccessTokenUrl() string {
	return s.renewAccessTokenUrl
}

func (s *endpointUrls) RevokeAccessTokenUrl() string {
	return s.revokeAccessTokenUrl
}

func (s *endpointUrls) ListAccountsUrl() string {
	return s.listAccountsUrl
}

func (s *endpointUrls) GetAccountBalancesUrl(accountIdKey string) string {
	return fmt.Sprintf(s.getAccountBalancesUrl, accountIdKey)
}

func (s *endpointUrls) ListTransactionsUrl(accountIdKey string) string {
	return fmt.Sprintf(s.listTransactionsUrl, accountIdKey)
}

func (s *endpointUrls) ListTransactionDetailsUrl(accountIdKey string, transactionId string) string {
	return fmt.Sprintf(s.listTransactionDetailsUrl, accountIdKey, transactionId)
}

func (s *endpointUrls) ViewPortfolioUrl(accountIdKey string) string {
	return fmt.Sprintf(s.viewPortfolioUrl, accountIdKey)
}

func (s *endpointUrls) ListAlertsUrl() string {
	return s.listAlertsUrl
}

func (s *endpointUrls) ListAlertDetailsUrl(alertId string) string {
	return fmt.Sprintf(s.listAlertDetailsUrl, alertId)
}

func (s *endpointUrls) DeleteAlertUrl(alertIdList string) string {
	return fmt.Sprintf(s.deleteAlertUrl, alertIdList)
}

func (s *endpointUrls) GetQuotesUrl(symbols string) string {
	return fmt.Sprintf(s.getQuotesUrl, symbols)
}

func (s *endpointUrls) LookUpProductUrl(search string) string {
	return fmt.Sprintf(s.lookUpProductUrl, search)
}

func (s *endpointUrls) GetOptionChainsUrl() string {
	return s.getOptionChainsUrl
}

func (s *endpointUrls) GetOptionExpireDatesUrl() string {
	return s.getOptionExpireDatesUrl
}

func (s *endpointUrls) ListOrdersUrl(accountIdKey string) string {
	return fmt.Sprintf(s.listOrdersUrl, accountIdKey)
}

func (s *endpointUrls) PreviewOrderUrl(accountIdKey string) string {
	return fmt.Sprintf(s.previewOrderUrl, accountIdKey)
}

func (s *endpointUrls) PlaceOrderUrl(accountIdKey string) string {
	return fmt.Sprintf(s.placeOrderUrl, accountIdKey)
}

func (s *endpointUrls) CancelOrderUrl(accountIdKey string) string {
	return fmt.Sprintf(s.cancelOrderUrl, accountIdKey)
}

func (s *endpointUrls) ChangePreviewedOrderUrl(accountIdKey string, orderId string) string {
	return fmt.Sprintf(s.changePreviewedOrderUrl, accountIdKey, orderId)
}

func (s *endpointUrls) PlaceChangedOrderUrl(accountIdKey string, orderId string) string {
	return fmt.Sprintf(s.placeChangedOrderUrl, accountIdKey, orderId)
}

const (
	productionUrlBase = "https://api.etrade.com"
	sandboxUrlBase    = "https://apisb.etrade.com"

	getRequestTokenUrlTemplate        = "{B}/oauth/request_token"
	authorizeApplicationUrlTemplate   = "https://us.etrade.com/e/t/etws/authorize"
	getAccessTokenUrlTemplate         = "{B}/oauth/access_token"
	renewAccessTokenUrlTemplate       = "{B}/oauth/renew_access_token"
	revokeAccessTokenUrlTemplate      = "{B}/oauth/revoke_access_token"
	listAccountsUrlTemplate           = "{B}/v1/accounts/list.json"
	getAccountBalancesUrlTemplate     = "{B}/v1/accounts/%s/balance.json"
	listTransactionsUrlTemplate       = "{B}/v1/accounts/%s/transactions.json"
	listTransactionDetailsUrlTemplate = "{B}/v1/accounts/%s/transactions/%s.json"
	viewPortfolioUrlTemplate          = "{B}/v1/accounts/%s/portfolio.json"
	listAlertsUrlTemplate             = "{B}/v1/user/alerts.json"
	listAlertDetailsUrlTemplate       = "{B}/v1/user/alerts/%s.json"
	deleteAlertUrlTemplate            = "{B}/v1/user/alerts/%s.json"
	getQuotesUrlTemplate              = "{B}/v1/market/quote/%s.json"
	lookUpProductUrlTemplate          = "{B}/v1/market/lookup/%s.json"
	getOptionChainsUrlTemplate        = "{B}/v1/market/optionchains.json"
	getOptionExpireDatesUrlTemplate   = "{B}/v1/market/optionexpiredate.json"
	listOrdersUrlTemplate             = "{B}/v1/accounts/%s/orders.json"
	previewOrderUrlTemplate           = "{B}/v1/accounts/%s/orders/preview.json"
	placeOrderUrlTemplate             = "{B}/v1/accounts/%s/orders/place.json"
	cancelOrderUrlTemplate            = "{B}/v1/accounts/%s/orders/cancel.json"
	changePreviewedOrderUrlTemplate   = "{B}/v1/accounts/%s/orders/%s/change/preview.json"
	placeChangedOrderUrlTemplate      = "{B}/v1/accounts/%s/orders/%s/change/place.json"
)

func GetEndpointUrls(production bool) EndpointUrls {
	var urlBase string
	if production {
		urlBase = productionUrlBase
	} else {
		urlBase = sandboxUrlBase
	}
	return &endpointUrls{
		getRequestTokenUrl:        renderUrlTemplateWithBase(getRequestTokenUrlTemplate, urlBase),
		authorizeApplicationUrl:   renderUrlTemplateWithBase(authorizeApplicationUrlTemplate, urlBase),
		getAccessTokenUrl:         renderUrlTemplateWithBase(getAccessTokenUrlTemplate, urlBase),
		renewAccessTokenUrl:       renderUrlTemplateWithBase(renewAccessTokenUrlTemplate, urlBase),
		revokeAccessTokenUrl:      renderUrlTemplateWithBase(revokeAccessTokenUrlTemplate, urlBase),
		listAccountsUrl:           renderUrlTemplateWithBase(listAccountsUrlTemplate, urlBase),
		getAccountBalancesUrl:     renderUrlTemplateWithBase(getAccountBalancesUrlTemplate, urlBase),
		listTransactionsUrl:       renderUrlTemplateWithBase(listTransactionsUrlTemplate, urlBase),
		listTransactionDetailsUrl: renderUrlTemplateWithBase(listTransactionDetailsUrlTemplate, urlBase),
		viewPortfolioUrl:          renderUrlTemplateWithBase(viewPortfolioUrlTemplate, urlBase),
		listAlertsUrl:             renderUrlTemplateWithBase(listAlertsUrlTemplate, urlBase),
		listAlertDetailsUrl:       renderUrlTemplateWithBase(listAlertDetailsUrlTemplate, urlBase),
		deleteAlertUrl:            renderUrlTemplateWithBase(deleteAlertUrlTemplate, urlBase),
		getQuotesUrl:              renderUrlTemplateWithBase(getQuotesUrlTemplate, urlBase),
		lookUpProductUrl:          renderUrlTemplateWithBase(lookUpProductUrlTemplate, urlBase),
		getOptionChainsUrl:        renderUrlTemplateWithBase(getOptionChainsUrlTemplate, urlBase),
		getOptionExpireDatesUrl:   renderUrlTemplateWithBase(getOptionExpireDatesUrlTemplate, urlBase),
		listOrdersUrl:             renderUrlTemplateWithBase(listOrdersUrlTemplate, urlBase),
		previewOrderUrl:           renderUrlTemplateWithBase(previewOrderUrlTemplate, urlBase),
		placeOrderUrl:             renderUrlTemplateWithBase(placeOrderUrlTemplate, urlBase),
		cancelOrderUrl:            renderUrlTemplateWithBase(cancelOrderUrlTemplate, urlBase),
		changePreviewedOrderUrl:   renderUrlTemplateWithBase(changePreviewedOrderUrlTemplate, urlBase),
		placeChangedOrderUrl:      renderUrlTemplateWithBase(placeChangedOrderUrlTemplate, urlBase),
	}
}

func renderUrlTemplateWithBase(urlTemplate string, base string) string {
	return strings.Replace(urlTemplate, "{B}", base, 1)
}
