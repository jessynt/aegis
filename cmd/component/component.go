package component

import (
	"os"
	"os/signal"
	"syscall"

	gkitLog "github.com/go-kit/kit/log"
	"github.com/pkg/errors"
)

type StartFunc func(errs chan error)

type StopFunc func()

type MakeComponentFunc func(gkitLog.Logger) ([]StartFunc, StopFunc)

func MakeComponents(
	rootLogger gkitLog.Logger,
	makers ...MakeComponentFunc,
) (starts []StartFunc, stops []StopFunc) {
	for _, maker := range makers {
		startFuncs, stopFunc := maker(rootLogger)
		starts = append(starts, startFuncs...)
		stops = append(stops, stopFunc)
	}

	starts = append(starts, func(errs chan error) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGABRT, syscall.SIGTERM)
		s := <-c
		errs <- errors.New(s.String())
	})

	return
}
