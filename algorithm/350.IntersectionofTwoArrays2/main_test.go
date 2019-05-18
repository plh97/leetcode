package intersect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two []int
}

type ans struct {
	one []int
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
				one: []int{1, 2, 2, 1},
				two: []int{2},
			},
			a: ans{
				one: []int{2},
			},
		},
		question{
			p: para{
				one: []int{4, 9, 5},
				two: []int{9, 4, 9, 8, 4},
			},
			a: ans{
				one: []int{4, 9},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, intersect(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
