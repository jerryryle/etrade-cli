package cmd

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCustomerConfigurationStoreLoadSucceedsWithGoodJson(t *testing.T) {
	jsonData := `{
  "TestCfg1": {
    "customerName": "TestName1",
    "customerProduction": true,
    "customerConsumerKey": "TestKey1",
    "customerConsumerSecret": "TestSecret1"
  },
  "TestCfg2": {
    "customerName": "TestName2",
    "customerProduction": true,
    "customerConsumerKey": "TestKey2",
    "customerConsumerSecret": "TestSecret2"
  }
}`
	reader := strings.NewReader(jsonData)
	ccs, err := LoadCustomerConfigurationStore(reader)
	assert.Nil(t, err)

	configuration, err := ccs.GetCustomerConfigurationById("TestCfg1")
	assert.Nil(t, err)
	assert.Equal(t, "TestName1", configuration.CustomerName)
	assert.Equal(t, true, configuration.CustomerProduction)
	assert.Equal(t, "TestKey1", configuration.CustomerConsumerKey)
	assert.Equal(t, "TestSecret1", configuration.CustomerConsumerSecret)

	configuration, err = ccs.GetCustomerConfigurationById("TestCfg2")
	assert.Nil(t, err)
	assert.Equal(t, "TestName2", configuration.CustomerName)
	assert.Equal(t, true, configuration.CustomerProduction)
	assert.Equal(t, "TestKey2", configuration.CustomerConsumerKey)
	assert.Equal(t, "TestSecret2", configuration.CustomerConsumerSecret)

	configuration, err = ccs.GetCustomerConfigurationById("TestCfg3")
	assert.Nil(t, configuration)
	assert.Error(t, err, "configuration not found")
}

func TestCustomerConfigurationStoreLoadSucceedsWithMissingFields(t *testing.T) {
	jsonData := `{
}`
	reader := strings.NewReader(jsonData)
	ccs, err := LoadCustomerConfigurationStore(reader)
	assert.Nil(t, err)
	assert.NotNil(t, ccs)
}

func TestCustomerConfigurationStoreLoadSucceedsWithExtraFields(t *testing.T) {
	jsonData := `{
  "TestCfg1": {
    "customerName": "TestName1",
    "customerProduction": true,
    "customerConsumerKey": "TestKey1",
    "customerConsumerSecret": "TestSecret1",
    "unexpectedField": "TestUnexpected1"
  },
  "TestCfg2": {
    "customerName": "TestName2",
    "customerProduction": true,
    "customerConsumerKey": "TestKey2",
    "customerConsumerSecret": "TestSecret2"
  }
}`
	reader := strings.NewReader(jsonData)
	ccs, err := LoadCustomerConfigurationStore(reader)
	assert.Nil(t, err)

	configuration, err := ccs.GetCustomerConfigurationById("TestCfg1")
	assert.Nil(t, err)
	assert.Equal(t, "TestName1", configuration.CustomerName)
	assert.Equal(t, true, configuration.CustomerProduction)
	assert.Equal(t, "TestKey1", configuration.CustomerConsumerKey)
	assert.Equal(t, "TestSecret1", configuration.CustomerConsumerSecret)

	configuration, err = ccs.GetCustomerConfigurationById("TestCfg2")
	assert.Nil(t, err)
	assert.Equal(t, "TestName2", configuration.CustomerName)
	assert.Equal(t, true, configuration.CustomerProduction)
	assert.Equal(t, "TestKey2", configuration.CustomerConsumerKey)
	assert.Equal(t, "TestSecret2", configuration.CustomerConsumerSecret)
}

func TestCustomerConfigurationStoreLoadFailsWithBadJson(t *testing.T) {
	// Malformed JSON
	jsonData := `{
  "BogusCfg1": {
    "customerName": TestName1,
    "customerProduction": 
    "customerConsumerKey":
    "customerConsumerSecret":
  }
}`
	reader := strings.NewReader(jsonData)
	ccs, err := LoadCustomerConfigurationStore(reader)
	assert.NotNil(t, err)
	assert.Nil(t, ccs)
}
