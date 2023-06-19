package cmd

import (
	"os"
	"path/filepath"
)

func getUserHomeFolder() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir, nil
}

func getCfgFilePath(baseFolder string) string {
	cfgFileName := ".etradecfg"
	cfgFilePath := filepath.Join(baseFolder, cfgFileName)
	return cfgFilePath
}

func getFileCachePathForCustomer(baseFolder string, customerConsumerKey string) string {
	cacheFileName := "." + customerConsumerKey
	cacheFilePath := filepath.Join(baseFolder, ".etrade", cacheFileName)
	return cacheFilePath
}
