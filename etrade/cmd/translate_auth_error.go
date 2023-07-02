package cmd

import (
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
)

func AddErrorHelp(err error) error {
	if client.IsAuthFailed(err) {
		return fmt.Errorf("%w; please authenticate with the 'auth login' command first", err)
	}
	return err
}
