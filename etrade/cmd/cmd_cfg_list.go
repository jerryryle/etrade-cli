package cmd

import (
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"github.com/spf13/cobra"
)

type CommandCfgList struct {
	context CommandContextWithStore
}

func (c *CommandCfgList) Command(globalFlags *globalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List configuration",
		Long:  "List all available customers in the app configuration",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			context, err := NewCommandContextWithStoreFromFlags(globalFlags)
			if err != nil {
				return err
			}
			c.context = *context
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return c.context.Close()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListConfig()
		},
	}
	return cmd
}

func (c *CommandCfgList) ListConfig() error {
	customerSlice := jsonmap.JsonSlice{}
	for customerId, customerConfig := range c.context.CustomerConfigurationStore.GetAllConfigurations() {
		customerMap := jsonmap.JsonMap{}
		if err := customerMap.SetString("customerId", customerId); err != nil {
			return err
		}
		if err := customerMap.SetString("customerName", customerConfig.CustomerName); err != nil {
			return err
		}
		if err := customerMap.SetBool("productionAccess", customerConfig.CustomerProduction); err != nil {
			return err
		}
		customerSlice = append(customerSlice, customerMap)
	}

	resultMap := jsonmap.JsonMap{
		"customers": customerSlice,
	}
	if err := c.context.Renderer.Render(resultMap, cfgListDescriptor); err != nil {
		return err
	}
	return nil
}

var cfgListDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".customers",
		Values: []RenderValue{
			{Header: "Customer Id", Path: ".customerId"},
			{Header: "Customer Name", Path: ".customerName"},
			{Header: "Production Access", Path: ".productionAccess"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
