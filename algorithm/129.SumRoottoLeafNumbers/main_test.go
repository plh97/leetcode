package sumNumbers

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *Helper.TreeNode
	ans int
}{
	// {
	// 	Helper.Ints2TreeNode([]int{1, 2, 3}),
	// 	25,
	// },
	{
		Helper.Ints2TreeNode([]int{4, 9, 0, 5, 1}),
		1026,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, sumNumbers(tc.N1), "输入:%v", tc)
	}
}
