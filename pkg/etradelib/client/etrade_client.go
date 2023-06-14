package client

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ETradeClient interface {
	ListAccounts() (*responses.AccountListResponse, error)
	ListAlerts() (*responses.AlertsResponse, error)
	GetQuotes(symbols []string, detailFlag QuoteDetailFlag) (*responses.QuoteResponse, error)
	LookupProduct(search string) (*responses.LookupResponse, error)
	GetOptionChains(symbol string,
		expiryYear int, expiryMonth int, expiryDay int,
		strikePriceNear int, noOfStrikes int, includeWeekly bool, skipAdjusted bool,
		optionCategory OptionCategory, chainType ChainType, priceType PriceType) (*responses.OptionChainResponse, error)
	GetOptionExpireDates(symbol string, expiryType ExpiryType) (*responses.OptionExpireDateResponse, error)
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

func (c *eTradeClient) ListAccounts() (*responses.AccountListResponse, error) {
	response := responses.AccountListResponse{}
	err := c.doRequest("GET", c.urls.ListAccountsUrl(), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *eTradeClient) ListAlerts() (*responses.AlertsResponse, error) {
	response := responses.AlertsResponse{}
	err := c.doRequest("GET", c.urls.ListAlertsUrl(), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *eTradeClient) GetQuotes(symbols []string, detailFlag QuoteDetailFlag) (*responses.QuoteResponse, error) {
	if len(symbols) > GetQuotesMaxSymbols {
		return nil, errors.New(fmt.Sprintf("%d symbols requested, which exceeds the maximum of %d symbols in a request", len(symbols), GetQuotesMaxSymbols))
	}
	symbolsList := strings.Join(symbols, ",")
	queryValues := url.Values{}
	queryValues.Add("requireEarningsDate", "true")
	queryValues.Add("overrideSymbolCount", "true")
	queryValues.Add("skipMiniOptionsCheck", "false")
	queryValues.Add("detailFlag", detailFlag.String())

	response := responses.QuoteResponse{}
	err := c.doRequest("GET", c.urls.GetQuotesUrl(symbolsList), queryValues, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *eTradeClient) LookupProduct(search string) (*responses.LookupResponse, error) {
	response := responses.LookupResponse{}
	err := c.doRequest("GET", c.urls.LookUpProductUrl(search), nil, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *eTradeClient) GetOptionChains(symbol string,
	expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes int,
	includeWeekly, skipAdjusted bool,
	optionCategory OptionCategory, chainType ChainType, priceType PriceType) (*responses.OptionChainResponse, error) {
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

	response := responses.OptionChainResponse{}
	err := c.doRequest("GET", c.urls.GetOptionChainsUrl(), queryValues, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *eTradeClient) GetOptionExpireDates(symbol string, expiryType ExpiryType) (*responses.OptionExpireDateResponse, error) {
	queryValues := url.Values{}
	queryValues.Add("symbol", symbol)
	queryValues.Add("expiryType", expiryType.String())

	response := responses.OptionExpireDateResponse{}
	err := c.doRequest("GET", c.urls.GetOptionExpireDatesUrl(), queryValues, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *eTradeClient) doRequest(method string, baseUrl string, queryValues url.Values, response interface{}) error {
	// Perform the request
	responseBytes, err := c.doRequestRaw(method, baseUrl, queryValues)
	if err != nil {
		return err
	}

	// Unmarshal the response into the provided structure
	err = xml.Unmarshal(responseBytes, &response)
	if err != nil {
		return err
	}
	return nil
}

func (c *eTradeClient) doRequestRaw(method string, baseUrl string, queryValues url.Values) ([]byte, error) {
	req, err := http.NewRequest(method, baseUrl, nil)
	if err != nil {
		return nil, err
	}

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
