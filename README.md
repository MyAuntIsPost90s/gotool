gotool
===
## Version
0.0.2
## Installation
```
go get github.com/MyAuntIsPost90s/gotool
```
## Demo
### Filter
```
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
```
### Map
```
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
```
### ListToMap
```
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
```
### Reduce
```
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
```
### Exists
```
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
```
### Stream
```
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

```
