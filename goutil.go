package goutil

import (
	"os"
	"strings"

	log "github.com/maerics/golog"
)

func MustEnv(varname string) string {
	value := strings.TrimSpace(os.Getenv(varname))
	if value == "" {
		log.Fatalf("the environment variable %q must be set", varname)
	}
	return value
}
