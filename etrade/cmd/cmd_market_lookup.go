package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type CommandMarketLookup struct {
	Context *CommandContext
}

func (c *CommandMarketLookup) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lookup [search]",
		Short: "Look up product",
		Long:  "Look up products based on a full or partial match of any part of the company name.",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Lookup(args[0])
		},
	}
	return cmd
}

func (c *CommandMarketLookup) Lookup(search string) error {
	response, err := c.Context.Client.LookupProduct(search)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintf(c.Context.OutputFile, string(response))
	return nil
}
