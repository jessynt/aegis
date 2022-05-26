package md

import (
	"net/http"
	"runtime/debug"

	gkitLog "github.com/go-kit/kit/log"

	"aegis/internal/response"
)

// 捕捉 panic 异常
func CapturePanic(printStack, returnStack bool, logger gkitLog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		f := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				rval := recover()
				if rval == nil {
					return
				}

				switch rval := rval.(type) {
				case error:
					_ = logger.Log("error", rval, "culprit", "api.panic")
					_, _ = response.FromError(rval).WriteTo(w)
				default:
					st := string(debug.Stack())
					_ = logger.Log("error", rval, "stacktrace", st, "culprit", "api.panic")
					if returnStack {
						_, _ = response.Build(500, "服务器开小差了", nil).WriteTo(w)
					} else {
						_, _ = response.Build(500, st, nil).WriteTo(w)
					}
				}

				if printStack {
					debug.PrintStack()
				}
			}()

			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(f)
	}
}
