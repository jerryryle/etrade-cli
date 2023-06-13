package main

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/etrade/cmd"
	"os"

	// Ensure the program has an embedded copy of tzdata in case the host system is missing this.
	_ "time/tzdata"
)

func main() {
	var appContext = cmd.ApplicationContext{}

	rootCmd := (&cmd.RootCommand{AppContext: &appContext}).Command()
	rootCmd.AddCommand((&cmd.ListAccountsCommand{AppContext: &appContext}).Command())
	rootCmd.AddCommand((&cmd.ListAlertsCommand{AppContext: &appContext}).Command())
	rootCmd.AddCommand((&cmd.GetQuotesCommand{AppContext: &appContext}).Command())
	rootCmd.AddCommand((&cmd.LookupCommand{AppContext: &appContext}).Command())
	rootCmd.AddCommand((&cmd.GetOptionChainsCommand{AppContext: &appContext}).Command())

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}
