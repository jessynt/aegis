package grpc

import (
	"context"

	gkitEndpoint "github.com/go-kit/kit/endpoint"
	gkitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	gkitGRPC "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	"aegis/internal/engine/service"
	"aegis/pkg/kit/endpoint"
)

type EngineServer struct {
	reload gkitGRPC.Handler
}

func SetupEngineServer(
	baseServer *grpc.Server,
	rootLogger gkitLog.Logger,
	reload gkitEndpoint.Endpoint,
) {
	logger := gkitLog.With(rootLogger, "component", "engine.grpc.server")
	opts := []gkitGRPC.ServerOption{
		gkitGRPC.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	s := &EngineServer{}
	var identityDecoder = func(ctx context.Context, r interface{}) (interface{}, error) {
		return r, nil
	}

	var e gkitEndpoint.Endpoint
	{
		e = gkitEndpoint.Chain(
			endpoint.LogEndpointDuration(logger, "reload"),
		)(reload)
		s.reload = gkitGRPC.NewServer(
			e,
			identityDecoder,
			identityDecoder,
			opts...,
		)
	}

	service.RegisterEngineServer(baseServer, s)
}

func (e EngineServer) Reload(ctx context.Context, request *service.ReloadRequest) (*service.ReloadResponse, error) {
	_, response, err := e.reload.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*service.ReloadResponse), nil
}
