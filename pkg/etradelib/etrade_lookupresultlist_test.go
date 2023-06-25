package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateETradeLookupResultList(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeLookupResultList
	}{
		{
			name: "CreateETradeLookupResultList Creates List With Valid Response",
			testJson: `
{
  "LookupResponse": {
    "Data": [
      {
        "key1": "value1"
      },
      {
        "key2": "value2"
      }
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeLookupResultList{
				results: []ETradeLookupResult{
					&eTradeLookupResult{
						jsonMap: jsonmap.JsonMap{
							"key1": "value1",
						},
					},
					&eTradeLookupResult{
						jsonMap: jsonmap.JsonMap{
							"key2": "value2",
						},
					},
				},
			},
		},
		{
			name: "CreateETradeLookupResultList Can Create Empty List",
			testJson: `
{
  "LookupResponse": {
    "Data": [
    ]
  }
}`,
			expectErr: false,
			expectValue: &eTradeLookupResultList{
				results: []ETradeLookupResult{},
			},
		},
		{
			name: "CreateETradeLookupResultList Fails With Invalid Response",
			// The "LookupResult" level is not an array in the following string
			testJson: `
{
  "LookupResponse": {
    "Data": {
      "key": "value"
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
				actualValue, err := CreateETradeLookupResultList(responseMap)
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

func TestETradeLookupResultList_GetAllLookupResults(t *testing.T) {
	tests := []struct {
		name                 string
		testLookupResultList ETradeLookupResultList
		expectValue          []ETradeLookupResult
	}{
		{
			name: "GetAllLookupResults Returns All LookupResults",
			testLookupResultList: &eTradeLookupResultList{
				results: []ETradeLookupResult{
					&eTradeLookupResult{
						jsonMap: jsonmap.JsonMap{
							"key1": "value1",
						},
					},
					&eTradeLookupResult{
						jsonMap: jsonmap.JsonMap{
							"key2": "value2",
						},
					},
				},
			},
			expectValue: []ETradeLookupResult{
				&eTradeLookupResult{
					jsonMap: jsonmap.JsonMap{
						"key1": "value1",
					},
				},
				&eTradeLookupResult{
					jsonMap: jsonmap.JsonMap{
						"key2": "value2",
					},
				},
			},
		},
		{
			name: "GetAllLookupResults Can Return Empty List",
			testLookupResultList: &eTradeLookupResultList{
				results: []ETradeLookupResult{},
			},
			expectValue: []ETradeLookupResult{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testLookupResultList.GetAllResults()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}
