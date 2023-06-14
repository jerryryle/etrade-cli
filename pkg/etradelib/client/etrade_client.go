package client

import (
	"errors"
	"fmt"
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
		accountIdKey string, startDate *time.Time, endDate *time.Time, sortOrder SortOrder, marker string, count int,
	) ([]byte, error)

	ListTransactionDetails(accountIdKey string, transactionId string) ([]byte, error)

	ViewPortfolio(
		accountIdKey string, count int, sortBy PortfolioSortBy, sortOrder SortOrder, pageNumber int,
		marketSession PortfolioMarketSession, totalsRequired bool, lotsRequired bool, view PortfolioView,
	) ([]byte, error)

	ListAlerts() ([]byte, error)

	GetQuotes(
		symbols []string, detailFlag QuoteDetailFlag, requireEarningsDate bool, skipMiniOptionsCheck bool,
	) ([]byte, error)

	LookupProduct(search string) ([]byte, error)

	GetOptionChains(
		symbol string, expiryYear int, expiryMonth int, expiryDay int, strikePriceNear int, noOfStrikes int,
		includeWeekly bool, skipAdjusted bool, optionCategory OptionCategory, chainType ChainType, priceType PriceType,
	) ([]byte, error)

	GetOptionExpireDates(symbol string, expiryType ExpiryType) ([]byte, error)
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

func (c *eTradeClient) ListAccounts() ([]byte, error) {
	response, err := c.doRequest("GET", c.urls.ListAccountsUrl(), nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) GetAccountBalances(accountIdKey string, realTimeNAV bool) ([]byte, error) {
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
	accountIdKey string, startDate *time.Time, endDate *time.Time, sortOrder SortOrder, marker string, count int,
) ([]byte, error) {
	dateLayout := "01022006"
	queryValues := url.Values{}
	if startDate != nil {
		queryValues.Add("startDate", startDate.Format(dateLayout))
	}
	if endDate != nil {
		queryValues.Add("endDate", endDate.Format(dateLayout))
	}
	queryValues.Add("sortOrder", sortOrder.String())
	if marker != "" {
		queryValues.Add("marker", marker)
	}
	if count > 0 {
		queryValues.Add("count", fmt.Sprintf("%d", count))
	}

	response, err := c.doRequest("GET", c.urls.ListTransactionsUrl(accountIdKey), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) ListTransactionDetails(accountIdKey string, transactionId string) ([]byte, error) {
	response, err := c.doRequest("GET", c.urls.ListTransactionDetailsUrl(accountIdKey, transactionId), nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) ViewPortfolio(
	accountIdKey string, count int, sortBy PortfolioSortBy, sortOrder SortOrder, pageNumber int,
	marketSession PortfolioMarketSession, totalsRequired bool, lotsRequired bool, view PortfolioView,
) ([]byte, error) {
	queryValues := url.Values{}
	if count > 0 {
		queryValues.Add("count", fmt.Sprintf("%d", count))
	}
	if pageNumber > 0 {
		queryValues.Add("pageNumber", fmt.Sprintf("%d", pageNumber))
	}
	queryValues.Add("totalsRequired", fmt.Sprintf("%t", totalsRequired))
	queryValues.Add("lotsRequired", fmt.Sprintf("%t", lotsRequired))
	queryValues.Add("sortBy", sortBy.String())
	queryValues.Add("sortOrder", sortOrder.String())
	queryValues.Add("marketSession", marketSession.String())
	queryValues.Add("view", view.String())

	response, err := c.doRequest("GET", c.urls.ViewPortfolioUrl(accountIdKey), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) ListAlerts() ([]byte, error) {
	response, err := c.doRequest("GET", c.urls.ListAlertsUrl(), nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) GetQuotes(
	symbols []string, detailFlag QuoteDetailFlag, requireEarningsDate bool, skipMiniOptionsCheck bool,
) ([]byte, error) {
	if len(symbols) > GetQuotesMaxSymbols {
		return nil, errors.New(
			fmt.Sprintf(
				"%d symbols requested, which exceeds the maximum of %d symbols in a request", len(symbols),
				GetQuotesMaxSymbols,
			),
		)
	}
	symbolsList := strings.Join(symbols, ",")
	queryValues := url.Values{}
	if len(symbols) > GetQuotesMaxSymbolsBeforeOverride {
		queryValues.Add("overrideSymbolCount", "true")
	}
	queryValues.Add("requireEarningsDate", fmt.Sprintf("%t", requireEarningsDate))
	queryValues.Add("skipMiniOptionsCheck", fmt.Sprintf("%t", skipMiniOptionsCheck))
	queryValues.Add("detailFlag", detailFlag.String())

	response, err := c.doRequest("GET", c.urls.GetQuotesUrl(symbolsList), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) LookupProduct(search string) ([]byte, error) {
	response, err := c.doRequest("GET", c.urls.LookUpProductUrl(search), nil)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) GetOptionChains(
	symbol string, expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes int,
	includeWeekly, skipAdjusted bool, optionCategory OptionCategory, chainType ChainType, priceType PriceType,
) ([]byte, error) {
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
	queryValues.Add("optionCategory", optionCategory.String())
	queryValues.Add("chainType", chainType.String())
	queryValues.Add("priceType", priceType.String())

	response, err := c.doRequest("GET", c.urls.GetOptionChainsUrl(), queryValues)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *eTradeClient) GetOptionExpireDates(symbol string, expiryType ExpiryType) ([]byte, error) {
	queryValues := url.Values{}
	queryValues.Add("symbol", symbol)
	queryValues.Add("expiryType", expiryType.String())

	response, err := c.doRequest("GET", c.urls.GetOptionExpireDatesUrl(), queryValues)
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
		return nil, errors.New(fmt.Sprintf("request failed: %s", httpResponse.Status))
	}
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	c.Logger.Debug(string(responseBytes))
	return responseBytes, nil
}
