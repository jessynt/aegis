package abstraction

import (
	"context"
	"encoding/json"
	"net/http"

	gkitLog "github.com/go-kit/kit/log"
	gkitTransport "github.com/go-kit/kit/transport"
	gkithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"aegis/internal/proto"
	"aegis/internal/response"
)

func MakeHandler(as Service, logger gkitLog.Logger) http.Handler {
	opts := []gkithttp.ServerOption{
		gkithttp.ServerErrorHandler(gkitTransport.NewLogErrorHandler(logger)),
		gkithttp.ServerErrorEncoder(func(_ context.Context, err error, _ http.ResponseWriter) { panic(err) }),
	}

	createAbstractionHandler := gkithttp.NewServer(
		makeCreateAbstractionEndpoint(as),
		decodeCreateAbstractionRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()
	r.Handle("/abstractions", createAbstractionHandler).Methods("POST")
	return r
}

func decodeCreateAbstractionRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body struct {
		Name                   string `json:"name"`
		Label                  string `json:"label"`
		ModelId                int64  `json:"model_id"`
		AggregateType          int8   `json:"aggregate_type"`
		AggregateField         string `json:"aggregate_field"`
		AggregateIntervalType  int8   `json:"aggregate_interval_type"`
		AggregateIntervalValue int64  `json:"aggregate_interval_value"`
		SearchField            string `json:"search_field"`
		FilterExpression       string `json:"filter_expression"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}

	return createAbstractionRequest{
		Name:                   body.Name,
		Label:                  body.Label,
		ModelId:                body.ModelId,
		AggregateType:          proto.AggregateType(body.AggregateType),
		AggregateField:         body.AggregateField,
		AggregateIntervalType:  proto.AggregateIntervalType(body.AggregateIntervalType),
		AggregateIntervalValue: body.AggregateIntervalValue,
		SearchField:            body.SearchField,
		FilterExpression:       body.FilterExpression,
	}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	_, err := response.Make().WithData(resp).WriteTo(w)
	return err
}
