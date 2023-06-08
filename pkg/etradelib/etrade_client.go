package etradelib

import "net/http"

type ETradeClient interface {
	Urls() EndpointUrls
	HttpClient() *http.Client
}

type eTradeClient struct {
	urls       EndpointUrls
	httpClient *http.Client
}

func (c *eTradeClient) Urls() EndpointUrls {
	return c.urls
}

func (c *eTradeClient) HttpClient() *http.Client {
	return c.httpClient
}
