package md

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Tomasen/realip"
	gkitLog "github.com/go-kit/kit/log"
	opentracing "github.com/opentracing/opentracing-go"
	opentracingLog "github.com/opentracing/opentracing-go/log"

	"aegis/pkg/kit/trace"
)

func LogRequestDuration(logger gkitLog.Logger, requestEndpoint string) Middleware {
	spanOperation := fmt.Sprintf("api.%s", requestEndpoint)
	return func(next http.Handler) http.Handler {
		f := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			span, ctx := opentracing.StartSpanFromContext(ctx, spanOperation)
			begin := time.Now()

			defer func() {
				reqip := realip.FromRequest(r)
				durationInMillisecond := time.Since(begin).Nanoseconds() / 1e6

				span.LogFields(
					opentracingLog.String("req_ip", reqip),
					opentracingLog.Int64("duration_in_millisecond", durationInMillisecond),
				)
				span.Finish()

				_ = logger.Log(
					"request_endpoint", requestEndpoint,
					"duration_in_millisecond", durationInMillisecond,
					"req_ip", reqip,
				)
			}()
			w.Header().Add("X-Aegis-Trace", trace.ExtractTraceId(span))
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		return http.HandlerFunc(f)
	}
}
