package test

import (
	"testing"

	"github.com/MyAuntIsPost90s/gotool/stream"
	"github.com/stretchr/testify/assert"
)

func TestCreateStream(t *testing.T) {
	s := stream.CreateStream([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	s = s.Filter(func(o int) bool { return o < 10 })
	result := s.Reduce(0, func(a int, b any) any { return a + b.(int) }).(int)
	assert.Equal(t, 45, result)
}
