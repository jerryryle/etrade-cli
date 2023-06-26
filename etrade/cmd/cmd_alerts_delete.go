package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type CommandAlertsDelete struct {
	Context *CommandContext
}

func (c *CommandAlertsDelete) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [Alert Id] ...",
		Short: "Delete alerts",
		Long:  "Delete one or more alerts by ID",
		Args:  cobra.MatchAll(cobra.RangeArgs(1, 25)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.DeleteAlerts(args)
		},
	}
	return cmd
}

func (c *CommandAlertsDelete) DeleteAlerts(alertIds []string) error {
	_, err := c.Context.Client.DeleteAlerts(alertIds)
	if err != nil {
		return fmt.Errorf("requested Alert Id(s) may not exist (%w)", err)
	}
	response, err := c.Context.Client.ListAlerts(
		constants.AlertsMaxCount, constants.AlertCategoryNil, constants.AlertStatusNil, constants.SortOrderNil, "",
	)
	if err != nil {
		return err
	}
	alertsList, err := etradelib.CreateETradeAlertListFromResponse(response)
	if err != nil {
		return err
	}
	err = c.Context.Renderer.Render(alertsList.AsJsonMap(), alertListDescriptor)
	if err != nil {
		return err
	}
	return nil
}
