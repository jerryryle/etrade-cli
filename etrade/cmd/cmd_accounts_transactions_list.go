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
	sortOrder enumFlagValue[constants.SortOrder]
}

type CommandAccountsTransactionsList struct {
	Context *CommandContextWithClient
	flags   accountsTransactionsListFlags
}

func (c *CommandAccountsTransactionsList) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list [account ID]",
		Short: "List transactions",
		Long:  "List transactions for account",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			accountId := args[0]
			var startDate, endDate *time.Time = nil, nil
			if c.flags.startDate != "" {
				var err error
				*startDate, err = time.Parse("01022006", c.flags.startDate)
				if err != nil {
					return errors.New("start date must be in format MMDDYYYY")
				}
			}
			if c.flags.endDate != "" {
				var err error
				*endDate, err = time.Parse("01022006", c.flags.startDate)
				if err != nil {
					return errors.New("end date must be in format MMDDYYYY")
				}
			}
			if response, err := ListTransactions(
				c.Context.Client, accountId, startDate, endDate, c.flags.sortOrder.Value(),
			); err == nil {
				return c.Context.Renderer.Render(response, transactionListDescriptor)
			} else {
				return err
			}
		},
	}

	// Add Flags
	cmd.Flags().StringVarP(&c.flags.startDate, "start-date", "s", "", "start date (MMDDYYYY)")
	cmd.Flags().StringVarP(&c.flags.endDate, "end-date", "e", "", "end date (MMDDYYYY)")

	// Initialize Enum Flag Values
	c.flags.sortOrder = *newEnumFlagValue(sortOrderMap, constants.SortOrderNil)

	// Add Enum Flags
	cmd.Flags().VarP(
		&c.flags.sortOrder, "sort-order", "o",
		fmt.Sprintf("sort order (%s)", c.flags.sortOrder.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"sort-order",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.sortOrder.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	return cmd
}

var transactionListDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".transactions",
		Values: []RenderValue{
			{Header: "Transaction ID", Path: ".transactionId"},
			{Header: "Account ID", Path: ".accountId"},
			{Header: "Transaction Date", Path: ".transactionDate", Transformer: dateTransformerMs},
			{Header: "Post Date", Path: ".postDate", Transformer: dateTransformerMs},
			{Header: "Amount", Path: ".amount"},
			{Header: "Description", Path: ".description"},
			{Header: "Transaction Type", Path: ".transactionType"},
			{Header: "Memo", Path: ".memo"},
			{Header: "Symbol", Path: ".brokerage.product.symbol"},
			{Header: "Security Type", Path: ".brokerage.product.securityType"},
			{Header: "Quantity", Path: ".brokerage.quantity"},
			{Header: "Price", Path: ".brokerage.price"},
			{Header: "Settlement Currency", Path: ".brokerage.settlementCurrency"},
			{Header: "Payment Currency", Path: ".brokerage.paymentCurrency"},
			{Header: "Fee", Path: ".brokerage.fee"},
			{Header: "Settlement Date", Path: ".brokerage.settlementDate"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
