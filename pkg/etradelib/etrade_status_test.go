package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeStatus(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeStatus
	}{
		{
			name: "Creates Success Status",
			testJson: `
{
  "status": "success"
}`,
			expectErr: false,
			expectValue: &eTradeStatus{
				isSuccess:    true,
				errorMessage: "",
				jsonMap: jsonmap.JsonMap{
					"status": "success",
				},
			},
		},
		{
			name: "Creates Error Status",
			testJson: `
{
  "status": "error",
  "error": "test error message"
}`,
			expectErr: false,
			expectValue: &eTradeStatus{
				isSuccess:    false,
				errorMessage: "test error message",
				jsonMap: jsonmap.JsonMap{
					"status": "error",
					"error":  "test error message",
				},
			},
		},
		{
			name: "Fails With Unexpected Status",
			testJson: `
{
  "status": "bad status"
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails Without Status",
			testJson: `
{
  "test key": "test value"
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Error Status Fails Without Error Message",
			testJson: `
{
  "status": "error"
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails On Bad JSON",
			testJson: `
{
  "status": 
}`,
			expectErr:   true,
			expectValue: nil,
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				// Call the Method Under Test
				actualValue, err := CreateETradeStatusFromResponse([]byte(tt.testJson))
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

func TestETradeStatus_IsSuccess(t *testing.T) {
	testSuccessStatus := eTradeStatus{
		isSuccess:    true,
		errorMessage: "",
		jsonMap: jsonmap.JsonMap{
			"status": "success",
		},
	}

	testErrorStatus := eTradeStatus{
		isSuccess:    false,
		errorMessage: "test error message",
		jsonMap: jsonmap.JsonMap{
			"status": "error",
			"error":  "test error message",
		},
	}

	assert.True(t, testSuccessStatus.IsSuccess())
	assert.False(t, testErrorStatus.IsSuccess())
}

func TestETradeStatus_GetErrorMessage(t *testing.T) {
	testErrorStatus := eTradeStatus{
		isSuccess:    false,
		errorMessage: "test error message",
		jsonMap: jsonmap.JsonMap{
			"status": "error",
			"error":  "test error message",
		},
	}

	assert.Equal(t, "test error message", testErrorStatus.GetErrorMessage())
}

func TestETradeStatus_AsJsonMap(t *testing.T) {
	testObject := &eTradeStatus{
		isSuccess:    false,
		errorMessage: "test error message",
		jsonMap: jsonmap.JsonMap{
			"status": "error",
			"error":  "test error message",
		},
	}

	expectedValue := jsonmap.JsonMap{
		"status": "error",
		"error":  "test error message",
	}

	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectedValue, actualValue)
}
