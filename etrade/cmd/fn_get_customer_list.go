package cmd

import "github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"

func GetCustomerList(cfgStore *CustomerConfigurationStore) jsonmap.JsonMap {
	customerSlice := jsonmap.JsonSlice{}
	for customerId, customerConfig := range cfgStore.GetAllConfigurations() {
		customerMap := jsonmap.JsonMap{}
		customerMap.SetString("customerId", customerId)
		customerMap.SetString("customerName", customerConfig.CustomerName)
		customerMap.SetBool("productionAccess", customerConfig.CustomerProduction)
		customerSlice = append(customerSlice, customerMap)
	}
	return jsonmap.JsonMap{
		"customers": customerSlice,
	}
}
