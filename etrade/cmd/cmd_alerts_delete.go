package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/spf13/cobra"
	"strings"
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
				return c.Context.Renderer.Render(response, deleteAlertsDescriptor)
			} else {
				return err
			}
		},
	}
	return cmd
}

var deleteAlertsDescriptor = []RenderDescriptor{
	{
		ObjectPath: "",
		Values: []RenderValue{
			{Header: "Status", Path: ".status"},
			{Header: "Error Message", Path: ".error"},
			{
				Header: "Failed to Delete", Path: ".failedAlerts", Transformer: transformAlertList,
			},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}

func transformAlertList(value interface{}) interface{} {
	if alertList, ok := value.(jsonmap.JsonSlice); ok {
		alertStringList := make([]string, len(alertList))
		for i, v := range alertList {
			alertStringList[i] = fmt.Sprintf("%d", v)
		}
		return strings.Join(alertStringList, ", ")
	} else {
		return value
	}
}
