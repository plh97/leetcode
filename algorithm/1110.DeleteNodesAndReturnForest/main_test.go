package delNodes

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	N2  []int
	ans []*TreeNode
}{
	//   1
	//  2 3
	//4,5 6,7
	//
	{
		Helper.Ints2TreeNode([]int{1, 2, 3, 4, 5, 6, 7}),
		[]int{3, 5},
		[]*TreeNode{
			Helper.Ints2TreeNode([]int{6}),
			Helper.Ints2TreeNode([]int{7}),
			Helper.Ints2TreeNode([]int{1, 2, -1 << 63, 4}),
		},
	},
	{
		Helper.Ints2TreeNode([]int{1, 2, -1 << 63, 4, 3}),
		[]int{2, 3},
		[]*TreeNode{
			Helper.Ints2TreeNode([]int{4}),
			Helper.Ints2TreeNode([]int{1}),
		},
	},
	{
		Helper.Ints2TreeNode([]int{1, 2, 3, -1 << 63, -1 << 63, -1 << 63, 4}),
		[]int{2, 1},
		[]*TreeNode{
			Helper.Ints2TreeNode([]int{3, -1 << 63, 4}),
		},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, delNodes(tc.N1, tc.N2), "输入:%v", tc)
	}
}
