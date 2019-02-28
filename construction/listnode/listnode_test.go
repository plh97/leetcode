package construction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one *ListNode
	two *ListNode
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
				one: MakeListNode([]int{2, 0, 8}),
			},
			a: ans{
				one: MakeListNode([]int{2, 0, 8}),
			},
		},
		question{
			p: para{
				one: MakeListNode([]int{0, 0, 0, 0, 0, 1}),
			},
			a: ans{
				one: MakeListNode([]int{0, 0, 0, 0, 0, 1}),
			},
		},
		question{
			p: para{
				one: MakeListNode([]int{1, 6, 4}),
			},
			a: ans{
				one: MakeListNode([]int{1, 6, 4}),
			},
		},
	};

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one.Val, p.one.Val, "输入:%v", p)
		ast.Equal(a.one.Next.Val, p.one.Next.Val, "输入:%v", p)
		ast.Equal(a.one.Next.Next.Val, p.one.Next.Next.Val, "输入:%v", p)
	}
}
