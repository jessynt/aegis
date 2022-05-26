package aegis

import (
	"fmt"
	"runtime"
)

var (
	Version   string
	BuildDate string
)

func VersionString() string {
	version := Version
	if version == "" {
		version = "unknown"
	}

	buildDate := BuildDate
	if buildDate == "" {
		buildDate = "unknown"
	}

	return fmt.Sprintf(
		"%s-%s-%s",
		runtime.Version(),
		version,
		buildDate,
	)
}
