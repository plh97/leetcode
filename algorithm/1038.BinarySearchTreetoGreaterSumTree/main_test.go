package bstToGst

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	ans *TreeNode
}{
	{
		Helper.Ints2TreeNode([]int{4, 1, 6, 0, 2, 5, 7, -1 << 63, -1 << 63, -1 << 63, 3, -1 << 63, -1 << 63, -1 << 63, 8}),
		Helper.Ints2TreeNode([]int{30, 36, 21, 36, 35, 26, 15, -1 << 63, -1 << 63, -1 << 63, 33, -1 << 63, -1 << 63, -1 << 63, 8}),
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, bstToGst(tc.N1), "输入:%v", tc)
	}
}
