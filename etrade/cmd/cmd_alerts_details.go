package cmd

import (
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
			alertId := args[0]
			if response, err := ListAlertDetails(c.Context.Client, alertId); err == nil {
				return c.Context.Renderer.Render(response, alertDetailsDescriptor)
			} else {
				return err
			}
		},
	}
	return cmd
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
