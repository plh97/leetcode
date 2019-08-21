package isCousins

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	N2  int
	N3  int
	ans bool
}{
	// {
	// 	Helper.Ints2TreeNode([]int{1, 2, 3, -1 << 63, 4, -1 << 63, 5}),
	// 	5,
	// 	4,
	// 	true,
	// },
	{
		Helper.Ints2TreeNode([]int{1, 2, 3, -1 << 63, 4}),
		2,
		3,
		false,
	},
	// {
	// 	Helper.Ints2TreeNode([]int{1, -1 << 63, 2, 3}),
	// 	4,
	// 	3,
	// 	false,
	// },
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, isCousins(tc.N1, tc.N2, tc.N3), "输入:%v", tc)
	}
}
