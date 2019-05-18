package nextPermutation

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
				one: []int{1, 2, 3},
			},
			a: ans{
				one: []int{1, 3, 2},
			},
		},
		question{
			p: para{
				one: []int{3, 2, 1},
			},
			a: ans{
				one: []int{1, 2, 3},
			},
		},
		question{
			p: para{
				one: []int{1, 1, 5},
			},
			a: ans{
				one: []int{1, 5, 1},
			},
		},
		question{
			p: para{
				one: []int{1, 5, 1},
			},
			a: ans{
				one: []int{5, 1, 1},
			},
		},
		question{
			p: para{
				one: []int{1, 5, 8, 4, 7, 6, 5, 3, 1},
			},
			a: ans{
				one: []int{1, 5, 8, 5, 1, 3, 4, 6, 7},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, nextPermutation(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
