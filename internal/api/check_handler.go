package api

import (
	"net/http"

	"github.com/go-kit/kit/log/level"

	"aegis/internal/api/md"
	"aegis/internal/engine"
	"aegis/internal/engine/bag"
	"aegis/internal/proto"
	"aegis/internal/response"
	"aegis/pkg/kit/log"

	gkitLog "github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

func makeCheckHandler(
	router *mux.Router,
	logger gkitLog.Logger,
	e *engine.Engine,
) http.Handler {
	logger = log.With(logger, "request_endpoint", "check")
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqPayload := md.PopulateJSONPayload(ctx)
		_ = level.Debug(logger).Log("body", reqPayload.String())

		ctx, b := bag.New(ctx)

		for k, v := range reqPayload.Map() {
			b.StoreProperty(k, v.Value())
		}

		err := e.Check(ctx)
		if err != nil {
			panic(err)
		}

		recommendedAction := proto.RiskPass

		activations := b.GetsActivations()
		for _, activation := range activations {
			if activation.RickType > recommendedAction {
				recommendedAction = activation.RickType
			}
		}

		_, _ = response.Make().
			MustSetData("risk_type", recommendedAction).
			MustSetData("activations", activations).
			WriteTo(w)
	}

	var h http.Handler
	h = http.HandlerFunc(f)
	h = md.Chain(
		md.LogRequestDuration(logger, "check"),
		md.RequireJSONPayload,
	)(h)

	router.Handle("/check", h).Methods("POST")

	return router
}

// func isValidTimestamp(timestamp int64, timeFree bool) bool {
// 	if timeFree {
// 		return true
// 	}
// 	if timestamp < 1000000000000 || timestamp > 10000000000000 {
// 		return false
// 	}
// 	start := time.Now().Add(-time.Duration(48) * time.Hour)
// 	end := time.Now().Add(time.Duration(1) * time.Hour)
// 	t := time.Unix(0, timestamp*int64(time.Millisecond))
// 	return inTimeSpan(start, end, t)
// }
//
// inTimeSpan 是否在时间区间
// func inTimeSpan(start, end, check time.Time) bool {
// 	return check.After(start) && check.Before(end)
// }
