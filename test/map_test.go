package test

import (
	"strconv"
	"testing"

	"github.com/MyAuntIsPost90s/gotool/stream"
	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	strNumbers := stream.Map(numbers, func(num int) string { return strconv.Itoa(num) })
	assert.Equal(t, strNumbers, []string{"1", "2", "3", "4", "5"})
}
