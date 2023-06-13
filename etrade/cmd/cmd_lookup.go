package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type LookupCommand struct {
	AppContext *ApplicationContext
}

func (c *LookupCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lookup",
		Short: "Look up product",
		Long:  "Look up products based on a full or partial match of any part of the company name.",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Lookup(args[0])
		},
	}
	return cmd
}

func (c *LookupCommand) Lookup(search string) error {
	result, err := c.AppContext.Client.LookupProduct(search)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", result)
	return nil
}
