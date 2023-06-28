package cmd

import (
	"errors"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
	"time"
)

type accountsTransactionsListFlags struct {
	startDate string
	endDate   string
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
	cmd.Flags().StringVarP(&c.flags.startDate, "start-date", "s", "", "start date (MMDDYYYY)")
	cmd.Flags().StringVarP(&c.flags.endDate, "end-date", "e", "", "end date (MMDDYYYY)")
	return cmd
}

func (c *CommandAccountsTransactionsList) ListTransactions(
	accountId string, startDate *time.Time, endDate *time.Time,
) error {
	// This determines how many transaction items will be retrieved in each
	// request. This should normally be set to the max for efficiency, but can
	// be lowered to test the pagination logic.
	const countPerRequest = constants.TransactionsMaxCount

	account, err := GetAccountById(c.Context.Client, accountId)
	if err != nil {
		return err
	}
	response, err := c.Context.Client.ListTransactions(
		account.GetIdKey(),
		startDate, endDate, constants.SortOrderAsc, "", countPerRequest,
	)
	if err != nil {
		return err
	}
	transactionList, err := etradelib.CreateETradeTransactionListFromResponse(response)
	if err != nil {
		return err
	}

	for transactionList.NextPage() != "" {
		response, err = c.Context.Client.ListTransactions(
			account.GetIdKey(),
			startDate, endDate, constants.SortOrderAsc, transactionList.NextPage(), countPerRequest,
		)
		if err != nil {
			return err
		}
		err = transactionList.AddPageFromResponse(response)
		if err != nil {
			return err
		}
	}

	err = c.Context.Renderer.Render(transactionList.AsJsonMap(), transactionListDescriptor)
	if err != nil {
		return err
	}
	return nil
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
