package cmd

import (
	"encoding/json"
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViewPortfolio(t *testing.T) {
	testAccountList := []byte(`
{
  "AccountListResponse": {
    "Accounts": {
      "Account": [
        {
          "accountId": "test id",
          "accountIdKey": "test key"
        }
      ]
    }
  }
}`)
	testPortfolioResponse1 := []byte(`
{
  "PortfolioResponse": {
    "AccountPortfolio": [
      {
        "nextPageNo": "test page no",
        "Position": [
          {
            "positionId": 1234
          }
        ]
      }
    ]
  }
}`)
	testPortfolioResponse2 := []byte(`
{
  "PortfolioResponse": {
    "AccountPortfolio": [
      {
        "Position": [
          {
            "positionId": 5678
          }
        ]
      }
    ]
  }
}`)
	testLotsDetails1 := []byte(`
{
  "PositionLotsResponse": {
    "PositionLot": [
      {
        "testKey1": "testValue1"
      }
    ]
  }
}`)
	testLotsDetails2 := []byte(`
{
  "PositionLotsResponse": {
    "PositionLot": [
      {
        "testKey2": "testValue2"
      }
    ]
  }
}`)

	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "View Portfolio With Pagination And Lots",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ViewPortfolio", "test key", 65535, constants.PortfolioSortByNil, constants.SortOrderNil, "",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				).Return(testPortfolioResponse1, nil)
				mockClient.On(
					"ViewPortfolio", "test key", 65535, constants.PortfolioSortByNil, constants.SortOrderNil,
					"test page no",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				).Return(testPortfolioResponse2, nil)
				mockClient.On("ListPositionLotsDetails", "test key", int64(1234)).Return(testLotsDetails1, nil)
				mockClient.On("ListPositionLotsDetails", "test key", int64(5678)).Return(testLotsDetails2, nil)

				return ViewPortfolio(
					mockClient, "test id", constants.PortfolioSortByNil, constants.SortOrderNil,
					constants.MarketSessionNil, true, constants.PortfolioViewNil, true,
				)
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"positions": jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"positionId": json.Number("1234"),
						"lots": jsonmap.JsonSlice{
							jsonmap.JsonMap{
								"testKey1": "testValue1",
							},
						},
					},
					jsonmap.JsonMap{
						"positionId": json.Number("5678"),
						"lots": jsonmap.JsonSlice{
							jsonmap.JsonMap{
								"testKey2": "testValue2",
							},
						},
					},
				},
			},
		},
		{
			name: "Fails With Bad Account ID",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On("ListAccounts").Return(testAccountList, nil)

				return ViewPortfolio(
					mockClient, "bad id", constants.PortfolioSortByNil, constants.SortOrderNil,
					constants.MarketSessionNil, true, constants.PortfolioViewNil, true,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On First Page ViewPortfolio Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ViewPortfolio", "test key", 65535, constants.PortfolioSortByNil, constants.SortOrderNil, "",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				).Return([]byte{}, errors.New("test error"))

				return ViewPortfolio(
					mockClient, "test id", constants.PortfolioSortByNil, constants.SortOrderNil,
					constants.MarketSessionNil, true, constants.PortfolioViewNil, true,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Subsequent Page ViewPortfolio Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ViewPortfolio", "test key", 65535, constants.PortfolioSortByNil, constants.SortOrderNil, "",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				).Return(testPortfolioResponse1, nil)
				mockClient.On(
					"ViewPortfolio", "test key", 65535, constants.PortfolioSortByNil, constants.SortOrderNil,
					"test page no",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				).Return([]byte{}, errors.New("test error"))

				return ViewPortfolio(
					mockClient, "test id", constants.PortfolioSortByNil, constants.SortOrderNil,
					constants.MarketSessionNil, true, constants.PortfolioViewNil, true,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad First Page ViewPortfolio Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testBadPortfolioResponse := []byte(`
{
  "PortfolioResponse": {
}`)

				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ViewPortfolio", "test key", 65535, constants.PortfolioSortByNil, constants.SortOrderNil, "",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				).Return(testBadPortfolioResponse, nil)

				return ViewPortfolio(
					mockClient, "test id", constants.PortfolioSortByNil, constants.SortOrderNil,
					constants.MarketSessionNil, true, constants.PortfolioViewNil, true,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad Subsequent Page ViewPortfolio Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testBadPortfolioResponse := []byte(`
{
  "PortfolioResponse": {
}`)

				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ViewPortfolio", "test key", 65535, constants.PortfolioSortByNil, constants.SortOrderNil, "",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				).Return(testPortfolioResponse1, nil)
				mockClient.On(
					"ViewPortfolio", "test key", 65535, constants.PortfolioSortByNil, constants.SortOrderNil,
					"test page no",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				).Return(testBadPortfolioResponse, nil)

				return ViewPortfolio(
					mockClient, "test id", constants.PortfolioSortByNil, constants.SortOrderNil,
					constants.MarketSessionNil, true, constants.PortfolioViewNil, true,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On ListPositionLotsDetails Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ViewPortfolio", "test key", 65535, constants.PortfolioSortByNil, constants.SortOrderNil, "",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				).Return(testPortfolioResponse1, nil)
				mockClient.On(
					"ViewPortfolio", "test key", 65535, constants.PortfolioSortByNil, constants.SortOrderNil,
					"test page no",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				).Return(testPortfolioResponse2, nil)
				mockClient.On("ListPositionLotsDetails", "test key", int64(1234)).Return(
					[]byte{}, errors.New("test error"),
				)

				return ViewPortfolio(
					mockClient, "test id", constants.PortfolioSortByNil, constants.SortOrderNil,
					constants.MarketSessionNil, true, constants.PortfolioViewNil, true,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad ListPositionLotsDetails Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testBadLotsDetails := []byte(`
{
  "PositionLotsResponse": {
`)

				mockClient.On("ListAccounts").Return(testAccountList, nil)
				mockClient.On(
					"ViewPortfolio", "test key", 65535, constants.PortfolioSortByNil, constants.SortOrderNil, "",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				).Return(testPortfolioResponse1, nil)
				mockClient.On(
					"ViewPortfolio", "test key", 65535, constants.PortfolioSortByNil, constants.SortOrderNil,
					"test page no",
					constants.MarketSessionNil, true, true, constants.PortfolioViewNil,
				).Return(testPortfolioResponse2, nil)
				mockClient.On("ListPositionLotsDetails", "test key", int64(1234)).Return(testBadLotsDetails, nil)

				return ViewPortfolio(
					mockClient, "test id", constants.PortfolioSortByNil, constants.SortOrderNil,
					constants.MarketSessionNil, true, constants.PortfolioViewNil, true,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				mockClient := client.ETradeClientMock{}
				// Call the Method Under Test
				actualValue, err := tt.testFn(&mockClient)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, actualValue)
				mockClient.AssertExpectations(t)
			},
		)
	}
}
