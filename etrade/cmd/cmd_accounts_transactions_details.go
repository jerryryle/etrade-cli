package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/spf13/cobra"
)

type CommandAccountsTransactionsDetails struct {
	Context *CommandContext
}

func (c *CommandAccountsTransactionsDetails) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "details [account ID] [transaction ID]",
		Short: "List transaction details",
		Long:  "List transaction details",
		Args:  cobra.MatchAll(cobra.ExactArgs(2)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListTransactionDetails(args[0], args[1])
		},
	}
	return cmd
}

func (c *CommandAccountsTransactionsDetails) ListTransactionDetails(accountId string, transactionId string) error {
	account, err := GetAccountById(c.Context.Client, accountId)
	if err != nil {
		return err
	}
	response, err := c.Context.Client.ListTransactionDetails(account.GetIdKey(), transactionId)
	if err != nil {
		return err
	}
	transactionDetails, err := etradelib.CreateETradeTransactionDetailsFromResponse(response)
	err = c.Context.Renderer.Render(transactionDetails.AsJsonMap(), transactionDetailsDescriptor)
	if err != nil {
		return err
	}
	return nil
}

var transactionDetailsDescriptor = []RenderDescriptor{
	{
		ObjectPath: "",
		Values: []RenderValue{
			{Header: "Transaction ID", Path: ".transactionId"},
			{Header: "Account ID", Path: ".accountId"},
			{Header: "Transaction Date", Path: ".transactionDate", Transformer: dateTransformer},
			{Header: "Post Date", Path: ".postDate", Transformer: dateTransformer},
			{Header: "Amount", Path: ".amount"},
			{Header: "Description", Path: ".description"},
		},
		DefaultValue: "",
		SpaceAfter:   true,
	},
	{
		ObjectPath: ".category",
		Values: []RenderValue{
			{Header: "Category ID", Path: ".categoryId"},
			{Header: "Parent ID", Path: ".parentId"},
			{Header: "Category Name", Path: ".categoryName"},
			{Header: "Parent Name", Path: ".parentName"},
		},
		DefaultValue: "",
		SpaceAfter:   true,
	},
	{
		ObjectPath: ".brokerage",
		Values: []RenderValue{
			{Header: "Transaction Type", Path: ".transactionType"},
			{Header: "Quantity", Path: ".quantity"},
			{Header: "Price", Path: ".price"},
			{Header: "Settlement Currency", Path: ".settlementCurrency"},
			{Header: "Payment Currency", Path: ".paymentCurrency"},
			{Header: "Fee", Path: ".fee"},
			{Header: "Memo", Path: ".memo"},
			{Header: "Check Number", Path: ".checkNo"},
			{Header: "Order Number", Path: ".orderNo"},
		},
		DefaultValue: "",
		SpaceAfter:   true,
	},
	{
		ObjectPath: ".brokerage.product",
		Values: []RenderValue{
			{Header: "Symbol", Path: ".symbol"},
			{Header: "Security Type", Path: ".securityType"},
			{Header: "Security Subtype", Path: ".securitySubType"},
			{Header: "Option Type", Path: ".callPut"},
			{Header: "Option Expiry Year", Path: ".expiryYear"},
			{Header: "Option Expiry Month", Path: ".expiryMonth"},
			{Header: "Option Expiry Day", Path: ".expiryDay"},
			{Header: "Option Strike Price", Path: ".strikePrice"},
			{Header: "Option Expiry Type", Path: ".expiryType"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
