package main

import (
	"context"

	"connectrpc.com/connect"
	"github.com/rs/zerolog"
)

func ServiceVersionInterceptor(serviceName, version string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			res, err := next(ctx, req)
			if err != nil {
				return nil, err
			}
			res.Header().Set(serviceName+"-Version", version)
			return res, nil
		}
	}
}

func LoggingInterceptor(logger *zerolog.Logger) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			method := req.Spec().Procedure
			logger.Debug().Str("method", method).Msg("request started")

			res, err := next(ctx, req)

			if err != nil {
				logger.Error().
					Str("method", method).
					Err(err).
					Msg("request failed")
			} else {
				logger.Debug().
					Str("method", method).
					Msg("request completed")
			}

			return res, err
		}
	}
}
