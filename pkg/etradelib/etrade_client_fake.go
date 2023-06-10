package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
)

type ListAccountsFn func() (*responses.AccountListResponse, error)
type ListAlertsFn func() (*responses.AlertsResponse, error)

type eTradeClientFake struct {
	ListAccountsFn ListAccountsFn
	ListAlertsFn   ListAlertsFn
}

func (c *eTradeClientFake) ListAccounts() (*responses.AccountListResponse, error) {
	return c.ListAccountsFn()
}

func (c *eTradeClientFake) ListAlerts() (*responses.AlertsResponse, error) {
	return c.ListAlertsFn()
}
