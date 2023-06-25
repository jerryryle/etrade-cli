package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	globalFlags GlobalFlags
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
	c.globalFlags.outputFormat = *newEnumFlagValue(outputFormatMap, OutputFormatCsv)

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
	cmd.AddCommand((&CommandAccounts{GlobalFlags: &c.globalFlags}).Command())
	cmd.AddCommand((&CommandAlerts{GlobalFlags: &c.globalFlags}).Command())
	cmd.AddCommand((&CommandMarket{GlobalFlags: &c.globalFlags}).Command())
	cmd.AddCommand((&CommandOrders{GlobalFlags: &c.globalFlags}).Command())
	cmd.AddCommand((&CommandCfg{}).Command())

	return cmd
}

var outputFormatMap = map[string]enumValueWithHelp[OutputFormat]{
	"csv":        {OutputFormatCsv, "CSV output"},
	"json":       {OutputFormatJson, "raw JSON output"},
	"jsonPretty": {OutputFormatJsonPretty, "formatted JSON output"},
}
