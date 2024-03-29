package main

import (
	gkitLog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"aegis"
	"aegis/cmd/component"
	"aegis/pkg/kit/log"
)

var rootLogger = log.NewJSONLogger(
	"server", "admin",
	"mode", "production",
)

func main() {
	_ = rootLogger.Log("message", "initializing", "version", aegis.VersionString())

	var logger gkitLog.Logger
	{
		logger = log.With(rootLogger, "component", "main")
	}

	starts, stops := component.MakeComponents(
		rootLogger,
		component.SetupAdmin,
	)
	errs := make(chan error, len(starts))
	for _, c := range starts {
		go c(errs)
	}

	err := <-errs
	_ = level.Info(logger).Log("message", "shutting down", "error", err)
	for _, stop := range stops {
		stop()
	}
	_ = level.Info(logger).Log("message", "terminated", "error", err)
}
