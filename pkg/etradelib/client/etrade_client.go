package client

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ETradeClient interface {
	ListAccounts() ([]byte, error)

	GetAccountBalances(accountIdKey string, realTimeNAV bool) ([]byte, error)

	ListTransactions(
		accountIdKey string, startDate *time.Time, endDate *time.Time, sortOrder constants.SortOrder, marker string,
		count int,
	) ([]byte, error)

	ListTransactionDetails(accountIdKey string, transactionId string) ([]byte, error)

	ViewPortfolio(
		accountIdKey string, count int, sortBy constants.PortfolioSortBy, sortOrder constants.SortOrder,
		pageNumber string, marketSession constants.MarketSession, totalsRequired bool, lotsRequired bool,
		view constants.PortfolioView,
	) ([]byte, error)

	ListPositionLotsDetails(accountIdKey string, positionId int64) ([]byte, error)

	ListAlerts(
		count int, category constants.AlertCategory, status constants.AlertStatus, sort constants.SortOrder,
		search string,
	) ([]byte, error)

	ListAlertDetails(alertId string, htmlTags bool) ([]byte, error)

	GetQuotes(
		symbols []string, detailFlag constants.QuoteDetailFlag, requireEarningsDate bool, skipMiniOptionsCheck bool,
	) ([]byte, error)

	LookupProduct(search string) ([]byte, error)

	GetOptionChains(
		symbol string, expiryYear int, expiryMonth int, expiryDay int, strikePriceNear int, noOfStrikes int,
		includeWeekly bool, skipAdjusted bool, optionCategory constants.OptionCategory,
		chainType constants.OptionChainType,
		priceType constants.OptionPriceType,
	) ([]byte, error)

	GetOptionExpireDates(symbol string, expiryType constants.OptionExpiryType) ([]byte, error)

	ListOrders(
		accountIdKey string, marker string, count int, status constants.OrderStatus, fromDate *time.Time,
		toDate *time.Time, symbols []string, securityType constants.OrderSecurityType,
		transactionType constants.OrderTransactionType, marketSession constants.MarketSession,
	) ([]byte, error)
}

type eTradeClient struct {
	urls       EndpointUrls
	httpClient *http.Client
	Logger     *slog.Logger
}

func CreateETradeClient(urls EndpointUrls, httpClient *http.Client, logger *slog.Logger) ETradeClient {
	return &eTradeClient{
		urls:       urls,
		httpClient: httpClient,
		Logger:     logger,
	}
}

const queryDateLayout = "01022006"

func (c *eTradeClient) ListAccounts() ([]byte, error) {
	response, err := c.doRequest("GET", c.urls.ListAccountsUrl(), nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) GetAccountBalances(accountIdKey string, realTimeNAV bool) ([]byte, error) {
	if accountIdKey == "" {
		return nil, errors.New("accountIdKey not provided")
	}
	queryValues := url.Values{}
	queryValues.Add("instType", "BROKERAGE")
	queryValues.Add("realTimeNAV", fmt.Sprintf("%t", realTimeNAV))

	response, err := c.doRequest("GET", c.urls.GetAccountBalancesUrl(accountIdKey), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) ListTransactions(
	accountIdKey string, startDate *time.Time, endDate *time.Time, sortOrder constants.SortOrder, marker string,
	count int,
) ([]byte, error) {
	if accountIdKey == "" {
		return nil, errors.New("accountIdKey not provided")
	}
	queryValues := url.Values{}
	if startDate != nil {
		queryValues.Add("startDate", startDate.Format(queryDateLayout))
	}
	if endDate != nil {
		queryValues.Add("endDate", endDate.Format(queryDateLayout))
	}
	if sortOrder != constants.SortOrderNil {
		queryValues.Add("sortOrder", sortOrder.String())
	}
	if marker != "" {
		queryValues.Add("marker", marker)
	}
	if count >= 0 {
		queryValues.Add("count", fmt.Sprintf("%d", count))
	}

	response, err := c.doRequest("GET", c.urls.ListTransactionsUrl(accountIdKey), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) ListTransactionDetails(accountIdKey string, transactionId string) ([]byte, error) {
	if accountIdKey == "" {
		return nil, errors.New("accountIdKey not provided")
	}
	if transactionId == "" {
		return nil, errors.New("transactionId not provided")
	}
	response, err := c.doRequest("GET", c.urls.ListTransactionDetailsUrl(accountIdKey, transactionId), nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) ViewPortfolio(
	accountIdKey string, count int, sortBy constants.PortfolioSortBy, sortOrder constants.SortOrder, pageNumber string,
	marketSession constants.MarketSession, totalsRequired bool, lotsRequired bool, view constants.PortfolioView,
) ([]byte, error) {
	if accountIdKey == "" {
		return nil, errors.New("accountIdKey not provided")
	}
	if count > constants.PortfolioMaxCount {
		return nil, fmt.Errorf(
			"count of %d requested, which exceeds the maximum of %d", count, constants.PortfolioMaxCount,
		)
	}
	queryValues := url.Values{}
	if count >= 0 {
		queryValues.Add("count", fmt.Sprintf("%d", count))
	}
	if pageNumber != "" {
		queryValues.Add("pageNumber", pageNumber)
	}
	queryValues.Add("totalsRequired", fmt.Sprintf("%t", totalsRequired))
	queryValues.Add("lotsRequired", fmt.Sprintf("%t", lotsRequired))
	if sortBy != constants.PortfolioSortByNil {
		queryValues.Add("sortBy", sortBy.String())
	}
	if sortOrder != constants.SortOrderNil {
		queryValues.Add("sortOrder", sortOrder.String())
	}
	if marketSession != constants.MarketSessionNil {
		queryValues.Add("marketSession", marketSession.String())
	}
	if view != constants.PortfolioViewNil {
		queryValues.Add("view", view.String())
	}

	response, err := c.doRequest("GET", c.urls.ViewPortfolioUrl(accountIdKey), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) ListPositionLotsDetails(accountIdKey string, positionId int64) ([]byte, error) {
	if accountIdKey == "" {
		return nil, errors.New("accountIdKey not provided")
	}

	response, err := c.doRequest("GET", c.urls.ListPositionLotsDetailsUrl(accountIdKey, positionId), nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) ListAlerts(
	count int, category constants.AlertCategory, status constants.AlertStatus, sortOrder constants.SortOrder,
	search string,
) (
	[]byte, error,
) {
	queryValues := url.Values{}
	if count > constants.AlertsMaxCount {
		return nil, fmt.Errorf(
			"count of %d requested, which exceeds the maximum of %d", count, constants.AlertsMaxCount,
		)
	}
	if count >= 0 {
		queryValues.Add("count", fmt.Sprintf("%d", count))
	}
	if category != constants.AlertCategoryNil {
		queryValues.Add("category", category.String())
	}
	if status != constants.AlertStatusNil {
		queryValues.Add("status", status.String())
	}
	if sortOrder != constants.SortOrderNil {
		queryValues.Add("direction", sortOrder.String())
	}
	if search != "" {
		queryValues.Add("search", search)
	}

	response, err := c.doRequest("GET", c.urls.ListAlertsUrl(), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) ListAlertDetails(alertId string, htmlTags bool) ([]byte, error) {
	queryValues := url.Values{}
	queryValues.Add("htmlTags", fmt.Sprintf("%t", htmlTags))

	response, err := c.doRequest("GET", c.urls.ListAlertDetailsUrl(alertId), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) GetQuotes(
	symbols []string, detailFlag constants.QuoteDetailFlag, requireEarningsDate bool, skipMiniOptionsCheck bool,
) ([]byte, error) {
	if len(symbols) < 1 {
		return nil, errors.New("no symbols provided")
	}
	if len(symbols) > constants.GetQuotesMaxSymbols {
		return nil, fmt.Errorf(
			"%d symbols requested, which exceeds the maximum of %d symbols in a request", len(symbols),
			constants.GetQuotesMaxSymbols,
		)
	}
	symbolsList := strings.Join(symbols, ",")
	queryValues := url.Values{}
	if len(symbols) > constants.GetQuotesMaxSymbolsBeforeOverride {
		queryValues.Add("overrideSymbolCount", "true")
	}
	queryValues.Add("requireEarningsDate", fmt.Sprintf("%t", requireEarningsDate))
	queryValues.Add("skipMiniOptionsCheck", fmt.Sprintf("%t", skipMiniOptionsCheck))
	if detailFlag != constants.QuoteDetailNil {
		queryValues.Add("detailFlag", detailFlag.String())
	}

	response, err := c.doRequest("GET", c.urls.GetQuotesUrl(symbolsList), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) LookupProduct(search string) ([]byte, error) {
	if search == "" {
		return nil, errors.New("no search string provided")
	}
	response, err := c.doRequest("GET", c.urls.LookUpProductUrl(search), nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) GetOptionChains(
	symbol string, expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes int,
	includeWeekly, skipAdjusted bool, optionCategory constants.OptionCategory, chainType constants.OptionChainType,
	priceType constants.OptionPriceType,
) ([]byte, error) {
	if symbol == "" {
		return nil, errors.New("no symbol provided")
	}
	queryValues := url.Values{}
	queryValues.Add("symbol", symbol)
	if expiryYear > 0 {
		queryValues.Add("expiryYear", fmt.Sprintf("%d", expiryYear))
	}
	if expiryMonth > 0 {
		queryValues.Add("expiryMonth", fmt.Sprintf("%d", expiryMonth))
	}
	if expiryDay > 0 {
		queryValues.Add("expiryDay", fmt.Sprintf("%d", expiryDay))
	}
	if strikePriceNear >= 0 {
		queryValues.Add("strikePriceNear", fmt.Sprintf("%d", strikePriceNear))
	}
	if noOfStrikes >= 0 {
		queryValues.Add("noOfStrikes", fmt.Sprintf("%d", noOfStrikes))
	}
	queryValues.Add("includeWeekly", fmt.Sprintf("%t", includeWeekly))
	queryValues.Add("skipAdjusted", fmt.Sprintf("%t", skipAdjusted))
	if optionCategory != constants.OptionCategoryNil {
		queryValues.Add("optionCategory", optionCategory.String())
	}
	if chainType != constants.OptionChainTypeNil {
		queryValues.Add("chainType", chainType.String())
	}
	if priceType != constants.OptionPriceTypeNil {
		queryValues.Add("priceType", priceType.String())
	}

	response, err := c.doRequest("GET", c.urls.GetOptionChainsUrl(), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) GetOptionExpireDates(symbol string, expiryType constants.OptionExpiryType) ([]byte, error) {
	if symbol == "" {
		return nil, errors.New("no symbol provided")
	}
	queryValues := url.Values{}
	queryValues.Add("symbol", symbol)
	if expiryType != constants.OptionExpiryTypeNil {
		queryValues.Add("expiryType", expiryType.String())
	}

	response, err := c.doRequest("GET", c.urls.GetOptionExpireDatesUrl(), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) ListOrders(
	accountIdKey string, marker string, count int, status constants.OrderStatus, fromDate *time.Time, toDate *time.Time,
	symbols []string, securityType constants.OrderSecurityType, transactionType constants.OrderTransactionType,
	marketSession constants.MarketSession,
) ([]byte, error) {
	if accountIdKey == "" {
		return nil, errors.New("accountIdKey not provided")
	}
	queryValues := url.Values{}
	if marker != "" {
		queryValues.Add("marker", marker)
	}
	if count >= 0 {
		queryValues.Add("count", fmt.Sprintf("%d", count))
	}
	if status != constants.OrderStatusNil {
		queryValues.Add("status", status.String())
	}
	if fromDate != nil {
		queryValues.Add("fromDate", fromDate.Format(queryDateLayout))
	}
	if toDate != nil {
		queryValues.Add("toDate", toDate.Format(queryDateLayout))
	}
	if len(symbols) > 0 {
		if len(symbols) > constants.ListOrdersMaxSymbols {
			return nil, fmt.Errorf(
				"%d symbols provided, which exceeds the limit of %d", len(symbols), constants.ListOrdersMaxSymbols,
			)
		}
		queryValues.Add("symbol", strings.Join(symbols, ","))
	}
	if securityType != constants.OrderSecurityTypeNil {
		queryValues.Add("securityType", securityType.String())
	}
	if transactionType != constants.OrderTransactionTypeNil {
		queryValues.Add("transactionType", transactionType.String())
	}
	if marketSession != constants.MarketSessionNil {
		queryValues.Add("marketSession", marketSession.String())
	}

	response, err := c.doRequest("GET", c.urls.ListOrdersUrl(accountIdKey), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) doRequest(method string, baseUrl string, queryValues url.Values) ([]byte, error) {
	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, err
	}

	// Request that the server respond with JSON
	req.Header.Add("Accept", `application/json`)

	// Parse any query parameters from the base URL and merge them with the provided query parameters and encode
	urlQueryValues, err := url.ParseQuery(req.URL.RawQuery)
	if err != nil {
		return nil, err
	}
	for key, values := range urlQueryValues {
		for _, value := range values {
			queryValues.Add(key, value)
		}
	}
	req.URL.RawQuery = queryValues.Encode()

	// Perform the request
	c.Logger.Debug(method + " " + req.URL.String())
	httpResponse, err := c.httpClient.Do(req)
	if httpResponse != nil {
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				c.Logger.Error(err.Error())
			}
		}(httpResponse.Body)
	}
	if err != nil {
		return nil, err
	}

	// Check the response for an error and return response bytes if none
	if httpResponse.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", httpResponse.Status)
	}
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	c.Logger.Debug(string(responseBytes))
	return responseBytes, nil
}
