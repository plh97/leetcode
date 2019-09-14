package findDiagonalOrder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  [][]int
	ans []int
}{
	{
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{7, 8, 9},
		},
		[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
	},
	{
		[][]int{
			[]int{2, 5},
			[]int{8, 4},
			[]int{0, -1},
		},
		[]int{2, 5, 8, 0, 4, -1},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, findDiagonalOrder(tc.N1), "输入:%v", tc)
	}
}
