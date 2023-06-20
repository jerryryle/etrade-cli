package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type commandAccountsBalancesFlags struct {
	realTimeBalance bool
}

type CommandAccountsBalances struct {
	Resources *CommandResources
	flags     commandAccountsBalancesFlags
}

func (c *CommandAccountsBalances) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "balances [account ID]",
		Short: "Get account balances",
		Long:  "Get account balances",
		Args:  cobra.MatchAll(cobra.ExactArgs(1)),
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.GetAccountBalances(args[0])
		},
	}
	cmd.Flags().BoolVarP(&c.flags.realTimeBalance, "realtime-balance", "r", true, "return real time balance")
	return cmd
}

func (c *CommandAccountsBalances) GetAccountBalances(accountKeyId string) error {
	response, err := c.Resources.Client.GetAccountBalances(accountKeyId, c.flags.realTimeBalance)
	if err != nil {
		return err
	}
	_, _ = fmt.Fprintln(c.Resources.OFile, string(response))
	return nil
}
