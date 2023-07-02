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
	realTimeBalance := r.URL.Query().Get("realTimeBalance") == "true"

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

	withLots := r.URL.Query().Get("withLots") == "true"
	totalsRequired := r.URL.Query().Get("totalsRequired") == "true"

	portfolioView := *newEnumFlagValue(portfolioViewMap, constants.PortfolioViewQuick)
	if err := portfolioView.Set(r.URL.Query().Get("view")); err != nil {
		s.WriteError(w, err)
		return
	}

	sortBy := *newEnumFlagValue(portfolioSortByMap, constants.PortfolioSortByNil)
	if err := sortBy.Set(r.URL.Query().Get("sortBy")); err != nil {
		s.WriteError(w, err)
		return
	}

	sortOrder := *newEnumFlagValue(sortOrderMap, constants.SortOrderNil)
	if err := sortOrder.Set(r.URL.Query().Get("sortOrder")); err != nil {
		s.WriteError(w, err)
		return
	}

	marketSession := *newEnumFlagValue(marketSessionMap, constants.MarketSessionNil)
	if err := marketSession.Set(r.URL.Query().Get("marketSession")); err != nil {
		s.WriteError(w, err)
		return
	}

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := ViewPortfolio(
			eTradeClient, accountId, sortBy.Value(), sortOrder.Value(), marketSession.Value(), totalsRequired,
			portfolioView.Value(), withLots,
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

	var startDate, endDate *time.Time = nil, nil
	startDateString := r.URL.Query().Get("startDate")
	endDateString := r.URL.Query().Get("endDate")
	if startDateString != "" {
		var err error
		*startDate, err = time.Parse("01022006", startDateString)
		if err != nil {
			s.WriteError(w, errors.New("start date must be in format MMDDYYYY"))
			return
		}
	}
	if endDateString != "" {
		var err error
		*endDate, err = time.Parse("01022006", endDateString)
		if err != nil {
			s.WriteError(w, errors.New("end date must be in format MMDDYYYY"))
			return
		}
	}

	sortOrder := *newEnumFlagValue(sortOrderMap, constants.SortOrderNil)
	if err := sortOrder.Set(r.URL.Query().Get("sortOrder")); err != nil {
		s.WriteError(w, err)
		return
	}

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := ListTransactions(
			eTradeClient, accountId, startDate, endDate, sortOrder.Value(),
		); err == nil {
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

	var fromDate, toDate *time.Time = nil, nil
	fromDateString := r.URL.Query().Get("fromDate")
	toDateString := r.URL.Query().Get("toDate")
	if fromDateString != "" {
		var err error
		*fromDate, err = time.Parse("01022006", fromDateString)
		if err != nil {
			s.WriteError(w, errors.New("start date must be in format MMDDYYYY"))
			return
		}
	}
	if toDateString != "" {
		var err error
		*toDate, err = time.Parse("01022006", toDateString)
		if err != nil {
			s.WriteError(w, errors.New("end date must be in format MMDDYYYY"))
			return
		}
	}

	status := *newEnumFlagValue(orderStatusMap, constants.OrderStatusNil)
	if err := status.Set(r.URL.Query().Get("status")); err != nil {
		s.WriteError(w, err)
		return
	}

	securityType := *newEnumFlagValue(orderSecurityTypeMap, constants.OrderSecurityTypeNil)
	if err := securityType.Set(r.URL.Query().Get("securityType")); err != nil {
		s.WriteError(w, err)
		return
	}

	transactionType := *newEnumFlagValue(orderTransactionTypeMap, constants.OrderTransactionTypeNil)
	if err := transactionType.Set(r.URL.Query().Get("transactionType")); err != nil {
		s.WriteError(w, err)
		return
	}

	marketSession := *newEnumFlagValue(marketSessionMap, constants.MarketSessionNil)
	if err := marketSession.Set(r.URL.Query().Get("marketSession")); err != nil {
		s.WriteError(w, err)
		return
	}

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := ListOrders(
			eTradeClient, accountId, status.Value(), fromDate, toDate, symbols, securityType.Value(),
			transactionType.Value(), marketSession.Value(),
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
	count, err := strconv.Atoi(r.URL.Query().Get("count"))
	if err != nil {
		s.WriteError(w, err)
		return
	}

	search := r.URL.Query().Get("search")

	category := *newEnumFlagValue(alertCategoryMap, constants.AlertCategoryNil)
	if err = category.Set(r.URL.Query().Get("category")); err != nil {
		s.WriteError(w, err)
		return
	}

	status := *newEnumFlagValue(alertStatusMap, constants.AlertStatusNil)
	if err = status.Set(r.URL.Query().Get("status")); err != nil {
		s.WriteError(w, err)
		return
	}

	sortOrder := *newEnumFlagValue(sortOrderMap, constants.SortOrderNil)
	if err = sortOrder.Set(r.URL.Query().Get("sortOrder")); err != nil {
		s.WriteError(w, err)
		return
	}

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := ListAlerts(
			eTradeClient, count, category.Value(), status.Value(), sortOrder.Value(), search,
		); err == nil {
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
	search := r.URL.Query().Get("search")

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

	detail := *newEnumFlagValue(quoteDetailMap, constants.QuoteDetailAll)
	if err := detail.Set(r.URL.Query().Get("detail")); err != nil {
		s.WriteError(w, err)
		return
	}

	requireEarningsDate := r.URL.Query().Get("requireEarningsDate") == "true"
	skipMiniOptionsCheck := r.URL.Query().Get("skipMiniOptionsCheck") == "true"

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := GetQuotes(
			eTradeClient, symbols, detail.Value(), requireEarningsDate, skipMiniOptionsCheck,
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
	//TODO: Fix this and other methods. Need ability to get value with default if it doesn't exist.
	symbol := r.URL.Query().Get("symbol")

	expiryYear, err := strconv.Atoi(r.URL.Query().Get("expiryYear"))
	if err != nil {
		s.WriteError(w, err)
		return
	}
	expiryMonth, err := strconv.Atoi(r.URL.Query().Get("expiryMonth"))
	if err != nil {
		s.WriteError(w, err)
		return
	}
	expiryDay, err := strconv.Atoi(r.URL.Query().Get("expiryDay"))
	if err != nil {
		s.WriteError(w, err)
		return
	}
	strikePriceNear, err := strconv.Atoi(r.URL.Query().Get("strikePriceNear"))
	if err != nil {
		s.WriteError(w, err)
		return
	}
	noOfStrikes, err := strconv.Atoi(r.URL.Query().Get("noOfStrikes"))
	if err != nil {
		s.WriteError(w, err)
		return
	}

	includeWeekly := r.URL.Query().Get("includeWeekly") == "true"
	skipAdjusted := r.URL.Query().Get("skipAdjusted") == "true"

	optionCategory := *newEnumFlagValue(optionCategoryMap, constants.OptionCategoryNil)
	if err = optionCategory.Set(r.URL.Query().Get("optionCategory")); err != nil {
		s.WriteError(w, err)
		return
	}

	chainType := *newEnumFlagValue(optionChainTypeMap, constants.OptionChainTypeNil)
	if err = chainType.Set(r.URL.Query().Get("chainType")); err != nil {
		s.WriteError(w, err)
		return
	}

	priceType := *newEnumFlagValue(optionPriceTypeMap, constants.OptionPriceTypeNil)
	if err = priceType.Set(r.URL.Query().Get("priceType")); err != nil {
		s.WriteError(w, err)
		return
	}

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := GetOptionChains(
			eTradeClient, symbol, expiryYear, expiryMonth, expiryDay,
			strikePriceNear, noOfStrikes, includeWeekly, skipAdjusted,
			optionCategory.Value(), chainType.Value(), priceType.Value(),
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
	symbol := r.URL.Query().Get("symbol")

	expiryType := *newEnumFlagValue(optionExpiryTypeMap, constants.OptionExpiryTypeNil)
	if err := expiryType.Set(r.URL.Query().Get("expiryType")); err != nil {
		s.WriteError(w, err)
		return
	}

	if eTradeClient, ok := r.Context().Value("eTradeClient").(client.ETradeClient); ok {
		if response, err := GetOptionExpireDates(eTradeClient, symbol, expiryType.Value()); err == nil {
			s.WriteJsonMap(w, response)
		} else {
			s.WriteError(w, err)
		}
	} else {
		s.WriteError(w, errors.New("unable to find ETrade client for customer"))
	}
}

func (s *eTradeServer) WriteJsonMap(w http.ResponseWriter, jsonMap jsonmap.JsonMap) {
	responseBytes, err := jsonMap.ToJsonBytes(false)
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
	responseBytes, err := responseMap.ToJsonBytes(false)
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
