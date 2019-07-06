package permute

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
	p para
	a ans
}

var tcs = []struct {
	N1  []int
	ans [][]int
}{
	{
		[]int{5, 4, 6, 2},
		[][]int{[]int{5, 4, 6, 2}, []int{5, 4, 2, 6}, []int{5, 6, 4, 2}, []int{5, 6, 2, 4}, []int{5, 2, 4, 6}, []int{5, 2, 6, 4}, []int{4, 5, 6, 2}, []int{4, 5, 2, 6}, []int{4, 6, 5, 2}, []int{4, 6, 2, 5}, []int{4, 2, 5, 6}, []int{4, 2, 6, 5}, []int{6, 5, 4, 2}, []int{6, 5, 2, 4}, []int{6, 4, 5, 2}, []int{6, 4, 2, 5}, []int{6, 2, 5, 4}, []int{6, 2, 4, 5}, []int{2, 5, 4, 6}, []int{2, 5, 6, 4}, []int{2, 4, 5, 6}, []int{2, 4, 6, 5}, []int{2, 6, 5, 4}, []int{2, 6, 4, 5}},
	},
	{
		[]int{0, 1},
		[][]int{
			[]int{0, 1},
			[]int{1, 0},
		},
	},
	{
		[]int{1},
		[][]int{
			[]int{1},
		},
	},
	{
		[]int{1, 2, 3},
		[][]int{[]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}, []int{2, 3, 1}, []int{3, 1, 2}, []int{3, 2, 1}},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, permute(tc.N1), "输入:%v", tc)
	}
}
