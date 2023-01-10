package goutil

import (
	"os/exec"

	"github.com/maerics/golog"
)

func MustExec(command string, args ...string) string {
	cmd := exec.Command(command, args...)
	return string(golog.Must1(cmd.Output()))
}
