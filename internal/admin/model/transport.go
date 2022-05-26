package model

import (
	"context"
	"encoding/json"
	"net/http"

	gkitLog "github.com/go-kit/kit/log"
	gkitTransport "github.com/go-kit/kit/transport"
	gkithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"aegis/internal/response"
)

func MakeHandler(ms Service, logger gkitLog.Logger) http.Handler {
	opts := []gkithttp.ServerOption{
		gkithttp.ServerErrorHandler(gkitTransport.NewLogErrorHandler(logger)),
		gkithttp.ServerErrorEncoder(func(_ context.Context, err error, _ http.ResponseWriter) { panic(err) }),
	}

	createModelHandler := gkithttp.NewServer(
		makeCreateModelEndpoint(ms),
		decodeCreateModelRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/models", createModelHandler).Methods("POST")
	return r
}

func decodeCreateModelRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		GUID  string `json:"guid"`
		Name  string `json:"name"`
		Label string `json:"label"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return createModelRequest{
		GUID:  body.GUID,
		Name:  body.Name,
		Label: body.Label,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	if e, ok := resp.(errorer); ok && e.error() != nil {
		panic(e.error())
	}
	_, err := response.Make().WithData(resp).WriteTo(w)
	return err
}

type errorer interface {
	error() error
}
