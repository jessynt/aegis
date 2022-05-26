package auto

import (
	"os"

	"aegis/internal/config"
)

func init() {
	config.Init("aegis_test")
	os.Setenv("TEST", "")
}
