package test

import (
	"testing"

	"github.com/MyAuntIsPost90s/gotool/stream"
	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	exists := stream.Exists(numbers, func(number int) bool { return number == 5 })
	assert.True(t, exists)

	exists = stream.Exists(numbers, func(number int) bool { return number == 15 })
	assert.False(t, exists)
}
