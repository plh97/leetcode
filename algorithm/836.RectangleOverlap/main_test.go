package isRectangleOverlap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two []int
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
				one: []int{
					0, 0, 2, 2,
				},
				two: []int{
					1, 1, 3, 3,
				},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{
					2, 17, 6, 20,
				},
				two: []int{
					3, 8, 6, 20,
				},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{
					0, 0, 1, 1,
				},
				two: []int{
					1, 0, 2, 1,
				},
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isRectangleOverlap(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
