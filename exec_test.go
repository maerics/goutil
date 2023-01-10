package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustExec(t *testing.T) {
	assert.Equal(t, "", MustExec("echo", "-n"))
	assert.Equal(t, "\n", MustExec("echo"))
	assert.Equal(t, "Hello", MustExec("echo", "-n", "Hello"))
	assert.Equal(t, "Hello\n", MustExec("echo", "Hello"))
}
