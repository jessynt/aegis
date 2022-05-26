package model

import (
	"context"

	gkitEndpoint "github.com/go-kit/kit/endpoint"
)

type createModelRequest struct {
	Name  string
	GUID  string
	Label string
}

type createModelResponse struct {
	Err error `json:"error,omitempty"`
}

func (r createModelResponse) error() error { return r.Err }

func makeCreateModelEndpoint(ms Service) gkitEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createModelRequest)
		err := ms.CreateModel(ctx, req.GUID, req.Name, req.Label)
		return createModelResponse{Err: err}, nil
	}
}
