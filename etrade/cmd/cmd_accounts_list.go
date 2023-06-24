package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/spf13/cobra"
)

type CommandAccountsList struct {
	Context *CommandContext
}

func (c *CommandAccountsList) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List accounts",
		Long:  "List all accounts for the current customer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListAccounts()
		},
	}
	return cmd
}

func (c *CommandAccountsList) ListAccounts() error {
	response, err := c.Context.Client.ListAccounts()
	if err != nil {
		return err
	}
	accountList, err := etradelib.CreateETradeAccountListFromResponse(response)
	if err != nil {
		return err
	}
	err = c.Context.Renderer.Render(accountList.AsJsonMap(), accountListDescriptor)
	if err != nil {
		return err
	}
	return nil
}

var accountListDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".accounts",
		Values: []RenderValue{
			{Header: "Account ID", Path: ".accountId"},
			{Header: "Account Mode", Path: ".accountMode"},
			{Header: "Account Description", Path: ".accountDesc"},
			{Header: "Account Nickname", Path: ".accountName"},
			{Header: "Account Type", Path: ".accountType"},
			{Header: "Institution Type", Path: ".institutionType"},
			{Header: "Account Status", Path: ".accountStatus"},
			{Header: "Account Closed Date", Path: ".closedDate"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
