package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func makeHealthzHandler(
	router *mux.Router,
) http.Handler {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}

	h := http.HandlerFunc(f)

	router.Handle("/healthz", h).Methods("GET")

	return router
}
