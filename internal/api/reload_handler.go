package api

import (
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"

	"aegis/internal/api/md"
	"aegis/internal/engine"
	"aegis/internal/response"

	gkitLog "github.com/go-kit/kit/log"
)

func makeReloadHandler(
	router *mux.Router,
	logger gkitLog.Logger,
	e *engine.Engine,
) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		_ = level.Info(logger).Log("message", fmt.Sprintf("engine reloading"))

		err := e.Reload(ctx)
		if err != nil {
			panic(fmt.Errorf("engine reload failed: %w", err))
		}
		_, _ = response.Make().WriteTo(w)
	}

	var h http.Handler
	h = http.HandlerFunc(f)
	h = md.Chain(
		md.LogRequestDuration(logger, "reload"),
	)(h)

	router.Handle("/reload", h).Methods("POST")

	return router
}
