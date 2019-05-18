package countPrimes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
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
				one: 10,
			},
			a: ans{
				one: 4,
			},
		},
		question{
			p: para{
				one: 5,
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: 2,
			},
			a: ans{
				one: 0,
			},
		},
		question{
			p: para{
				one: 3,
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: 499979,
			},
			a: ans{
				one: 41537,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, countPrimes(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
