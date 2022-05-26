package api

import (
	"net/http"

	gkitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"

	"aegis/internal/api/md"
	"aegis/internal/engine"
	"aegis/internal/engine/bag"
	"aegis/internal/response"
	"aegis/pkg/kit/log"
)

func makeReportHandler(
	router *mux.Router,
	logger gkitLog.Logger,
	e *engine.Engine,
) http.Handler {
	logger = log.With(logger, "request_endpoint", "report")
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqPayload := md.PopulateJSONPayload(ctx)
		_ = level.Debug(logger).Log("body", reqPayload.String())

		ctx, b := bag.New(ctx)

		for k, v := range reqPayload.Map() {
			b.StoreProperty(k, v.Value())
		}

		err := e.Report(ctx)
		if err != nil {
			panic(err)
		}

		_, _ = response.Make().WriteTo(w)
	}

	var h http.Handler
	h = http.HandlerFunc(f)
	h = md.Chain(
		md.LogRequestDuration(logger, "report"),
		md.RequireJSONPayload,
	)(h)

	router.Handle("/report", h).Methods("POST")

	return router
}
