package checkstraightline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  [][]int
	ans bool
}{
	{
		[][]int{
			[]int{1, 2},
			[]int{2, 3},
			[]int{3, 4},
			[]int{4, 5},
			[]int{5, 6},
			[]int{6, 7},
		},
		true,
	},
	{
		[][]int{
			[]int{1, 1},
			[]int{2, 2},
			[]int{3, 4},
			[]int{4, 5},
			[]int{5, 6},
			[]int{7, 7},
		},
		false,
	},
	{
		[][]int{
			[]int{-3, -2},
			[]int{-1, -2},
			[]int{2, -2},
			[]int{-2, -2},
			[]int{0, -2},
		},
		true,
	},
	{
		[][]int{
			[]int{-3, -2},
		},
		true,
	},
}

func Test(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, checkStraightLine(tc.N1), "输入:%v", tc)
	}
}
