package maxIncreaseKeepingSkyline

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
			[]int{3, 0, 8, 4},
			[]int{2, 4, 5, 7},
			[]int{9, 2, 6, 3},
			[]int{0, 3, 1, 0},
		},
		35,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, maxIncreaseKeepingSkyline(tc.N1), "输入:%v", tc)
	}
}
