# etrade-cli
E*TRADE Command-line Interface

This is a hobby project to create a command-line E*TRADE client in Go.

Quick Start:
1. Install Go 1.20 or later: https://go.dev/doc/install
2. `make install` - Build and install the binary to your Go install path.
3. `export PATH=$PATH:/path/to/your/install/directory` - Ensure Go install path is in your system path.
4. `etrade cfg create` - Create a default config file (the command will print the config file path)
5. Edit the default config file to choose a Customer Id and add your keys/secrets.
6. `etrade --customer-id <your customer ID> accounts list` - List all accounts for customer.
7. `etrade --customer-id <your customer ID> accounts portfolio <account ID>` - Get portfolio for an account in CSV format
8. `etrade --customer-id --format json <your customer ID> accounts portfolio <account ID>` - Get portfolio for an account in JSON format
