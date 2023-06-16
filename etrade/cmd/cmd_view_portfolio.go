package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/spf13/cobra"
)

type ViewPortfolioCommand struct {
	AppContext *ApplicationContext
}

func (c *ViewPortfolioCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "viewportfolio [account ID]",
		Short: "View Portfolio",
		Long:  "View Portfolio",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ViewPortfolio(args[0])
		},
	}
	return cmd
}

func (c *ViewPortfolioCommand) ViewPortfolio(accountKeyId string) error {
	response, err := c.AppContext.Client.ViewPortfolio(
		accountKeyId, 0, client.PortfolioSortBySymbol, client.SortOrderAsc, 0, client.PortfolioMarketSessionRegular,
		true, true, client.PortfolioViewComplete,
	)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
