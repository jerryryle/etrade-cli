# etrade-cli
ETrade Command-line Interface

This is a hobby project to create a command-line ETrade client in Go.

USE THIS AT YOUR OWN RISK. The ETrade API is poorly-documented and ETrade's test environment is inadequate; therefore, I cannot guarantee the accuracy of this client implementation. Using it with your ETrade account could result in financial disaster.   

Quick Start:
1. Install Go 1.20 or later: https://go.dev/doc/install
2. `make install` - Build and install the binary to your Go install path.
3. `export PATH=$PATH:/path/to/your/install/directory` - Ensure Go install path is in your system path.
4. `etrade cfg create` - Create a default config file (the command will print the config file path)
5. Edit the default config file to choose a Customer Id and add your keys/secrets.
6. `etrade --customer-id <your customer ID> auth login` - Log in (required before any other commands)
7. `etrade --customer-id <your customer ID> accounts list` - List all accounts for customer.
8. `etrade --customer-id <your customer ID> accounts portfolio <account ID>` - Get portfolio for an account in CSV format
9. `etrade --customer-id --format json <your customer ID> accounts portfolio <account ID>` - Get portfolio for an account in JSON format

Server Mode:
* You can run the application in server mode with: `etrade server`
* In this mode, the server listens for HTTP requests on port 8888. You can change the listen IP address and port using the --addr flag (e.g. --addr=:4444 to listen on all interfaces with port 4444 or --addr=192.168.1.2:4444 to listen on the interface with the IP address 192.168.1.2).
* Stop the server with SIGINT (ctrl-C).
* To quickly test the server using curl:
  1. `curl -X POST http://127.0.0.1:8888/customers/[CUSTOMER_ID]/auth` - Begin authentication. This will either return success (if cached credentials are still valid, in which case you can skip step 2) or a URL for authorization. Visit the URL to get an auth code.
  2. `curl -X POST http://127.0.0.1:8888/customers/[CUSTOMER_ID]/auth -d 'verifyCode=[VERIFY_CODE]'` - Verify using the code obtained from the authorization URL.
  3. `curl http://127.0.0.1:8888/customers/[CUSTOMER_ID]/accounts` - List accounts 
  4. `curl -X DELETE http://127.0.0.1:8888/customers/[CUSTOMER_ID]/auth` - Delete authentication. 

The following documents the server's API:
* /customers
    * GET - Get Customer List
        * No parameters
* /customers/[CUSTOMER ID]/auth
    * POST - Begin/Complete authentication
        * Form Parameters:
            * No Form Parameters - Begin authentication
            * verifyCode=[VERIFY CODE] - Complete authentication
    * DELETE
        * No Form Parameters - Clear cached credentials
* /customers/[CUSTOMER ID]/accounts
    * GET - Get customer account list
        * No Query Parameters
* /customers/[CUSTOMER ID]/accounts/[ACCOUNT ID]/balance
    * GET - Get customer account balance
        * Optional Query Parameters:
            * realTimeBalance=[true, false] - Whether to include real time balance
* /customers/[CUSTOMER ID]/accounts/[ACCOUNT ID]/portfolio
    * GET - Get customer account portfolio
        * Optional Query Parameters:
            * totalsRequired=[true, false] - Whether to include totals
            * view=[performance, fundamental, optionsWatch, quick, complete] - The portfolio view to return (see [this page](https://apisb.etrade.com/docs/api/account/api-portfolio-v1.html#/definitions/Position) for documentation on what's in the various views).
            * sortBy=[symbol, typeName, exchangeName, currency, quantity, longOrShort, dateAcquired, pricePaid, totalGain, totalGainPct, marketValue, bid, ask, priceChange, priceChangePct, volume, week52High, week52Low, eps, peRatio, optionType, strikePrice, premium, expiration, daysGain, commission, marketCap, prevClose, open, daysRange, totalCost, daysGainPct, pctOfPortfolio, lastTradeTime, baseSymbolPrice, week52Range, lastTrade, symbolDesc, bidSize, askSize, otherFees, heldAs, optionMultiplier, deliverables, costPerShare, dividend, divYield, divPayDate, estEarn, exDivDate, tenDayAvgVol, beta, bidAskSpread, marginable, delta52wkHi, delta52WkLow, perf1Mon, annualDiv, perf12Mon, perf3Mon, perf6Mon, preDayVol, sv1MonAvg, sv10DayAvg, sv20DayAvg, sv2MonAvg, sv3MonAvg, sv4MonAvg, sv6MonAvg, delta, gamma, ivPct, theta, vega, adjNonadjFlag, daysExpiration, openInterest, intrinsicValue, rho, typeCode, displaySymbol, afterHoursPctChange, preMarketPctChange, expandCollapseFlag] - The value by which to sort the portfolio results.
            * sortOrder=[ascending, descending] - The sort order
            * marketSession=[regular, extended] - The market session from which to return results
            * withLots=[true, false] - Whether to include lot information for each position (including lot information will make this query take significantly longer)
* /customers/[CUSTOMER ID]/accounts/[ACCOUNT ID]/transactions
    * GET - Get customer account transactions list
        * Optional Query Parameters:
            * startDate=[MMDDYYYY] - The earliest date to include in the date range, formatted as MMDDYYYY. History is available for two years.
            * endDate=[MMDDYYYY] - The latest date to include in the date range, formatted as MMDDYYYY
            * sortOrder=[ascending, descending] - The sort order
* /customers/[CUSTOMER ID]/accounts/[ACCOUNT ID]/transactions/[TRANSACTION ID]
    * GET - Get customer account transaction detail
        * No Query Parameters
* /customers/[CUSTOMER ID]/accounts/[ACCOUNT ID]/orders
    * GET - List customer account orders
        * Optional Query Parameters:
            * symbol=[SYMBOL] - The symbol(s) for which to list orders. This parameter may be repeated to include up to 25 symbols (eg "?symbol=GOOG&symbol=AAPL")
            * fromDate=[MMDDYYYY] - The earliest date to include in the date range, formatted as MMDDYYYY. History is available for two years. If using fromDate, both fromDate and toDate should be provided and toDate should be greater than fromDate
            * toDate=[MMDDYYYY] - The latest date to include in the date range, formatted as MMDDYYYY. If using toDate, both fromDate and toDate should be provided, toDate should be greater than fromDate.
            * status=[open, executed, canceled, individualFills, cancelRequested, expired, rejected] - List only orders with this status
            * securityType=[equity, option, mutualFund, moneyMarketFund] - List only orders for securities of this type
            * transactionType=[extendedHours, buy, sell, short, buyToCover, mutualFundExchange] - List only orders with this transaction type
            * marketSession=[regular, extended] - The market session from which to return results
* /customers/[CUSTOMER ID]/alerts
    * GET - Get customer alert list
        * Optional Query Parameters:
            * count=[COUNT] - Maximum number of alerts to list
            * search=[SEARCH STRING] - Return alerts whose subjects include the search string
            * category=[stock, account] - Return only alerts pertaining to either stocks or the customer account
            * status=[read, unread, deleted] - Return only alerts with this status
            * sortOrder=[ascending, descending] - The sort order
* /customers/[CUSTOMER ID]/alerts/[ALERT ID]
    * GET - Get customer alert details
        * No Query Parameters
    * DELETE - Delete customer alert
        * No Query Parameters
* /customers/[CUSTOMER ID]/market/lookup
    * GET - Search for a company and get matching symbols
        * Required Query Parameters:
            * search=[SEARCH STRING] - Return symbols that match the search string (ETrade doesn't document how the search is performed)
* /customers/[CUSTOMER ID]/market/quote
    * GET - Get quotes for one or more symbols
        * Required Query Parameters:
            * symbol=[SYMBOL] - The symbol(s) for which to quote. This parameter may be repeated to include up to 25 symbols (eg "?symbol=GOOG&symbol=AAPL")
        * Optional Query Parameters:
            * detail=[all, fundamental, intraday, options, week52, mutualFund] - The quote detail to return (see [this page](https://apisb.etrade.com/docs/api/market/api-quote-v1.html#/definitions/QuoteData) for documentation on what's in the various detail types).
            * requireEarningsDate=[true, false] - If value is true, then nextEarningDate will be provided in the output. If value is false or if the field is not passed, nextEarningDate will be returned with no value.
            * skipMiniOptionsCheck=[true, false] - If value is true, no call is made to the service to check whether the symbol has mini options. If value is false or if the field is not specified, a service call is made to check if the symbol has mini options.
* /customers/[CUSTOMER ID]/market/optionchains
    * GET - Get option chains for a symbol
        * Required Query Parameters:
            * symbol=[SYMBOL] - The symbol for which to get option chains.
        * Optional Query Parameters:
            * expiryYear=[YYYY] - Fetch option chains with this expiry year
            * expiryMonth=[MM] - Fetch option chains with this expiry month (1-12)
            * expiryDay=[DD] - Fetch option chains with this expiry day (1-31)
            * strikePriceNear=[NEAR PRICE] - Return option chains with a strike price nearer to this integer value
            * noOfStrikes=[NUMBER OF STRIKES] - Return option chains with this many strikes
            * includeWeekly=[true, false] - Whether to include weekly options
            * skipAdjusted=[true, false] - Whether to skip adjusted options
            * optionCategory=[standard, all, mini] - Return only options in this category
            * chainType=[call, put, callPut] - Return only options with this chain type
            * priceType=[extendedHours, all] - Return only options with this price type
* /customers/[CUSTOMER ID]/market/optionexpire
    * GET - Get option expire dates for a symbol
        * Required Query Parameters:
            * symbol=[SYMBOL] - The symbol for which to get option chains.
        * Optional Query Parameters:
            * expiryType=[unspecified, daily, weekly, monthly, quarterly, vix, all, monthEnd] - Return only options with this expiration type
