package isSameTree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one *TreeNode
	two *TreeNode
}

type ans struct {
	one bool
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
				one: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
				two: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 0,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
					},
				},
				two: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isSameTree(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
