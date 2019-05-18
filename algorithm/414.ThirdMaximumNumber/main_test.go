package thirdMax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
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
				one: []int{-2147483648, -2147483648, -2147483648, -2147483648, 1, 1, 1},
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: []int{1, 2},
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: []int{3, 2, 1},
			},
			a: ans{
				one: 1,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, thirdMax(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
