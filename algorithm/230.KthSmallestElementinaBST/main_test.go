package kthSmallest

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	N2  int
	ans int
}{
	{
		Helper.Ints2TreeNode([]int{3, 1, 4, -1 << 63, 2}),
		1,
		1,
	},
	{
		Helper.Ints2TreeNode([]int{5, 3, 6, 2, 4, -1 << 63, -1 << 63, 1}),
		3,
		3,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, kthSmallest(tc.N1, tc.N2), "输入:%v", tc)
	}
}
