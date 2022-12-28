package goutil

import (
	"bytes"
	"testing"
)

func assertEqual(t *testing.T, exampleIndex int, expected, actual any) {
	if actual != expected {
		t.Errorf("example %#v, wanted %#v, got %#v",
			exampleIndex+1, expected, actual)
	}
}

func newlines(s string) string {
	return string(bytes.ReplaceAll(
		[]byte(s), []byte{'\\'}, []byte{'\n'}))
}

type JsonExample struct {
	x        any
	expected string
}

func TestMustJson(t *testing.T) {
	for i, eg := range []JsonExample{
		{nil, "null"},
		{123, "123"},
		{1.2, "1.2"},
		{true, "true"},
		{[]any{0, true, "foo"}, `[0,true,"foo"]`},
		{[]any{0, true, "foo"}, `[0,true,"foo"]`},
		{map[string]any{"foo": 1, "bar": []int{1, 2, 3}},
			`{"bar":[1,2,3],"foo":1}`},
	} {
		assertEqual(t, i, eg.expected, MustJson(eg.x))
	}
}

func TestMustJsonPretty(t *testing.T) {
	for i, eg := range []JsonExample{
		{true, "true"},
		{[]any{0, true, "foo"},
			newlines(`[\  0,\  true,\  "foo"\]`)},
		{map[string]any{"foo": 1, "bar": []int{1, 2, 3}},
			newlines(`{\  "bar": [\    1,\    2,\    3\  ],\  "foo": 1\}`)},
	} {
		assertEqual(t, i, eg.expected, MustJson(eg.x, true))
	}
}

func ToSnakeTest(t *testing.T) {
	for i, eg := range []struct{ x, expected string }{
		{"Id", "id"},
		{"CreatedAt", "created_at"},
	} {
		assertEqual(t, i, eg.expected, ToSnake(eg.x))
	}
}
