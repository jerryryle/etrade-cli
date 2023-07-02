package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type marketOptionExpireFlags struct {
	expiryType enumFlagValue[constants.OptionExpiryType]
}

type CommandMarketOptionExpire struct {
	Context *CommandContextWithClient
	flags   marketOptionExpireFlags
}

func (c *CommandMarketOptionExpire) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "optionexpire [symbol]",
		Short: "Get option expire dates",
		Long:  "Get option expire dates for a specific underlying instrument",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			symbol := args[0]
			if response, err := GetOptionExpireDates(c.Context.Client, symbol, c.flags.expiryType.Value()); err == nil {
				return c.Context.Renderer.Render(response, optionExpireDatesDescriptor)
			} else {
				return err
			}
		},
	}

	// Initialize Enum Flag Values
	c.flags.expiryType = *newEnumFlagValue(optionExpiryTypeMap, constants.OptionExpiryTypeNil)

	// Add Enum Flags
	cmd.Flags().VarP(
		&c.flags.expiryType, "expiry-type", "e",
		fmt.Sprintf("expiry type (%s)", c.flags.expiryType.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"expiry-type",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.expiryType.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	return cmd
}

var optionExpireDatesDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".optionExpireDates",
		Values: []RenderValue{
			{Header: "Expiry Year", Path: ".year"},
			{Header: "Expiry Month", Path: ".month"},
			{Header: "Expiry Day", Path: ".day"},
			{Header: "Expiry Type", Path: ".expiryType"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
