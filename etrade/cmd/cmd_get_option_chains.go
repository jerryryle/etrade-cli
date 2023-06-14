package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/spf13/cobra"
)

type GetOptionChainsCommandFlags struct {
	expiryYear, expiryMonth, expiryDay int
	strikePriceNear, noOfStrikes       int
	includeWeekly, skipAdjusted        bool
	optionCategory                     optionCategory
	chainType                          chainType
	priceType                          priceType
}

type GetOptionChainsCommand struct {
	AppContext *ApplicationContext
	flags      GetOptionChainsCommandFlags
}

func (c *GetOptionChainsCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getoptionchains [symbol]",
		Short: "Get option chains",
		Long:  "Get option chains for a specific underlying instrument",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.GetOptionChains(args[0])
		},
	}
	cmd.Flags().IntVarP(&c.flags.expiryYear, "expiryYear", "y", -1, "expiration year")
	cmd.Flags().IntVarP(&c.flags.expiryMonth, "expiryMonth", "m", -1, "expiration month")
	cmd.Flags().IntVarP(&c.flags.expiryDay, "expiryDay", "d", -1, "expiration day")
	cmd.Flags().IntVarP(&c.flags.strikePriceNear, "strikePriceNear", "s", -1, "strike price near")
	cmd.Flags().IntVarP(&c.flags.noOfStrikes, "noOfStrikes", "n", -1, "number of strikes")
	cmd.Flags().BoolVarP(&c.flags.includeWeekly, "includeWeekly", "w", false, "include weekly options")
	cmd.Flags().BoolVarP(&c.flags.includeWeekly, "skipAdjusted", "a", true, "skip adjusted")
	cmd.Flags().VarP(&c.flags.optionCategory, "optionCategory", "c", fmt.Sprintf("option category (%s, %s, %s)", optionCategoryStandard, optionCategoryAll, optionCategoryMini))
	cmd.Flags().VarP(&c.flags.chainType, "chainType", "t", fmt.Sprintf("chain type (%s, %s, %s)", chainTypeCall, chainTypePut, chainTypeCallPut))
	cmd.Flags().VarP(&c.flags.priceType, "priceType", "p", fmt.Sprintf("price type (%s, %s)", priceTypeAtnm, priceTypeAll))
	return cmd
}

func (c *GetOptionChainsCommand) GetOptionChains(symbol string) error {
	result, err := c.AppContext.Client.GetOptionChains(symbol,
		c.flags.expiryYear, c.flags.expiryMonth, c.flags.expiryDay,
		c.flags.strikePriceNear, c.flags.noOfStrikes, c.flags.includeWeekly, c.flags.skipAdjusted,
		c.flags.optionCategory.OptionCategory(), c.flags.chainType.ChainType(), c.flags.priceType.PriceType())
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", result)
	return nil
}

type optionCategory string

const (
	optionCategoryStandard optionCategory = "standard"
	optionCategoryAll      optionCategory = "all"
	optionCategoryMini     optionCategory = "mini"
)

func (e *optionCategory) String() string {
	return string(*e)
}

func (e *optionCategory) Set(v string) error {
	switch optionCategory(v) {
	case optionCategoryStandard, optionCategoryAll, optionCategoryMini:
		*e = optionCategory(v)
		return nil
	default:
		return errors.New(fmt.Sprintf("must be %s, %s, or %s", optionCategoryStandard, optionCategoryAll, optionCategoryMini))
	}
}

func (e *optionCategory) Type() string {
	return "optionCategory"
}

func (e *optionCategory) OptionCategory() client.OptionCategory {
	switch *e {
	case optionCategoryStandard:
		return client.OptionCategoryStandard
	case optionCategoryAll:
		return client.OptionCategoryAll
	case optionCategoryMini:
		return client.OptionCategoryMini
	}
	return client.OptionCategoryAll
}

type chainType string

const (
	chainTypeCall    chainType = "call"
	chainTypePut     chainType = "put"
	chainTypeCallPut chainType = "callput"
)

func (e *chainType) String() string {
	return string(*e)
}

func (e *chainType) Set(v string) error {
	switch chainType(v) {
	case chainTypeCall, chainTypePut, chainTypeCallPut:
		*e = chainType(v)
		return nil
	default:
		return errors.New(fmt.Sprintf("must be %s, %s, or %s", chainTypeCall, chainTypePut, chainTypeCallPut))
	}
}

func (e *chainType) Type() string {
	return "chainType"
}

func (e *chainType) ChainType() client.ChainType {
	switch *e {
	case chainTypeCall:
		return client.ChainTypeCall
	case chainTypePut:
		return client.ChainTypePut
	case chainTypeCallPut:
		return client.ChainTypeCallPut
	}
	return client.ChainTypeCallPut
}

type priceType string

const (
	priceTypeAtnm priceType = "atnm"
	priceTypeAll  priceType = "all"
)

func (e *priceType) String() string {
	return string(*e)
}

func (e *priceType) Set(v string) error {
	switch priceType(v) {
	case priceTypeAtnm, priceTypeAll:
		*e = priceType(v)
		return nil
	default:
		return errors.New(fmt.Sprintf("must be %s or %s", priceTypeAtnm, priceTypeAll))
	}
}

func (e *priceType) Type() string {
	return "priceType"
}

func (e *priceType) PriceType() client.PriceType {
	switch *e {
	case priceTypeAtnm:
		return client.PriceTypeAtnm
	case priceTypeAll:
		return client.PriceTypeAll
	}
	return client.PriceTypeAll
}