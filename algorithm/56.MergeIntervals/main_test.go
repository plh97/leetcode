package merge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  [][]int
	ans [][]int
}{
	{
		[][]int{
			[]int{2, 3},
			[]int{4, 5},
			[]int{6, 7},
			[]int{8, 9},
			[]int{1, 10},
		},
		[][]int{
			[]int{1, 10},
		},
	},
	{
		[][]int{
			[]int{1, 3},
			[]int{2, 6},
			[]int{8, 10},
			[]int{15, 18},
		},
		[][]int{
			[]int{1, 6},
			[]int{8, 10},
			[]int{15, 18},
		},
	},
	{
		[][]int{
			[]int{1, 4},
			[]int{0, 4},
		},
		[][]int{
			[]int{0, 4},
		},
	},
	{
		[][]int{
			[]int{1, 4},
			[]int{0, 0},
		},
		[][]int{
			[]int{0, 0},
			[]int{1, 4},
		},
	},
	{
		[][]int{
			[]int{1, 4},
			[]int{0, 2},
			[]int{3, 5},
		},
		[][]int{
			[]int{0, 5},
		},
	},
	{
		[][]int{
			[]int{2, 3},
			[]int{2, 2},
			[]int{3, 3},
			[]int{1, 3},
			[]int{5, 7},
			[]int{2, 2},
			[]int{4, 6},
		},
		[][]int{
			[]int{1, 3},
			[]int{4, 7},
		},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, merge(tc.N1), "输入:%v", tc)
	}
}
