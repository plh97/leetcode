package minDepth

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
	ans int
}{
	{
		Helper.Ints2TreeNode([]int{3, 9, 20, -1 << 63, -1 << 63, 15, 7}),
		2,
	},
	{
		Helper.Ints2TreeNode([]int{3, 9, 20, -1 << 63, -1 << 63, 15, 7}),
		2,
	},
	{
		Helper.Ints2TreeNode([]int{1, 2}),
		2,
	},
	{
		Helper.Ints2TreeNode([]int{1, -1 << 63, 2}),
		2,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, minDepth(tc.N), "输入:%v", tc)
	}
}

func Benchmark_bitwiseComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			minDepth(tc.N)
		}
	}
}
