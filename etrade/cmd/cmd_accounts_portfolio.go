package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type CommandAccountsPortfolio struct {
	Resources *CommandResources
}

func (c *CommandAccountsPortfolio) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "portfolio [account ID]",
		Short: "View Portfolio",
		Long:  "View Portfolio",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ViewPortfolio(args[0])
		},
	}
	return cmd
}

func (c *CommandAccountsPortfolio) ViewPortfolio(accountKeyId string) error {
	response, err := c.Resources.Client.ViewPortfolio(
		accountKeyId, 0, constants.PortfolioSortBySymbol, constants.SortOrderAsc, 0,
		constants.MarketSessionRegular,
		true, true, constants.PortfolioViewComplete,
	)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
