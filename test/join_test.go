package test

import (
	"strconv"
	"testing"

	"github.com/MyAuntIsPost90s/gotool/stream"
	"github.com/stretchr/testify/assert"
)

func TestJoin(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	str := stream.Join(numbers, ",", func(number int) string { return strconv.Itoa(number) })
	assert.Equal(t, "1,2,3,4,5", str)
}
