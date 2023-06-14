package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/spf13/cobra"
	"time"
)

type listTransactionsFlags struct {
	startDate string
	endDate   string
}

type ListTransactionsCommand struct {
	AppContext *ApplicationContext
	flags      listTransactionsFlags
}

func (c *ListTransactionsCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listtransactions [account ID]",
		Short: "List transactions",
		Long:  "List transactions",
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

func (c *ListTransactionsCommand) ListTransactions(accountKeyId string, startDate *time.Time, endDate *time.Time) error {
	response, err := c.AppContext.Client.ListTransactions(accountKeyId,
		startDate, endDate, client.TransactionSortOrderAsc, "", 0)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", response)
	return nil
}
