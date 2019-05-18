package checkPossibility

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
				one: []int{2, 3, 3, 2, 4},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{1, 1, 1},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{3, 4, 2, 3},
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: []int{4, 2, 1},
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: []int{4, 2, 3},
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, checkPossibility(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
