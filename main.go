package main

import (
	"net/http"
	"park/project/2019_2_Premium-Harbor/controller"
)

func main() {
	r := controller.InitAPIRouter()
	http.ListenAndServe(":8080", r)
}
