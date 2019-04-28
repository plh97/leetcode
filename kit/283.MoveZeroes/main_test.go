package moveZeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
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
				one: []int{2, 1},
			},
			a: ans{
				one: []int{2, 1},
			},
		},
		question{
			p: para{
				one: []int{0, 1, 0, 1},
			},
			a: ans{
				one: []int{1, 1, 0, 0},
			},
		},
		question{
			p: para{
				one: []int{1, 0, 1, 0},
			},
			a: ans{
				one: []int{1, 1, 0, 0},
			},
		},
		question{
			p: para{
				one: []int{0, 1, 0, 3, 12},
			},
			a: ans{
				one: []int{1, 3, 12, 0, 0},
			},
		},
		question{
			p: para{
				one: []int{0, 0, 1},
			},
			a: ans{
				one: []int{1, 0, 0},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, moveZeroes(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
