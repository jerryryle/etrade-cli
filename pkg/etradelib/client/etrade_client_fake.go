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

	defaultJson []byte
	defaultErr  error
}

func NewClientFake(defaultJson string, defaultError error) *ETradeClientFake {
	clientFake := ETradeClientFake{defaultJson: []byte(defaultJson), defaultErr: defaultError}
	return &clientFake
}

func (c *ETradeClientFake) ListAccounts() ([]byte, error) {
	if c.ListAccountsFn != nil {
		return c.ListAccountsFn()
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) GetAccountBalances(accountIdKey string, realTimeNAV bool) ([]byte, error) {
	if c.GetAccountBalancesFn != nil {
		return c.GetAccountBalancesFn(accountIdKey, realTimeNAV)
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) ListTransactions(
	accountIdKey string, startDate *time.Time, endDate *time.Time, sortOrder SortOrder, marker string, count int,
) ([]byte, error) {
	if c.ListTransactionsFn != nil {
		return c.ListTransactionsFn(accountIdKey, startDate, endDate, sortOrder, marker, count)
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) ListTransactionDetails(accountIdKey string, transactionId string) ([]byte, error) {
	if c.ListTransactionDetailsFn != nil {
		return c.ListTransactionDetailsFn(accountIdKey, transactionId)
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) ViewPortfolio(
	accountIdKey string, count int, sortBy PortfolioSortBy, sortOrder SortOrder, pageNumber int,
	marketSession PortfolioMarketSession, totalsRequired bool, lotsRequired bool, view PortfolioView,
) ([]byte, error) {
	if c.ViewPortfolioFn != nil {
		return c.ViewPortfolioFn(
			accountIdKey, count, sortBy, sortOrder, pageNumber, marketSession, totalsRequired, lotsRequired, view,
		)
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) ListAlerts() ([]byte, error) {
	if c.ListAlertsFn != nil {
		return c.ListAlertsFn()
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) GetQuotes(
	symbols []string, detailFlag QuoteDetailFlag, requireEarningsDate bool, skipMiniOptionsCheck bool,
) ([]byte, error) {
	if c.GetQuotesFn != nil {
		return c.GetQuotesFn(symbols, detailFlag, requireEarningsDate, skipMiniOptionsCheck)
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) LookupProduct(search string) ([]byte, error) {
	if c.LookupProductFn != nil {
		return c.LookupProductFn(search)
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) GetOptionChains(
	symbol string, expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes int,
	includeWeekly, skipAdjusted bool, optionCategory OptionCategory, chainType ChainType, priceType PriceType,
) ([]byte, error) {
	if c.GetOptionChainsFn != nil {
		return c.GetOptionChainsFn(
			symbol,
			expiryYear, expiryMonth, expiryDay,
			strikePriceNear, noOfStrikes, includeWeekly, skipAdjusted,
			optionCategory, chainType, priceType,
		)
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) GetOptionExpireDates(symbol string, expiryType ExpiryType) ([]byte, error) {
	if c.GetOptionExpireDatesFn != nil {
		return c.GetOptionExpireDatesFn(symbol, expiryType)
	} else {
		return c.defaultJson, c.defaultErr
	}
}
