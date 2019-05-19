package isBalanced

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
	ans bool
}{
	{
		Helper.Ints2TreeNode([]int{3, 9, 20, -1 << 63, -1 << 63, 15, 7}),
		true,
	},
	{
		Helper.Ints2TreeNode([]int{1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, -1 << 63, -1 << 63, 5, 5}),
		true,
	},
	{
		Helper.Ints2TreeNode([]int{1, 2, 2, 3, 3, -1 << 63, -1 << 63, 4, 4}),
		false,
	},
	{
		Helper.Ints2TreeNode([]int{}),
		true,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, isBalanced(tc.N), "输入:%v", tc)
	}
}

func Benchmark_bitwiseComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			isBalanced(tc.N)
		}
	}
}
