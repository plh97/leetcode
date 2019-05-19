package levelOrderBottom

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

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
	N   *TreeNode
	ans [][]int
}{
	{
		Helper.Ints2TreeNode([]int{3, 9, 20, -1 << 63, -1 << 63, 15, 7}),
		[][]int{
			[]int{15, 7},
			[]int{9, 20},
			[]int{3},
		},
	},
	{
		nil,
		[][]int{},
	},
	{
		Helper.Ints2TreeNode([]int{1, 2, -1 << 63, 3, -1 << 63, 4, -1 << 63, 5}),
		[][]int{
			[]int{5},
			[]int{4},
			[]int{3},
			[]int{2},
			[]int{1},
		},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, levelOrderBottom(tc.N), "输入:%v", tc)
	}
}

func Benchmark_bitwiseComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			levelOrderBottom(tc.N)
		}
	}
}
