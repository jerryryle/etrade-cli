package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
)

func ClearAuth(customerId string, cfgFolder ConfigurationFolder, cfgStore *CustomerConfigurationStore) (
	jsonmap.JsonMap, error,
) {
	customerConfig, err := cfgStore.GetCustomerConfigurationById(customerId)
	if err != nil {
		return nil, fmt.Errorf("customer id '%s' not found in config file", customerId)
	}
	err = cfgFolder.RemoveCachedCredentialsFile(customerConfig.CustomerConsumerKey)
	if err != nil {
		return nil, fmt.Errorf("unable to remove auth cache for %s (%w)", customerId, err)
	}
	return jsonmap.JsonMap{
		"status": "success",
	}, nil
}
