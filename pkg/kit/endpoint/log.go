package endpoint

import (
	"context"
	"time"

	gkitEndpoint "github.com/go-kit/kit/endpoint"
	gkitLog "github.com/go-kit/kit/log"
)

func LogEndpointDuration(logger gkitLog.Logger, requestEndpoint string) gkitEndpoint.Middleware {
	return func(next gkitEndpoint.Endpoint) gkitEndpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			begin := time.Now()
			resp, err := next(ctx, request)
			logger.Log(
				"request_endpoint", requestEndpoint,
				"duration_in_millisecond", time.Since(begin).Nanoseconds()/1e6,
			)
			return resp, err
		}
	}
}
