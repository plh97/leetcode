package invertTree

import (
	"testing"
	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one *TreeNode
}

type ans struct {
	one *TreeNode
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
				one: Helper.Ints2TreeNode([]int{4, 2, 7, 1, 3, 6, 9}),
			},
			a: ans{
				one: Helper.Ints2TreeNode([]int{4, 7, 2, 9, 6, 3, 1}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, invertTree(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
