package cmd

import (
	"fmt"
	"os"
)

func abort(reason string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, reason+"\n", args...)
	os.Exit(1)
}
