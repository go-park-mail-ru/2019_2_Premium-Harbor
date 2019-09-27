package main

import (
	"net/http"
	"premium-harbor/controller"
)

func main() {
	r := controller.InitAPIRouter()
	http.ListenAndServe(":8080", r)
}
