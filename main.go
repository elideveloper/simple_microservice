package main

import (
	"net/http"
	"simple_microservice/service"
)

func main() {
	svc := service.NewService()
	eps := service.MakeEndpoints(svc)
	h := service.MakeHTTPHandler(eps)

	http.ListenAndServe(":8080", h)
}
