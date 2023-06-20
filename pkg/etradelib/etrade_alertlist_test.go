package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeAlertList(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeAlertList
	}{
		{
			name: "CreateETradeAlertList Creates List With Valid Response",
			testJson: `
{
  "AlertsResponse": {
    "Alert": [
      {
        "id": 1234
      },
      {
        "id": 5678
      }
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeAlertList{
				alerts: []ETradeAlert{
					&eTradeAlert{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"id": json.Number("1234"),
						},
					},
					&eTradeAlert{
						id: 5678,
						jsonMap: jsonmap.JsonMap{
							"id": json.Number("5678"),
						},
					},
				},
			},
		},
		{
			name: "CreateETradeAlertList Can Create Empty List",
			testJson: `
{
  "AlertsResponse": {
    "Alert": [
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeAlertList{
				alerts: []ETradeAlert{},
			},
		},
		{
			name: "CreateETradeAlertList Fails With Invalid Response",
			// The "Alert" level is not an array in the following string
			testJson: `
{
  "AlertsResponse": {
    "Alert": {
      "id": 1234
    }
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				responseMap, err := NewNormalizedJsonMap([]byte(tt.testJson))
				require.Nil(t, err)
				// Call the Method Under Test
				actualValue, err := CreateETradeAlertList(responseMap)
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeAlertList_GetAllAlerts(t *testing.T) {
	tests := []struct {
		name          string
		testAlertList ETradeAlertList
		expectValue   []ETradeAlert
	}{
		{
			name: "GetAllAlerts Returns All Alerts",
			testAlertList: &eTradeAlertList{
				alerts: []ETradeAlert{
					&eTradeAlert{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"id": json.Number("1234"),
						},
					},
					&eTradeAlert{
						id: 5678,
						jsonMap: jsonmap.JsonMap{
							"id": json.Number("5678"),
						},
					},
				},
			},
			expectValue: []ETradeAlert{
				&eTradeAlert{
					id: 1234,
					jsonMap: jsonmap.JsonMap{
						"id": json.Number("1234"),
					},
				},
				&eTradeAlert{
					id: 5678,
					jsonMap: jsonmap.JsonMap{
						"id": json.Number("5678"),
					},
				},
			},
		},
		{
			name: "GetAllAccounts Can Return Empty List",
			testAlertList: &eTradeAlertList{
				alerts: []ETradeAlert{},
			},
			expectValue: []ETradeAlert{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testAlertList.GetAllAlerts()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeAlertList_GetAlertById(t *testing.T) {
	tests := []struct {
		name          string
		testAlertList ETradeAlertList
		testAlertID   int64
		expectValue   ETradeAlert
	}{
		{
			name: "GetAccountById Returns Account For Valid ID",
			testAlertList: &eTradeAlertList{
				alerts: []ETradeAlert{
					&eTradeAlert{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"id": json.Number("1234"),
						},
					},
				},
			},
			testAlertID: 1234,
			expectValue: &eTradeAlert{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"id": json.Number("1234"),
				},
			},
		},
		{
			name: "GetAccountById Returns Nil For Invalid ID",
			testAlertList: &eTradeAlertList{
				alerts: []ETradeAlert{
					&eTradeAlert{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"id": json.Number("1234"),
						},
					},
				},
			},
			testAlertID: 5678,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testAlertList.GetAlertById(tt.testAlertID)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}
