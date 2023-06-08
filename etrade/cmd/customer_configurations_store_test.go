package cmd

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCustomerConfigurationsStoreLoadSucceedsWithGoodJson(t *testing.T) {
	jsonData := `{
  "TestCfg1": {
    "customerNickname": "TestName1",
    "customerProduction": true,
    "customerConsumerKey": "TestKey1",
    "customerConsumerSecret": "TestSecret1"
  },
  "TestCfg2": {
    "customerNickname": "TestName2",
    "customerProduction": true,
    "customerConsumerKey": "TestKey2",
    "customerConsumerSecret": "TestSecret2"
  }
}`
	reader := strings.NewReader(jsonData)
	ccs, err := LoadCustomerConfigurationsStore(reader)
	assert.Nil(t, err)

	configuration, err := ccs.GetCustomerConfigurationById("TestCfg1")
	assert.Nil(t, err)
	assert.Equal(t, "TestName1", configuration.CustomerNickname)
	assert.Equal(t, true, configuration.CustomerProduction)
	assert.Equal(t, "TestKey1", configuration.CustomerConsumerKey)
	assert.Equal(t, "TestSecret1", configuration.CustomerConsumerSecret)

	configuration, err = ccs.GetCustomerConfigurationById("TestCfg2")
	assert.Nil(t, err)
	assert.Equal(t, "TestName2", configuration.CustomerNickname)
	assert.Equal(t, true, configuration.CustomerProduction)
	assert.Equal(t, "TestKey2", configuration.CustomerConsumerKey)
	assert.Equal(t, "TestSecret2", configuration.CustomerConsumerSecret)

	configuration, err = ccs.GetCustomerConfigurationById("TestCfg3")
	assert.Nil(t, configuration)
	assert.Error(t, err, "configuration not found")
}

func TestCustomerConfigurationsStoreLoadSucceedsWithMissingFields(t *testing.T) {
	jsonData := `{
}`
	reader := strings.NewReader(jsonData)
	ccs, err := LoadCustomerConfigurationsStore(reader)
	assert.Nil(t, err)
	assert.NotNil(t, ccs)
}

func TestCustomerConfigurationsStoreLoadSucceedsWithExtraFields(t *testing.T) {
	jsonData := `{
  "TestCfg1": {
    "customerNickname": "TestName1",
    "customerProduction": true,
    "customerConsumerKey": "TestKey1",
    "customerConsumerSecret": "TestSecret1",
    "unexpectedField": "TestUnexpected1"
  },
  "TestCfg2": {
    "customerNickname": "TestName2",
    "customerProduction": true,
    "customerConsumerKey": "TestKey2",
    "customerConsumerSecret": "TestSecret2"
  }
}`
	reader := strings.NewReader(jsonData)
	ccs, err := LoadCustomerConfigurationsStore(reader)
	assert.Nil(t, err)

	configuration, err := ccs.GetCustomerConfigurationById("TestCfg1")
	assert.Nil(t, err)
	assert.Equal(t, "TestName1", configuration.CustomerNickname)
	assert.Equal(t, true, configuration.CustomerProduction)
	assert.Equal(t, "TestKey1", configuration.CustomerConsumerKey)
	assert.Equal(t, "TestSecret1", configuration.CustomerConsumerSecret)

	configuration, err = ccs.GetCustomerConfigurationById("TestCfg2")
	assert.Nil(t, err)
	assert.Equal(t, "TestName2", configuration.CustomerNickname)
	assert.Equal(t, true, configuration.CustomerProduction)
	assert.Equal(t, "TestKey2", configuration.CustomerConsumerKey)
	assert.Equal(t, "TestSecret2", configuration.CustomerConsumerSecret)
}

func TestCustomerConfigurationsStoreLoadFailsWithBadJson(t *testing.T) {
	// Malformed JSON
	jsonData := `{
  "BogusCfg1": {
    "customerNickname": TestName1,
    "customerProduction": 
    "customerConsumerKey":
    "customerConsumerSecret":
  }
}`
	reader := strings.NewReader(jsonData)
	ccs, err := LoadCustomerConfigurationsStore(reader)
	assert.NotNil(t, err)
	assert.Nil(t, ccs)
}
