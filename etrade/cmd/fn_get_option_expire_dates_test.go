package cmd

import (
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetOptionExpireDates(t *testing.T) {
	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Gets Option Expire Dates",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testGetOptionExpireDatesResponse := []byte(`
{
  "OptionExpireDateResponse": {
    "ExpirationDate": [
      {
        "testKey": "testValue"
      }
    ]
  }
}`)
				mockClient.On(
					"GetOptionExpireDates", "TestSymbol", constants.OptionExpiryTypeNil,
				).Return(testGetOptionExpireDatesResponse, nil)

				return GetOptionExpireDates(mockClient, "TestSymbol", constants.OptionExpiryTypeNil)
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"optionExpireDates": jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"testKey": "testValue",
					},
				},
			},
		},
		{
			name: "Fails On GetOptionExpireDates Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On(
					"GetOptionExpireDates", "TestSymbol", constants.OptionExpiryTypeNil,
				).Return([]byte{}, errors.New("test error"))

				return GetOptionExpireDates(mockClient, "TestSymbol", constants.OptionExpiryTypeNil)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testGetOptionExpireDatesResponse := []byte(`
{
  "MISSING": {
    "ExpirationDate": [
      {
        "testKey": "testValue"
      }
    ]
  }
}`)
				mockClient.On(
					"GetOptionExpireDates", "TestSymbol", constants.OptionExpiryTypeNil,
				).Return(testGetOptionExpireDatesResponse, nil)

				return GetOptionExpireDates(mockClient, "TestSymbol", constants.OptionExpiryTypeNil)
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
