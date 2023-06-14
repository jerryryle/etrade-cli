package client

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
)

type ListAccountsFn func() (*responses.AccountListResponse, error)
type ListAlertsFn func() (*responses.AlertsResponse, error)
type GetQuotesFn func(symbols []string, detailFlag QuoteDetailFlag) (*responses.QuoteResponse, error)
type LookupProductFn func(search string) (*responses.LookupResponse, error)
type GetOptionChainsFn func(symbol string,
	expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes int,
	includeWeekly, skipAdjusted bool,
	optionCategory OptionCategory, chainType ChainType, priceType PriceType) (*responses.OptionChainResponse, error)
type GetOptionExpireDatesFn func(symbol string, expiryType ExpiryType) (*responses.OptionExpireDateResponse, error)

type ETradeClientFake struct {
	ListAccountsFn         ListAccountsFn
	ListAlertsFn           ListAlertsFn
	GetQuotesFn            GetQuotesFn
	LookupProductFn        LookupProductFn
	GetOptionChainsFn      GetOptionChainsFn
	GetOptionExpireDatesFn GetOptionExpireDatesFn
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

func (c *ETradeClientFake) GetOptionChains(symbol string,
	expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes int,
	includeWeekly, skipAdjusted bool,
	optionCategory OptionCategory, chainType ChainType, priceType PriceType) (*responses.OptionChainResponse, error) {
	return c.GetOptionChainsFn(
		symbol,
		expiryYear, expiryMonth, expiryDay,
		strikePriceNear, noOfStrikes, includeWeekly, skipAdjusted,
		optionCategory, chainType, priceType)
}

func (c *ETradeClientFake) GetOptionExpireDates(symbol string, expiryType ExpiryType) (*responses.OptionExpireDateResponse, error) {
	return c.GetOptionExpireDatesFn(symbol, expiryType)
}
