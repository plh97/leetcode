package diameterOfBinaryTree

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
		Helper.Ints2TreeNode([]int{1, 2, 3, 4, 5}),
		3,
	},
	{
		Helper.Ints2TreeNode([]int{1}),
		0,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, diameterOfBinaryTree(tc.N), "输入:%v", tc)
	}
}

func Benchmark_bitwiseComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			diameterOfBinaryTree(tc.N)
		}
	}
}
