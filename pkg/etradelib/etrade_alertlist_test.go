package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeAlertListFromResponse(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeAlertList
	}{
		{
			name: "Creates Alert List",
			testJson: `
{
  "AlertsResponse": {
    "Alert": [
      {
        "id": 1234
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
				},
			},
		},
		{
			name: "Creates Empty List",
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
			name: "Fails With Bad JSON",
			testJson: `
{
  "AlertsResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails With Missing AlertsResponse",
			testJson: `
{
  "MISSING": {
    "Alert": [
      {
        "id": 1234
      }
    ]
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails With Missing Alert Id",
			testJson: `
{
  "AlertsResponse": {
    "Alert": [
      {
        "MISSING": 1234
      }
    ]
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue, err := CreateETradeAlertListFromResponse([]byte(tt.testJson))
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
		name        string
		testObject  ETradeAlertList
		expectValue []ETradeAlert
	}{
		{
			name: "Returns All Alerts",
			testObject: &eTradeAlertList{
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
			name: "Can Return Empty List",
			testObject: &eTradeAlertList{
				alerts: []ETradeAlert{},
			},
			expectValue: []ETradeAlert{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testObject.GetAllAlerts()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeAlertList_GetAlertById(t *testing.T) {
	tests := []struct {
		name        string
		testObject  ETradeAlertList
		testId      int64
		expectValue ETradeAlert
	}{
		{
			name: "Returns Alert For Valid ID",
			testObject: &eTradeAlertList{
				alerts: []ETradeAlert{
					&eTradeAlert{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"id": json.Number("1234"),
						},
					},
				},
			},
			testId: 1234,
			expectValue: &eTradeAlert{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"id": json.Number("1234"),
				},
			},
		},
		{
			name: "Returns Nil For Invalid ID",
			testObject: &eTradeAlertList{
				alerts: []ETradeAlert{
					&eTradeAlert{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"id": json.Number("1234"),
						},
					},
				},
			},
			testId:      5678,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testObject.GetAlertById(tt.testId)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeAlertList_AsJsonMap(t *testing.T) {
	testObject := &eTradeAlertList{
		alerts: []ETradeAlert{
			&eTradeAlert{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"id": "1234",
				},
			},
		},
	}

	expectValue := jsonmap.JsonMap{
		"alerts": jsonmap.JsonSlice{
			jsonmap.JsonMap{
				"id": "1234",
			},
		},
	}

	// Call the Method Under Test
	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectValue, actualValue)
}
