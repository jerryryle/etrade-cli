package cmd

import (
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetQuotes(t *testing.T) {
	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Gets Quotes",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testGetQuotesResponse := []byte(`
{
  "QuoteResponse": {
    "QuoteData": [
      {
        "testKey": "testValue"
      }
    ]
  }
}`)
				mockClient.On(
					"GetQuotes", []string{"TestSymbol"}, constants.QuoteDetailFlagNil, true, true,
				).Return(testGetQuotesResponse, nil)

				return GetQuotes(
					mockClient, []string{"TestSymbol"}, constants.QuoteDetailFlagNil, true,
					true,
				)
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"quotes": jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"testKey": "testValue",
					},
				},
			},
		},
		{
			name: "Fails On GetQuotes Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {

				mockClient.On(
					"GetQuotes", []string{"TestSymbol"}, constants.QuoteDetailFlagNil, true, true,
				).Return([]byte{}, errors.New("test error"))

				return GetQuotes(
					mockClient, []string{"TestSymbol"}, constants.QuoteDetailFlagNil, true,
					true,
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testGetQuotesResponse := []byte(`
{
  "QuoteResponse": {
}`)
				mockClient.On(
					"GetQuotes", []string{"TestSymbol"}, constants.QuoteDetailFlagNil, true, true,
				).Return(testGetQuotesResponse, nil)

				return GetQuotes(
					mockClient, []string{"TestSymbol"}, constants.QuoteDetailFlagNil, true,
					true,
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
