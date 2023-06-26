package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeLookupResultListFromResponse(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeLookupResultList
	}{
		{
			name: "Creates List With Valid Response",
			testJson: `
{
  "LookupResponse": {
    "Data": [
      {
        "key1": "value1"
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
				},
			},
		},
		{
			name: "Creates Empty List",
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
			name: "Fails With Invalid JSON",
			testJson: `
{
  "LookupResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails With Missing LookupResponse",
			testJson: `
{
  "MISSING": {
    "Data": [
      {
        "key1": "value1"
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
				actualValue, err := CreateETradeLookupResultListFromResponse([]byte(tt.testJson))
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
		name        string
		testObject  ETradeLookupResultList
		expectValue []ETradeLookupResult
	}{
		{
			name: "Returns All LookupResults",
			testObject: &eTradeLookupResultList{
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
			name: "Can Return Empty List",
			testObject: &eTradeLookupResultList{
				results: []ETradeLookupResult{},
			},
			expectValue: []ETradeLookupResult{},
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue := tt.testObject.GetAllResults()
				assert.Equal(t, tt.expectValue, actualValue)
			},
		)
	}
}

func TestETradeLookupResultList_AsJsonMap(t *testing.T) {
	testObject := &eTradeLookupResultList{
		results: []ETradeLookupResult{
			&eTradeLookupResult{
				jsonMap: jsonmap.JsonMap{
					"key1": "value1",
				},
			},
		},
	}

	expectValue := jsonmap.JsonMap{
		"results": jsonmap.JsonSlice{
			jsonmap.JsonMap{
				"key1": "value1",
			},
		},
	}

	// Call the Method Under Test
	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectValue, actualValue)
}
