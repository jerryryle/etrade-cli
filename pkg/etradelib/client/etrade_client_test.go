package client

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/etradelibtest"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestETradeClient_ListAccounts(t *testing.T) {
	type args struct {
		httpClientFakeXml string
	}
	tests := []struct {
		name      string
		args      args
		expectErr bool
		expect    *responses.AccountListResponse
	}{
		{
			name: "List Accounts With Results",
			args: args{
				httpClientFakeXml: listAccountsTestXml,
			},
			expectErr: false,
			expect:    &listAccountsTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
				}
			})

			client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
			response, err := client.ListAccounts()
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expect, response)
		})
	}
}

func TestETradeClient_ListAlerts(t *testing.T) {
	type args struct {
		httpClientFakeXml string
	}
	tests := []struct {
		name      string
		args      args
		expectErr bool
		expect    *responses.AlertsResponse
	}{
		{
			name: "List Accounts With Results",
			args: args{
				httpClientFakeXml: listAlertsTestXml,
			},
			expectErr: false,
			expect:    &listAlertsTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
				}
			})

			client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
			response, err := client.ListAlerts()
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expect, response)
		})
	}
}

func TestETradeClient_GetQuotes(t *testing.T) {
	type args struct {
		symbols           []string
		detailFlag        QuoteDetailFlag
		httpClientFakeXml string
	}
	tests := []struct {
		name      string
		args      args
		expectErr bool
		expect    *responses.QuoteResponse
	}{
		{
			name: "Valid QuoteDetailAll XML",
			args: args{
				symbols:           []string{"GOOG"},
				detailFlag:        QuoteDetailAll,
				httpClientFakeXml: quoteDetailAllTestXml,
			},
			expectErr: false,
			expect:    &quoteDetailAllTestResponse,
		},
		{
			name: "Valid QuoteDetailFundamental XML",
			args: args{
				symbols:           []string{"GOOG"},
				detailFlag:        QuoteDetailFundamental,
				httpClientFakeXml: quoteDetailFundamentalTestXml,
			},
			expectErr: false,
			expect:    &quoteDetailFundamentalTestResponse,
		},
		{
			name: "Valid QuoteDetailIntraday XML",
			args: args{
				symbols:           []string{"GOOG"},
				detailFlag:        QuoteDetailIntraday,
				httpClientFakeXml: quoteDetailIntradayTestXml,
			},
			expectErr: false,
			expect:    &quoteDetailIntradayTestResponse,
		},
		{
			name: "Valid QuoteDetailOptions XML",
			args: args{
				symbols:           []string{"GOOG"},
				detailFlag:        QuoteDetailOptions,
				httpClientFakeXml: quoteDetailOptionsTestXml,
			},
			expectErr: false,
			expect:    &quoteDetailOptionsTestResponse,
		},
		{
			name: "Valid QuoteDetailWeek52 XML",
			args: args{
				symbols:           []string{"GOOG"},
				detailFlag:        QuoteDetailWeek52,
				httpClientFakeXml: quoteDetailWeek52TestXml,
			},
			expectErr: false,
			expect:    &quoteDetailWeek52TestResponse,
		},
		{
			name: "Valid QuoteDetailMutualFund XML",
			args: args{
				symbols:           []string{"VFIAX"},
				detailFlag:        QuoteDetailMutualFund,
				httpClientFakeXml: quoteDetailMutualFundTestXml,
			},
			expectErr: false,
			expect:    &quoteDetailMutualFundTestResponse,
		},
		{
			name: "Too Many Symbols",
			args: args{
				symbols:           []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51"},
				detailFlag:        QuoteDetailAll,
				httpClientFakeXml: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><QuoteResponse></QuoteResponse>`,
			},
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
				}
			})

			client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
			response, err := client.GetQuotes(tt.args.symbols, tt.args.detailFlag)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expect, response)
		})
	}
}

func TestETradeClient_LookupProduct(t *testing.T) {
	type args struct {
		search            string
		httpClientFakeXml string
	}
	tests := []struct {
		name      string
		args      args
		expectErr bool
		expect    *responses.LookupResponse
	}{
		{
			name: "Valid Search With Results",
			args: args{
				search:            "A",
				httpClientFakeXml: lookupProductResultsTestXml,
			},
			expectErr: false,
			expect:    &lookupProductResultsTestResponse,
		},
		{
			name: "Valid Search With No Results",
			args: args{
				search:            "A",
				httpClientFakeXml: lookupProductNoResultsTestXml,
			},
			expectErr: false,
			expect:    &lookupProductNoResultsTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
				}
			})

			client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
			response, err := client.LookupProduct(tt.args.search)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expect, response)
		})
	}
}

func TestETradeClient_GetOptionChains(t *testing.T) {
	type args struct {
		symbol                             string
		expiryYear, expiryMonth, expiryDay int
		strikePriceNear, noOfStrikes       int
		includeWeekly, skipAdjusted        bool
		optionCategory                     OptionCategory
		chainType                          ChainType
		priceType                          PriceType
		httpClientFakeXml                  string
	}
	tests := []struct {
		name      string
		args      args
		expectErr bool
		expect    *responses.OptionChainResponse
	}{
		{
			name: "Get Option Chains With Results",
			args: args{
				httpClientFakeXml: getOptionChainsTestXml,
			},
			expectErr: false,
			expect:    &getOptionChainsTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpClient := NewHttpClientFake(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
				}
			})

			client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
			response, err := client.GetOptionChains(tt.args.symbol,
				tt.args.expiryYear, tt.args.expiryMonth, tt.args.expiryDay,
				tt.args.strikePriceNear, tt.args.noOfStrikes,
				tt.args.includeWeekly, tt.args.skipAdjusted,
				tt.args.optionCategory, tt.args.chainType, tt.args.priceType)
			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.expect, response)
		})
	}
}
