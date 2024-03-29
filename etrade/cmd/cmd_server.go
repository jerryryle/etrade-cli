package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"os/signal"
)

type commandServerFlags struct {
	listenAddr string
}

type CommandServer struct {
	context CommandContextWithStore
	flags   commandServerFlags
}

func (c *CommandServer) Command(globalFlags *globalFlags) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Run as server",
		Long:  "Run as server",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			cmdContext, err := NewCommandContextWithStoreFromFlags(globalFlags)
			if err != nil {
				return err
			}
			c.context = *cmdContext
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			return c.context.Close()
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			_, _ = fmt.Fprintf(os.Stderr, "Starting server on: \"%s\"\n", c.flags.listenAddr)

			server := NewETradeServer(
				c.flags.listenAddr, c.context.Logger, c.context.ConfigurationFolder,
				c.context.CustomerConfigurationStore,
			)

			idleConnsClosed := make(chan struct{})
			go func() {
				sigint := make(chan os.Signal, 1)
				signal.Notify(sigint, os.Interrupt)
				<-sigint

				// Shut down upon receiving an interrupt signal
				if err := server.Shutdown(context.Background()); err != nil {
					// Error from closing listeners, or context timeout:
					c.context.Logger.Error(fmt.Errorf("http server Shutdown() failed (%w)", err).Error())
				}
				close(idleConnsClosed)
			}()

			if err := server.ListenAndServe(); err != http.ErrServerClosed {
				// Error starting or closing listener:
				c.context.Logger.Error(fmt.Errorf("http server ListenAndServe() failed (%w)", err).Error())
				return err
			}

			<-idleConnsClosed
			return nil
		},
	}
	// Add Flags
	cmd.Flags().StringVarP(&c.flags.listenAddr, "addr", "a", ":8888", "server listen address:port")
	return cmd
}
