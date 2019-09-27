package controller

import (
	"github.com/gorilla/mux"
	"net/http"
)

const (
	ApiV1UserListPath     = "/api/v1/user/list"
	ApiV1UserUpdatePath   = "/api/v1/user/update"
	ApiV1UserRegisterPath = "/api/v1/user/register"
	ApiV1UserLoginPath    = "/api/v1/user/login"
	ApiV1UserLogoutPath   = "/api/v1/user/logout"

	ApiV1CurrencyRatesPath = "/api/v1/currency/rates"
)

func InitAPIRouter() *mux.Router {
	r := mux.NewRouter()

	userController := NewUserController()
	r.HandleFunc(ApiV1UserListPath, userController.HandleUserList).Methods(http.MethodGet)
	r.HandleFunc(ApiV1UserUpdatePath, userController.HandleUserUpdate).Methods(http.MethodPost)
	r.HandleFunc(ApiV1UserRegisterPath, userController.HandleUserRegister).Methods(http.MethodPost)
	r.HandleFunc(ApiV1UserLoginPath, userController.HandleUserLogin).Methods(http.MethodPost)
	r.HandleFunc(ApiV1UserLogoutPath, userController.HandleUserLogout).Methods(http.MethodPost)

	currencyController := NewCurrencyController()
	r.HandleFunc(ApiV1CurrencyRatesPath, currencyController.HandleCurrencyRates).Methods(http.MethodGet)

	return r
}
