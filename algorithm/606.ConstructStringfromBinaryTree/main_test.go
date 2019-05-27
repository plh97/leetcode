package tree2str

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	ans string
}{
	{
		Helper.Ints2TreeNode([]int{1, 2, 3, 4}),
		"1(2(4))(3)",
	},
	{
		Helper.Ints2TreeNode([]int{1, 2, 3, -1 << 63, 4}),
		"1(2()(4))(3)",
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, tree2str(tc.N1), "输入:%v", tc)
	}
}
