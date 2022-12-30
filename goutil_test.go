package goutil

import (
	"bytes"
)

func newlines(s string) string {
	return string(bytes.ReplaceAll(
		[]byte(s), []byte{'\\'}, []byte{'\n'}))
}
