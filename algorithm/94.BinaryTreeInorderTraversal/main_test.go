package inorderTraversal

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	ans []int
}{
	{
		Helper.Ints2TreeNode([]int{1, -1 << 63, 2, 3}),
		[]int{1, 3, 2},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, inorderTraversal(tc.N1), "输入:%v", tc)
	}
}
