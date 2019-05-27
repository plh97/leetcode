package isSubtree

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
	N1  *TreeNode
	N2  *TreeNode
	ans bool
}{
	{
		Helper.Ints2TreeNode([]int{3, 4, 5, 1, 2}),
		Helper.Ints2TreeNode([]int{3, 4, 5}),
		false,
	},
	{
		Helper.Ints2TreeNode([]int{1}),
		Helper.Ints2TreeNode([]int{0}),
		false,
	},
	{
		Helper.Ints2TreeNode([]int{-1, -4, 8, -6, -2, 3, 9, -1 << 63, -5, -1 << 63, -1 << 63, 0, 7}),
		Helper.Ints2TreeNode([]int{3, 0, 5848}),
		false,
	},
	{
		Helper.Ints2TreeNode([]int{3, 4, 5, 1, 2, -1 << 63, -1 << 63, 0}),
		Helper.Ints2TreeNode([]int{4, 1, 2}),
		false,
	},
	{
		Helper.Ints2TreeNode([]int{1, 1}),
		Helper.Ints2TreeNode([]int{1}),
		true,
	},
	{
		Helper.Ints2TreeNode([]int{1, 2, 3}),
		Helper.Ints2TreeNode([]int{1, 2}),
		false,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, isSubtree(tc.N1, tc.N2), "输入:%v", tc)
	}
}

// func Benchmark_bitwiseComplement(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, tc := range tcs {
// 			isSubtree(tc.N)
// 		}
// 	}
// }
