package etradelib

import "fmt"

type EndpointUrlSet interface {
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

type endpointUrlSet struct {
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

func (s *endpointUrlSet) GetRequestTokenUrl() string {
	return s.getRequestTokenUrl
}

func (s *endpointUrlSet) AuthorizeApplicationUrl() string {
	return s.authorizeApplicationUrl
}

func (s *endpointUrlSet) GetAccessTokenUrl() string {
	return s.getAccessTokenUrl
}

func (s *endpointUrlSet) RenewAccessTokenUrl() string {
	return s.renewAccessTokenUrl
}

func (s *endpointUrlSet) RevokeAccessTokenUrl() string {
	return s.revokeAccessTokenUrl
}

func (s *endpointUrlSet) ListAccountsUrl() string {
	return s.listAccountsUrl
}

func (s *endpointUrlSet) GetAccountBalancesUrl(accountIdKey string) string {
	return fmt.Sprintf(s.getAccountBalancesUrl, accountIdKey)
}

func (s *endpointUrlSet) ListTransactionsUrl(accountIdKey string) string {
	return fmt.Sprintf(s.listTransactionsUrl, accountIdKey)
}

func (s *endpointUrlSet) ListTransactionDetailsUrl(accountIdKey string, transactionId string) string {
	return fmt.Sprintf(s.listTransactionDetailsUrl, accountIdKey, transactionId)
}

func (s *endpointUrlSet) ViewPortfolioUrl(accountIdKey string) string {
	return fmt.Sprintf(s.viewPortfolioUrl, accountIdKey)
}

func (s *endpointUrlSet) ListAlertsUrl() string {
	return s.listAlertsUrl
}

func (s *endpointUrlSet) ListAlertDetailsUrl(alertId string) string {
	return fmt.Sprintf(s.listAlertDetailsUrl, alertId)
}

func (s *endpointUrlSet) DeleteAlertUrl(alertIdList string) string {
	return fmt.Sprintf(s.deleteAlertUrl, alertIdList)
}

func (s *endpointUrlSet) GetQuotesUrl(symbols string) string {
	return fmt.Sprintf(s.getQuotesUrl, symbols)
}

func (s *endpointUrlSet) LookUpProductUrl(search string) string {
	return fmt.Sprintf(s.lookUpProductUrl, search)
}

func (s *endpointUrlSet) GetOptionChainsUrl() string {
	return s.getOptionChainsUrl
}

func (s *endpointUrlSet) GetOptionExpireDatesUrl() string {
	return s.getOptionExpireDatesUrl
}

func (s *endpointUrlSet) ListOrdersUrl(accountIdKey string) string {
	return fmt.Sprintf(s.listOrdersUrl, accountIdKey)
}

func (s *endpointUrlSet) PreviewOrderUrl(accountIdKey string) string {
	return fmt.Sprintf(s.previewOrderUrl, accountIdKey)
}

func (s *endpointUrlSet) PlaceOrderUrl(accountIdKey string) string {
	return fmt.Sprintf(s.placeOrderUrl, accountIdKey)
}

func (s *endpointUrlSet) CancelOrderUrl(accountIdKey string) string {
	return fmt.Sprintf(s.cancelOrderUrl, accountIdKey)
}

func (s *endpointUrlSet) ChangePreviewedOrderUrl(accountIdKey string, orderId string) string {
	return fmt.Sprintf(s.changePreviewedOrderUrl, accountIdKey, orderId)
}

func (s *endpointUrlSet) PlaceChangedOrderUrl(accountIdKey string, orderId string) string {
	return fmt.Sprintf(s.placeChangedOrderUrl, accountIdKey, orderId)
}

var sandboxEndpoints = endpointUrlSet{
	getRequestTokenUrl:        "https://api.etrade.com/oauth/request_token",
	authorizeApplicationUrl:   "https://us.etrade.com/e/t/etws/authorize",
	getAccessTokenUrl:         "https://api.etrade.com/oauth/access_token",
	renewAccessTokenUrl:       "https://api.etrade.com/oauth/renew_access_token",
	revokeAccessTokenUrl:      "https://api.etrade.com/oauth/revoke_access_token",
	listAccountsUrl:           "https://apisb.etrade.com/v1/accounts/list",
	getAccountBalancesUrl:     "https://apisb.etrade.com/v1/accounts/%s/balance",
	listTransactionsUrl:       "https://apisb.etrade.com/v1/accounts/%s/transactions",
	listTransactionDetailsUrl: "https://apisb.etrade.com/v1/accounts/%s/transactions/%s",
	viewPortfolioUrl:          "https://apisb.etrade.com/v1/accounts/%s/portfolio",
	listAlertsUrl:             "https://apisb.etrade.com/v1/user/alerts",
	listAlertDetailsUrl:       "https://apisb.etrade.com/v1/user/alerts/%s",
	deleteAlertUrl:            "https://apisb.etrade.com/v1/user/alerts/%s",
	getQuotesUrl:              "https://apisb.etrade.com/v1/market/quote/%s",
	lookUpProductUrl:          "https://apisb.etrade.com/v1/market/lookup/%s",
	getOptionChainsUrl:        "https://apisb.etrade.com/v1/market/optionchains",
	getOptionExpireDatesUrl:   "https://apisb.etrade.com/v1/market/optionexpiredate",
	listOrdersUrl:             "https://apisb.etrade.com/v1/accounts/%s/orders",
	previewOrderUrl:           "https://apisb.etrade.com/v1/accounts/%s/orders/preview",
	placeOrderUrl:             "https://apisb.etrade.com/v1/accounts/%s/orders/place",
	cancelOrderUrl:            "https://apisb.etrade.com/v1/accounts/%s/orders/cancel",
	changePreviewedOrderUrl:   "https://apisb.etrade.com/v1/accounts/%s/orders/%s/change/preview",
	placeChangedOrderUrl:      "https://apisb.etrade.com/v1/accounts/%s/orders/%s/change/place",
}

var prodEndpoints = endpointUrlSet{
	getRequestTokenUrl:        "https://api.etrade.com/oauth/request_token",
	authorizeApplicationUrl:   "https://us.etrade.com/e/t/etws/authorize",
	getAccessTokenUrl:         "https://api.etrade.com/oauth/access_token",
	renewAccessTokenUrl:       "https://api.etrade.com/oauth/renew_access_token",
	revokeAccessTokenUrl:      "https://api.etrade.com/oauth/revoke_access_token",
	listAccountsUrl:           "https://api.etrade.com/v1/accounts/list",
	getAccountBalancesUrl:     "https://api.etrade.com/v1/accounts/%s/balance",
	listTransactionsUrl:       "https://api.etrade.com/v1/accounts/%s/transactions",
	listTransactionDetailsUrl: "https://api.etrade.com/v1/accounts/%s/transactions/%s",
	viewPortfolioUrl:          "https://api.etrade.com/v1/accounts/%s/portfolio",
	listAlertsUrl:             "https://api.etrade.com/v1/user/alerts",
	listAlertDetailsUrl:       "https://api.etrade.com/v1/user/alerts/%s",
	deleteAlertUrl:            "https://api.etrade.com/v1/user/alerts/%s",
	getQuotesUrl:              "https://api.etrade.com/v1/market/quote/%s",
	lookUpProductUrl:          "https://api.etrade.com/v1/market/lookup/%s",
	getOptionChainsUrl:        "https://api.etrade.com/v1/market/optionchains",
	getOptionExpireDatesUrl:   "https://api.etrade.com/v1/market/optionexpiredate",
	listOrdersUrl:             "https://api.etrade.com/v1/accounts/%s/orders",
	previewOrderUrl:           "https://api.etrade.com/v1/accounts/%s/orders/preview",
	placeOrderUrl:             "https://api.etrade.com/v1/accounts/%s/orders/place",
	cancelOrderUrl:            "https://api.etrade.com/v1/accounts/%s/orders/cancel",
	changePreviewedOrderUrl:   "https://api.etrade.com/v1/accounts/%s/orders/%s/change/preview",
	placeChangedOrderUrl:      "https://api.etrade.com/v1/accounts/%s/orders/%s/change/place",
}

func NewEndpointUrlSet(sandbox bool) EndpointUrlSet {
	if sandbox {
		return &sandboxEndpoints
	} else {
		return &prodEndpoints
	}
}
