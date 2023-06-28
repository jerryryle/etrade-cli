package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

type CommandCfgList struct {
	Context *CommandContext
}

func (c *CommandCfgList) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List configuration",
		Long:  "List all available customers in the app configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.ListConfig()
		},
	}
	return cmd
}

func (c *CommandCfgList) ListConfig() error {
	w := tabwriter.NewWriter(os.Stdout, 1, 0, 4, ' ', 0)
	_, _ = fmt.Fprintln(w, "Customer Id\tCustomer Name\tProduction Access")
	_, _ = fmt.Fprintln(w, "-----------\t-------------\t-----------------")

	for customerId, customerConfig := range c.Context.CustomerConfigurationStore.GetAllConfigurations() {
		_, _ = fmt.Fprintf(
			w,
			"%s\t%s\t%t\n", customerId, customerConfig.CustomerName, customerConfig.CustomerProduction,
		)
	}

	_ = w.Flush()
	return nil
}
