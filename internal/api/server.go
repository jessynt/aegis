package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/viper"

	"aegis/internal/api/md"
	"aegis/internal/config"
	"aegis/internal/engine"
	"aegis/internal/engine/store"

	gkitLog "github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

type HttpServer struct {
	config   *viper.Viper
	stopChan chan chan struct{}
	logger   gkitLog.Logger
}

func NewHttpServer(config *viper.Viper, logger gkitLog.Logger) *HttpServer {
	return &HttpServer{
		config:   config,
		logger:   logger,
		stopChan: make(chan chan struct{}),
	}
}

func (s *HttpServer) Start(
	dbStore *store.Store,
	e *engine.Engine,
) error {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/v1").Subrouter()

	makeHealthzHandler(router)
	makeCheckHandler(apiRouter, s.logger, e)
	makeReportHandler(apiRouter, s.logger, e)
	makeAddToCollectionHandler(apiRouter, s.logger, dbStore.MysqlConn())
	makeReloadHandler(apiRouter, s.logger, e)

	var httpHandler http.Handler
	{
		httpHandler = router
		httpHandler = md.Chain(
			md.CapturePanic(!config.IsProduction(), !config.IsProduction(), s.logger),
		)(httpHandler)
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.config.GetString("http.host"), s.config.GetInt("http.port")),
		Handler: httpHandler,
	}

	httpServerErrChan := make(chan error)
	go func(httpServerErrChan chan error, httpServer *http.Server) {
		httpServerErrChan <- httpServer.ListenAndServe()
	}(httpServerErrChan, httpServer)

	_ = s.logger.Log(
		"message",
		fmt.Sprintf("http server started at %s", httpServer.Addr),
	)

	select {
	case c := <-s.stopChan:
		httpServer.Shutdown(context.Background())
		c <- struct{}{}
		return nil
	case err := <-httpServerErrChan:
		return err
	}
}

func (s *HttpServer) Stop() {
	c := make(chan struct{}, 1)
	s.stopChan <- c
	<-c
}
