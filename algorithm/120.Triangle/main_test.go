package minimumTotal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  [][]int
	ans int
}{
	{
		[][]int{
			[]int{2},
			[]int{3, 4},
			[]int{6, 5, 7},
			[]int{4, 1, 8, 3},
		},
		11,
	},
	{
		[][]int{
			[]int{-1},
			[]int{-2, -3},
		},
		-4,
	},
	{
		[][]int{
				[]int{-1},
			 []int{2, 3},
		 []int{1, -1, -3},
		},
		-1,
	},
	{
		[][]int{
			[]int{2},
			[]int{3, 4},
			[]int{6, 5, 9},
			[]int{4, 4, 8, 0},
		},
		14,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, minimumTotal(tc.N1), "输入:%v", tc)
	}
}
