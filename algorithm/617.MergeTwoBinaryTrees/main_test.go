package mergeTrees

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	N2  *TreeNode
	ans *TreeNode
}{
	{
		Helper.Ints2TreeNode([]int{1, 3, 2, 5}),
		Helper.Ints2TreeNode([]int{2, 1, 3, -1 << 63, 4, -1 << 63, 7}),
		Helper.Ints2TreeNode([]int{3, 4, 5, 5, 4, -1 << 63, 7}),
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, mergeTrees(tc.N1, tc.N2), "输入:%v", tc)
	}
}
