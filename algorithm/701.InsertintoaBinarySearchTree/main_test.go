package insertIntoBST

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	N2  int
	ans *TreeNode
}{
	{
		Helper.Ints2TreeNode([]int{4, 2, 7, 1, 3}),
		5,
		Helper.Ints2TreeNode([]int{4, 2, 7, 1, 3, 5}),
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, insertIntoBST(tc.N1, tc.N2), "输入:%v", tc)
	}
}
