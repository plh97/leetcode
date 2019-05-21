package maxSlidingWindow

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	one []int
	two int
}

var tcs = []struct {
	N   question
	ans []int
}{
	{
		question{
			one: []int{7, 6, 5, 4, 3, 2, 1},
			two: 3,
		},
		[]int{7, 6, 5, 4, 3},
	},
	{
		question{
			one: []int{1, 2, 3, 4, 5, 6, 7},
			two: 3,
		},
		[]int{3, 4, 5, 6, 7},
	},
	{
		question{
			one: []int{1, 3, -1, -3, 5, 3, 6, 7},
			two: 3,
		},
		[]int{3, 3, 5, 5, 6, 7},
	},
	{
		question{
			one: []int{7, 2, 4},
			two: 2,
		},
		[]int{7, 4},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, maxSlidingWindow(tc.N.one, tc.N.two), "输入:%v", tc)
	}
}

func Benchmark_bitwiseComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			maxSlidingWindow(tc.N.one, tc.N.two)
		}
	}
}
