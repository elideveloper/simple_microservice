package service

import (
	"context"
	"simple_microservice/requests"
	//"simple_microservice/responses"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Get    endpoint.Endpoint
	Add    endpoint.Endpoint
	Delete endpoint.Endpoint
	Update endpoint.Endpoint
}

func MakeEndpoints(svc ExampleServer) *Endpoints {
	eps := &Endpoints{}
	eps.Add = func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(requests.Add)
		return svc.Add(&r)
	}
	eps.Get = func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(requests.Get)
		return svc.Get(&r)
	}
	eps.Update = func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(requests.Update)
		return svc.Update(&r)
	}
	eps.Delete = func(ctx context.Context, req interface{}) (interface{}, error) {
		r := req.(requests.Delete)
		return svc.Delete(&r)
	}
	return eps
}