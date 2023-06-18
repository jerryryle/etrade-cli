package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type marketOptionExpireFlags struct {
	expiryType expiryType
}

type CommandMarketOptionexpire struct {
	AppContext *ApplicationContext
	flags      marketOptionExpireFlags
}

func (c *CommandMarketOptionexpire) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "optionexpire [symbol]",
		Short: "Get option expire dates",
		Long:  "Get option expire dates for a specific underlying instrument",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.GetOptionExpireDates(args[0])
		},
	}
	cmd.Flags().VarP(
		&c.flags.expiryType, "expiryType", "e", fmt.Sprintf(
			"expiry type (%s, %s, %s, %s, %s, %s, %s, %s)", expiryTypeUnspecified, expiryTypeDaily, expiryTypeWeekly,
			expiryTypeMonthly, expiryTypeQuarterly, expiryTypeVix, expiryTypeAll, expiryTypeMonthEnd,
		),
	)
	return cmd
}

func (c *CommandMarketOptionexpire) GetOptionExpireDates(symbol string) error {
	response, err := c.AppContext.Client.GetOptionExpireDates(symbol, c.flags.expiryType.ExpiryType())
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}

type expiryType string

const (
	expiryTypeUnspecified expiryType = "unspecified"
	expiryTypeDaily                  = "daily"
	expiryTypeWeekly                 = "weekly"
	expiryTypeMonthly                = "monthly"
	expiryTypeQuarterly              = "quarterly"
	expiryTypeVix                    = "vix"
	expiryTypeAll                    = "all"
	expiryTypeMonthEnd               = "monthEnd"
)

func (e *expiryType) String() string {
	return string(*e)
}

func (e *expiryType) Set(v string) error {
	switch expiryType(v) {
	case expiryTypeUnspecified, expiryTypeDaily, expiryTypeWeekly, expiryTypeMonthly, expiryTypeQuarterly, expiryTypeVix, expiryTypeAll, expiryTypeMonthEnd:
		*e = expiryType(v)
		return nil
	default:
		return fmt.Errorf(
			"must be %s, %s, %s, %s, %s, %s, %s, or %s", expiryTypeUnspecified, expiryTypeDaily, expiryTypeWeekly,
			expiryTypeMonthly, expiryTypeQuarterly, expiryTypeVix, expiryTypeAll, expiryTypeMonthEnd,
		)
	}
}

func (e *expiryType) Type() string {
	return "expiryType"
}

func (e *expiryType) ExpiryType() constants.OptionExpiryType {
	switch *e {
	case expiryTypeUnspecified:
		return constants.OptionExpiryTypeUnspecified
	case expiryTypeDaily:
		return constants.OptionExpiryTypeDaily
	case expiryTypeWeekly:
		return constants.OptionExpiryTypeWeekly
	case expiryTypeMonthly:
		return constants.OptionExpiryTypeMonthly
	case expiryTypeQuarterly:
		return constants.OptionExpiryTypeQuarterly
	case expiryTypeVix:
		return constants.OptionExpiryTypeVix
	case expiryTypeAll:
		return constants.OptionExpiryTypeAll
	case expiryTypeMonthEnd:
		return constants.OptionExpiryTypeMonthEnd
	}
	return constants.OptionExpiryTypeAll
}
