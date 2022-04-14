package test

import (
	"testing"

	"github.com/MyAuntIsPost90s/gotool/stream"
	"github.com/stretchr/testify/assert"
)

func TestReduce(t *testing.T) {
	var tests = []struct {
		input  []int
		output int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 120},
	}

	for _, test := range tests {
		output := stream.Reduce(test.input, 0, func(a, b int) int { return a + b })
		assert.Equal(t, test.output, output)
	}
}
