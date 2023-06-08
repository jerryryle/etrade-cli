package etradelib

import "net/http"

type eTradeClient interface {
	Urls() EndpointUrls
	HttpClient() *http.Client
}

type eTradeClientStruct struct {
	urls       EndpointUrls
	httpClient *http.Client
}

func (c *eTradeClientStruct) Urls() EndpointUrls {
	return c.urls
}

func (c *eTradeClientStruct) HttpClient() *http.Client {
	return c.httpClient
}
