package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCustomerList(t *testing.T) {
	testCfgStore := CustomerConfigurationStore{
		customerConfigMap: map[string]CustomerConfiguration{
			"TestCustomerId": {
				CustomerName:           "Test Customer Name",
				CustomerProduction:     true,
				CustomerConsumerKey:    "Test Customer Consumer Key",
				CustomerConsumerSecret: "Test Customer Consumer Secret",
			},
		},
	}

	expectedResult := jsonmap.JsonMap{
		"customers": jsonmap.JsonSlice{
			jsonmap.JsonMap{
				"customerId":       "TestCustomerId",
				"customerName":     "Test Customer Name",
				"productionAccess": true,
			},
		},
	}

	actualResult, err := GetCustomerList(&testCfgStore)
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, actualResult)
}
