package etradelib

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

type ETradeClient interface {
	ListAccounts() (string, error)
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

func (c *eTradeClient) ListAccounts() (string, error) {
	response, err := c.httpClient.Get(c.urls.ListAccountsUrl())
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("request failed: %s", response.Status))
	}
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(responseBytes), nil
}
