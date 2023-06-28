package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/spf13/cobra"
)

type CommandAlertsDetails struct {
	Context *CommandContextWithClient
}

func (c *CommandAlertsDetails) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "details [Alert ID]",
		Short: "Show alert details",
		Long:  "Show alert details for alert with ID",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListAlertDetails(args[0])
		},
	}
	return cmd
}

func (c *CommandAlertsDetails) ListAlertDetails(alertId string) error {
	response, err := c.Context.Client.ListAlertDetails(alertId, false)
	if err != nil {
		return err
	}
	alertDetails, err := etradelib.CreateETradeAlertDetailsFromResponse(response)
	if err != nil {
		return err
	}
	err = c.Context.Renderer.Render(alertDetails.AsJsonMap(), alertDetailsDescriptor)
	if err != nil {
		return err
	}
	return nil
}

var alertDetailsDescriptor = []RenderDescriptor{
	{
		ObjectPath: "",
		Values: []RenderValue{
			{Header: "Alert ID", Path: ".id"},
			{Header: "Create Time", Path: ".createTime", Transformer: dateTimeTransformer},
			{Header: "Subject", Path: ".subject"},
			{Header: "Message Text", Path: ".msgText"},
			{Header: "Read Time", Path: ".readTime", Transformer: dateTimeTransformer},
			{Header: "Delete Time", Path: ".deleteTime", Transformer: dateTimeTransformer},
			{Header: "Symbol", Path: ".symbol"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
