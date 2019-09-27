package controller_test

import (
	"encoding/json"
	"net/http/httptest"
	"park/project/2019_2_Premium-Harbor/component"
	"park/project/2019_2_Premium-Harbor/controller"
	"park/project/2019_2_Premium-Harbor/test"
	"sort"
	"testing"
)

var currencySuite = NewCurrencyControllerTestSuite()

func TestCurrencyRates(t *testing.T) {
	currencySuite.SetTesting(t)
	currencySuite.ExpectCurrencyRates([]component.Currency{
		{Name: "Доллар США", ShortName: "USD"},
		{Name: "Евро", ShortName: "EUR"},
	})
}

type CurrencyControllerTestSuite struct {
	test.ControllerTestSuite
	currencyController *controller.CurrencyController
}

func NewCurrencyControllerTestSuite() *CurrencyControllerTestSuite {
	return &CurrencyControllerTestSuite{
		currencyController: controller.NewCurrencyController(),
	}
}

func (s CurrencyControllerTestSuite) ExpectCurrencyRates(expectedCurrencies []component.Currency) {
	s.Request = httptest.NewRequest("GET", controller.ApiV1CurrencyRatesPath, nil)
	s.Response = httptest.NewRecorder()
	s.currencyController.HandleCurrencyRates(s.Response, s.Request)
	s.TestResponseStatus()
	responseBody, err := s.GetResponseBody()
	if err != nil {
		s.T.Error("invalid response")
		return
	}
	var currencies []component.Currency
	err = json.Unmarshal(*responseBody["currencies"], &currencies)
	if err != nil {
		s.T.Error("invalid response")
		return
	}
	sort.Slice(currencies, func(i, j int) bool {
		return currencies[i].ShortName < currencies[j].ShortName
	})
	sort.Slice(expectedCurrencies, func(i, j int) bool {
		return expectedCurrencies[i].ShortName < expectedCurrencies[j].ShortName
	})
	for i := range expectedCurrencies {
		expectedCurrency := expectedCurrencies[i]
		if i >= len(currencies) {
			s.T.Errorf("expected currency %v, but not found", expectedCurrency.ShortName)
			return
		}
		currency := currencies[i]
		if currency.ShortName != expectedCurrency.ShortName || currency.Name != expectedCurrency.Name {
			s.T.Errorf("expected %v (%v), \ngot %v (%v)", expectedCurrency.Name, expectedCurrency.ShortName, currency.Name, currency.ShortName)
		}
		if currency.Rate == 0 {
			s.T.Errorf("got zero %v rate", currency.ShortName)
		}
	}
}
