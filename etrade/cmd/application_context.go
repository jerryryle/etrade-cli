package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"golang.org/x/exp/slog"
)

type ApplicationContext struct {
	Logger   *slog.Logger
	Client   etradelib.ETradeClient
	Customer etradelib.ETradeCustomer
}
