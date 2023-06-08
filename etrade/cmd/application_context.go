package cmd

import "github.com/jerryryle/etrade-cli/pkg/etradelib"

type ApplicationContext struct {
	Customer etradelib.ETradeCustomer
}
