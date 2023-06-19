package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type CommandAlertsList struct {
	Resources *CommandResources
}

func (c *CommandAlertsList) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List alerts",
		Long:  "List all alerts for the current customer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListAlerts()
		},
	}
	return cmd
}

func (c *CommandAlertsList) ListAlerts() error {
	response, err := c.Resources.Client.ListAlerts(
		-1, constants.AlertCategoryNil, constants.AlertStatusNil, constants.SortOrderNil, "",
	)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
