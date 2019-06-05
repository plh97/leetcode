package swapPairs

import (
	"testing"
	"www/leetcode/Helper"

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
				one: Helper.Ints2LinkList([]int{1, 2, 3, 4}),
			},
			a: ans{
				one: Helper.Ints2LinkList([]int{2, 1, 4, 3}),
			},
		},

		question{
			p: para{
				one: Helper.Ints2LinkList([]int{1}),
			},
			a: ans{
				one: Helper.Ints2LinkList([]int{1}),
			},
		},

		question{
			p: para{
				one: Helper.Ints2LinkList([]int{}),
			},
			a: ans{
				one: Helper.Ints2LinkList([]int{}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, swapPairs(p.one), "输入:%v", p)
	}
}
