package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
)

type ListAccountsFn func() (*responses.AccountListResponse, error)
type ListAlertsFn func() (*responses.AlertsResponse, error)
type GetQuotesFn func(symbols []string, detailFlag QuoteDetailFlag) (*responses.QuoteResponse, error)

type eTradeClientFake struct {
	ListAccountsFn ListAccountsFn
	ListAlertsFn   ListAlertsFn
	GetQuotesFn    GetQuotesFn
}

func (c *eTradeClientFake) ListAccounts() (*responses.AccountListResponse, error) {
	return c.ListAccountsFn()
}

func (c *eTradeClientFake) ListAlerts() (*responses.AlertsResponse, error) {
	return c.ListAlertsFn()
}

func (c *eTradeClientFake) GetQuotes(symbols []string, detailFlag QuoteDetailFlag) (*responses.QuoteResponse, error) {
	return c.GetQuotesFn(symbols, detailFlag)
}
