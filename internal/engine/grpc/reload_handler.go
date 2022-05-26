package grpc

import (
	"context"

	gkitEndpoint "github.com/go-kit/kit/endpoint"

	"aegis/internal/engine/service"
)

type ReloadableT interface {
	Reload(ctx context.Context) error
}

func MakeReloadHandler(engine ReloadableT) gkitEndpoint.Endpoint {
	failed := func(err error) (*service.ReloadResponse, error) {
		rv := &service.ReloadResponse{}
		rv.SetError(err)
		return rv, nil
	}

	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if err := engine.Reload(ctx); err != nil {
			return failed(err)
		}
		return &service.ReloadResponse{}, nil
	}
}
