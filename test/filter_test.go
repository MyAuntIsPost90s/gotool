package test

import (
	"testing"

	"github.com/MyAuntIsPost90s/gotool/stream"
	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr = stream.Filter(arr, func(i int) bool {
		return i%2 == 0
	})
	assert.Equal(t, arr, []int{2, 4, 6, 8, 10})
}
