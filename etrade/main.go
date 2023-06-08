package main

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/etrade/cmd"
	"os"
)

func main() {
	rootCmd := (&cmd.RootCommand{}).Command()

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	os.Exit(0)
}
