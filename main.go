package main

import (
	"net/http"
	"os"

	"github.com/elideveloper/simple_microservice/service"

	"github.com/go-kit/kit/log"
)

func main() {
	logger := log.NewJSONLogger(os.Stdout)
	svc := service.NewService()
	eps := service.MakeEndpoints(svc, logger)
	h := service.MakeHTTPHandler(eps)

	http.ListenAndServe(":8080", h)
}
