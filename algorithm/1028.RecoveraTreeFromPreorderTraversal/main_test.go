package recoverFromPreorder

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  string
	ans *TreeNode
}{
	{
		"1-2--3--4-5--6--7",
		Helper.Ints2TreeNode([]int{1, 2, 5, 3, 4, 6, 7}),
	},
	{
		"1-401--349---90--88",
		Helper.Ints2TreeNode([]int{1, 401, -1 << 63, 349, 88, 90, -1 << 63}),
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, recoverFromPreorder(tc.N1), "输入:%v", tc)
	}
}
