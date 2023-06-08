package cmd

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type CustomerConfiguration struct {
	CustomerName           string `json:"customerName"`
	CustomerProduction     bool   `json:"customerProduction"`
	CustomerConsumerKey    string `json:"customerConsumerKey"`
	CustomerConsumerSecret string `json:"customerConsumerSecret"`
}

type CustomerConfigurationsStore struct {
	customerConfigMap map[string]CustomerConfiguration
}

func CreateCustomerConfigurationsStore() *CustomerConfigurationsStore {
	return &CustomerConfigurationsStore{make(map[string]CustomerConfiguration)}
}

func LoadCustomerConfigurationsStore(reader io.Reader) (*CustomerConfigurationsStore, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var cc = CustomerConfigurationsStore{}
	if err := json.Unmarshal(bytes, &cc.customerConfigMap); err != nil {
		return nil, err
	}
	return &cc, nil
}

func LoadCustomerConfigurationsStoreFromFile(filename string) (*CustomerConfigurationsStore, error) {
	file, err := os.Open(filename)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		return nil, err
	}
	return LoadCustomerConfigurationsStore(file)
}

func SaveCustomerConfigurationsStore(writer io.Writer, cc *CustomerConfigurationsStore) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(&cc.customerConfigMap); err != nil {
		return err
	}
	return nil
}

func SaveCustomerConfigurationsStoreToFile(filename string, cc *CustomerConfigurationsStore) error {
	file, err := os.Create(filename)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		return err
	}
	return SaveCustomerConfigurationsStore(file, cc)
}

func (c *CustomerConfigurationsStore) GetCustomerConfigurationById(configId string) (*CustomerConfiguration, error) {
	configItem, exists := c.customerConfigMap[configId]
	if !exists {
		return nil, errors.New("configuration not found")
	}
	return &configItem, nil
}

func (c *CustomerConfigurationsStore) SetCustomerConfigurationForId(configId string, configuration *CustomerConfiguration) {
	c.customerConfigMap[configId] = *configuration
}
