package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAccountsList struct {
	Context *CommandContextWithClient
}

func (c *CommandAccountsList) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List accounts",
		Long:  "List all accounts for the current customer",
		RunE: func(cmd *cobra.Command, args []string) error {
			if response, err := ListAccounts(c.Context.Client); err == nil {
				return c.Context.Renderer.Render(response, accountListDescriptor)
			} else {
				return err
			}
		},
	}
	return cmd
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
