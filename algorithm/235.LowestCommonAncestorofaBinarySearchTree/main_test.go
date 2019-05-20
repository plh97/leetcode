package lowestCommonAncestor

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	one   *TreeNode
	two   *TreeNode
	three *TreeNode
}

var tcs = []struct {
	N   question
	ans *TreeNode
}{
	{
		question{
			Helper.IntsPreOrder2TreeNode([]int{-10, -3, 0, 5, 9}),
			&TreeNode{
				Val: 2,
			},
			&TreeNode{
				Val: 8,
			},
		},
		&TreeNode{
			Val: 5,
		},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, lowestCommonAncestor(tc.N.one, tc.N.two, tc.N.three), "输入:%v", tc)
	}
}

func Benchmark_bitwiseComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			lowestCommonAncestor(tc.N.one, tc.N.two, tc.N.three)
		}
	}
}
