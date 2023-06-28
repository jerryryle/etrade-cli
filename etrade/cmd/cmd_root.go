package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	globalFlags globalFlags
}

func (c *RootCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "etrade",
		Short: "E*TRADE CLI",
		Long:  "E*TRADE Command Line Interface",
	}
	// Add Global Flags
	cmd.PersistentFlags().StringVar(&c.globalFlags.customerId, "customer-id", "", "customer identifier")
	cmd.PersistentFlags().BoolVar(&c.globalFlags.debug, "debug", false, "debug output")
	cmd.PersistentFlags().StringVar(
		&c.globalFlags.outputFileName, "output-file", "", "write output to specified file instead of stdout",
	)

	// Initialize Global Enum Flag Values
	c.globalFlags.outputFormat = *newEnumFlagValue(outputFormatMap, outputFormatCsv)

	// Add Global Enum Flags
	cmd.PersistentFlags().Var(
		&c.globalFlags.outputFormat, "format",
		fmt.Sprintf("output format (%s)", c.globalFlags.outputFormat.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"format",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.globalFlags.outputFormat.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	// Add Subcommands
	cmd.AddCommand((&CommandAccounts{}).Command(&c.globalFlags))
	cmd.AddCommand((&CommandAlerts{}).Command(&c.globalFlags))
	cmd.AddCommand((&CommandMarket{}).Command(&c.globalFlags))
	cmd.AddCommand((&CommandOrders{}).Command(&c.globalFlags))
	cmd.AddCommand((&CommandAuth{}).Command(&c.globalFlags))
	cmd.AddCommand((&CommandCfg{}).Command(&c.globalFlags))

	return cmd
}

var outputFormatMap = map[string]enumValueWithHelp[outputFormat]{
	"csv":        {outputFormatCsv, "CSV output"},
	"json":       {outputFormatJson, "raw JSON output"},
	"jsonPretty": {outputFormatJsonPretty, "formatted JSON output"},
}
