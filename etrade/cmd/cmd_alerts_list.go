package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type CommandAlertsList struct {
	Context *CommandContext
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
	response, err := c.Context.Client.ListAlerts(
		-1, constants.AlertCategoryNil, constants.AlertStatusNil, constants.SortOrderNil, "",
	)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintln(c.Context.OutputFile, string(response))
	return nil
}
