package cmd

import "github.com/jerryryle/etrade-cli/pkg/etradelib"

type ApplicationContext struct {
	Client   etradelib.ETradeClient
	Customer etradelib.ETradeCustomer
}
