package etradelib

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateETradeBalancesFromResponse(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue ETradeBalances
	}{
		{
			name: "Creates Balances",
			testJson: `
{
  "BalanceResponse": {
    "balanceKey": "BalanceValue"
  }
}`,
			expectErr: false,
			expectValue: &eTradeBalances{
				balancesMap: jsonmap.JsonMap{
					"balanceKey": "BalanceValue",
				},
			},
		},
		{
			name: "Fails With Bad JSON",
			testJson: `
{
  "BalanceResponse": {
}`,
			expectErr:   true,
			expectValue: nil,
		},
		{
			name: "Fails With Missing BalanceResponse",
			testJson: `
{
  "MISSING": {
    "balanceKey": "BalanceValue"
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
				actualValue, err := CreateETradeBalancesFromResponse([]byte(tt.testJson))
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

func TestETradeBalances_AsJsonMap(t *testing.T) {
	testObject := &eTradeBalances{
		balancesMap: jsonmap.JsonMap{
			"balanceKey": "BalanceValue",
		},
	}

	expectValue := jsonmap.JsonMap{
		"balanceKey": "BalanceValue",
	}

	// Call the Method Under Test
	actualValue := testObject.AsJsonMap()
	assert.Equal(t, expectValue, actualValue)
}
