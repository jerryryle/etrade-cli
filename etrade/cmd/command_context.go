package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"golang.org/x/exp/slog"
	"os"
)

type CommandContext struct {
	Logger              *slog.Logger
	Renderer            Renderer
	ConfigurationFolder string
}

type CommandContextWithStore struct {
	Logger                     *slog.Logger
	Renderer                   Renderer
	ConfigurationFolder        string
	CustomerConfigurationStore *CustomerConfigurationStore
}

type CommandContextWithClient struct {
	Logger   *slog.Logger
	Renderer Renderer
	Client   client.ETradeClient
}

func NewCommandContextFromFlags(flags *globalFlags) (*CommandContext, error) {
	var err error

	// Set the default log level, based on the debug flag.
	var logLevel = slog.LevelError
	if flags.debug {
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
	if flags.outputFileName != "" {
		outputFile, err = os.Create(flags.outputFileName)
		if err != nil {
			return nil, err
		}
	}

	// Set up output renderer
	renderer := Renderer(nil)
	switch flags.outputFormat.Value() {
	case outputFormatJson:
		renderer = &jsonRenderer{
			outputFile: outputFile,
			pretty:     false,
		}
	case outputFormatJsonPretty:
		renderer = &jsonRenderer{
			outputFile: outputFile,
			pretty:     true,
		}
	default:
		renderer = &csvRenderer{
			outputFile: outputFile,
			pretty:     true,
		}
	}

	// Locate the configuration folder
	configurationFolder, err := getUserHomeFolder()
	if err != nil {
		return nil, fmt.Errorf("unable to locate the current user's home folder: %w", err)
	}

	return &CommandContext{
		Logger:              logger,
		Renderer:            renderer,
		ConfigurationFolder: configurationFolder,
	}, nil
}

func (c *CommandContext) Close() error {
	return c.Renderer.Close()
}

func NewCommandContextWithStoreFromFlags(flags *globalFlags) (*CommandContextWithStore, error) {
	context, err := NewCommandContextFromFlags(flags)
	if err != nil {
		return nil, err
	}

	// Load the configuration file
	cfgFilePath := getCfgFilePath(context.ConfigurationFolder)
	customerConfigurationStore, err := LoadCustomerConfigurationStoreFromFile(cfgFilePath, context.Logger)
	if err != nil {
		return nil, fmt.Errorf(
			"configuration file %s is missing or corrupt (error: %w). you can create a default configuration file with the command 'cfg create'",
			cfgFilePath, err,
		)
	}

	return &CommandContextWithStore{
		Logger:                     context.Logger,
		Renderer:                   context.Renderer,
		ConfigurationFolder:        context.ConfigurationFolder,
		CustomerConfigurationStore: customerConfigurationStore,
	}, nil
}

func (c *CommandContextWithStore) Close() error {
	return c.Renderer.Close()
}

func NewCommandContextWithClientFromFlags(flags *globalFlags) (*CommandContextWithClient, error) {
	context, err := NewCommandContextWithStoreFromFlags(flags)
	if err != nil {
		return nil, err
	}

	eTradeClient, err := NewETradeClientForCustomer(
		flags.customerId, context.ConfigurationFolder, context.CustomerConfigurationStore, context.Logger,
	)
	if err != nil {
		return nil, err
	}
	return &CommandContextWithClient{
		Logger:   context.Logger,
		Renderer: context.Renderer,
		Client:   eTradeClient,
	}, nil
}

func NewETradeClientForCustomer(
	customerId string, cfgFolder string, cfgStore *CustomerConfigurationStore, logger *slog.Logger,
) (client.ETradeClient, error) {
	if customerId == "" {
		return nil, errors.New("customer id must be specified with --customer-id flag")
	}
	customerConfig, err := cfgStore.GetCustomerConfigurationById(customerId)
	if err != nil {
		return nil, fmt.Errorf("customer id '%s' not found in config file", customerId)
	}
	cacheFilePath := getFileCachePathForCustomer(cfgFolder, customerConfig.CustomerConsumerKey)

	// Try loading cached credentials
	cachedCredentials, err := LoadCachedCredentialsFromFile(cacheFilePath, logger)
	if err != nil {
		// If loading cached credentials fails, then create a new, empty
		// credential cache. It will yield empty strings for the cached token,
		// which will indicate that there are no cached credentials for this
		// customer.
		cachedCredentials = &CachedCredentials{}
	}
	return client.CreateETradeClient(
		logger, customerConfig.CustomerProduction, customerConfig.CustomerConsumerKey,
		customerConfig.CustomerConsumerSecret, cachedCredentials.AccessToken, cachedCredentials.AccessSecret,
	)
}

func (c *CommandContextWithClient) Close() error {
	return c.Renderer.Close()
}
