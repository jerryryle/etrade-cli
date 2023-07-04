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

func TestListAlerts(t *testing.T) {
	type testFn func(mockClient *client.ETradeClientMock) (interface{}, error)
	tests := []struct {
		name        string
		testFn      testFn
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Lists Alerts",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAlertsResponse := []byte(`
{
  "AlertsResponse": {
    "Alert": [
      {
        "id": 1234
      }
    ]
  }
}`)
				mockClient.On(
					"ListAlerts", 5678, constants.AlertCategoryNil, constants.AlertStatusNil, constants.SortOrderNil,
					"test string",
				).Return(testAlertsResponse, nil)
				return ListAlerts(
					mockClient, 5678, constants.AlertCategoryNil, constants.AlertStatusNil, constants.SortOrderNil,
					"test string",
				)
			},
			expectErr: false,
			expectValue: jsonmap.JsonMap{
				"alerts": jsonmap.JsonSlice{
					jsonmap.JsonMap{
						"id": json.Number("1234"),
					},
				},
			},
		},
		{
			name: "Fails On ListAlerts Error",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				mockClient.On(
					"ListAlerts", 5678, constants.AlertCategoryNil, constants.AlertStatusNil, constants.SortOrderNil,
					"test string",
				).Return([]byte{}, errors.New("test error"))
				return ListAlerts(
					mockClient, 5678, constants.AlertCategoryNil, constants.AlertStatusNil, constants.SortOrderNil,
					"test string",
				)
			},
			expectErr:   true,
			expectValue: jsonmap.JsonMap(nil),
		},
		{
			name: "Fails On Bad Response",
			testFn: func(mockClient *client.ETradeClientMock) (interface{}, error) {
				testAlertsResponse := []byte(`
{
  "AlertsResponse": {
}`)
				mockClient.On(
					"ListAlerts", 5678, constants.AlertCategoryNil, constants.AlertStatusNil, constants.SortOrderNil,
					"test string",
				).Return(testAlertsResponse, nil)
				return ListAlerts(
					mockClient, 5678, constants.AlertCategoryNil, constants.AlertStatusNil, constants.SortOrderNil,
					"test string",
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
