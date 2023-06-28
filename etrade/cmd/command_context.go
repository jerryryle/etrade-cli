package cmd

import (
	"errors"
	"fmt"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/session"
	"golang.org/x/exp/slog"
	"os"
	"time"
)

type CommandContext struct {
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

	// Load the configuration file and locate the configuration for the requested customer ID
	configurationFolder, err := getUserHomeFolder()
	if err != nil {
		return nil, fmt.Errorf("unable to locate the current user's home folder: %w", err)
	}
	cfgFilePath := getCfgFilePath(configurationFolder)
	customerConfigurationStore, err := LoadCustomerConfigurationStoreFromFile(cfgFilePath, logger)
	if err != nil {
		return nil, fmt.Errorf(
			"configuration file %s is missing or corrupt (error: %w). you can create a default configuration file with the command 'cfg create'",
			cfgFilePath, err,
		)
	}

	return &CommandContext{
		Logger:                     logger,
		Renderer:                   renderer,
		ConfigurationFolder:        configurationFolder,
		CustomerConfigurationStore: customerConfigurationStore,
	}, nil
}

func (c *CommandContext) Close() error {
	return c.Renderer.Close()
}

func NewCommandContextWithClientFromFlags(flags *globalFlags) (*CommandContextWithClient, error) {
	context, err := NewCommandContextFromFlags(flags)
	if err != nil {
		return nil, err
	}

	if flags.customerId == "" {
		return nil, errors.New("customer id must be specified with --customer-id flag")
	}
	customerConfig, err := context.CustomerConfigurationStore.GetCustomerConfigurationById(flags.customerId)
	if err != nil {
		return nil, fmt.Errorf("customer id '%s' not found in config file", flags.customerId)
	}
	cacheFilePath := getFileCachePathForCustomer(context.ConfigurationFolder, customerConfig.CustomerConsumerKey)

	// Create an ETrade client that's authorized for the customer
	eTradeClient, err := createClientWithCredentialCache(
		customerConfig.CustomerProduction,
		customerConfig.CustomerConsumerKey,
		customerConfig.CustomerConsumerSecret,
		cacheFilePath,
		context.Logger,
	)

	return &CommandContextWithClient{
		Logger:   context.Logger,
		Renderer: context.Renderer,
		Client:   eTradeClient,
	}, nil
}

func (c *CommandContextWithClient) Close() error {
	return c.Renderer.Close()
}

func createClientWithCredentialCache(
	production bool, consumerKey string, consumerSecret string, cacheFilePath string, logger *slog.Logger,
) (client.ETradeClient, error) {
	cachedCredentials, err := LoadCachedCredentialsFromFile(cacheFilePath, logger)
	if err != nil {
		// Create a new, empty credential cache. It will yield empty strings for the cached token, which
		// will indicate that there are no cached credentials for this customer
		cachedCredentials = &CachedCredentials{}
	}

	var eTradeClient client.ETradeClient
	authSession, err := session.CreateSession(production, consumerKey, consumerSecret, logger)
	if err != nil {
		return nil, err
	}
	var accessToken = cachedCredentials.AccessToken
	var accessSecret = cachedCredentials.AccessSecret
	eTradeClient, err = authSession.Renew(accessToken, accessSecret)
	if err != nil {
		authUrl, err := authSession.Begin()
		if err != nil {
			return nil, err
		}
		_, _ = fmt.Fprintf(os.Stderr, "Visit this URL to get a validation code:\n%s\n\n", authUrl)

		var validationCode string
		_, _ = fmt.Fprintf(os.Stderr, "Enter validation code: ")
		_, err = fmt.Scanln(&validationCode)
		if err != nil {
			return nil, err
		}
		if validationCode == "" {
			return nil, errors.New("no validation code provided")
		}

		eTradeClient, accessToken, accessSecret, err = authSession.Verify(validationCode)
		if err != nil {
			return nil, err
		}
	}
	err = SaveCachedCredentialsToFile(
		cacheFilePath,
		&CachedCredentials{accessToken, accessSecret, time.Now()},
		logger,
	)
	if err != nil {
		return nil, err
	}
	return eTradeClient, nil
}
