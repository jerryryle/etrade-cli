package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
)

type ListAccountsFn func() (*responses.AccountListResponse, error)

type eTradeClientFake struct {
	ListAccountsFn ListAccountsFn
}

func (c *eTradeClientFake) ListAccounts() (*responses.AccountListResponse, error) {
	return c.ListAccountsFn()
}
