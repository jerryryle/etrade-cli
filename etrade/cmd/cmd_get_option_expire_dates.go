package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/spf13/cobra"
)

type getOptionExpireDatesFlags struct {
	expiryType expiryType
}

type GetOptionExpireDatesCommand struct {
	AppContext *ApplicationContext
	flags      getOptionExpireDatesFlags
}

func (c *GetOptionExpireDatesCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getoptionexpiredates [symbol]",
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

func (c *GetOptionExpireDatesCommand) GetOptionExpireDates(symbol string) error {
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
		return errors.New(
			fmt.Sprintf(
				"must be %s, %s, %s, %s, %s, %s, %s, or %s", expiryTypeUnspecified, expiryTypeDaily, expiryTypeWeekly,
				expiryTypeMonthly, expiryTypeQuarterly, expiryTypeVix, expiryTypeAll, expiryTypeMonthEnd,
			),
		)
	}
}

func (e *expiryType) Type() string {
	return "expiryType"
}

func (e *expiryType) ExpiryType() client.ExpiryType {
	switch *e {
	case expiryTypeUnspecified:
		return client.ExpiryTypeUnspecified
	case expiryTypeDaily:
		return client.ExpiryTypeDaily
	case expiryTypeWeekly:
		return client.ExpiryTypeWeekly
	case expiryTypeMonthly:
		return client.ExpiryTypeMonthly
	case expiryTypeQuarterly:
		return client.ExpiryTypeQuarterly
	case expiryTypeVix:
		return client.ExpiryTypeVix
	case expiryTypeAll:
		return client.ExpiryTypeAll
	case expiryTypeMonthEnd:
		return client.ExpiryTypeMonthEnd
	}
	return client.ExpiryTypeAll
}
