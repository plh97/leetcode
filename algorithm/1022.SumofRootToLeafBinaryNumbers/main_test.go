package sumRootToLeaf

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	ans int
}{
	{
		Helper.Ints2TreeNode([]int{1, 0, 1, 0, 1, 0, 1}),
		22,
	},
	{
		Helper.Ints2TreeNode([]int{1}),
		1,
	},
	{
		Helper.Ints2TreeNode([]int{}),
		0,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, sumRootToLeaf(tc.N1), "输入:%v", tc)
	}
}
