package goutil

import (
	"bytes"
	"testing"
)

func TestMustJson(t *testing.T) {
	examples := []struct {
		x        any
		expected string
	}{
		{nil, "null"},
		{123, "123"},
		{1.2, "1.2"},
		{true, "true"},
		{[]any{0, true, "foo"}, `[0,true,"foo"]`},
		{[]any{0, true, "foo"}, `[0,true,"foo"]`},
		{map[string]any{"foo": 1, "bar": []int{1, 2, 3}},
			`{"bar":[1,2,3],"foo":1}`},
	}

	for i, eg := range examples {
		actual := MustJson(eg.x)
		if actual != eg.expected {
			t.Errorf("example %#v, wanted %q, got %q", i+1, eg.expected, actual)
		}
	}
}

func TestMustJsonPretty(t *testing.T) {
	examples := []struct {
		x        any
		expected string
	}{
		{[]any{0, true, "foo"},
			newlines(`[\  0,\  true,\  "foo"\]`)},
		{map[string]any{"foo": 1, "bar": []int{1, 2, 3}},
			newlines(`{\  "bar": [\    1,\    2,\    3\  ],\  "foo": 1\}`)},
	}

	for i, eg := range examples {
		actual := MustJson(eg.x, true)
		if actual != eg.expected {
			t.Errorf("example %#v\nwanted %q,\n   got %q", i+1, eg.expected, actual)
		}
	}
}

func newlines(s string) string {
	return string(bytes.ReplaceAll(
		[]byte(s), []byte{'\\'}, []byte{'\n'}))
}
