package bstFromPreorder

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	ans *TreeNode
}{
	{
		[]int{8, 5, 1, 7, 10, 12},
		Helper.Ints2TreeNode([]int{8, 5, 10, 1, 7, -1 << 63, 12}),
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, bstFromPreorder(tc.N1), "输入:%v", tc)
	}
}
