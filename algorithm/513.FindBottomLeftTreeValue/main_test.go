package findBottomLeftValue

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *Helper.TreeNode
	ans int
}{
	{
		Helper.Ints2TreeNode([]int{2, 1, 3}),
		1,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, findBottomLeftValue(tc.N1), "输入:%v", tc)
	}
}
