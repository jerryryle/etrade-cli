package client

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/etradelibtest"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/responses"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestETradeClient_ListAccounts(t *testing.T) {
	type args struct {
		httpClientFakeXml string
	}
	tests := []struct {
		name      string
		args      args
		expectUrl string
		expectErr bool
		expect    *responses.AccountListResponse
	}{
		{
			name: "List Accounts With Results",
			args: args{
				httpClientFakeXml: listAccountsTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/accounts/list",
			expectErr: false,
			expect:    &listAccountsTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				httpClient := NewHttpClientFake(
					func(req *http.Request) *http.Response {
						assert.Equal(t, tt.expectUrl, req.URL.String())
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
						}
					},
				)

				client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
				response, err := client.ListAccounts()
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expect, response)
			},
		)
	}
}

func TestETradeClient_GetAccountBalances(t *testing.T) {
	type args struct {
		accountIdKey      string
		realTimeNAV       bool
		httpClientFakeXml string
	}
	tests := []struct {
		name      string
		args      args
		expectUrl string
		expectErr bool
		expect    *responses.BalanceResponse
	}{
		{
			name: "Get Account Balances With Results",
			args: args{
				accountIdKey:      "1234",
				realTimeNAV:       true,
				httpClientFakeXml: getAccountBalancesTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/accounts/1234/balance?instType=BROKERAGE&realTimeNAV=true",
			expectErr: false,
			expect:    &getAccountBalancesTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				httpClient := NewHttpClientFake(
					func(req *http.Request) *http.Response {
						assert.Equal(t, tt.expectUrl, req.URL.String())
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
						}
					},
				)

				client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
				response, err := client.GetAccountBalances(tt.args.accountIdKey, tt.args.realTimeNAV)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expect, response)
			},
		)
	}
}

func TestETradeClient_ListTransactions(t *testing.T) {
	type args struct {
		accountIdKey      string
		startDate         *time.Time
		endDate           *time.Time
		sortOrder         TransactionSortOrder
		marker            string
		count             int
		httpClientFakeXml string
	}
	tests := []struct {
		name      string
		args      args
		expectUrl string
		expectErr bool
		expect    *responses.TransactionListResponse
	}{
		{
			name: "Get Transactions With Results",
			args: args{
				accountIdKey:      "1234",
				startDate:         etradelibtest.CreateTime(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
				endDate:           etradelibtest.CreateTime(2023, time.January, 2, 0, 0, 0, 0, time.UTC),
				sortOrder:         TransactionSortOrderAsc,
				marker:            "FOO",
				count:             6,
				httpClientFakeXml: listTransactionsTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/accounts/1234/transactions?count=6&endDate=01022023&marker=FOO&sortOrder=ASC&startDate=01012023",
			expectErr: false,
			expect:    &listTransactionsTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				httpClient := NewHttpClientFake(
					func(req *http.Request) *http.Response {
						assert.Equal(t, tt.expectUrl, req.URL.String())
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
						}
					},
				)

				client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
				response, err := client.ListTransactions(
					tt.args.accountIdKey,
					tt.args.startDate, tt.args.endDate, tt.args.sortOrder, tt.args.marker, tt.args.count,
				)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expect, response)
			},
		)
	}
}

func TestETradeClient_ListTransactionDetails(t *testing.T) {
	type args struct {
		accountIdKey      string
		transactionId     string
		httpClientFakeXml string
	}
	tests := []struct {
		name      string
		args      args
		expectUrl string
		expectErr bool
		expect    *responses.TransactionDetailsResponse
	}{
		{
			name: "Get Transaction Details With Results",
			args: args{
				accountIdKey:      "1234",
				transactionId:     "5678",
				httpClientFakeXml: listTransactionDetailsTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/accounts/1234/transactions/5678",
			expectErr: false,
			expect:    &listTransactionDetailsTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				httpClient := NewHttpClientFake(
					func(req *http.Request) *http.Response {
						assert.Equal(t, tt.expectUrl, req.URL.String())
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
						}
					},
				)

				client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
				response, err := client.ListTransactionDetails(tt.args.accountIdKey, tt.args.transactionId)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expect, response)
			},
		)
	}
}

func TestETradeClient_ListAlerts(t *testing.T) {
	type args struct {
		httpClientFakeXml string
	}
	tests := []struct {
		name      string
		args      args
		expectUrl string
		expectErr bool
		expect    *responses.AlertsResponse
	}{
		{
			name: "List Alerts With Results",
			args: args{
				httpClientFakeXml: listAlertsTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/user/alerts",
			expectErr: false,
			expect:    &listAlertsTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				httpClient := NewHttpClientFake(
					func(req *http.Request) *http.Response {
						assert.Equal(t, tt.expectUrl, req.URL.String())
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
						}
					},
				)

				client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
				response, err := client.ListAlerts()
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expect, response)
			},
		)
	}
}

func TestETradeClient_GetQuotes(t *testing.T) {
	type args struct {
		symbols              []string
		detailFlag           QuoteDetailFlag
		requireEarningsDate  bool
		skipMiniOptionsCheck bool
		httpClientFakeXml    string
	}
	tests := []struct {
		name      string
		args      args
		expectUrl string
		expectErr bool
		expect    *responses.QuoteResponse
	}{
		{
			name: "Valid QuoteDetailAll XML",
			args: args{
				symbols:              []string{"GOOG"},
				detailFlag:           QuoteDetailAll,
				requireEarningsDate:  true,
				skipMiniOptionsCheck: false,
				httpClientFakeXml:    quoteDetailAllTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/market/quote/GOOG?detailFlag=ALL&requireEarningsDate=true&skipMiniOptionsCheck=false",
			expectErr: false,
			expect:    &quoteDetailAllTestResponse,
		},
		{
			name: "Valid QuoteDetailFundamental XML",
			args: args{
				symbols:              []string{"GOOG"},
				detailFlag:           QuoteDetailFundamental,
				requireEarningsDate:  true,
				skipMiniOptionsCheck: false,
				httpClientFakeXml:    quoteDetailFundamentalTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/market/quote/GOOG?detailFlag=FUNDAMENTAL&requireEarningsDate=true&skipMiniOptionsCheck=false",
			expectErr: false,
			expect:    &quoteDetailFundamentalTestResponse,
		},
		{
			name: "Valid QuoteDetailIntraday XML",
			args: args{
				symbols:              []string{"GOOG"},
				detailFlag:           QuoteDetailIntraday,
				requireEarningsDate:  true,
				skipMiniOptionsCheck: false,
				httpClientFakeXml:    quoteDetailIntradayTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/market/quote/GOOG?detailFlag=INTRADAY&requireEarningsDate=true&skipMiniOptionsCheck=false",
			expectErr: false,
			expect:    &quoteDetailIntradayTestResponse,
		},
		{
			name: "Valid QuoteDetailOptions XML",
			args: args{
				symbols:              []string{"GOOG"},
				detailFlag:           QuoteDetailOptions,
				requireEarningsDate:  true,
				skipMiniOptionsCheck: false,
				httpClientFakeXml:    quoteDetailOptionsTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/market/quote/GOOG?detailFlag=OPTIONS&requireEarningsDate=true&skipMiniOptionsCheck=false",
			expectErr: false,
			expect:    &quoteDetailOptionsTestResponse,
		},
		{
			name: "Valid QuoteDetailWeek52 XML",
			args: args{
				symbols:              []string{"GOOG"},
				detailFlag:           QuoteDetailWeek52,
				requireEarningsDate:  true,
				skipMiniOptionsCheck: false,
				httpClientFakeXml:    quoteDetailWeek52TestXml,
			},
			expectUrl: "https://api.etrade.com/v1/market/quote/GOOG?detailFlag=WEEK_52&requireEarningsDate=true&skipMiniOptionsCheck=false",
			expectErr: false,
			expect:    &quoteDetailWeek52TestResponse,
		},
		{
			name: "Valid QuoteDetailMutualFund XML",
			args: args{
				symbols:              []string{"VFIAX"},
				detailFlag:           QuoteDetailMutualFund,
				requireEarningsDate:  true,
				skipMiniOptionsCheck: false,
				httpClientFakeXml:    quoteDetailMutualFundTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/market/quote/VFIAX?detailFlag=MF_DETAIL&requireEarningsDate=true&skipMiniOptionsCheck=false",
			expectErr: false,
			expect:    &quoteDetailMutualFundTestResponse,
		},
		{
			name: "Override Symbols When More Than 25",
			args: args{
				symbols: []string{
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18",
					"19", "20", "21", "22", "23", "24", "25", "26",
				},
				detailFlag:           QuoteDetailAll,
				requireEarningsDate:  true,
				skipMiniOptionsCheck: false,
				httpClientFakeXml:    quoteDetailAllTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/market/quote/1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26?detailFlag=ALL&overrideSymbolCount=true&requireEarningsDate=true&skipMiniOptionsCheck=false",
			expectErr: false,
			expect:    &quoteDetailAllTestResponse,
		},
		{
			name: "Too Many Symbols",
			args: args{
				symbols: []string{
					"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18",
					"19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34",
					"35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50",
					"51",
				},
				detailFlag:           QuoteDetailAll,
				requireEarningsDate:  true,
				skipMiniOptionsCheck: false,
				httpClientFakeXml:    `<?xml version="1.0" encoding="UTF-8" standalone="yes"?><QuoteResponse></QuoteResponse>`,
			},
			expectUrl: "",
			expectErr: true,
			expect:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				httpClient := NewHttpClientFake(
					func(req *http.Request) *http.Response {
						assert.Equal(t, tt.expectUrl, req.URL.String())
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
						}
					},
				)

				client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
				response, err := client.GetQuotes(
					tt.args.symbols,
					tt.args.detailFlag, tt.args.requireEarningsDate, tt.args.skipMiniOptionsCheck,
				)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expect, response)
			},
		)
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
		expectUrl string
		expectErr bool
		expect    *responses.LookupResponse
	}{
		{
			name: "Valid Search With Results",
			args: args{
				search:            "A",
				httpClientFakeXml: lookupProductResultsTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/market/lookup/A",
			expectErr: false,
			expect:    &lookupProductResultsTestResponse,
		},
		{
			name: "Valid Search With No Results",
			args: args{
				search:            "A",
				httpClientFakeXml: lookupProductNoResultsTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/market/lookup/A",
			expectErr: false,
			expect:    &lookupProductNoResultsTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				httpClient := NewHttpClientFake(
					func(req *http.Request) *http.Response {
						assert.Equal(t, tt.expectUrl, req.URL.String())
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
						}
					},
				)

				client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
				response, err := client.LookupProduct(tt.args.search)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expect, response)
			},
		)
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
		expectUrl string
		expectErr bool
		expect    *responses.OptionChainResponse
	}{
		{
			name: "Get Option Chains With Results",
			args: args{
				symbol:            "GOOG",
				expiryYear:        1,
				expiryMonth:       2,
				expiryDay:         3,
				strikePriceNear:   4,
				noOfStrikes:       5,
				includeWeekly:     true,
				skipAdjusted:      true,
				optionCategory:    OptionCategoryAll,
				chainType:         ChainTypeCall,
				priceType:         PriceTypeAll,
				httpClientFakeXml: getOptionChainsTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/market/optionchains?chainType=CALL&expiryDay=3&expiryMonth=2&expiryYear=1&includeWeekly=true&noOfStrikes=5&optionCategory=ALL&priceType=ALL&skipAdjusted=true&strikePriceNear=4&symbol=GOOG",
			expectErr: false,
			expect:    &getOptionChainsTestResponse,
		},
		{
			name: "Get Option Chains And Omit Some Parameters",
			args: args{
				symbol:            "GOOG",
				expiryYear:        -1,
				expiryMonth:       -1,
				expiryDay:         -1,
				strikePriceNear:   -1,
				noOfStrikes:       -1,
				includeWeekly:     true,
				skipAdjusted:      true,
				optionCategory:    OptionCategoryAll,
				chainType:         ChainTypeCall,
				priceType:         PriceTypeAll,
				httpClientFakeXml: getOptionChainsTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/market/optionchains?chainType=CALL&includeWeekly=true&optionCategory=ALL&priceType=ALL&skipAdjusted=true&symbol=GOOG",
			expectErr: false,
			expect:    &getOptionChainsTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				httpClient := NewHttpClientFake(
					func(req *http.Request) *http.Response {
						assert.Equal(t, tt.expectUrl, req.URL.String())
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
						}
					},
				)

				client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
				response, err := client.GetOptionChains(
					tt.args.symbol,
					tt.args.expiryYear, tt.args.expiryMonth, tt.args.expiryDay,
					tt.args.strikePriceNear, tt.args.noOfStrikes,
					tt.args.includeWeekly, tt.args.skipAdjusted,
					tt.args.optionCategory, tt.args.chainType, tt.args.priceType,
				)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expect, response)
			},
		)
	}
}

func TestETradeClient_GetOptionExpireDates(t *testing.T) {
	type args struct {
		symbol            string
		expiryType        ExpiryType
		httpClientFakeXml string
	}
	tests := []struct {
		name      string
		args      args
		expectUrl string
		expectErr bool
		expect    *responses.OptionExpireDateResponse
	}{
		{
			name: "Get Option Expire Date With Results",
			args: args{
				symbol:            "GOOG",
				expiryType:        ExpiryTypeAll,
				httpClientFakeXml: getOptionExpireDateTestXml,
			},
			expectUrl: "https://api.etrade.com/v1/market/optionexpiredate?expiryType=ALL&symbol=GOOG",
			expectErr: false,
			expect:    &getOptionExpireDateTestResponse,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				httpClient := NewHttpClientFake(
					func(req *http.Request) *http.Response {
						assert.Equal(t, tt.expectUrl, req.URL.String())
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(tt.args.httpClientFakeXml)),
						}
					},
				)

				client := CreateETradeClient(GetEndpointUrls(true), httpClient, etradelibtest.CreateNullLogger())
				response, err := client.GetOptionExpireDates(tt.args.symbol, tt.args.expiryType)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expect, response)
			},
		)
	}
}
