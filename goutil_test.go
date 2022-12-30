package goutil

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ToSnakeTest(t *testing.T) {
	for _, eg := range []struct{ x, expected string }{
		{"Id", "id"},
		{"CreatedAt", "created_at"},
	} {
		assert.Equal(t, eg.expected, ToSnake(eg.x))
	}
}

func newlines(s string) string {
	return string(bytes.ReplaceAll(
		[]byte(s), []byte{'\\'}, []byte{'\n'}))
}
