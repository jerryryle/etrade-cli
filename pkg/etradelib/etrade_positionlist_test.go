package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradePositionListFromResponse(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradePositionList
	}{
		{
			name: "Creates List",
			testJson: `
{
  "PortfolioResponse": {
    "Totals": {
      "bogusTotal": 9999
    },
    "AccountPortfolio": [
      {
        "nextPageNo": "2",
        "Position": [
          {
            "positionId": 1234
          }
        ]
      }
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradePositionList{
				positions: []ETradePosition{
					&eTradePosition{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"positionId": json.Number("1234"),
						},
					},
				},
				totalsMap: jsonmap.JsonMap{
					"bogusTotal": json.Number("9999"),
				},
				nextPage: "2",
			},
		},
		{
			name: "Can Create Empty List and Nil Totals Map",
			testJson: `
{
  "PortfolioResponse": {
    "AccountPortfolio": [
      {
        "Position": [
        ]
      }
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradePositionList{
				positions: []ETradePosition{},
				totalsMap: nil,
				nextPage:  "",
			},
		},
		{
			name: "Fails With Invalid JSON",
			testJson: `
{
  "PortfolioResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails With Missing AccountPortfolio",
			testJson: `
{
  "PortfolioResponse": {
    "Totals": {
      "bogusTotal": 9999
    },
    "MISSING": [
      {
        "nextPageNo": "2",
        "Position": [
          {
            "positionId": 1234
          }
        ]
      }
    ]
  }
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails With Missing positionId",
			testJson: `
{
  "PortfolioResponse": {
    "Totals": {
      "bogusTotal": 9999
    },
    "AccountPortfolio": [
      {
        "nextPageNo": "2",
        "Position": [
          {
            "MISSING": 1234
          }
        ]
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
				actualValue, err := CreateETradePositionListFromResponse([]byte(tt.testJson))
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

func TestETradePositionList_GetAllPositions(t *testing.T) {
	tests := []struct {
		name        string
		testObject  ETradePositionList
		expectValue []ETradePosition
	}{
		{
			name: "Returns All Positions",
			testObject: &eTradePositionList{
				positions: []ETradePosition{
					&eTradePosition{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"positionId": json.Number("1234"),
						},
					},
					&eTradePosition{
						id: 5678,
						jsonMap: jsonmap.JsonMap{
							"positionId": json.Number("5678"),
						},
					},
				},
			},
			expectValue: []ETradePosition{
				&eTradePosition{
					id: 1234,
					jsonMap: jsonmap.JsonMap{
						"positionId": json.Number("1234"),
					},
				},
				&eTradePosition{
					id: 5678,
					jsonMap: jsonmap.JsonMap{
						"positionId": json.Number("5678"),
					},
				},
			},
		},
		{
			name: "Can Return Empty List",
			testObject: &eTradePositionList{
				positions: []ETradePosition{},
			},
			expectValue: []ETradePosition{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testObject.GetAllPositions()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradePositionList_GetPositionById(t *testing.T) {
	tests := []struct {
		name        string
		testObject  ETradePositionList
		testId      int64
		expectValue ETradePosition
	}{
		{
			name: "Returns Account For Valid ID",
			testObject: &eTradePositionList{
				positions: []ETradePosition{
					&eTradePosition{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"positionId": json.Number("1234"),
						},
					},
				},
			},
			testId: 1234,
			expectValue: &eTradePosition{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"positionId": json.Number("1234"),
				},
			},
		},
		{
			name: "Returns Nil For Invalid ID",
			testObject: &eTradePositionList{
				positions: []ETradePosition{
					&eTradePosition{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"positionId": json.Number("1234"),
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
				actualValue := tt.testObject.GetPositionById(tt.testId)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradePositionList_AddPageFromResponse(t *testing.T) {
	startingObject := &eTradePositionList{
		positions: []ETradePosition{
			&eTradePosition{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"positionId": json.Number("1234"),
				},
			},
		},
		totalsMap: jsonmap.JsonMap{
			"bogusTotal": json.Number("9999"),
		},
		nextPage: "2",
	}

	tests := []struct {
		name        string
		startValue  ETradePositionList
		testJson    string
		expectErr   bool
		expectValue ETradePositionList
	}{
		{
			name:       "Can Add Pages",
			startValue: startingObject,
			testJson: `
{
  "PortfolioResponse": {
    "Totals": {
      "bogusTotal": 8888
    },
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
}`,
			expectErr: false,
			expectValue: &eTradePositionList{
				positions: []ETradePosition{
					&eTradePosition{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"positionId": json.Number("1234"),
						},
					},
					// Positions in subsequent pages are appended to
					// the position list.
					&eTradePosition{
						id: 5678,
						jsonMap: jsonmap.JsonMap{
							"positionId": json.Number("5678"),
						},
					},
				},
				totalsMap: jsonmap.JsonMap{
					// The totals come from the first page only. Totals
					// in subsequent pages are ignored.
					"bogusTotal": json.Number("9999"),
				},
				nextPage: "",
			},
		},
		{
			name:       "Fails on Invalid JSON",
			startValue: startingObject,
			testJson: `
{
  "PortfolioResponse": {
}`,
			expectErr:   true,
			expectValue: startingObject,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				positionList := tt.startValue

				// Call the Method Under Test
				err := positionList.AddPageFromResponse([]byte(tt.testJson))
				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.Nil(t, err)
				}
				assert.Equal(t, tt.expectValue, positionList)
			},
		)
	}
}

func TestETradePositionList_NextPage(t *testing.T) {
	testObject := &eTradePositionList{
		positions: []ETradePosition{},
		totalsMap: jsonmap.JsonMap{},
		nextPage:  "1234",
	}
	assert.Equal(t, "1234", testObject.NextPage())

	testObject = &eTradePositionList{
		positions: []ETradePosition{},
		totalsMap: jsonmap.JsonMap{},
		nextPage:  "",
	}
	assert.Equal(t, "", testObject.NextPage())
}

func TestETradePositionList_AsJsonMap(t *testing.T) {
	testObject := &eTradePositionList{
		positions: []ETradePosition{
			&eTradePosition{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"positionId": json.Number("1234"),
				},
			},
		},
		totalsMap: jsonmap.JsonMap{
			"testTotal": "testValue",
		},
	}

	expectValue := jsonmap.JsonMap{
		"positions": jsonmap.JsonSlice{
			jsonmap.JsonMap{
				"positionId": json.Number("1234"),
			},
		},
		"totals": jsonmap.JsonMap{
			"testTotal": "testValue",
		},
	}

	// Call the Method Under Test
	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectValue, actualValue)
}
