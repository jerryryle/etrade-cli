package client

import (
	"time"
)

type ListAccountsFn func() ([]byte, error)

type GetAccountBalancesFn func(accountIdKey string, realTimeNAV bool) ([]byte, error)

type ListTransactionsFn func(
	accountIdKey string, startDate *time.Time, endDate *time.Time, sortOrder SortOrder, marker string, count int,
) ([]byte, error)

type ListTransactionDetailsFn func(accountIdKey string, transactionId string) ([]byte, error)

type ViewPortfolioFn func(
	accountIdKey string, count int, sortBy PortfolioSortBy, sortOrder SortOrder, pageNumber int,
	marketSession PortfolioMarketSession, totalsRequired bool, lotsRequired bool, view PortfolioView,
) ([]byte, error)

type ListAlertsFn func() ([]byte, error)

type GetQuotesFn func(
	symbols []string, detailFlag QuoteDetailFlag, requireEarningsDate bool, skipMiniOptionsCheck bool,
) ([]byte, error)

type LookupProductFn func(search string) ([]byte, error)

type GetOptionChainsFn func(
	symbol string, expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes int,
	includeWeekly, skipAdjusted bool, optionCategory OptionCategory, chainType ChainType, priceType PriceType,
) ([]byte, error)

type GetOptionExpireDatesFn func(symbol string, expiryType ExpiryType) ([]byte, error)

type ETradeClientFake struct {
	ListAccountsFn           ListAccountsFn
	GetAccountBalancesFn     GetAccountBalancesFn
	ListTransactionsFn       ListTransactionsFn
	ListTransactionDetailsFn ListTransactionDetailsFn
	ViewPortfolioFn          ViewPortfolioFn
	ListAlertsFn             ListAlertsFn
	GetQuotesFn              GetQuotesFn
	LookupProductFn          LookupProductFn
	GetOptionChainsFn        GetOptionChainsFn
	GetOptionExpireDatesFn   GetOptionExpireDatesFn
}

func (c *ETradeClientFake) ListAccounts() ([]byte, error) {
	return c.ListAccountsFn()
}

func (c *ETradeClientFake) GetAccountBalances(accountIdKey string, realTimeNAV bool) ([]byte, error) {
	return c.GetAccountBalancesFn(accountIdKey, realTimeNAV)
}

func (c *ETradeClientFake) ListTransactions(
	accountIdKey string, startDate *time.Time, endDate *time.Time, sortOrder SortOrder, marker string, count int,
) ([]byte, error) {
	return c.ListTransactionsFn(accountIdKey, startDate, endDate, sortOrder, marker, count)
}

func (c *ETradeClientFake) ListTransactionDetails(accountIdKey string, transactionId string) ([]byte, error) {
	return c.ListTransactionDetailsFn(accountIdKey, transactionId)
}

func (c *ETradeClientFake) ViewPortfolio(
	accountIdKey string, count int, sortBy PortfolioSortBy, sortOrder SortOrder, pageNumber int,
	marketSession PortfolioMarketSession, totalsRequired bool, lotsRequired bool, view PortfolioView,
) ([]byte, error) {
	return c.ViewPortfolioFn(
		accountIdKey, count, sortBy, sortOrder, pageNumber, marketSession, totalsRequired, lotsRequired, view,
	)
}

func (c *ETradeClientFake) ListAlerts() ([]byte, error) {
	return c.ListAlertsFn()
}

func (c *ETradeClientFake) GetQuotes(
	symbols []string, detailFlag QuoteDetailFlag, requireEarningsDate bool, skipMiniOptionsCheck bool,
) ([]byte, error) {
	return c.GetQuotesFn(symbols, detailFlag, requireEarningsDate, skipMiniOptionsCheck)
}

func (c *ETradeClientFake) LookupProduct(search string) ([]byte, error) {
	return c.LookupProductFn(search)
}

func (c *ETradeClientFake) GetOptionChains(
	symbol string, expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes int,
	includeWeekly, skipAdjusted bool, optionCategory OptionCategory, chainType ChainType, priceType PriceType,
) ([]byte, error) {
	return c.GetOptionChainsFn(
		symbol,
		expiryYear, expiryMonth, expiryDay,
		strikePriceNear, noOfStrikes, includeWeekly, skipAdjusted,
		optionCategory, chainType, priceType,
	)
}

func (c *ETradeClientFake) GetOptionExpireDates(symbol string, expiryType ExpiryType) ([]byte, error) {
	return c.GetOptionExpireDatesFn(symbol, expiryType)
}
