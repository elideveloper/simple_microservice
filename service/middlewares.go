package service

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
)

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func LoggingMiddleware() Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			resp, err := next(ctx, request)
			fmt.Println(err)
			return resp, err
		}
	}
}
