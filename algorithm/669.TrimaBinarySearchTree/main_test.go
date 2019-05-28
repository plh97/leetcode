package trimBST

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	N2  int
	N3  int
	ans *TreeNode
}{
	{
		Helper.Ints2TreeNode([]int{1, 0, 2}),
		1,
		2,
		Helper.Ints2TreeNode([]int{1, -1 << 63, 2}),
	},
	{
		Helper.Ints2TreeNode([]int{3, 0, 4, -1 << 63, 2, -1 << 63, -1 << 63, 1}),
		1,
		3,
		Helper.Ints2TreeNode([]int{3, 2, -1 << 63, 1}),
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, trimBST(tc.N1, tc.N2, tc.N3), "输入:%v", tc)
	}
}
