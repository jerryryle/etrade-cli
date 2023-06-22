package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type marketQuotesFlags struct {
	requireEarningsDate, skipMiniOptionsCheck bool
}

type CommandMarketQuote struct {
	Context *CommandContext
	flags   marketQuotesFlags
}

func (c *CommandMarketQuote) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "quote [symbol] ...",
		Short: "Get quotes",
		Long:  "Get quotes for one or more symbols",
		Args:  cobra.MatchAll(cobra.MaximumNArgs(50)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.GetQuotes(args)
		},
	}
	cmd.Flags().BoolVarP(
		&c.flags.requireEarningsDate, "require-earnings-date", "r", true, "include next earning date in output",
	)
	cmd.Flags().BoolVarP(
		&c.flags.skipMiniOptionsCheck, "skip-mini-check", "s", false,
		"skip the check for whether the symbol has mini options",
	)

	return cmd
}

func (c *CommandMarketQuote) GetQuotes(symbols []string) error {
	response, err := c.Context.Client.GetQuotes(
		symbols, constants.QuoteDetailAll, c.flags.requireEarningsDate, c.flags.skipMiniOptionsCheck,
	)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(c.Context.OutputFile, string(response))
	return nil
}
