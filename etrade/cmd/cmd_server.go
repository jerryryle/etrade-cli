package cmd

import (
	"context"
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
					c.context.Logger.Error("HTTP server Shutdown: %v", err)
				}
				close(idleConnsClosed)
			}()

			if err := server.ListenAndServe(); err != http.ErrServerClosed {
				// Error starting or closing listener:
				c.context.Logger.Error("HTTP server ListenAndServe: %v", err)
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
