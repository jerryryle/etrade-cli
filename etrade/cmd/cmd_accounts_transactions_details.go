package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAccountsTransactionsDetails struct {
	Context *CommandContextWithClient
}

func (c *CommandAccountsTransactionsDetails) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "details [account ID] [transaction ID]",
		Short: "List transaction details",
		Long:  "List transaction details",
		Args:  cobra.MatchAll(cobra.ExactArgs(2)),
		RunE: func(cmd *cobra.Command, args []string) error {
			accountId := args[0]
			transactionId := args[1]
			if response, err := ListTransactionDetails(c.Context.Client, accountId, transactionId); err == nil {
				return c.Context.Renderer.Render(response, transactionDetailsDescriptor)
			} else {
				return err
			}
		},
	}
	return cmd
}

var transactionDetailsDescriptor = []RenderDescriptor{
	{
		ObjectPath: "",
		Values: []RenderValue{
			{Header: "Transaction ID", Path: ".transactionId"},
			{Header: "Account ID", Path: ".accountId"},
			{Header: "Transaction Date", Path: ".transactionDate", Transformer: dateTransformerMs},
			{Header: "Post Date", Path: ".postDate", Transformer: dateTransformerMs},
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
