package maxDepth

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one *TreeNode
	two *TreeNode
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: Helper.Ints2TreeNode([]int{3, 9, 20, -1 << 63, -1 << 63, 15, 7}),
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: Helper.Ints2TreeNode([]int{1,2}),
			},
			a: ans{
				one: 2,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, maxDepth(p.one), "input: %v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
