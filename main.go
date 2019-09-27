package main

import (
	"net/http"
	"os"
	"premium-harbor/controller"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r := controller.InitAPIRouter()
	http.ListenAndServe(":"+port, r)
}
