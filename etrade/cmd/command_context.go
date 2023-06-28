package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"golang.org/x/exp/slog"
	"os"
)

type CommandContext struct {
	Logger   *slog.Logger
	Client   client.ETradeClient
	Renderer Renderer
}

func NewCommandContext(customerId string, debug bool, outputFileName string, format OutputFormat) (
	*CommandContext, error,
) {
	var err error

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

	// Set the command output destination
	outputFile := os.Stdout
	if outputFileName != "" {
		outputFile, err = os.Create(outputFileName)
		if err != nil {
			return nil, err
		}
	}

	// Set up output renderer
	renderer := Renderer(nil)
	switch format {
	case OutputFormatJson:
		renderer = &jsonRenderer{
			outputFile: outputFile,
			pretty:     false,
		}
	case OutputFormatJsonPretty:
		renderer = &jsonRenderer{
			outputFile: outputFile,
			pretty:     true,
		}
	case OutputFormatCsv:
		renderer = &csvRenderer{
			outputFile: outputFile,
			pretty:     true,
		}
	}

	// Load the configuration file and locate the configuration for the requested customer ID
	userHomeFolder, err := getUserHomeFolder()
	if err != nil {
		return nil, fmt.Errorf("unable to locate the current user's home folder: %w", err)
	}
	cfgFilePath := getCfgFilePath(userHomeFolder)
	customerConfigStore, err := LoadCustomerConfigurationStoreFromFile(cfgFilePath, logger)
	if err != nil {
		return nil, fmt.Errorf(
			"configuration file %s is missing or corrupt (error: %w). you can create a default configuration file with the command 'cfg create'",
			cfgFilePath, err,
		)
	}
	if customerId == "" {
		return nil, errors.New("customer id must be specified with --customer-id flag")
	}
	customerConfig, err := customerConfigStore.GetCustomerConfigurationById(customerId)
	if err != nil {
		return nil, fmt.Errorf("customer id '%s' not found in config file at %s", customerId, cfgFilePath)
	}
	cacheFilePath := getFileCachePathForCustomer(userHomeFolder, customerConfig.CustomerConsumerKey)

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
		Logger:   logger,
		Client:   eTradeClient,
		Renderer: renderer,
	}, nil
}

func CleanupCommandContext(context *CommandContext) error {
	err := context.Renderer.Close()
	if err != nil {
		return err
	}
	return nil
}
