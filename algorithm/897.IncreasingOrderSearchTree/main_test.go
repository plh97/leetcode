package increasingBST

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	ans *TreeNode
}{
	// {
	// 	Helper.Ints2TreeNode([]int{383, 453, -1 << 63, -1 << 63, 469, 619, 867, 723, 752}),
	// 	Helper.Ints2TreeNode([]int{1, -1 << 63, 2, -1 << 63, 3, -1 << 63, 4, -1 << 63, 5, -1 << 63, 6, -1 << 63, 7, -1 << 63, 8, -1 << 63, 9}),
	// },
	{
		Helper.Ints2TreeNode([]int{5, 3, 6, 2, 4, -1 << 63, 8, 1, -1 << 63, -1 << 63, -1 << 63, 7, 9}),
		Helper.Ints2TreeNode([]int{1, -1 << 63, 2, -1 << 63, 3, -1 << 63, 4, -1 << 63, 5, -1 << 63, 6, -1 << 63, 7, -1 << 63, 8, -1 << 63, 9}),
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, increasingBST(tc.N1), "输入:%v", tc)
	}
}
