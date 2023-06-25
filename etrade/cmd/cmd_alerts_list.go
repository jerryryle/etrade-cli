package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/spf13/cobra"
)

type alertsListFlags struct {
	count     int
	category  enumFlagValue[constants.AlertCategory]
	status    enumFlagValue[constants.AlertStatus]
	sortOrder enumFlagValue[constants.SortOrder]
	search    string
}

type CommandAlertsList struct {
	Context *CommandContext
	flags   alertsListFlags
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
	// Add Flags
	cmd.Flags().IntVarP(&c.flags.count, "count", "n", constants.AlertsMaxCount, "max number of alerts to return")
	cmd.Flags().StringVarP(&c.flags.search, "search", "q", "", "alert subject search string")

	// Initialize Enum Flag Values
	c.flags.category = *newEnumFlagValue(alertCategoryMap, constants.AlertCategoryNil)
	c.flags.status = *newEnumFlagValue(alertStatusMap, constants.AlertStatusNil)
	c.flags.sortOrder = *newEnumFlagValue(sortOrderMap, constants.SortOrderNil)

	// Add Enum Flags
	cmd.Flags().VarP(
		&c.flags.category, "category", "c",
		fmt.Sprintf("alert category (%s)", c.flags.category.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"category",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.category.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.category, "status", "s",
		fmt.Sprintf("alert status (%s)", c.flags.status.JoinAllowedValues(", ")),
	)
	_ = cmd.RegisterFlagCompletionFunc(
		"status",
		func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
			return c.flags.status.AllowedValuesWithHelp(), cobra.ShellCompDirectiveDefault
		},
	)

	cmd.Flags().VarP(
		&c.flags.category, "sort-order", "o",
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

func (c *CommandAlertsList) ListAlerts() error {
	response, err := c.Context.Client.ListAlerts(
		c.flags.count, c.flags.category.Value(), c.flags.status.Value(), c.flags.sortOrder.Value(), c.flags.search,
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

var alertListDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".alerts",
		Values: []RenderValue{
			{Header: "Alert ID", Path: ".id"},
			{Header: "Create Time", Path: ".createTime", Transformer: dateTimeTransformer},
			{Header: "Subject", Path: ".subject"},
			{Header: "Status", Path: ".status"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}

var alertCategoryMap = map[string]enumValueWithHelp[constants.AlertCategory]{
	"stock":   {constants.AlertCategoryStock, "only stock-related alerts"},
	"account": {constants.AlertCategoryAccount, "only account-related alerts"},
}

var alertStatusMap = map[string]enumValueWithHelp[constants.AlertStatus]{
	"read":    {constants.AlertStatusRead, "only read alerts"},
	"unread":  {constants.AlertStatusUnread, "only unread alerts"},
	"deleted": {constants.AlertStatusDeleted, "only deleted alerts"},
}
