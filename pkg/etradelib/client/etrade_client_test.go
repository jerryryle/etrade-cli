package client

import (
	"errors"
	"github.com/dghubble/oauth1"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/etradelibtest"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func createMockClient(
	httpClient HttpClient, config OAuthConfig, production bool,
	consumerKey, consumerSecret, requestToken, requestSecret, accessToken, accessSecret string,
) ETradeClient {

	return &eTradeClient{
		urls:           GetEndpointUrls(production),
		httpClient:     httpClient,
		logger:         etradelibtest.CreateNullLogger(),
		config:         config,
		consumerKey:    consumerKey,
		consumerSecret: consumerSecret,
		requestToken:   requestToken,
		requestSecret:  requestSecret,
		accessToken:    accessToken,
		accessSecret:   accessSecret,
	}
}

func TestETradeClient_Authenticate(t *testing.T) {
	type testFn func(testClient ETradeClient, clientMock *httpClientMock, configMock *oAuthConfigMock) ([]byte, error)

	tests := []struct {
		name            string
		testConsumerKey string
		testFn          testFn
		expectResponse  []byte
		expectErr       bool
	}{
		{
			name:            "Authenticate Returns Empty Auth URL On Successful Renewal",
			testConsumerKey: "TestConsumerKey",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock, configMock *oAuthConfigMock) (
				[]byte, error,
			) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/oauth/renew_access_token",
				).Return(http.StatusOK, "", nil)
				return testClient.Authenticate()
			},
			expectResponse: []byte(`{"status":"success"}` + "\n"),
			expectErr:      false,
		},
		{
			name:            "Authenticate Returns Auth URL On Unsuccessful Renewal",
			testConsumerKey: "TestConsumerKey",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock, configMock *oAuthConfigMock) (
				[]byte, error,
			) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/oauth/renew_access_token",
				).Return(http.StatusUnauthorized, "", nil)
				configMock.On("RequestToken").Return("TestToken", "TestKey", nil)
				return testClient.Authenticate()
			},
			expectResponse: []byte(`{"authorizationUrl":"https://us.etrade.com/e/t/etws/authorize?key=TestConsumerKey&token=TestToken","status":"authorize"}` + "\n"),
			expectErr:      false,
		},
		{
			name:            "Authenticate Fails On HTTP Error",
			testConsumerKey: "TestConsumerKey",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock, configMock *oAuthConfigMock) (
				[]byte, error,
			) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/oauth/renew_access_token",
				).Return(http.StatusBadRequest, "", nil)
				return testClient.Authenticate()
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name:            "Authenticate Returns Error On RequestToken Failure",
			testConsumerKey: "TestConsumerKey",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock, configMock *oAuthConfigMock) (
				[]byte, error,
			) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/oauth/renew_access_token",
				).Return(http.StatusUnauthorized, "", nil)
				configMock.On("RequestToken").Return("", "", errors.New("test error"))
				return testClient.Authenticate()
			},
			expectResponse: nil,
			expectErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				configMock := new(oAuthConfigMock)
				clientMock := new(httpClientMock)
				testClient := createMockClient(clientMock, configMock, true, tt.testConsumerKey, "", "", "", "", "")
				// Call the Method Under Test
				actualResponse, err := tt.testFn(testClient, clientMock, configMock)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectResponse, actualResponse)
				clientMock.AssertExpectations(t)
				configMock.AssertExpectations(t)
			},
		)
	}
}

func TestETradeClient_Verify(t *testing.T) {
	type testFn func(testClient ETradeClient, clientMock *httpClientMock, configMock *oAuthConfigMock) ([]byte, error)

	tests := []struct {
		name              string
		testRequestToken  string
		testRequestSecret string
		testFn            testFn
		expectResponse    []byte
		expectErr         bool
	}{
		{
			name:              "Verify Succeeds",
			testRequestToken:  "TestRequestToken",
			testRequestSecret: "TestRequestSecret",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock, configMock *oAuthConfigMock) (
				[]byte, error,
			) {
				configMock.On(
					"AccessToken", "TestRequestToken", "TestRequestSecret", "TestVerifyKey",
				).Return("TestAccessToken", "TestAccessSecret", nil)
				configMock.On(
					"Client", oauth1.NoContext, oauth1.NewToken("TestAccessToken", "TestAccessSecret"),
				).Return(&http.Client{}, nil)
				return testClient.Verify("TestVerifyKey")
			},
			expectResponse: []byte(`{"status":"success"}` + "\n"),
			expectErr:      false,
		},
		{
			name:              "Verify Fails On Access Token Error",
			testRequestToken:  "TestRequestToken",
			testRequestSecret: "TestRequestSecret",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock, configMock *oAuthConfigMock) (
				[]byte, error,
			) {
				configMock.On(
					"AccessToken", "TestRequestToken", "TestRequestSecret", "TestVerifyKey",
				).Return("", "", errors.New("test error"))
				return testClient.Verify("TestVerifyKey")
			},
			expectResponse: nil,
			expectErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				configMock := new(oAuthConfigMock)
				clientMock := new(httpClientMock)

				expectedConsumerKey := "TestConsumerKey"
				expectedConsumerSecret := "TestConsumerSecret"
				testClient := createMockClient(
					clientMock, configMock, true, expectedConsumerKey, expectedConsumerSecret, tt.testRequestToken,
					tt.testRequestSecret, "", "",
				)
				// Call the Method Under Test
				actualResponse, err := tt.testFn(testClient, clientMock, configMock)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectResponse, actualResponse)
				clientMock.AssertExpectations(t)
				configMock.AssertExpectations(t)
			},
		)
	}
}

func TestETradeClient_GetKeys(t *testing.T) {
	expectedConsumerKey := "TestConsumerKey"
	expectedConsumerSecret := "TestConsumerSecret"
	expectedAccessToken := "TestAccessToken"
	expectedAccessSecret := "TestAccessSecret"

	testClient := createMockClient(
		nil, nil, true, expectedConsumerKey, expectedConsumerSecret, "",
		"", expectedAccessToken, expectedAccessSecret,
	)
	actualConsumerKey, actualConsumerSecret, actualAccessToken, actualAccessSecret := testClient.GetKeys()
	assert.Equal(t, expectedConsumerKey, actualConsumerKey)
	assert.Equal(t, expectedConsumerSecret, actualConsumerSecret)
	assert.Equal(t, expectedAccessToken, actualAccessToken)
	assert.Equal(t, expectedAccessSecret, actualAccessSecret)
}

func TestETradeClient(t *testing.T) {
	type testFn func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error)

	// testResponseData is bogus JSON that's only used to ensure the client
	// returns the exact response from the server
	const testResponseData = `{"testResponse": true}`

	tests := []struct {
		name           string
		testFn         testFn
		expectResponse []byte
		expectErr      bool
	}{
		{
			name: "List Accounts",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On("Do", "GET", "https://api.etrade.com/v1/accounts/list").Return(
					http.StatusOK, testResponseData, nil,
				)
				return testClient.ListAccounts()
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "List Accounts Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On("Do", "GET", "https://api.etrade.com/v1/accounts/list").Return(
					0, "", errors.New("test error"),
				)
				return testClient.ListAccounts()
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "Get Account Balances",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/accounts/1234/balance?instType=BROKERAGE&realTimeNAV=true",
				).Return(http.StatusOK, testResponseData, nil)
				return testClient.GetAccountBalances("1234", true)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "Get Account Balances Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/accounts/1234/balance?instType=BROKERAGE&realTimeNAV=true",
				).Return(0, "", errors.New("test error"))
				return testClient.GetAccountBalances("1234", true)
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "Get Account Balances Fails Without Account ID Key",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.GetAccountBalances("", true)
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "List Transactions With All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/accounts/1234/transactions?count=6&endDate=01022023&marker=5&sortOrder=ASC&startDate=01012023",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.ListTransactions(
					"1234",
					etradelibtest.CreateTime(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
					etradelibtest.CreateTime(2023, time.January, 2, 0, 0, 0, 0, time.UTC),
					constants.SortOrderAsc, "5", 6,
				)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "List Transactions Can Omit All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/accounts/1234/transactions",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.ListTransactions("1234", nil, nil, constants.SortOrderNil, "", -1)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "List Transactions Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/accounts/1234/transactions",
				).Return(0, "", errors.New("test error"))

				return testClient.ListTransactions("1234", nil, nil, constants.SortOrderNil, "", -1)
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "List Transactions Fails Without Account ID Key",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.ListTransactions("", nil, nil, constants.SortOrderNil, "", -1)
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "List Transaction Details",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/accounts/1234/transactions/5678",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.ListTransactionDetails("1234", "5678")
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "List Transaction Details Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/accounts/1234/transactions/5678",
				).Return(0, "", errors.New("test error"))

				return testClient.ListTransactionDetails("1234", "5678")
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "List Transaction Details Fails Without Account ID Key",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.ListTransactionDetails("", "5678")
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "List Transaction Details Fails Without Transaction ID",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.ListTransactionDetails("1234", "")
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "View Portfolio With All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/accounts/1234/portfolio?count=5&lotsRequired=true&marketSession=REGULAR&pageNumber=6&sortBy=SYMBOL&sortOrder=ASC&totalsRequired=true&view=COMPLETE",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.ViewPortfolio(
					"1234", 5, constants.PortfolioSortBySymbol, constants.SortOrderAsc, "6",
					constants.MarketSessionRegular, true, true, constants.PortfolioViewComplete,
				)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "View Portfolio Can Omit All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/accounts/1234/portfolio?lotsRequired=true&totalsRequired=true",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.ViewPortfolio(
					"1234", -1, constants.PortfolioSortByNil, constants.SortOrderNil, "",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "View Portfolio Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/accounts/1234/portfolio?lotsRequired=true&totalsRequired=true",
				).Return(0, "", errors.New("test error"))

				return testClient.ViewPortfolio(
					"1234", -1, constants.PortfolioSortByNil, constants.SortOrderNil, "",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				)
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "View Portfolio Fails Without Account ID Key",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.ViewPortfolio(
					"", -1, constants.PortfolioSortByNil, constants.SortOrderNil, "",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				)
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "View Portfolio Fails If Count Is Too Big",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.ViewPortfolio(
					"1234", constants.PortfolioMaxCount+1, constants.PortfolioSortByNil, constants.SortOrderNil, "",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				)
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "List Position Lots Details",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/accounts/1234/portfolio/5678",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.ListPositionLotsDetails(
					"1234", 5678,
				)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "List Position Lots Details Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/accounts/1234/portfolio/5678",
				).Return(0, "", errors.New("test error"))

				return testClient.ListPositionLotsDetails(
					"1234", 5678,
				)
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "List Position Lots Details Fails With Empty Account Key",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.ListPositionLotsDetails("", 5678)
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "List Alerts With All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/user/alerts?category=ACCOUNT&count=1&direction=ASC&search=FOO&status=UNREAD",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.ListAlerts(
					1, constants.AlertCategoryAccount, constants.AlertStatusUnread, constants.SortOrderAsc, "FOO",
				)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "List Alerts Can Omit All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/user/alerts",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.ListAlerts(
					-1, constants.AlertCategoryNil, constants.AlertStatusNil, constants.SortOrderNil, "",
				)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "List Alerts Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/user/alerts",
				).Return(0, "", errors.New("test error"))

				return testClient.ListAlerts(
					-1, constants.AlertCategoryNil, constants.AlertStatusNil, constants.SortOrderNil, "",
				)
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "List Alerts Fails With Count Too Big",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.ListAlerts(
					301, constants.AlertCategoryAccount, constants.AlertStatusUnread, constants.SortOrderAsc, "FOO",
				)
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "List Alert Details",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/user/alerts/1234?htmlTags=true",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.ListAlertDetails("1234", true)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "List Alert Details Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/user/alerts/1234?htmlTags=true",
				).Return(0, "", errors.New("test error"))

				return testClient.ListAlertDetails("1234", true)
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "Delete Alerts",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "DELETE", "https://api.etrade.com/v1/user/alerts/1234,5678",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.DeleteAlerts([]string{"1234", "5678"})
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "Delete Alerts Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "DELETE", "https://api.etrade.com/v1/user/alerts/1234,5678",
				).Return(0, "", errors.New("test error"))

				return testClient.DeleteAlerts([]string{"1234", "5678"})
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "Get Quotes With All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/market/quote/GOOG?detailFlag=ALL&requireEarningsDate=true&skipMiniOptionsCheck=false",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.GetQuotes([]string{"GOOG"}, constants.QuoteDetailAll, true, false)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "Get Quotes Can Omit All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/market/quote/GOOG?requireEarningsDate=true&skipMiniOptionsCheck=false",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.GetQuotes([]string{"GOOG"}, constants.QuoteDetailNil, true, false)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "Get Quotes Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/market/quote/GOOG?requireEarningsDate=true&skipMiniOptionsCheck=false",
				).Return(0, "", errors.New("test error"))

				return testClient.GetQuotes([]string{"GOOG"}, constants.QuoteDetailNil, true, false)
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "Get Quotes Fails Without Symbols",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.GetQuotes([]string{}, constants.QuoteDetailNil, true, false)
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "Get Quotes Overrides When More Than 25 Symbols",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/market/quote/1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26?detailFlag=ALL&overrideSymbolCount=true&requireEarningsDate=true&skipMiniOptionsCheck=false",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.GetQuotes(
					[]string{
						"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17",
						"18", "19", "20", "21", "22", "23", "24", "25", "26",
					}, constants.QuoteDetailAll, true, false,
				)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "Get Quotes Fails With More Than 50 Symbols",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.GetQuotes(
					[]string{
						"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17",
						"18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33",
						"34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49",
						"50", "51",
					}, constants.QuoteDetailAll, true, false,
				)
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "Lookup Product",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/market/lookup/A",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.LookupProduct("A")
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "Lookup Product Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/market/lookup/A",
				).Return(0, "", errors.New("test error"))

				return testClient.LookupProduct("A")
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "Lookup Product Fails With Empty Search String",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.LookupProduct("")
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "Get Option Chains With All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/market/optionchains?chainType=CALL&expiryDay=3&expiryMonth=2&expiryYear=1&includeWeekly=true&noOfStrikes=5&optionCategory=ALL&priceType=ALL&skipAdjusted=true&strikePriceNear=4&symbol=GOOG",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.GetOptionChains(
					"GOOG",
					1, 2, 3,
					4, 5,
					true, true,
					constants.OptionCategoryAll, constants.OptionChainTypeCall, constants.OptionPriceTypeAll,
				)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "Get Option Chains Can Omit All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/market/optionchains?includeWeekly=true&skipAdjusted=true&symbol=GOOG",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.GetOptionChains(
					"GOOG",
					-1, -1, -1,
					-1, -1,
					true, true,
					constants.OptionCategoryNil, constants.OptionChainTypeNil, constants.OptionPriceTypeNil,
				)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "Get Option Chains Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/market/optionchains?includeWeekly=true&skipAdjusted=true&symbol=GOOG",
				).Return(0, "", errors.New("test error"))

				return testClient.GetOptionChains(
					"GOOG",
					-1, -1, -1,
					-1, -1,
					true, true,
					constants.OptionCategoryNil, constants.OptionChainTypeNil, constants.OptionPriceTypeNil,
				)
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "Get Option Chains Fails Without Symbol",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.GetOptionChains(
					"",
					-1, -1, -1,
					-1, -1,
					true, true,
					constants.OptionCategoryNil, constants.OptionChainTypeNil, constants.OptionPriceTypeNil,
				)
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "Get Option Expire Dates With All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/market/optionexpiredate?expiryType=ALL&symbol=GOOG",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.GetOptionExpireDates("GOOG", constants.OptionExpiryTypeAll)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "Get Option Expire Dates Can Omit All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/market/optionexpiredate?symbol=GOOG",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.GetOptionExpireDates("GOOG", constants.OptionExpiryTypeNil)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "Get Option Expire Dates Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/market/optionexpiredate?symbol=GOOG",
				).Return(0, "", errors.New("test error"))

				return testClient.GetOptionExpireDates("GOOG", constants.OptionExpiryTypeNil)
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "Get Option Expire Dates Fails Without Symbol",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.GetOptionExpireDates("", constants.OptionExpiryTypeNil)
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "List Orders With All Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET",
					"https://api.etrade.com/v1/accounts/1234/orders?count=5&fromDate=01012023&marker=TestMarker&marketSession=REGULAR&securityType=EQ&status=OPEN&symbol=A%2CB&toDate=01022023&transactionType=BUY",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.ListOrders(
					"1234", "TestMarker", 5, constants.OrderStatusOpen,
					etradelibtest.CreateTime(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
					etradelibtest.CreateTime(2023, time.January, 2, 0, 0, 0, 0, time.UTC),
					[]string{"A", "B"},
					constants.OrderSecurityTypeEquity, constants.OrderTransactionTypeBuy,
					constants.MarketSessionRegular,
				)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "List Orders Can Omit All Optional Arguments",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/accounts/1234/orders",
				).Return(http.StatusOK, testResponseData, nil)

				return testClient.ListOrders(
					"1234", "", -1, constants.OrderStatusNil, nil, nil, nil, constants.OrderSecurityTypeNil,
					constants.OrderTransactionTypeNil, constants.MarketSessionNil,
				)
			},
			expectResponse: []byte(testResponseData),
			expectErr:      false,
		},
		{
			name: "List Orders Fails On HTTP Error",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				clientMock.On(
					"Do", "GET", "https://api.etrade.com/v1/accounts/1234/orders",
				).Return(0, "", errors.New("test error"))

				return testClient.ListOrders(
					"1234", "", -1, constants.OrderStatusNil, nil, nil, nil, constants.OrderSecurityTypeNil,
					constants.OrderTransactionTypeNil, constants.MarketSessionNil,
				)
			},
			expectResponse: []byte(nil),
			expectErr:      true,
		},
		{
			name: "List Orders Fails With Too Many Symbols",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.ListOrders(
					"1234", "", -1, constants.OrderStatusNil, nil, nil, []string{
						"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17",
						"18", "19", "20", "21", "22", "23", "24", "25", "26",
					}, constants.OrderSecurityTypeNil,
					constants.OrderTransactionTypeNil, constants.MarketSessionNil,
				)
			},
			expectResponse: nil,
			expectErr:      true,
		},
		{
			name: "List Orders Fails Without Account ID Key",
			testFn: func(testClient ETradeClient, clientMock *httpClientMock) ([]byte, error) {
				return testClient.ListOrders(
					"", "", -1, constants.OrderStatusNil, nil, nil, nil, constants.OrderSecurityTypeNil,
					constants.OrderTransactionTypeNil, constants.MarketSessionNil,
				)
			},
			expectResponse: nil,
			expectErr:      true,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				configMock := new(oAuthConfigMock)
				clientMock := new(httpClientMock)
				testClient := createMockClient(clientMock, configMock, true, "", "", "", "", "", "")
				// Call the Method Under Test
				actualResponse, err := tt.testFn(testClient, clientMock)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectResponse, actualResponse)
				clientMock.AssertExpectations(t)
				configMock.AssertExpectations(t)
			},
		)
	}
}
