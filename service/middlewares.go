package service

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Middleware func(endpoint.Endpoint) endpoint.Endpoint

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (resp interface{}, err error) {
			defer func() {
				r := recover()
				if r != nil {
					logger.Log("panic:", r)
				} else if err != nil {
					logger.Log("error:", err)
				}
			}()
			return next(ctx, request)
		}
	}
}
