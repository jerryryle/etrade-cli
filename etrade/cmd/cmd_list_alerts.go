package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type ListAlertsCommand struct {
	AppContext *ApplicationContext
}

func (c *ListAlertsCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "listalerts",
		Short: "List alerts",
		Long:  "List all alerts for the current customer",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListAlerts()
		},
	}
	return cmd
}

func (c *ListAlertsCommand) ListAlerts() error {
	response, err := c.AppContext.Client.ListAlerts()
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
