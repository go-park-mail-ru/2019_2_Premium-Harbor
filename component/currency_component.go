package component

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type CurrencyComponent struct {
}

func NewCurrencyComponent() *CurrencyComponent {
	return &CurrencyComponent{}
}

type Currency struct {
	Name      string  `json:"name"`
	ShortName string  `json:"short_name"`
	Rate      float32 `json:"rate"`
}

type exchangeRatesAPIResponse struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float32 `json:"rates"`
}

func (c CurrencyComponent) GetCurrencyRates() ([]Currency, error) {
	ratesAPIResponse, err := http.Get("https://api.exchangeratesapi.io/latest?base=RUB")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	ratesJson, err := ioutil.ReadAll(ratesAPIResponse.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	rates := exchangeRatesAPIResponse{}
	err = json.Unmarshal(ratesJson, &rates)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	currencies := c.getSupportedCurrencies()
	for i, currency := range currencies {
		if rate, ok := rates.Rates[currency.ShortName]; ok {
			currencies[i].Rate = 1 / rate
		}
	}
	return currencies, nil
}

func (c CurrencyComponent) getSupportedCurrencies() []Currency {
	return []Currency{
		{Name: "Доллар США", ShortName: "USD"},
		{Name: "Евро", ShortName: "EUR"},
	}
}
