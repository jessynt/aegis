package activation

import (
	"context"
	"net/http"

	gkitLog "github.com/go-kit/kit/log"
	gkitTransport "github.com/go-kit/kit/transport"
	gkithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"aegis/internal/response"
)

func MakeHandler(as Service, logger gkitLog.Logger) http.Handler {
	opts := []gkithttp.ServerOption{
		gkithttp.ServerErrorHandler(gkitTransport.NewLogErrorHandler(logger)),
		gkithttp.ServerErrorEncoder(func(_ context.Context, err error, _ http.ResponseWriter) { panic(err) }),
	}

	createActivationHandler := gkithttp.NewServer(
		makeCreateActivationEndpoint(as),
		decodeCreateActivationRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/activations", createActivationHandler).Methods("POST")
	return r
}

func decodeCreateActivationRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	_, err := response.Make().WithData(resp).WriteTo(w)
	return err
}
