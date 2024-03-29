package controller

import (
	"github.com/go-park-mail-ru/2019_2_Premium-Harbor/component"
	"net/http"
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
