package component

import (
	"context"
	"fmt"
	"net"
	"time"

	gkitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"aegis/internal/api"
	"aegis/internal/config"
	_ "aegis/internal/config/auto"
	"aegis/internal/engine"
	"aegis/internal/engine/abstraction"
	"aegis/internal/engine/activation"
	"aegis/internal/engine/aggregation"
	"aegis/internal/engine/collection"
	"aegis/internal/engine/contract"
	engineGRPC "aegis/internal/engine/grpc"
	"aegis/internal/engine/model"
	"aegis/internal/engine/property"
	"aegis/internal/engine/rule"
	"aegis/internal/engine/store"
	"aegis/internal/mysql"
)

func SetupEngine(rootLogger gkitLog.Logger) ([]StartFunc, StopFunc) {
	logger := gkitLog.With(rootLogger, "component", "engine")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)

	var withGRPCServerTracer grpc.ServerOption
	{
		tracer := opentracing.GlobalTracer()
		withGRPCServerTracer = grpc.UnaryInterceptor(otgrpc.OpenTracingServerInterceptor(tracer))
	}

	var mysqlConn *sqlx.DB
	{
		var err error
		mysqlConn, err = mysql.CreateConnx(config.MySQL)
		if err != nil {
			panic(err)
		}
	}

	var clickHouseConn *sqlx.DB
	{
		var err error
		clickHouseConn, err = sqlx.Open("clickhouse", fmt.Sprintf(
			"tcp://%s:%s?database=%s&debug=%s",
			config.ClickHouse.GetString("conn.host"),
			config.ClickHouse.GetString("conn.port"),
			config.ClickHouse.GetString("conn.dbname"),
			config.ClickHouse.GetString("conn.debug"),
		))
		if err != nil {
			panic(err)
		}
	}

	var dbStore *store.Store
	{
		dbStore = store.NewStore(mysqlConn, clickHouseConn)
	}

	var modelManager contract.ModelManager
	{
		modelManager = model.NewManager(model.WithStore(dbStore))
	}

	var propertyManager contract.PropertyManager
	{
		propertyManager = property.NewManager(property.WithStore(dbStore))
	}

	var abstractionManager contract.AbstractionManager
	{
		abstractionManager = abstraction.NewManager(abstraction.WithStore(dbStore))
	}

	var activationManager contract.ActivationManager
	{
		activationManager = activation.NewManager(activation.WithStore(dbStore))
	}

	var ruleManager contract.RuleManager
	{
		ruleManager = rule.NewManager(rule.WithStore(dbStore))
	}

	var aggregator contract.Aggregator
	{
		aggregator = aggregation.NewAggregator(aggregation.WithClickHouseConn(clickHouseConn))
	}

	var collectionManager contract.CollectionManager
	{
		collectionManager = collection.NewManager(collection.WithStore(dbStore))
	}

	var e *engine.Engine
	{
		e = engine.NewEngine(
			engine.WithLogger(logger),
			engine.WithModelManager(modelManager),
			engine.WithPropertyManager(propertyManager),
			engine.WithAbstractionManager(abstractionManager),
			engine.WithActivationManager(activationManager),
			engine.WithRuleManager(ruleManager),
			engine.WithAggregator(aggregator),
			engine.WithCollectionManager(collectionManager),
			engine.WithStore(dbStore),
		)
		if err := e.Init(ctx); err != nil {
			panic(err)
		}
	}

	var apiHTTPServer *api.HttpServer
	{
		apiHTTPServer = api.NewHttpServer(config.Api, logger)
	}

	var grpcServer = grpc.NewServer(withGRPCServerTracer)
	{
		engineGRPC.SetupEngineServer(
			grpcServer,
			rootLogger,
			engineGRPC.MakeReloadHandler(e),
		)
	}

	cancel()

	return []StartFunc{
			func(errs chan error) {
				errs <- apiHTTPServer.Start(dbStore, e)
			},
			func(errs chan error) {
				addr := fmt.Sprintf(
					"%s:%d",
					config.Engine.GetString("grpc.host"),
					config.Engine.GetInt("grpc.port"),
				)
				listener, err := net.Listen("tcp", addr)
				if err != nil {
					errs <- err
					return
				}
				_ = level.Info(logger).Log("message", fmt.Sprintf("engine grpc server started at %s", addr))
				errs <- grpcServer.Serve(listener)
			},
		}, func() {
			_ = level.Info(logger).Log("message", "http server stopping")
			apiHTTPServer.Stop()
			_ = level.Info(logger).Log("message", "http server stopped")

			_ = level.Info(logger).Log("message", "engine grpc server stopping")
			grpcServer.GracefulStop()
			_ = level.Info(logger).Log("message", "engine grpc server stopped")
		}
}
