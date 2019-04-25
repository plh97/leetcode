package searchInsert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
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
				one: []int{1, 3, 5, 6},
				two: 5,
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: []int{1, 3, 5, 6},
				two: 2,
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: []int{1, 3, 5, 6},
				two: 7,
			},
			a: ans{
				one: 4,
			},
		},
		question{
			p: para{
				one: []int{1, 3, 5, 6},
				two: 0,
			},
			a: ans{
				one: 0,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, searchInsert(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
