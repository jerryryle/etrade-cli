package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type getQuotesCommandFlags struct {
	requireEarningsDate, skipMiniOptionsCheck bool
}

type GetQuotesCommand struct {
	AppContext *ApplicationContext
	flags      getQuotesCommandFlags
}

func (c *GetQuotesCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getquotes [symbol] ...",
		Short: "Get quotes",
		Long:  "Get quotes for one or more symbols",
		Args:  cobra.MatchAll(cobra.MinimumNArgs(50)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.GetQuotes(args)
		},
	}
	cmd.Flags().BoolVarP(
		&c.flags.requireEarningsDate, "requireEarningsDate", "r", true, "include next earning date in output",
	)
	cmd.Flags().BoolVarP(
		&c.flags.skipMiniOptionsCheck, "skipMiniOptionsCheck", "s", false,
		"skip checking whether the symbol has mini options",
	)

	return cmd
}

func (c *GetQuotesCommand) GetQuotes(symbols []string) error {
	response, err := c.AppContext.Client.GetQuotes(
		symbols, constants.QuoteDetailAll, c.flags.requireEarningsDate, c.flags.skipMiniOptionsCheck,
	)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
