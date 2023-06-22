package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"golang.org/x/exp/slog"
	"os"
)

type CommandContext struct {
	Logger             *slog.Logger
	Client             client.ETradeClient
	OutputFile         *os.File
	OutputFormatJson   bool
	OutputFormatPretty bool

	closeOFile bool
}

func NewCommandContext(customerId string, debug bool, outputFileName string, format OutputFormat) (
	*CommandContext, error,
) {
	// Set the default log level, based on the verbose flag.
	var logLevel = slog.LevelError
	if debug {
		logLevel = slog.LevelDebug
	}

	// Create a logger.
	logHandlerOptions := slog.HandlerOptions{
		AddSource: false,
		Level:     logLevel,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &logHandlerOptions))

	// Load the configuration file and locate the configuration for the requested customer ID
	userHomeFolder, err := getUserHomeFolder()
	if err != nil {
		return nil, fmt.Errorf("unable to locate the current user's home folder: %w", err)
	}
	cfgFilePath := getCfgFilePath(userHomeFolder)
	customerConfigStore, err := LoadCustomerConfigurationsStoreFromFile(cfgFilePath, logger)
	if err != nil {
		return nil, fmt.Errorf(
			"configuration file %s is missing or corrupt (error: %w). you can create a default configuration file with the command 'cfg create'",
			cfgFilePath, err,
		)
	}
	if customerId == "" {
		return nil, errors.New("customer id must be specified with --customerId flag")
	}
	customerConfig, err := customerConfigStore.GetCustomerConfigurationById(customerId)
	if err != nil {
		return nil, fmt.Errorf("customer id '%s' not found in config file at %s", customerId, cfgFilePath)
	}
	cacheFilePath := getFileCachePathForCustomer(userHomeFolder, customerConfig.CustomerConsumerKey)

	// Set the command output destination
	OutputFile := os.Stdout
	closeOFile := false
	if outputFileName != "" {
		OutputFile, err = os.Create(outputFileName)
		if err != nil {
			return nil, err
		}
		closeOFile = true
	}

	// Set up output format variables.
	// TODO: Revisit this to see if creating an output renderer makes more sense.
	var json, pretty bool
	switch format {
	case OutputFormatText:
		json = false
	case OutputFormatJson:
		json = true
		pretty = false
	case OutputFormatJsonPretty:
		json = true
		pretty = true
	}

	// Create an ETrade client that's authorized for the customer
	eTradeClient, err := createClientWithCredentialCache(
		customerConfig.CustomerProduction,
		customerConfig.CustomerConsumerKey,
		customerConfig.CustomerConsumerSecret,
		cacheFilePath,
		logger,
	)
	if err != nil {
		return nil, err
	}

	return &CommandContext{
		Logger:             logger,
		Client:             eTradeClient,
		OutputFile:         OutputFile,
		OutputFormatJson:   json,
		OutputFormatPretty: pretty,
		closeOFile:         closeOFile,
	}, nil
}

func CleanupCommandContext(context *CommandContext) error {
	if context.closeOFile {
		err := context.OutputFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
