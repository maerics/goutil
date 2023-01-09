package goutil

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMustJson(t *testing.T) {
	for _, eg := range []struct {
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
	} {
		assert.Equal(t, eg.expected, MustJson(eg.x))
	}
}

func TestMustJsonPretty(t *testing.T) {
	for _, eg := range []struct {
		x        any
		expected string
	}{
		{true, "true"},
		{[]any{0, true, "foo"},
			newlines(`[\  0,\  true,\  "foo"\]`)},
		{map[string]any{"foo": 1, "bar": []int{1, 2, 3}},
			newlines(`{\  "bar": [\    1,\    2,\    3\  ],\  "foo": 1\}`)},
	} {
		assert.Equal(t, eg.expected, MustJson(eg.x, true))
	}
}

func TestOrderedJsonMap(t *testing.T) {
	for _, eg := range []struct {
		m        map[string]any
		v        []any
		keys     []string
		nulls    bool
		expected string
	}{
		// Maps
		{m: map[string]any{"name": "Mike", "alpha": "bravo"},
			keys:     []string{"name", "alpha"},
			expected: `{"name":"Mike","alpha":"bravo"}`},
		{m: map[string]any{"name": "Mike"},
			keys:     []string{"name", "alpha"},
			expected: `{"name":"Mike"}`},
		{m: map[string]any{"name": "Mike"},
			keys:     []string{"name", "alpha"},
			nulls:    true,
			expected: `{"name":"Mike","alpha":null}`},
		{m: map[string]any{"name": "Mike", "settings": json.RawMessage(`{"darkmode":true}`)},
			keys:     []string{"name", "settings"},
			expected: `{"name":"Mike","settings":{"darkmode":true}}`},
		// Slices
		{v: []any{"Mike", "bravo"},
			keys:     []string{"name", "alpha"},
			expected: `{"name":"Mike","alpha":"bravo"}`},
		{v: []any{"Mike"},
			keys:     []string{"name", "alpha"},
			expected: `{"name":"Mike"}`},
		{v: []any{"Mike"},
			keys:     []string{"name", "alpha"},
			nulls:    true,
			expected: `{"name":"Mike","alpha":null}`},
	} {
		assert.Equal(t, eg.expected, MustJson(OrderedJsonObj{
			Keys:   eg.keys,
			Values: eg.v,
			Map:    eg.m,
			Nulls:  eg.nulls,
		}))
	}
}
