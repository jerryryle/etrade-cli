package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
	"github.com/spf13/cobra"
)

type GetQuotesCommand struct {
	AppContext *ApplicationContext
}

func (c *GetQuotesCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getquotes",
		Short: "Get quotes",
		Long:  "Get quotes for one or more symbols",
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.GetQuotes()
		},
	}
	return cmd
}

func (c *GetQuotesCommand) GetQuotes() error {
	quotes, err := c.AppContext.Client.GetQuotes([]string{"GSPC"}, etradelib.QuoteDetailMutualFund)
	if err != nil {
		return err
	}
	fmt.Printf("%#v\n", quotes)
	return nil
}
