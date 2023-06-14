package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type getAccountBalancesCommandFlags struct {
	realTimeNAV bool
}

type GetAccountBalancesCommand struct {
	AppContext *ApplicationContext
	flags      getAccountBalancesCommandFlags
}

func (c *GetAccountBalancesCommand) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "getaccountbalances [account ID]",
		Short: "Get account balances",
		Long:  "Get account balances",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.GetAccountBalances(args[0])
		},
	}
	cmd.Flags().BoolVarP(&c.flags.realTimeNAV, "realTimeNAV", "r", true, "return real time balance")
	return cmd
}

func (c *GetAccountBalancesCommand) GetAccountBalances(accountKeyId string) error {
	response, err := c.AppContext.Client.GetAccountBalances(accountKeyId, c.flags.realTimeNAV)
	if err != nil {
		return err
	}
	fmt.Println(string(response))
	return nil
}
