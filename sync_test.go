package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoBatchWork(t *testing.T) {
	xs := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10}
	for workerCount := 1; workerCount < 10; workerCount++ {
		doublexs := make([]int, len(xs))
		DoBatchWork(workerCount, xs, func(i int, x int) { doublexs[i] = x * 2 })
		assert.Equal(t, expected, doublexs)
	}
}
