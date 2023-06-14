package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"golang.org/x/exp/slog"
)

type ApplicationContext struct {
	Logger *slog.Logger
	Client client.ETradeClient
}
