package cmd

import (
	"encoding/json"
	"errors"
	"golang.org/x/exp/slog"
	"io"
	"os"
)

type CustomerConfiguration struct {
	CustomerName           string `json:"customerName"`
	CustomerProduction     bool   `json:"customerProduction"`
	CustomerConsumerKey    string `json:"customerConsumerKey"`
	CustomerConsumerSecret string `json:"customerConsumerSecret"`
}

type CustomerConfigurationStore struct {
	customerConfigMap map[string]CustomerConfiguration
}

func LoadCustomerConfigurationStore(reader io.Reader) (*CustomerConfigurationStore, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var cc = CustomerConfigurationStore{}
	if err := json.Unmarshal(bytes, &cc.customerConfigMap); err != nil {
		return nil, err
	}
	return &cc, nil
}

func LoadCustomerConfigurationStoreFromFile(filename string, logger *slog.Logger) (
	*CustomerConfigurationStore, error,
) {
	file, err := os.Open(filename)
	if file != nil {
		defer func(file *os.File) {
			err := file.Close()
			if err != nil && logger != nil {
				logger.Error(err.Error())
			}
		}(file)
	}
	if err != nil {
		return nil, err
	}
	return LoadCustomerConfigurationStore(file)
}

func SaveCustomerConfigurationStore(writer io.Writer, cc *CustomerConfigurationStore) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(&cc.customerConfigMap); err != nil {
		return err
	}
	return nil
}

func SaveCustomerConfigurationStoreToFile(
	filename string, overwriteExisting bool, cc *CustomerConfigurationStore, logger *slog.Logger,
) error {
	openFlags := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	if !overwriteExisting {
		openFlags |= os.O_EXCL
	}
	file, err := os.OpenFile(filename, openFlags, 0666)
	if file != nil {
		defer func(file *os.File) {
			err := file.Close()
			if err != nil && logger != nil {
				logger.Error(err.Error())
			}
		}(file)
	}
	if err != nil {
		return err
	}
	return SaveCustomerConfigurationStore(file, cc)
}

func (c *CustomerConfigurationStore) GetCustomerConfigurationById(configId string) (*CustomerConfiguration, error) {
	configItem, exists := c.customerConfigMap[configId]
	if !exists {
		return nil, errors.New("configuration not found")
	}
	return &configItem, nil
}

func (c *CustomerConfigurationStore) SetCustomerConfigurationForId(
	configId string, configuration *CustomerConfiguration,
) {
	c.customerConfigMap[configId] = *configuration
}

func (c *CustomerConfigurationStore) GetAllConfigurations() map[string]CustomerConfiguration {
	return c.customerConfigMap
}
