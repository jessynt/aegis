package component

import (
	"fmt"
	"net/http"

	gkitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/jmoiron/sqlx"

	"aegis/internal/admin/abstraction"
	"aegis/internal/admin/activation"
	"aegis/internal/admin/md"
	"aegis/internal/admin/model"
	"aegis/internal/config"
	_ "aegis/internal/config/auto"
	"aegis/internal/mysql"
	"aegis/pkg/kit/log"
)

func SetupAdmin(rootLogger gkitLog.Logger) ([]StartFunc, StopFunc) {
	logger := gkitLog.With(rootLogger, "component", "admin")

	// var engineClient engineService.EngineClient
	// {
	// 	grpcUrl := config.Admin.GetString("engine.grpc_url")
	// 	conn, err := grpc.Dial(
	// 		grpcUrl,
	// 		grpc.WithInsecure(),
	// 	)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	engineClient = engineService.NewEngineClient(conn)
	// }

	var mysqlConn *sqlx.DB
	{
		var err error
		mysqlConn, err = mysql.CreateConnx(config.MySQL)
		if err != nil {
			panic(err)
		}
	}

	var ms model.Service
	{
		ms = model.NewService(mysqlConn)
		ms = model.NewLoggingService(log.With(logger, "component", "model"), ms)
	}

	var as abstraction.Service
	{
		as = abstraction.NewService(mysqlConn)
	}

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()
	// RESTful is being a pain
	// Please have mercy on me =(
	mux.Handle("/models", model.MakeHandler(ms, httpLogger))
	mux.Handle("/abstractions", abstraction.MakeHandler(as, httpLogger))
	mux.Handle("/activations", activation.MakeHandler(as, httpLogger))

	http.Handle("/", md.Chain(
		md.CapturePanic(true, false, logger),
		md.AccessControl(),
	)(mux))

	return []StartFunc{
			func(errs chan error) {
				httpAddr := fmt.Sprintf("%s:%d", config.Admin.GetString("http.host"), config.Admin.GetInt("http.port"))
				_ = level.Info(logger).Log("transport", "http", "address", httpAddr, "msg", "listening")
				errs <- http.ListenAndServe(httpAddr, nil)
			},
		}, func() {

		}
}
