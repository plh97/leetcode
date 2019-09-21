package isValidBST

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	ans bool
}{
	{
		Helper.Ints2TreeNode([]int{5, 1, 4, -1 << 63, -1 << 63, 3, 6}),
		false,
	},
	{
		Helper.Ints2TreeNode([]int{}),
		true,
	},
	{
		Helper.Ints2TreeNode([]int{1, 1}),
		false,
	},
	{
		Helper.Ints2TreeNode([]int{3, 1, 5, 0, 2, 4, 6}),
		true,
	},
	{
		Helper.Ints2TreeNode([]int{10, 5, 15, -1 << 63, -1 << 63, 6, 20}),
		false,
	},
	{
		Helper.Ints2TreeNode([]int{3, -1 << 63, 30, 10, -1 << 63, -1 << 63, 15, -1 << 63, 45}),
		false,
	},
	{
		Helper.Ints2TreeNode([]int{2147483647}),
		true,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, isValidBST(tc.N1), "输入:%v", tc)
	}
}
