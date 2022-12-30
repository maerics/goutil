package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type JsonExample struct {
	x        any
	expected string
}

func TestMustJson(t *testing.T) {
	for _, eg := range []JsonExample{
		{nil, "null"},
		{123, "123"},
		{1.2, "1.2"},
		{true, "true"},
		{[]any{0, true, "foo"}, `[0,true,"foo"]`},
		{[]any{0, true, "foo"}, `[0,true,"foo"]`},
		{map[string]any{"foo": 1, "bar": []int{1, 2, 3}},
			`{"bar":[1,2,3],"foo":1}`},
	} {
		assert.Equal(t, eg.expected, MustJson(eg.x))
	}
}

func TestMustJsonPretty(t *testing.T) {
	for _, eg := range []JsonExample{
		{true, "true"},
		{[]any{0, true, "foo"},
			newlines(`[\  0,\  true,\  "foo"\]`)},
		{map[string]any{"foo": 1, "bar": []int{1, 2, 3}},
			newlines(`{\  "bar": [\    1,\    2,\    3\  ],\  "foo": 1\}`)},
	} {
		assert.Equal(t, eg.expected, MustJson(eg.x, true))
	}
}
