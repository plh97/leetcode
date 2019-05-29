package leafSimilar

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	N2  *TreeNode
	ans bool
}{
	{
		Helper.Ints2TreeNode([]int{3, 5, 1, 6, 2, 9, 8, -1 << 63, -1 << 63, 7, 4}),
		Helper.Ints2TreeNode([]int{3, 5, 1, 6, 7, 4, 2, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, 9, 8}),
		true,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, leafSimilar(tc.N1, tc.N2), "输入:%v", tc)
	}
}
