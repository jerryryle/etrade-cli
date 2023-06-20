package etradelib

import (
	"encoding/json"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradePositionList(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradePositionList
	}{
		{
			name: "CreateETradePositionList Creates List With Valid Response",
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
          },
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
					&eTradePosition{
						id: 5678,
						jsonMap: jsonmap.JsonMap{
							"positionId": json.Number("5678"),
						},
					},
				},
				totalsJsonMap: jsonmap.JsonMap{
					"bogusTotal": json.Number("9999"),
				},
				nextPage: "2",
			},
		},
		{
			name: "CreateETradePositionList Can Create Empty List and Nil Totals Map",
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
				positions:     []ETradePosition{},
				totalsJsonMap: nil,
				nextPage:      "",
			},
		},
		{
			name: "CreateETradePositionList Fails With Invalid Response",
			// The "AccountPortfolio" key holds a map value instead of a slice
			// value in the following string
			testJson: `
{
  "PortfolioResponse": {
    "AccountPortfolio": {
      "Position": [
      ]
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
				actualValue, err := CreateETradePositionList(responseMap)
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
		name             string
		testPositionList ETradePositionList
		expectValue      []ETradePosition
	}{
		{
			name: "GetAllPositions Returns All Positions",
			testPositionList: &eTradePositionList{
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
			name: "GetAllPositions Can Return Empty List",
			testPositionList: &eTradePositionList{
				positions: []ETradePosition{},
			},
			expectValue: []ETradePosition{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testPositionList.GetAllPositions()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradePositionList_GetPositionById(t *testing.T) {
	tests := []struct {
		name             string
		testPositionList ETradePositionList
		testPositionID   int64
		expectValue      ETradePosition
	}{
		{
			name: "GetPositionById Returns Account For Valid ID",
			testPositionList: &eTradePositionList{
				positions: []ETradePosition{
					&eTradePosition{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"positionId": json.Number("1234"),
						},
					},
				},
			},
			testPositionID: 1234,
			expectValue: &eTradePosition{
				id: 1234,
				jsonMap: jsonmap.JsonMap{
					"positionId": json.Number("1234"),
				},
			},
		},
		{
			name: "GetPositionById Returns Nil For Invalid ID",
			testPositionList: &eTradePositionList{
				positions: []ETradePosition{
					&eTradePosition{
						id: 1234,
						jsonMap: jsonmap.JsonMap{
							"positionId": json.Number("1234"),
						},
					},
				},
			},
			testPositionID: 5678,
			expectValue:    nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testPositionList.GetPositionById(tt.testPositionID)
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradePositionList_AddPage(t *testing.T) {
	type pageTest struct {
		testJson    string
		expectErr   bool
		expectValue ETradePositionList
	}
	tests := []struct {
		name      string
		pageTests []pageTest
	}{
		{
			name: "AddPage Can Add Pages",
			pageTests: []pageTest{
				{
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
						totalsJsonMap: jsonmap.JsonMap{
							"bogusTotal": json.Number("9999"),
						},
						nextPage: "2",
					},
				},
				{
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
						totalsJsonMap: jsonmap.JsonMap{
							// The totals come from the first page only. Totals
							// in subsequent pages are ignored.
							"bogusTotal": json.Number("9999"),
						},
						nextPage: "",
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				var positionList ETradePositionList
				for testIndex, pt := range tt.pageTests {
					responseMap, err := NewNormalizedJsonMap([]byte(pt.testJson))
					require.Nil(t, err)

					if testIndex == 0 {
						positionList, err = CreateETradePositionList(responseMap)
					} else {
						// Call the Method Under Test
						err = positionList.AddPage(responseMap)
					}
					if pt.expectErr {
						assert.Error(t, err)
					} else {
						assert.Nil(t, err)
					}
					assert.Equal(t, pt.expectValue, positionList)
				}
			},
		)
	}
}

func TestETradePositionList_NextPage(t *testing.T) {
	testPositionList := &eTradePositionList{
		positions:     []ETradePosition{},
		totalsJsonMap: jsonmap.JsonMap{},
		nextPage:      "1234",
	}
	assert.Equal(t, "1234", testPositionList.NextPage())

	testPositionList = &eTradePositionList{
		positions:     []ETradePosition{},
		totalsJsonMap: jsonmap.JsonMap{},
		nextPage:      "",
	}
	assert.Equal(t, "", testPositionList.NextPage())
}
