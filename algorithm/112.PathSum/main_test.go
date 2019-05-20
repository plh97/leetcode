package hasPathSum

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
	one *TreeNode
	two int
}

var tcs = []struct {
	N   question
	ans bool
}{
	{
		question{
			one: Helper.Ints2TreeNode([]int{5, 4, 8, 11, 13, 4, 7, 2, 1}),
			two: 22,
		},
		true,
	},
	{
		question{
			one: Helper.Ints2TreeNode([]int{}),
			two: 22,
		},
		false,
	},
	{
		question{
			one: Helper.Ints2TreeNode([]int{1,2}),
			two: 0,
		},
		false,
	},
	{
		question{
			one: Helper.Ints2TreeNode([]int{5, 4, 8, 11, -1 << 63, 13, 4, 7, 2, -1 << 63, -1 << 63, -1 << 63, 1}),
			two: 22,
		},
		true,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, hasPathSum(tc.N.one, tc.N.two), "输入:%v", tc)
	}
}

func Benchmark_bitwiseComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			hasPathSum(tc.N.one, tc.N.two)
		}
	}
}
