package cmd

import (
	"github.com/spf13/cobra"
)

type CommandAlertsDelete struct {
	Context *CommandContextWithClient
}

func (c *CommandAlertsDelete) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [Alert Id] ...",
		Short: "Delete alerts",
		Long:  "Delete one or more alerts by ID",
		Args:  cobra.MatchAll(cobra.RangeArgs(1, 25)),
		RunE: func(cmd *cobra.Command, args []string) error {
			alertIds := args
			if response, err := DeleteAlerts(c.Context.Client, alertIds); err == nil {
				return c.Context.Renderer.Render(response, alertListDescriptor)
			} else {
				return err
			}
		},
	}
	return cmd
}
