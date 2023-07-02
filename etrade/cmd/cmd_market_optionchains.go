package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type marketOptionChainsFlags struct {
	expiryYear, expiryMonth, expiryDay int
	strikePriceNear, noOfStrikes       int
	includeWeekly, skipAdjusted        bool
	optionCategory                     enumFlagValue[constants.OptionCategory]
	chainType                          enumFlagValue[constants.OptionChainType]
	priceType                          enumFlagValue[constants.OptionPriceType]
}

type CommandMarketOptionChains struct {
	Context *CommandContextWithClient
	flags   marketOptionChainsFlags
}

func (c *CommandMarketOptionChains) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "optionchains [symbol]",
		Short: "Get option chains",
		Long:  "Get option chains for a specific underlying instrument",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			symbol := args[0]
			if response, err := GetOptionChains(
				c.Context.Client, symbol, c.flags.expiryYear, c.flags.expiryMonth, c.flags.expiryDay,
				c.flags.strikePriceNear, c.flags.noOfStrikes, c.flags.includeWeekly, c.flags.skipAdjusted,
				c.flags.optionCategory.Value(), c.flags.chainType.Value(), c.flags.priceType.Value(),
			); err == nil {
				return c.Context.Renderer.Render(response, optionChainsDescriptor)
			} else {
				return err
			}
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
	c.flags.chainType = *newEnumFlagValue(optionChainTypeMap, constants.OptionChainTypeNil)
	c.flags.priceType = *newEnumFlagValue(optionPriceTypeMap, constants.OptionPriceTypeNil)

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

var optionChainsDescriptor = []RenderDescriptor{
	{
		ObjectPath: "",
		Values: []RenderValue{
			{Header: "Timestamp", Path: ".timeStamp", Transformer: dateTimeTransformer},
			{Header: "Quote Type", Path: ".quoteType"},
			{Header: "Near Price", Path: ".nearPrice"},
			{Header: "Selected Month", Path: ".selected.month"},
			{Header: "Selected Year", Path: ".selected.year"},
			{Header: "Selected Day", Path: ".selected.day"},
		},
		DefaultValue: "",
		SpaceAfter:   true,
	},
	{
		ObjectPath: ".optionChainPairs",
		Values: []RenderValue{
			{Header: "Call Option Category", Path: ".call.optionCategory"},
			{Header: "Call Option Root Symbol", Path: ".call.optionRootSymbol"},
			{Header: "Call Time Stamp", Path: ".call.timeStamp", Transformer: dateTimeTransformer},
			{Header: "Call Adjusted", Path: ".call.adjustedFlag"},
			{Header: "Call Display Symbol", Path: ".call.displaySymbol"},
			{Header: "Call Option Type", Path: ".call.optionType"},
			{Header: "Call Strike Price", Path: ".call.strikePrice"},
			{Header: "Call Symbol", Path: ".call.symbol"},
			{Header: "Call Bid", Path: ".call.bid"},
			{Header: "Call Ask", Path: ".call.ask"},
			{Header: "Call Bid Size", Path: ".call.bidSize"},
			{Header: "Call Ask Size", Path: ".call.askSize"},
			{Header: "Call In The Money", Path: ".call.inTheMoney"},
			{Header: "Call Volume", Path: ".call.volume"},
			{Header: "Call Open Interest", Path: ".call.openInterest"},
			{Header: "Call Net Change", Path: ".call.netChange"},
			{Header: "Call Last Price", Path: ".call.lastPrice"},
			{Header: "Call Options Symbology Initiative Key", Path: ".call.osiKey"},
			{Header: "Call Rho", Path: ".call.optionGreeks.rho"},
			{Header: "Call Vega", Path: ".call.optionGreeks.vega"},
			{Header: "Call Theta", Path: ".call.optionGreeks.theta"},
			{Header: "Call Delta", Path: ".call.optionGreeks.delta"},
			{Header: "Call Gamma", Path: ".call.optionGreeks.gamma"},
			{Header: "Call Implied Volatility", Path: ".call.optionGreeks.iv"},
			{Header: "Call Current Value", Path: ".call.optionGreeks.currentValue"},
			{Header: "Put Option Category", Path: ".put.optionCategory"},
			{Header: "Put Option Root Symbol", Path: ".put.optionRootSymbol"},
			{Header: "Put Time Stamp", Path: ".put.timeStamp", Transformer: dateTimeTransformer},
			{Header: "Put Adjusted", Path: ".put.adjustedFlag"},
			{Header: "Put Display Symbol", Path: ".put.displaySymbol"},
			{Header: "Put Option Type", Path: ".put.optionType"},
			{Header: "Put Strike Price", Path: ".put.strikePrice"},
			{Header: "Put Symbol", Path: ".put.symbol"},
			{Header: "Put Bid", Path: ".put.bid"},
			{Header: "Put Ask", Path: ".put.ask"},
			{Header: "Put Bid Size", Path: ".put.bidSize"},
			{Header: "Put Ask Size", Path: ".put.askSize"},
			{Header: "Put In The Money", Path: ".put.inTheMoney"},
			{Header: "Put Volume", Path: ".put.volume"},
			{Header: "Put Open Interest", Path: ".put.openInterest"},
			{Header: "Put Net Change", Path: ".put.netChange"},
			{Header: "Put Last Price", Path: ".put.lastPrice"},
			{Header: "Put Options Symbology Initiative Key", Path: ".put.osiKey"},
			{Header: "Put Rho", Path: ".put.optionGreeks.rho"},
			{Header: "Put Vega", Path: ".put.optionGreeks.vega"},
			{Header: "Put Theta", Path: ".put.optionGreeks.theta"},
			{Header: "Put Delta", Path: ".put.optionGreeks.delta"},
			{Header: "Put Gamma", Path: ".put.optionGreeks.gamma"},
			{Header: "Put Implied Volatility", Path: ".put.optionGreeks.iv"},
			{Header: "Put Current Value", Path: ".put.optionGreeks.currentValue"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
