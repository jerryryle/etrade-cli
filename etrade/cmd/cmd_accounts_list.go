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
	responseMap, err := etradelib.NewNormalizedJsonMap(response)
	if err != nil {
		return err
	}
	accountList, err := etradelib.CreateETradeAccountList(responseMap)
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
		ValueHeaders: []string{
			"Account ID", "Account Mode", "Account Description", "Account Nickname", "Account Type", "Institution Type",
			"Account Status", "Account Closed Date",
		},
		ValuePaths: []string{
			".accountId", ".accountMode", ".accountDesc", ".accountName", ".accountType", ".institutionType",
			".accountStatus", ".closedDate",
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
