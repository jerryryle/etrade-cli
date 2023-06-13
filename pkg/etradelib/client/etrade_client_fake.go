package client

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
)

type ListAccountsFn func() (*responses.AccountListResponse, error)
type ListAlertsFn func() (*responses.AlertsResponse, error)
type GetQuotesFn func(symbols []string, detailFlag QuoteDetailFlag) (*responses.QuoteResponse, error)
type LookupProductFn func(search string) (*responses.LookupResponse, error)

type ETradeClientFake struct {
	ListAccountsFn  ListAccountsFn
	ListAlertsFn    ListAlertsFn
	GetQuotesFn     GetQuotesFn
	LookupProductFn LookupProductFn
}

func (c *ETradeClientFake) ListAccounts() (*responses.AccountListResponse, error) {
	return c.ListAccountsFn()
}

func (c *ETradeClientFake) ListAlerts() (*responses.AlertsResponse, error) {
	return c.ListAlertsFn()
}

func (c *ETradeClientFake) GetQuotes(symbols []string, detailFlag QuoteDetailFlag) (*responses.QuoteResponse, error) {
	return c.GetQuotesFn(symbols, detailFlag)
}

func (c *ETradeClientFake) LookupProduct(search string) (*responses.LookupResponse, error) {
	return c.LookupProductFn(search)
}
