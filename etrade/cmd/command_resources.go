package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"golang.org/x/exp/slog"
	"os"
)

type CommandResources struct {
	Logger *slog.Logger
	Client client.ETradeClient
	OFile  *os.File

	closeOFile bool
}

func NewCommandResources(customerId string, debug bool, outputFile string) (*CommandResources, error) {
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
	oDest := os.Stdout
	closeOFile := false
	if outputFile != "" {
		oDest, err = os.Create(outputFile)
		if err != nil {
			return nil, err
		}
		closeOFile = true
	}

	// Create an ETrade client that's authorized for the customer
	etradeClient, err := createClientWithCredentialCache(
		customerConfig.CustomerProduction,
		customerConfig.CustomerConsumerKey,
		customerConfig.CustomerConsumerSecret,
		cacheFilePath,
		logger,
	)
	if err != nil {
		return nil, err
	}

	return &CommandResources{
		Logger:     logger,
		Client:     etradeClient,
		OFile:      oDest,
		closeOFile: closeOFile,
	}, nil
}

func CleanupCommandResources(resources *CommandResources) error {
	if resources.closeOFile {
		err := resources.OFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
