package main

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/etrade/cmd"
	"os"

	// Ensure the program has an embedded copy of tzdata in case the host system is missing this.
	_ "time/tzdata"
)

func main() {
	rootCmd := (&cmd.RootCommand{}).Command()

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, cmd.AddErrorHelp(err))
		os.Exit(1)
	}
	os.Exit(0)
}
