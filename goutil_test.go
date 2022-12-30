package goutil

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, Sort([]string{"b", "c", "a"}))
	assert.Equal(t, []int{1, 2, 3}, Sort([]int{3, 1, 2}))
	assert.Equal(t, []float64{1.0, 2.0, 3.0}, Sort([]float64{2.0, 1.0, 3.0}))
}

func TestKeys(t *testing.T) {
	assert.Equal(t, []string{"a", "b"},
		Sort(Keys(map[string]int{"b": 1, "a": 2})))
	assert.Equal(t, []string{},
		Sort(Keys(map[string]any{})))

	assert.Equal(t, []int{1, 2, 3},
		Sort(Keys(map[int]int{2: 3, 1: 2, 3: 4})))
	assert.Equal(t, []int{},
		Sort(Keys(map[int]any{})))

	assert.Equal(t, []float64{1.2, 2.3},
		Sort(Keys(map[float64]any{1.2: nil, 2.3: nil})))
}

func newlines(s string) string {
	return string(bytes.ReplaceAll(
		[]byte(s), []byte{'\\'}, []byte{'\n'}))
}
