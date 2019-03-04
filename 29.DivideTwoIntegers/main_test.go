package divide

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one, two int
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
				one: math.MinInt32,
				two: -1,
			},
			a: ans{
				one: math.MaxInt32,
			},
		},
		question{
			p: para{
				one: math.MaxInt32,
				two: -1,
			},
			a: ans{
				one: math.MinInt32,
			},
		},
		question{
			p: para{
				one: -1,
				two: -1,
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: -1,
				two: 1,
			},
			a: ans{
				one: -1,
			},
		},
		question{
			p: para{
				one: 1,
				two: 1,
			},
			a: ans{
				one: 1,
			},
		},

		question{
			p: para{
				one: 7,
				two: -3,
			},
			a: ans{
				one: -2,
			},
		},

		question{
			p: para{
				one: 10,
				two: 3,
			},
			a: ans{
				one: 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, divide(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
