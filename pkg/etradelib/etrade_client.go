package etradelib

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
	"io"
	"net/http"
)

type ETradeClient interface {
	ListAccounts() (*responses.AccountListResponse, error)
	ListAlerts() (*responses.AlertsResponse, error)
}

type eTradeClient struct {
	urls       EndpointUrls
	httpClient *http.Client
}

func CreateETradeClient(urls EndpointUrls, httpClient *http.Client) ETradeClient {
	return &eTradeClient{
		urls:       urls,
		httpClient: httpClient,
	}
}

func (c *eTradeClient) ListAccounts() (*responses.AccountListResponse, error) {
	httpResponse, err := c.httpClient.Get(c.urls.ListAccountsUrl())
	if httpResponse != nil {
		defer httpResponse.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request failed: %s", httpResponse.Status))
	}
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	response := responses.AccountListResponse{}
	err = xml.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *eTradeClient) ListAlerts() (*responses.AlertsResponse, error) {
	httpResponse, err := c.httpClient.Get(c.urls.ListAlertsUrl())
	if httpResponse != nil {
		defer httpResponse.Body.Close()
	}
	if err != nil {
		return nil, err
	}
	if httpResponse.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("request failed: %s", httpResponse.Status))
	}
	responseBytes, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	response := responses.AlertsResponse{}
	err = xml.Unmarshal(responseBytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
