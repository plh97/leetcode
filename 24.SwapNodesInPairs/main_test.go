package swapPairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one *ListNode
}

type ans struct {
	one *ListNode
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
				one: MakeListNode([]int{1, 2, 3, 4}),
			},
			a: ans{
				one: MakeListNode([]int{2, 1, 4, 3}),
			},
		},

		question{
			p: para{
				one: MakeListNode([]int{1}),
			},
			a: ans{
				one: MakeListNode([]int{1}),
			},
		},

		question{
			p: para{
				one: MakeListNode([]int{}),
			},
			a: ans{
				one: MakeListNode([]int{}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, swapPairs(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
