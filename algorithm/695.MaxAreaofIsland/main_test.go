package maxAreaOfIsland

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
			[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			[]int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		},
		6,
	},
	{
		[][]int{
			[]int{0, 0, 0, 0, 0, 0, 0, 0},
		},
		0,
	},
	{
		[][]int{
			[]int{1, 1, 0, 0, 0},
			[]int{1, 1, 0, 0, 0},
			[]int{0, 0, 0, 1, 1},
			[]int{0, 0, 0, 1, 1},
		},
		4,
	},
	{
		[][]int{
			[]int{1},
		},
		1,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, maxAreaOfIsland(tc.N1), "输入:%v", tc)
	}
}
