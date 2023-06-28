package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"text/tabwriter"
)

type CommandCfgList struct {
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
	userHomeFolder, err := getUserHomeFolder()
	if err != nil {
		return fmt.Errorf("unable to locate the current user's home folder: %w", err)
	}
	cfgFilePath := getCfgFilePath(userHomeFolder)
	customerConfigStore, err := LoadCustomerConfigurationStoreFromFile(cfgFilePath, nil)
	if err != nil {
		return fmt.Errorf(
			"configuration file %s is missing or corrupt (error: %w). you can create a default configuration file with the command 'cfg create'",
			cfgFilePath, err,
		)
	}
	w := tabwriter.NewWriter(os.Stdout, 1, 0, 4, ' ', 0)
	_, _ = fmt.Fprintln(w, "Customer Id\tCustomer Name\tProduction Access")
	_, _ = fmt.Fprintln(w, "-----------\t-------------\t-----------------")

	customerConfigStore.ForEachCustomerConfig(
		func(configId string, customerName string, production bool) {
			_, _ = fmt.Fprintf(
				w,
				"%s\t%s\t%t\n", configId, customerName, production,
			)
		},
	)
	_ = w.Flush()
	return nil
}
