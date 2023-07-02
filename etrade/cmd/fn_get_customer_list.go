package cmd

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

func GetCustomerList(cfgStore *CustomerConfigurationStore) (jsonmap.JsonMap, error) {
	customerSlice := jsonmap.JsonSlice{}
	for customerId, customerConfig := range cfgStore.GetAllConfigurations() {
		customerMap := jsonmap.JsonMap{}
		if err := customerMap.SetString("customerId", customerId); err != nil {
			return nil, err
		}
		if err := customerMap.SetString("customerName", customerConfig.CustomerName); err != nil {
			return nil, err
		}
		if err := customerMap.SetBool("productionAccess", customerConfig.CustomerProduction); err != nil {
			return nil, err
		}
		customerSlice = append(customerSlice, customerMap)
	}
	return jsonmap.JsonMap{
		"customers": customerSlice,
	}, nil
}
