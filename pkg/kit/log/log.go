package log

import (
	"os"

	gkitLog "github.com/go-kit/kit/log"
)

var With = gkitLog.With

func NewJSONLogger(kvs ...interface{}) gkitLog.Logger {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "(unknown)"
	}

	kvs = append(
		[]interface{}{
			"time", gkitLog.DefaultTimestampUTC,
			"hostname", hostname,
		},
		kvs...,
	)

	logger := gkitLog.NewJSONLogger(os.Stdout)
	return With(
		logger,
		kvs...,
	)
}
