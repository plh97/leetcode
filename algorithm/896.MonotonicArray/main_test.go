package isMonotonic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
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
				one: []int{1, 2, 2, 3},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{6, 5, 4, 4},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{1, 3, 2},
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isMonotonic(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
