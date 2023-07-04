package client

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/stretchr/testify/mock"
	"time"
)

type ETradeClientMock struct {
	mock.Mock
}

func (c *ETradeClientMock) Authenticate() ([]byte, error) {
	args := c.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) Verify(verifyKey string) ([]byte, error) {
	args := c.Called(verifyKey)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) GetKeys() (
	consumerKey string, consumerSecret string, accessToken string, accessSecret string,
) {
	args := c.Called()
	return args.String(0), args.String(1), args.String(2), args.String(3)
}

func (c *ETradeClientMock) ListAccounts() ([]byte, error) {
	args := c.Called()
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) GetAccountBalances(accountIdKey string, realTimeNAV bool) ([]byte, error) {
	args := c.Called(accountIdKey, realTimeNAV)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) ListTransactions(
	accountIdKey string, startDate *time.Time, endDate *time.Time, sortOrder constants.SortOrder, marker string,
	count int,
) ([]byte, error) {
	args := c.Called(accountIdKey, startDate, endDate, sortOrder, marker, count)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) ListTransactionDetails(accountIdKey string, transactionId string) ([]byte, error) {
	args := c.Called(accountIdKey, transactionId)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) ViewPortfolio(
	accountIdKey string, count int, sortBy constants.PortfolioSortBy, sortOrder constants.SortOrder, pageNumber string,
	marketSession constants.MarketSession, totalsRequired bool, lotsRequired bool, view constants.PortfolioView,
) ([]byte, error) {
	args := c.Called(
		accountIdKey, count, sortBy, sortOrder, pageNumber, marketSession, totalsRequired, lotsRequired, view,
	)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) ListPositionLotsDetails(accountIdKey string, positionId int64) ([]byte, error) {
	args := c.Called(accountIdKey, positionId)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) ListAlerts(
	count int, category constants.AlertCategory, status constants.AlertStatus, sort constants.SortOrder, search string,
) ([]byte, error) {
	args := c.Called(count, category, status, sort, search)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) ListAlertDetails(alertId string, htmlTags bool) ([]byte, error) {
	args := c.Called(alertId, htmlTags)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) DeleteAlerts(alertIds []string) ([]byte, error) {
	args := c.Called(alertIds)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) GetQuotes(
	symbols []string, detailFlag constants.QuoteDetailFlag, requireEarningsDate bool, skipMiniOptionsCheck bool,
) ([]byte, error) {
	args := c.Called(symbols, detailFlag, requireEarningsDate, skipMiniOptionsCheck)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) LookupProduct(search string) ([]byte, error) {
	args := c.Called(search)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) GetOptionChains(
	symbol string, expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes int,
	includeWeekly, skipAdjusted bool, optionCategory constants.OptionCategory, chainType constants.OptionChainType,
	priceType constants.OptionPriceType,
) ([]byte, error) {
	args := c.Called(
		symbol, expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes, includeWeekly, skipAdjusted,
		optionCategory, chainType, priceType,
	)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) GetOptionExpireDates(symbol string, expiryType constants.OptionExpiryType) ([]byte, error) {
	args := c.Called(symbol, expiryType)
	return args.Get(0).([]byte), args.Error(1)
}

func (c *ETradeClientMock) ListOrders(
	accountIdKey string, marker string, count int, status constants.OrderStatus, fromDate *time.Time, toDate *time.Time,
	symbols []string, securityType constants.OrderSecurityType, transactionType constants.OrderTransactionType,
	marketSession constants.MarketSession,
) ([]byte, error) {
	args := c.Called(
		accountIdKey, marker, count, status, fromDate, toDate, symbols, securityType, transactionType, marketSession,
	)
	return args.Get(0).([]byte), args.Error(1)
}
