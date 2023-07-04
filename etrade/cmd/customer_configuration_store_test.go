package cmd

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLoadCustomerConfigurationStore(t *testing.T) {
	tests := []struct {
		name        string
		testJson    string
		expectErr   bool
		expectValue interface{}
	}{
		{
			name: "Can Load Store From Json",
			testJson: `{
  "TestCustomerId": {
    "customerName": "TestName",
    "customerProduction": true,
    "customerConsumerKey": "TestKey",
    "customerConsumerSecret": "TestSecret"
  }
}`,
			expectErr: false,
			expectValue: &CustomerConfigurationStore{
				customerConfigMap: map[string]CustomerConfiguration{
					"TestCustomerId": {
						CustomerName:           "TestName",
						CustomerProduction:     true,
						CustomerConsumerKey:    "TestKey",
						CustomerConsumerSecret: "TestSecret",
					},
				},
			},
		},
		{
			name: "Load Succeeds With Missing Fields",
			testJson: `{
}`,
			expectErr: false,
			expectValue: &CustomerConfigurationStore{
				customerConfigMap: map[string]CustomerConfiguration{},
			},
		},
		{
			name: "Load Succeeds With Extra Fields",
			testJson: `{
  "TestCustomerId": {
    "customerName": "TestName",
    "customerProduction": true,
    "customerConsumerKey": "TestKey",
    "customerConsumerSecret": "TestSecret",
    "unexpectedField": "TestUnexpected"
  }
}`,
			expectErr: false,
			expectValue: &CustomerConfigurationStore{
				customerConfigMap: map[string]CustomerConfiguration{
					"TestCustomerId": {
						CustomerName:           "TestName",
						CustomerProduction:     true,
						CustomerConsumerKey:    "TestKey",
						CustomerConsumerSecret: "TestSecret",
					},
				},
			},
		},
		{
			name: "Load Fails With Bad JSON",
			testJson: `{
  "TestCustomerId": {
}`,
			expectErr:   true,
			expectValue: (*CustomerConfigurationStore)(nil),
		},
	}

	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				reader := strings.NewReader(tt.testJson)
				// Call the Method Under Test
				actualValue, err := LoadCustomerConfigurationStore(reader)
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

func TestSaveCustomerConfigurationStore(t *testing.T) {
	testStore := CustomerConfigurationStore{
		customerConfigMap: map[string]CustomerConfiguration{
			"TestCustomerId": {
				CustomerName:           "TestName",
				CustomerProduction:     true,
				CustomerConsumerKey:    "TestKey",
				CustomerConsumerSecret: "TestSecret",
			},
		},
	}

	expectedJson := `{
  "TestCustomerId": {
    "customerName": "TestName",
    "customerProduction": true,
    "customerConsumerKey": "TestKey",
    "customerConsumerSecret": "TestSecret"
  }
}` + "\n"

	actualJson := strings.Builder{}
	err := SaveCustomerConfigurationStore(&actualJson, &testStore)
	assert.Nil(t, err)
	assert.Equal(t, expectedJson, actualJson.String())
}

func TestCustomerConfigurationStore_GetCustomerConfigurationById(t *testing.T) {
	expectedConfig := CustomerConfiguration{
		CustomerName:           "TestName",
		CustomerProduction:     true,
		CustomerConsumerKey:    "TestKey",
		CustomerConsumerSecret: "TestSecret",
	}

	testStore := CustomerConfigurationStore{
		customerConfigMap: map[string]CustomerConfiguration{
			"TestCustomerId": expectedConfig,
		},
	}

	actualConfig, err := testStore.GetCustomerConfigurationById("TestCustomerId")
	assert.Nil(t, err)
	assert.Equal(t, &expectedConfig, actualConfig)
}

func TestCustomerConfigurationStore_SetCustomerConfigurationForId(t *testing.T) {
	testConfig := CustomerConfiguration{
		CustomerName:           "TestName",
		CustomerProduction:     true,
		CustomerConsumerKey:    "TestKey",
		CustomerConsumerSecret: "TestSecret",
	}

	actualStore := CustomerConfigurationStore{
		customerConfigMap: map[string]CustomerConfiguration{},
	}

	expectedStore := CustomerConfigurationStore{
		customerConfigMap: map[string]CustomerConfiguration{
			"TestCustomerId": testConfig,
		},
	}

	actualStore.SetCustomerConfigurationForId("TestCustomerId", &testConfig)
	assert.Equal(t, expectedStore, actualStore)
}

func TestCustomerConfigurationStore_GetAllConfigurations(t *testing.T) {
	expectedMap := map[string]CustomerConfiguration{
		"TestCustomerId1": {
			CustomerName:           "TestName1",
			CustomerProduction:     true,
			CustomerConsumerKey:    "TestKey1",
			CustomerConsumerSecret: "TestSecret1",
		},
	}

	testStore := CustomerConfigurationStore{
		customerConfigMap: expectedMap,
	}

	actualMap := testStore.GetAllConfigurations()
	assert.Equal(t, expectedMap, actualMap)
}
