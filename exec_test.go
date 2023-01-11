package goutil

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustExec(t *testing.T) {
	assert.Equal(t, "", MustExec("echo -n"))
	assert.Equal(t, "\n", MustExec("echo"))
	assert.Equal(t, "Hello", MustExec("echo -n Hello"))
	assert.Equal(t, "Hello\n", MustExec("echo Hello"))
	assert.Equal(t, "d41d8cd98f00b204e9800998ecf8427e  -\n", MustExec("echo -n | md5sum"))
}

func TestMustExecArgs(t *testing.T) {
	assert.Equal(t, "", MustExecArgs("echo", "-n"))
	assert.Equal(t, "\n", MustExecArgs("echo"))
	assert.Equal(t, "Hello", MustExecArgs("echo", "-n", "Hello"))
	assert.Equal(t, "Hello\n", MustExecArgs("echo", "Hello"))
}

func TestMustExecStdin(t *testing.T) {
	assert.Equal(t, "", MustExecArgsStdin("cat", nil, nil))
	assert.Equal(t, "hello", MustExecArgsStdin("cat", nil, strings.NewReader("hello")))
	assert.Equal(t, "hello\n", MustExecArgsStdin("cat", nil, strings.NewReader("hello\n")))
}
