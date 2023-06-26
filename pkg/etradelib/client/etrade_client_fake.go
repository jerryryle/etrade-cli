package client

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"time"
)

type ListAccountsFn func() ([]byte, error)

type GetAccountBalancesFn func(accountIdKey string, realTimeNAV bool) ([]byte, error)

type ListTransactionsFn func(
	accountIdKey string, startDate *time.Time, endDate *time.Time, sortOrder constants.SortOrder, marker string,
	count int,
) ([]byte, error)

type ListTransactionDetailsFn func(accountIdKey string, transactionId string) ([]byte, error)

type ViewPortfolioFn func(
	accountIdKey string, count int, sortBy constants.PortfolioSortBy, sortOrder constants.SortOrder, pageNumber string,
	marketSession constants.MarketSession, totalsRequired bool, lotsRequired bool, view constants.PortfolioView,
) ([]byte, error)

type ListPositionLotsDetailsFn func(accountIdKey string, positionId int64) ([]byte, error)

type ListAlertsFn func(
	count int, category constants.AlertCategory, status constants.AlertStatus, sort constants.SortOrder, search string,
) ([]byte, error)

type ListAlertDetailsFn func(alertId string, htmlTags bool) ([]byte, error)

type GetQuotesFn func(
	symbols []string, detailFlag constants.QuoteDetailFlag, requireEarningsDate bool, skipMiniOptionsCheck bool,
) ([]byte, error)

type LookupProductFn func(search string) ([]byte, error)

type GetOptionChainsFn func(
	symbol string, expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes int,
	includeWeekly, skipAdjusted bool, optionCategory constants.OptionCategory, chainType constants.OptionChainType,
	priceType constants.OptionPriceType,
) ([]byte, error)

type GetOptionExpireDatesFn func(symbol string, expiryType constants.OptionExpiryType) ([]byte, error)

type ListOrdersFn func(
	accountIdKey string, marker string, count int, status constants.OrderStatus, fromDate *time.Time, toDate *time.Time,
	symbols []string, securityType constants.OrderSecurityType, transactionType constants.OrderTransactionType,
	marketSession constants.MarketSession,
) ([]byte, error)

type ETradeClientFake struct {
	ListAccountsFn            ListAccountsFn
	GetAccountBalancesFn      GetAccountBalancesFn
	ListTransactionsFn        ListTransactionsFn
	ListTransactionDetailsFn  ListTransactionDetailsFn
	ViewPortfolioFn           ViewPortfolioFn
	ListPositionLotsDetailsFn ListPositionLotsDetailsFn
	ListAlertsFn              ListAlertsFn
	ListAlertDetailsFn        ListAlertDetailsFn
	GetQuotesFn               GetQuotesFn
	LookupProductFn           LookupProductFn
	GetOptionChainsFn         GetOptionChainsFn
	GetOptionExpireDatesFn    GetOptionExpireDatesFn
	ListOrdersFn              ListOrdersFn

	defaultJson []byte
	defaultErr  error
}

func NewClientFake(defaultJson string, defaultError error) ETradeClient {
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
	accountIdKey string, startDate *time.Time, endDate *time.Time, sortOrder constants.SortOrder, marker string,
	count int,
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
	accountIdKey string, count int, sortBy constants.PortfolioSortBy, sortOrder constants.SortOrder, pageNumber string,
	marketSession constants.MarketSession, totalsRequired bool, lotsRequired bool, view constants.PortfolioView,
) ([]byte, error) {
	if c.ViewPortfolioFn != nil {
		return c.ViewPortfolioFn(
			accountIdKey, count, sortBy, sortOrder, pageNumber, marketSession, totalsRequired, lotsRequired, view,
		)
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) ListPositionLotsDetails(accountIdKey string, positionId int64) ([]byte, error) {
	if c.ListPositionLotsDetailsFn != nil {
		return c.ListPositionLotsDetailsFn(accountIdKey, positionId)
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) ListAlerts(
	count int, category constants.AlertCategory, status constants.AlertStatus, sort constants.SortOrder, search string,
) ([]byte, error) {
	if c.ListAlertsFn != nil {
		return c.ListAlertsFn(count, category, status, sort, search)
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) ListAlertDetails(alertId string, htmlTags bool) ([]byte, error) {
	if c.ListAlertDetailsFn != nil {
		return c.ListAlertDetailsFn(alertId, htmlTags)
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) GetQuotes(
	symbols []string, detailFlag constants.QuoteDetailFlag, requireEarningsDate bool, skipMiniOptionsCheck bool,
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
	includeWeekly, skipAdjusted bool, optionCategory constants.OptionCategory, chainType constants.OptionChainType,
	priceType constants.OptionPriceType,
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

func (c *ETradeClientFake) GetOptionExpireDates(symbol string, expiryType constants.OptionExpiryType) ([]byte, error) {
	if c.GetOptionExpireDatesFn != nil {
		return c.GetOptionExpireDatesFn(symbol, expiryType)
	} else {
		return c.defaultJson, c.defaultErr
	}
}

func (c *ETradeClientFake) ListOrders(
	accountIdKey string, marker string, count int, status constants.OrderStatus, fromDate *time.Time, toDate *time.Time,
	symbols []string, securityType constants.OrderSecurityType, transactionType constants.OrderTransactionType,
	marketSession constants.MarketSession,
) ([]byte, error) {
	if c.ListOrdersFn != nil {
		return c.ListOrdersFn(
			accountIdKey, marker, count, status, fromDate, toDate, symbols, securityType, transactionType,
			marketSession,
		)
	} else {
		return c.defaultJson, c.defaultErr
	}
}
