package controller

import (
	"net/http"
	"premium-harbor/component"
)

type CurrencyController struct {
	Controller
	currencyComponent *component.CurrencyComponent
}

func NewCurrencyController() *CurrencyController {
	return &CurrencyController{
		currencyComponent: component.NewCurrencyComponent(),
	}
}

func (c CurrencyController) HandleCurrencyRates(w http.ResponseWriter, r *http.Request) {
	c.writeCommonHeaders(w)
	currencies, err := c.currencyComponent.GetCurrencyRates()
	if err != nil {
		c.writeError(w, err)
		return
	}
	c.writeOkWithBody(w, map[string]interface{}{
		"currencies": currencies,
	})
}
