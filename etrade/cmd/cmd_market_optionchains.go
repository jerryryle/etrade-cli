package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type marketOptionchainsFlags struct {
	expiryYear, expiryMonth, expiryDay int
	strikePriceNear, noOfStrikes       int
	includeWeekly, skipAdjusted        bool
	optionCategory                     enumFlagValue[constants.OptionCategory]
	chainType                          enumFlagValue[constants.OptionChainType]
	priceType                          enumFlagValue[constants.OptionPriceType]
}

type CommandMarketOptionchains struct {
	Resources *CommandResources
	flags     marketOptionchainsFlags
}

func (c *CommandMarketOptionchains) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "optionchains [symbol]",
		Short: "Get option chains",
		Long:  "Get option chains for a specific underlying instrument",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.GetOptionChains(args[0])
		},
	}

	// Add Flags
	cmd.Flags().IntVarP(&c.flags.expiryYear, "expiry-year", "y", -1, "expiration year")
	cmd.Flags().IntVarP(&c.flags.expiryMonth, "expiry-month", "m", -1, "expiration month")
	cmd.Flags().IntVarP(&c.flags.expiryDay, "expiry-day", "d", -1, "expiration day")
	cmd.Flags().IntVarP(&c.flags.strikePriceNear, "strike-price-near", "s", -1, "strike price near")
	cmd.Flags().IntVarP(&c.flags.noOfStrikes, "strikes", "n", -1, "number of strikes")
	cmd.Flags().BoolVarP(&c.flags.includeWeekly, "include-weekly", "w", false, "include weekly options")
	cmd.Flags().BoolVarP(&c.flags.skipAdjusted, "skip-adjusted", "a", true, "skip adjusted")

	// Initialize Enum Flag Values
	c.flags.optionCategory = *newEnumFlagValue(optionCategoryMap, constants.OptionCategoryNil)
	c.flags.chainType = *newEnumFlagValue(chainTypeMap, constants.OptionChainTypeNil)
	c.flags.priceType = *newEnumFlagValue(priceTypeMap, constants.OptionPriceTypeNil)

	// Add Enum Flags
	cmd.Flags().VarP(
		&c.flags.optionCategory, "category", "c",
		fmt.Sprintf("option category (%s)", c.flags.optionCategory.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"category",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.optionCategory.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.chainType, "chain-type", "t",
		fmt.Sprintf("chain type (%s)", c.flags.chainType.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"chain-type",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.chainType.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.priceType, "price-type", "p",
		fmt.Sprintf("price type (%s)", c.flags.priceType.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"price-type",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.priceType.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	return cmd
}

func (c *CommandMarketOptionchains) GetOptionChains(symbol string) error {
	response, err := c.Resources.Client.GetOptionChains(
		symbol,
		c.flags.expiryYear, c.flags.expiryMonth, c.flags.expiryDay,
		c.flags.strikePriceNear, c.flags.noOfStrikes, c.flags.includeWeekly, c.flags.skipAdjusted,
		c.flags.optionCategory.Value(), c.flags.chainType.Value(), c.flags.priceType.Value(),
	)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(c.Resources.OFile, string(response))

	return nil
}

var optionCategoryMap = map[string]enumValueWithHelp[constants.OptionCategory]{
	"standard": {constants.OptionCategoryStandard, "only standard options"},
	"all":      {constants.OptionCategoryAll, "all options"},
	"mini":     {constants.OptionCategoryMini, "only mini options"},
}

var chainTypeMap = map[string]enumValueWithHelp[constants.OptionChainType]{
	"call":    {constants.OptionChainTypeCall, "only call options"},
	"put":     {constants.OptionChainTypePut, "only put options"},
	"callPut": {constants.OptionChainTypeCallPut, "call and put options"},
}

var priceTypeMap = map[string]enumValueWithHelp[constants.OptionPriceType]{
	"extendedHours": {constants.OptionPriceTypeExtendedHours, "only extended hours price types"},
	"all":           {constants.OptionPriceTypeAll, "all price types"},
}
