package kClosest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  [][]int
	N2  int
	ans [][]int
}{
	{
		[][]int{
			[]int{1, 3},
			[]int{-2, 2},
		},
		1,
		[][]int{
			[]int{-2, 2},
		},
	},
	{
		[][]int{
			[]int{3, 3},
			[]int{5, -1},
			[]int{-2, 4},
		},
		2,
		[][]int{
			[]int{3, 3},
			[]int{-2, 4},
		},
	},
	{
		[][]int{
			[]int{1, 3},
			[]int{-2, 2},
			[]int{2, -2},
		},
		2,
		[][]int{
			[]int{-2, 2},
			[]int{2, -2},
		},
	},
	{
		[][]int{
			[]int{6, 10},
			[]int{-3, 3},
			[]int{-2, 5},
			[]int{0, 2},
		},
		3,
		[][]int{
			[]int{0, 2},
			[]int{-3, 3},
			[]int{-2, 5},
		},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, kClosest(tc.N1, tc.N2), "输入:%v", tc)
	}
}
