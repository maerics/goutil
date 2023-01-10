package goutil

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/maerics/golog"
)

func MustExec(command string, args ...string) string {
	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}
	cmd := exec.Command(command, args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		cmdline := strings.Join(append([]string{command}, args...), " ")
		format := "command %q failed: %v"
		args := []any{cmdline, err}
		if strings.TrimSpace(stderr.String()) != "" {
			format += "\n\n%v"
			args = append(args, stderr.String())
		}
		golog.Fatalf(format, args...)
	}
	return stdout.String()
}
