package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jerryryle/etrade-cli/pkg/etradelib"
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
	responseMap := GetCustomerList(s.cfgStore)
	s.WriteJsonMap(w, responseMap)
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
		response, err := eTradeClient.Authenticate()
		if err != nil {
			s.WriteError(w, err)
			return
		}
		authStatus, err := etradelib.CreateETradeAuthenticationStatusFromResponse(response)
		if err != nil {
			s.WriteError(w, err)
			return
		}
		if !authStatus.NeedAuthorization() {
			// Authentication has succeeded, so update the credential cache
			consumerKey, _, accessToken, accessSecret := eTradeClient.GetKeys()
			if err = s.cfgFolder.SaveCachedCredentialsToFile(
				consumerKey, &CachedCredentials{accessToken, accessSecret, time.Now()}, s.logger,
			); err != nil {
				s.logger.Error(fmt.Errorf("saving credential cache to file failed (%w)", err).Error())
			}
		}
		// Respond with the authentication status. If authentication requires
		// authorization, then this status will contain the authorization URL.
		s.WriteJsonMap(w, authStatus.AsJsonMap())
		return
	} else {
		// If the form includes "verifyCode" then perform verification.
		response, err := eTradeClient.Verify(r.Form.Get("verifyCode"))
		if err != nil {
			// Verification failed. Respond with the error.
			s.WriteError(w, err)
			return
		}
		verifyStatus, err := etradelib.CreateETradeStatusFromResponse(response)
		if err != nil {
			s.WriteError(w, err)
			return
		}

		// Verification has succeeded, so update the credential cache
		consumerKey, _, accessToken, accessSecret := eTradeClient.GetKeys()
		if err = s.cfgFolder.SaveCachedCredentialsToFile(
			consumerKey, &CachedCredentials{accessToken, accessSecret, time.Now()}, s.logger,
		); err != nil {
			s.logger.Error(fmt.Errorf("saving credential cache to file failed (%w)", err).Error())
		}
		// Respond with the verification status.
		s.WriteJsonMap(w, verifyStatus.AsJsonMap())
	}
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
	realTimeBalance, err := getBoolWithDefaultFromValues(r.URL.Query(), "realTimeBalance", true)
	if err != nil {
		s.WriteError(w, err)
		return
	}

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

	withLots, err := getBoolWithDefaultFromValues(r.URL.Query(), "withLots", false)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	totalsRequired, err := getBoolWithDefaultFromValues(r.URL.Query(), "totalsRequired", true)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	portfolioView, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "view", portfolioViewMap, constants.PortfolioViewQuick,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	sortBy, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "sortBy", portfolioSortByMap, constants.PortfolioSortByNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	sortOrder, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "sortOrder", sortOrderMap, constants.SortOrderNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	marketSession, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "marketSession", marketSessionMap, constants.MarketSessionNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

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

	startDate, err := getDateWithDefaultFromValues(r.URL.Query(), "startDate", "01022006", nil)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	endDate, err := getDateWithDefaultFromValues(r.URL.Query(), "endDate", "01022006", nil)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	sortOrder, err := getEnumFlagWithDefaultFromValues(r.URL.Query(), "sortOrder", sortOrderMap, constants.SortOrderNil)
	if err != nil {
		s.WriteError(w, err)
		return
	}

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
	symbols := r.URL.Query()["symbol"]

	fromDate, err := getDateWithDefaultFromValues(r.URL.Query(), "fromDate", "01022006", nil)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	toDate, err := getDateWithDefaultFromValues(r.URL.Query(), "toDate", "01022006", nil)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	status, err := getEnumFlagWithDefaultFromValues(r.URL.Query(), "status", orderStatusMap, constants.OrderStatusNil)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	securityType, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "securityType", orderSecurityTypeMap, constants.OrderSecurityTypeNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	transactionType, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "transactionType", orderTransactionTypeMap, constants.OrderTransactionTypeNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	marketSession, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "marketSession", marketSessionMap, constants.MarketSessionNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

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
	count, err := getIntWithDefaultFromValues(r.URL.Query(), "count", -1)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	search := getStringWithDefaultFromValues(r.URL.Query(), "search", "")

	category, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "category", alertCategoryMap, constants.AlertCategoryNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	status, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "status", alertStatusMap, constants.AlertStatusNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	sortOrder, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "sortOrder", sortOrderMap, constants.SortOrderNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

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

	detail, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "detail", quoteDetailMap, constants.QuoteDetailFlagAll,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	requireEarningsDate, err := getBoolWithDefaultFromValues(r.URL.Query(), "requireEarningsDate", true)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	skipMiniOptionsCheck, err := getBoolWithDefaultFromValues(r.URL.Query(), "skipMiniOptionsCheck", false)
	if err != nil {
		s.WriteError(w, err)
		return
	}

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

	expiryYear, err := getIntWithDefaultFromValues(r.URL.Query(), "expiryYear", -1)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	expiryMonth, err := getIntWithDefaultFromValues(r.URL.Query(), "expiryMonth", -1)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	expiryDay, err := getIntWithDefaultFromValues(r.URL.Query(), "expiryDay", -1)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	strikePriceNear, err := getIntWithDefaultFromValues(r.URL.Query(), "strikePriceNear", -1)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	noOfStrikes, err := getIntWithDefaultFromValues(r.URL.Query(), "noOfStrikes", -1)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	includeWeekly, err := getBoolWithDefaultFromValues(r.URL.Query(), "includeWeekly", true)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	skipAdjusted, err := getBoolWithDefaultFromValues(r.URL.Query(), "skipAdjusted", false)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	optionCategory, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "optionCategory", optionCategoryMap, constants.OptionCategoryNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	chainType, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "chainType", optionChainTypeMap, constants.OptionChainTypeNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

	priceType, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "priceType", optionPriceTypeMap, constants.OptionPriceTypeNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

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
	expiryType, err := getEnumFlagWithDefaultFromValues(
		r.URL.Query(), "expiryType", optionExpiryTypeMap, constants.OptionExpiryTypeNil,
	)
	if err != nil {
		s.WriteError(w, err)
		return
	}

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
		s.logger.Error(fmt.Errorf("marshaling JSON response failed (%w)", err).Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(responseBytes); err != nil {
		s.logger.Error(fmt.Errorf("writing JSON response failed (%w)", err).Error())
	}
}

func (s *eTradeServer) WriteError(w http.ResponseWriter, err error) {
	s.logger.Error(fmt.Errorf("server encountered an error processing request (%w)", err).Error())
	responseMap := client.NewStatusMap("error", "error", err.Error())
	responseBytes, err := responseMap.ToJsonBytes(false, false)
	if err != nil {
		s.logger.Error(fmt.Errorf("marshaling JSON error response failed (%w)", err).Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	if _, err = w.Write(responseBytes); err != nil {
		s.logger.Error(fmt.Errorf("writing JSON error response failed (%w)", err).Error())
	}
}

func getStringWithDefaultFromValues(v url.Values, key string, defaultValue string) string {
	if !v.Has(key) {
		return defaultValue
	}
	return v.Get(key)
}

func getIntWithDefaultFromValues(v url.Values, key string, defaultValue int) (int, error) {
	if !v.Has(key) {
		return defaultValue, nil
	}
	stringValue := v.Get(key)
	if value, err := strconv.Atoi(stringValue); err != nil {
		return value, nil
	} else {
		return 0, fmt.Errorf("%s is not a valid integer (%w)", stringValue, err)
	}
}

func getBoolWithDefaultFromValues(v url.Values, key string, defaultValue bool) (bool, error) {
	if !v.Has(key) {
		return defaultValue, nil
	}
	stringValue := v.Get(key)
	if value, err := strconv.ParseBool(stringValue); err == nil {
		return value, nil
	} else {
		return false, fmt.Errorf("%s is not a valid boolean (%w)", stringValue, err)
	}
}

func getDateWithDefaultFromValues(v url.Values, key string, layout string, defaultValue *time.Time) (
	*time.Time, error,
) {
	if !v.Has(key) {
		return defaultValue, nil
	}
	stringValue := v.Get(key)
	if value, err := time.Parse(layout, stringValue); err == nil {
		return &value, nil
	} else {
		return nil, fmt.Errorf("%s is not a valid time/date (%w)", stringValue, err)
	}
}

func getEnumFlagWithDefaultFromValues[T comparable](
	v url.Values, key string, enumMap enumValueWithHelpMap[T], defaultValue T,
) (T, error) {
	if !v.Has(key) {
		return defaultValue, nil
	}
	enumString := v.Get(key)
	if value, err := enumMap.GetEnumValue(enumString); err == nil {
		return value, nil
	} else {
		var retVal T
		return retVal, fmt.Errorf("%s is not a valid value for %s (%w)", enumString, key, err)
	}
}
