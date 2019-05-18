package mergeklists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []*ListNode
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
				one: []*ListNode{
					MakeListNode([]int{1, 4, 5}),
					MakeListNode([]int{1, 3, 4}),
					MakeListNode([]int{2, 6}),
				},
			},
			a: ans{
				one: MakeListNode([]int{1, 1, 2, 3, 4, 4, 5, 6}),
			},
		},

		question{
			p: para{
				one: []*ListNode{
					MakeListNode([]int{2}),
					MakeListNode([]int{}),
					MakeListNode([]int{-1}),
				},
			},
			a: ans{
				one: MakeListNode([]int{-1, 2}),
			},
		},

		question{
			p: para{
				one: []*ListNode{
					MakeListNode([]int{}),
					MakeListNode([]int{-1, 5, 11}),
					MakeListNode([]int{}),
					MakeListNode([]int{6, 10}),
				},
			},
			a: ans{
				one: MakeListNode([]int{-1, 5, 6, 10, 11}),
			},
		},

		question{
			p: para{
				one: []*ListNode{
					MakeListNode([]int{}),
				},
			},
			a: ans{
				one: MakeListNode([]int{}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, mergeKLists(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
