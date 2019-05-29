package isUnivalTree

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
		Helper.Ints2TreeNode([]int{1, 1, 1, 1, 1, -1 << 63, 1}),
		true,
	},
	{
		Helper.Ints2TreeNode([]int{2, 2, 2, 5, 2}),
		false,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, isUnivalTree(tc.N1), "输入:%v", tc)
	}
}
