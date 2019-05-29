package rangeSumBST

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	N2  int
	N3  int
	ans int
}{
	{
		Helper.Ints2TreeNode([]int{10, 5, 15, 3, 7, -1 << 63, 18}),
		7,
		15,
		32,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, rangeSumBST(tc.N1, tc.N2, tc.N3), "输入:%v", tc)
	}
}
