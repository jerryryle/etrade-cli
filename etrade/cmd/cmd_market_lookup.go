package cmd

import (
	"github.com/spf13/cobra"
)

type CommandMarketLookup struct {
	Context *CommandContextWithClient
}

func (c *CommandMarketLookup) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "lookup [search]",
		Short: "Look up product",
		Long:  "Look up products based on a full or partial match of any part of the company name.",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			search := args[0]
			if response, err := Lookup(c.Context.Client, search); err == nil {
				return c.Context.Renderer.Render(response, resultListDescriptor)
			} else {
				return err
			}
		},
	}
	return cmd
}

var resultListDescriptor = []RenderDescriptor{
	{
		ObjectPath: ".results",
		Values: []RenderValue{
			{Header: "Symbol", Path: ".symbol"},
			{Header: "Description", Path: ".description"},
			{Header: "Type", Path: ".type"},
		},
		DefaultValue: "",
		SpaceAfter:   false,
	},
}
