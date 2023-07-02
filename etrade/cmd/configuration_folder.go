package cmd

import (
	"golang.org/x/exp/slog"
	"os"
	"path/filepath"
)

type ConfigurationFolder string

func NewConfigurationFolder(cfgFolder string) ConfigurationFolder {
	return ConfigurationFolder(cfgFolder)
}

func (f ConfigurationFolder) LoadCustomerConfiguration(logger *slog.Logger) (*CustomerConfigurationStore, error) {
	return LoadCustomerConfigurationStoreFromFile(f.GetConfigurationFilePath(), logger)
}

func (f ConfigurationFolder) SaveCustomerConfiguration(
	cfgStore *CustomerConfigurationStore, overwriteExisting bool, logger *slog.Logger,
) error {
	return SaveCustomerConfigurationStoreToFile(f.GetConfigurationFilePath(), overwriteExisting, cfgStore, logger)
}

func (f ConfigurationFolder) GetConfigurationFilePath() string {
	cfgFileName := ".etradecfg"
	cfgFilePath := filepath.Join(string(f), cfgFileName)
	return cfgFilePath
}

func (f ConfigurationFolder) LoadCachedCredentialsFromFile(
	customerConsumerKey string, logger *slog.Logger,
) (*CachedCredentials, error) {
	return LoadCachedCredentialsFromFile(f.GetFileCachePathForCustomer(customerConsumerKey), logger)
}

func (f ConfigurationFolder) SaveCachedCredentialsToFile(
	customerConsumerKey string, credentials *CachedCredentials, logger *slog.Logger,
) error {
	return SaveCachedCredentialsToFile(f.GetFileCachePathForCustomer(customerConsumerKey), credentials, logger)
}

func (f ConfigurationFolder) RemoveCachedCredentialsFile(customerConsumerKey string) error {
	err := os.Remove(f.GetFileCachePathForCustomer(customerConsumerKey))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func (f ConfigurationFolder) GetFileCachePathForCustomer(customerConsumerKey string) string {
	cacheFileName := "." + customerConsumerKey
	cacheFilePath := filepath.Join(string(f), ".etrade", cacheFileName)
	return cacheFilePath
}
