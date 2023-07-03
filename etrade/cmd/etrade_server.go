package cmd

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/client/constants"
	"github.com/jerryryle/etrade-cli/pkg/etradelib/jsonmap"
	"golang.org/x/exp/slog"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type eTradeServer struct {
	logger        *slog.Logger
	cfgFolder     ConfigurationFolder
	cfgStore      *CustomerConfigurationStore
	eTradeClients map[string]client.ETradeClient
}

func NewETradeServer(
	addr string, logger *slog.Logger, cfgFolder ConfigurationFolder, cfgStore *CustomerConfigurationStore,
) *http.Server {
	server := &eTradeServer{
		logger:        logger,
		cfgFolder:     cfgFolder,
		cfgStore:      cfgStore,
		eTradeClients: map[string]client.ETradeClient{},
	}

	r := chi.NewRouter()
	r.Get("/customers", server.GetCustomerList)
	r.Route(
		"/customers/{customerId}", func(r chi.Router) {
			r.Use(server.CustomerCtx)
			r.Post("/auth", server.Login)
			r.Delete("/auth", server.Logout)
			r.Get("/accounts", server.ListAccounts)
			r.Route(
				"/accounts/{accountId}", func(r chi.Router) {
					r.Get("/balance", server.GetAccountBalances)
					r.Get("/portfolio", server.ViewPortfolio)
					r.Get("/transactions", server.ListTransactions)
					r.Get("/transactions/{transactionId}", server.ListTransactionDetails)
					r.Get("/transactions/orders", server.ListOrders)
				},
			)
			r.Get("/alerts", server.ListAlerts)
			r.Route(
				"/alerts/{alertId}", func(r chi.Router) {
					r.Get("/", server.GetAlertDetails)
					r.Delete("/", server.DeleteAlert)
				},
			)
			r.Get("/market/lookup", server.Lookup)
			r.Get("/market/quote", server.GetQuote)
			r.Get("/market/optionchains", server.GetOptionChains)
			r.Get("/market/optionexpire", server.GetOptionExpire)
		},
	)
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

func (s *eTradeServer) CustomerCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			customerId := chi.URLParam(r, "customerId")
			eTradeClient, err := s.GetClientForCustomer(customerId)
			if err != nil {
				http.Error(w, http.StatusText(404), 404)
				return
			}
			ctx := context.WithValue(r.Context(), "eTradeClient", eTradeClient)
			next.ServeHTTP(w, r.WithContext(ctx))
		},
	)
}

func (s *eTradeServer) GetClientForCustomer(customerId string) (client.ETradeClient, error) {
	// See if there's already a cached client for this customerId
	if eTradeClient, ok := s.eTradeClients[customerId]; ok {
		return eTradeClient, nil
	}
	// If there's not a cached client, create a new one
	if eTradeClient, err := NewETradeClientForCustomer(customerId, s.cfgFolder, s.cfgStore, s.logger); err == nil {
		// Add the new client to the cache and return it
		s.eTradeClients[customerId] = eTradeClient
		return eTradeClient, nil
	} else {
		return nil, err
	}
}

func (s *eTradeServer) RemoveClientForCustomer(customerId string) {
	delete(s.eTradeClients, customerId)
}

func (s *eTradeServer) GetCustomerList(w http.ResponseWriter, _ *http.Request) {
	if responseMap, err := GetCustomerList(s.cfgStore); err == nil {
		s.WriteJsonMap(w, responseMap)
	} else {
		s.WriteError(w, err)
	}
}

func (s *eTradeServer) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		s.WriteError(w, err)
		return
	}
	eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient)
	if !ok {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
		return
	}

	if !r.Form.Has("verifyCode") {
		// If the form does not include "verifyCode" then begin authentication.
		if authUrl, err := eTradeClient.Authenticate(); err != nil {
			// Authentication failed. Respond with the error.
			s.WriteError(w, err)
			return
		} else {
			if authUrl != "" {
				// Authentication succeeded, but a verification code is needed.
				// Respond with the URL for getting the code.
				s.WriteJsonMap(w, NewStatusResponseMap("verify", "verifyUrl", authUrl))
				return
			}
		}
	} else {
		// If the form does not include "verifyCode" then perform verification.
		if err = eTradeClient.Verify(r.Form.Get("verifyCode")); err != nil {
			// Verification failed. Respond with the error.
			s.WriteError(w, err)
			return
		}
	}

	// If we get here, then authentication has succeeded.
	// Update the credential cache
	consumerKey, _, accessToken, accessSecret := eTradeClient.GetKeys()
	if err = s.cfgFolder.SaveCachedCredentialsToFile(
		consumerKey, &CachedCredentials{accessToken, accessSecret, time.Now()}, s.logger,
	); err != nil {
		s.logger.Error(err.Error())
	}
	// Authentication succeeded. Respond with success.
	s.WriteJsonMap(w, NewStatusResponseMap("success"))
}

func (s *eTradeServer) Logout(w http.ResponseWriter, r *http.Request) {
	customerId := chi.URLParam(r, "customerId")
	// Remove cached ETradeClient
	s.RemoveClientForCustomer(customerId)
	// Remove credential cache
	if response, err := ClearAuth(
		customerId, s.cfgFolder, s.cfgStore,
	); err == nil {
		s.WriteJsonMap(w, response)
	} else {
		s.WriteError(w, err)
	}
}

func (s *eTradeServer) ListAccounts(w http.ResponseWriter, r *http.Request) {
	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := ListAccounts(eTradeClient); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) GetAccountBalances(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountId")
	realTimeBalance := getBoolWithDefaultFromValues(r.URL.Query(), "realTimeBalance", true)

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := GetAccountBalances(eTradeClient, accountId, realTimeBalance); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) ViewPortfolio(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountId")

	withLots := getBoolWithDefaultFromValues(r.URL.Query(), "withLots", false)
	totalsRequired := getBoolWithDefaultFromValues(r.URL.Query(), "totalsRequired", true)
	portfolioView := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "view", portfolioViewMap, constants.PortfolioViewQuick,
	)
	sortBy := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "sortBy", portfolioSortByMap, constants.PortfolioSortByNil,
	)
	sortOrder := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "sortOrder", sortOrderMap, constants.SortOrderNil,
	)
	marketSession := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "marketSession", marketSessionMap, constants.MarketSessionNil,
	)

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := ViewPortfolio(
			eTradeClient, accountId, sortBy, sortOrder, marketSession, totalsRequired, portfolioView, withLots,
		); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) ListTransactions(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountId")

	startDate := getDateWithDefaultFromValues(r.URL.Query(), "startDate", "01022006", nil)
	endDate := getDateWithDefaultFromValues(r.URL.Query(), "endDate", "01022006", nil)
	sortOrder := getEnumFlagWithDefaultFromValues(r.URL.Query(), "sortOrder", sortOrderMap, constants.SortOrderNil)

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := ListTransactions(eTradeClient, accountId, startDate, endDate, sortOrder); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) ListTransactionDetails(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountId")
	transactionId := chi.URLParam(r, "transactionId")

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := ListTransactionDetails(eTradeClient, accountId, transactionId); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) ListOrders(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "accountId")
	symbols := r.URL.Query()["symbols"]

	fromDate := getDateWithDefaultFromValues(r.URL.Query(), "fromDate", "01022006", nil)
	toDate := getDateWithDefaultFromValues(r.URL.Query(), "toDate", "01022006", nil)
	status := getEnumFlagWithDefaultFromValues(r.URL.Query(), "status", orderStatusMap, constants.OrderStatusNil)
	securityType := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "securityType", orderSecurityTypeMap, constants.OrderSecurityTypeNil,
	)
	transactionType := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "transactionType", orderTransactionTypeMap, constants.OrderTransactionTypeNil,
	)
	marketSession := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "marketSession", marketSessionMap, constants.MarketSessionNil,
	)

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := ListOrders(
			eTradeClient, accountId, status, fromDate, toDate, symbols, securityType, transactionType, marketSession,
		); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) ListAlerts(w http.ResponseWriter, r *http.Request) {
	count := getIntWithDefaultFromValues(r.URL.Query(), "search", -1)
	search := getStringWithDefaultFromValues(r.URL.Query(), "search", "")
	category := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "category", alertCategoryMap, constants.AlertCategoryNil,
	)
	status := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "status", alertStatusMap, constants.AlertStatusNil,
	)
	sortOrder := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "sortOrder", sortOrderMap, constants.SortOrderNil,
	)

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := ListAlerts(eTradeClient, count, category, status, sortOrder, search); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) GetAlertDetails(w http.ResponseWriter, r *http.Request) {
	alertId := chi.URLParam(r, "alertId")

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := ListAlertDetails(eTradeClient, alertId); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) DeleteAlert(w http.ResponseWriter, r *http.Request) {
	alertId := chi.URLParam(r, "alertId")

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := DeleteAlerts(eTradeClient, []string{alertId}); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) Lookup(w http.ResponseWriter, r *http.Request) {
	search := getStringWithDefaultFromValues(r.URL.Query(), "search", "")
	if search == "" {
		s.WriteError(w, errors.New("missing search query"))
		return
	}

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := Lookup(eTradeClient, search); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) GetQuote(w http.ResponseWriter, r *http.Request) {
	symbols := r.URL.Query()["symbol"]

	detail := getEnumFlagWithDefaultFromValues(r.URL.Query(), "detail", quoteDetailMap, constants.QuoteDetailAll)
	requireEarningsDate := getBoolWithDefaultFromValues(r.URL.Query(), "requireEarningsDate", true)
	skipMiniOptionsCheck := getBoolWithDefaultFromValues(r.URL.Query(), "skipMiniOptionsCheck", false)

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := GetQuotes(
			eTradeClient, symbols, detail, requireEarningsDate, skipMiniOptionsCheck,
		); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) GetOptionChains(w http.ResponseWriter, r *http.Request) {
	symbol := getStringWithDefaultFromValues(r.URL.Query(), "symbol", "")
	if symbol == "" {
		s.WriteError(w, errors.New("missing symbol"))
		return
	}

	expiryYear := getIntWithDefaultFromValues(r.URL.Query(), "expiryYear", -1)
	expiryMonth := getIntWithDefaultFromValues(r.URL.Query(), "expiryMonth", -1)
	expiryDay := getIntWithDefaultFromValues(r.URL.Query(), "expiryDay", -1)
	strikePriceNear := getIntWithDefaultFromValues(r.URL.Query(), "strikePriceNear", -1)
	noOfStrikes := getIntWithDefaultFromValues(r.URL.Query(), "noOfStrikes", -1)
	includeWeekly := getBoolWithDefaultFromValues(r.URL.Query(), "includeWeekly", true)
	skipAdjusted := getBoolWithDefaultFromValues(r.URL.Query(), "skipAdjusted", false)
	optionCategory := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "optionCategory", optionCategoryMap, constants.OptionCategoryNil,
	)
	chainType := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "chainType", optionChainTypeMap, constants.OptionChainTypeNil,
	)
	priceType := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "chainType", optionPriceTypeMap, constants.OptionPriceTypeNil,
	)

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := GetOptionChains(
			eTradeClient, symbol, expiryYear, expiryMonth, expiryDay, strikePriceNear, noOfStrikes, includeWeekly,
			skipAdjusted, optionCategory, chainType, priceType,
		); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) GetOptionExpire(w http.ResponseWriter, r *http.Request) {
	symbol := getStringWithDefaultFromValues(r.URL.Query(), "symbol", "")
	if symbol == "" {
		s.WriteError(w, errors.New("missing symbol"))
		return
	}
	expiryType := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "expiryType", optionExpiryTypeMap, constants.OptionExpiryTypeNil,
	)

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := GetOptionExpireDates(eTradeClient, symbol, expiryType); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) WriteJsonMap(w http.ResponseWriter, jsonMap jsonmap.JsonMap) {
	responseBytes, err := jsonMap.ToJsonBytes(false, false)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(responseBytes); err != nil {
		s.logger.Error(err.Error())
	}
}

func (s *eTradeServer) WriteError(w http.ResponseWriter, err error) {
	s.logger.Error(err.Error())
	responseMap := NewStatusResponseMap("error", "error", err.Error())
	responseBytes, err := responseMap.ToJsonBytes(false, false)
	if err != nil {
		s.logger.Error(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	if _, err = w.Write(responseBytes); err != nil {
		s.logger.Error(err.Error())
	}
}

func NewStatusResponseMap(status string, keysAndValues ...string) jsonmap.JsonMap {
	responseMap := jsonmap.JsonMap{
		"status": status,
	}
	if len(keysAndValues)%2 == 0 {
		for i := 0; i < len(keysAndValues); i += 2 {
			key := keysAndValues[i]
			value := keysAndValues[i+1]
			_ = responseMap.SetString(key, value)
		}
	}

	return responseMap
}

func getStringWithDefaultFromValues(v url.Values, key string, defaultValue string) string {
	if !v.Has(key) {
		return defaultValue
	}
	return v.Get(key)
}

func getIntWithDefaultFromValues(v url.Values, key string, defaultValue int) int {
	if !v.Has(key) {
		return defaultValue
	}
	stringValue := v.Get(key)
	if value, err := strconv.Atoi(stringValue); err != nil {
		return value
	} else {
		return defaultValue
	}
}

func getBoolWithDefaultFromValues(v url.Values, key string, defaultValue bool) bool {
	if !v.Has(key) {
		return defaultValue
	}
	stringValue := v.Get(key)
	if value, err := strconv.ParseBool(stringValue); err == nil {
		return value
	} else {
		return defaultValue
	}
}

func getDateWithDefaultFromValues(v url.Values, key string, layout string, defaultValue *time.Time) *time.Time {
	if !v.Has(key) {
		return defaultValue
	}
	stringValue := v.Get(key)
	if value, err := time.Parse(layout, stringValue); err == nil {
		return &value
	} else {
		return defaultValue
	}
}

func getEnumFlagWithDefaultFromValues[T comparable](
	v url.Values, key string, enumMap enumValueWithHelpMap[T], defaultValue T,
) T {
	if !v.Has(key) {
		return defaultValue
	}
	return enumMap.GetEnumValueWithDefault(v.Get(key), defaultValue)
}
