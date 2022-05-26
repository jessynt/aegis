package abstraction

import (
	"context"

	gkitEndpoint "github.com/go-kit/kit/endpoint"

	"aegis/internal/proto"
)

type createAbstractionRequest struct {
	Name                   string
	Label                  string
	ModelId                int64
	AggregateType          proto.AggregateType
	AggregateField         string
	AggregateIntervalType  proto.AggregateIntervalType
	AggregateIntervalValue int64
	SearchField            string
	FilterExpression       string
}

type createAbstractionResponse struct{}

func makeCreateAbstractionEndpoint(as Service) gkitEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createAbstractionRequest)
		if err := as.CreateAbstraction(ctx, req); err != nil {
			panic(err)
		}
		return createAbstractionResponse{}, nil
	}
}
