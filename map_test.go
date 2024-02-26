package goutil

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	assert.Equal(t,
		Map([]int{1, 2, 3}, func(x int) int { return x + 1 }),
		[]int{2, 3, 4})

	assert.Equal(t,
		Map([]int{1, 2, 3}, func(s int) string { return strconv.Itoa(s) }),
		[]string{"1", "2", "3"})

	assert.Equal(t,
		Map([]string{"foo", "bar", "gah"}, func(s string) string { return string(s[0]) }),
		[]string{"f", "b", "g"})
}
