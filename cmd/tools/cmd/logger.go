package cmd

import (
	"log"
)

type logger struct{}

func (l *logger) Infof(format string, a ...interface{}) {
	log.Printf(format, a...)
}

func (l *logger) Warnf(format string, a ...interface{}) {
	log.Printf(format, a...)
}

func (l *logger) Errorf(format string, a ...interface{}) {
	log.Printf(format, a...)
}
