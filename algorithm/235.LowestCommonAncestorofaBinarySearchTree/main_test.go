package binaryTreePaths

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

var tcs = []struct {
	N   *TreeNode
	ans []string
}{
	{
		Helper.Ints2TreeNode([]int{1, 2, 3, -1 << 63, 5}),
		[]string{
			"1->2->5",
			"1->3",
		},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, binaryTreePaths(tc.N), "输入:%v", tc)
	}
}

func Benchmark_bitwiseComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			binaryTreePaths(tc.N)
		}
	}
}
