package test

import (
	"testing"

	"github.com/MyAuntIsPost90s/gotool/stream"
	"github.com/stretchr/testify/assert"
)

func TestListToMap(t *testing.T) {
	type People struct {
		id   int
		name string
	}

	peopleArr := []People{{1, "name"}, {2, "name2"}, {3, "name3"}}
	pMap := stream.ListToMap(peopleArr, func(p People) int { return p.id }, func(p People) string { return p.name })
	assert.Equal(t, pMap[1], "name")
	assert.Equal(t, pMap[2], "name2")
}
