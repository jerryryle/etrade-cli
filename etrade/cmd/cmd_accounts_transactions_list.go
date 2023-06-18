package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
	"time"
)

type accountsTransactionsListFlags struct {
	startDate string
	endDate   string
}

type CommandAccountsTransactionsList struct {
	AppContext *ApplicationContext
	flags      accountsTransactionsListFlags
}

func (c *CommandAccountsTransactionsList) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [account ID]",
		Short: "List transactions",
		Long:  "List transactions for account",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			var startDate, endDate *time.Time = nil, nil
			var err error
			if c.flags.startDate != "" {
				*startDate, err = time.Parse("01022006", c.flags.startDate)
				if err != nil {
					return errors.New("start date must be in format MMDDYYYY")
				}
			}
			if c.flags.endDate != "" {
				*endDate, err = time.Parse("01022006", c.flags.startDate)
				if err != nil {
					return errors.New("end date must be in format MMDDYYYY")
				}
			}
			return c.ListTransactions(args[0], startDate, endDate)
		},
	}
	cmd.Flags().StringVarP(&c.flags.startDate, "startDate", "s", "", "start date (MMDDYYYY)")
	cmd.Flags().StringVarP(&c.flags.endDate, "endDate", "e", "", "end date (MMDDYYYY)")
	return cmd
}

func (c *CommandAccountsTransactionsList) ListTransactions(
	accountKeyId string, startDate *time.Time, endDate *time.Time,
) error {
	response, err := c.AppContext.Client.ListTransactions(
		accountKeyId,
		startDate, endDate, constants.SortOrderAsc, "", 0,
	)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
