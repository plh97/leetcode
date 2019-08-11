package numEnclaves

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  [][]int
	ans int
}{
	// {
	// 	[][]int{
	// 		[]int{0,0,0,0},
	// 		[]int{1,0,1,0},
	// 		[]int{0,1,1,0},
	// 		[]int{0,0,0,0},
	// 	},
	// 	3,
	// },
	// {
	// 	[][]int{
	// 		[]int{0, 1, 1, 0},
	// 		[]int{0, 0, 1, 0},
	// 		[]int{0, 0, 1, 0},
	// 		[]int{0, 0, 0, 0},
	// 	},
	// 	0,
	// },
	{
		[][]int{
			[]int{0, 0, 0, 1, 1, 1, 0, 1, 0, 0, },
			[]int{1, 1, 0, 0, 0, 1, 0, 1, 1, 1, },
			[]int{0, 0, 0, 1, 1, 1, 0, 1, 0, 0, },
			[]int{0, 1, 1, 0, 0, 0, 1, 0, 1, 0, },
			[]int{0, 1, 1, 1, 1, 1, 0, 0, 1, 0, },
			[]int{0, 0, 1, 0, 1, 1, 1, 1, 0, 1, },
			[]int{0, 1, 1, 0, 0, 0, 1, 1, 1, 1, },
			[]int{0, 0, 1, 0, 0, 1, 0, 1, 0, 1, },
			[]int{1, 0, 1, 0, 1, 1, 0, 0, 0, 0, },
			[]int{0, 0, 0, 0, 1, 1, 0, 0, 0, 1, },
		},
		3,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, numEnclaves(tc.N1), "输入:%v", tc)
	}
}
