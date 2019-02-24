package intToRoman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	str     string
	num     int
	boolean bool
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
				one: 3,
			},
			a: ans{
				str: "III",
			},
		},

		question{
			p: para{
				one: 4,
			},
			a: ans{
				str: "IV",
			},
		},

		question{
			p: para{
				one: 9,
			},
			a: ans{
				str: "IX",
			},
		},

		question{
			p: para{
				one: 58,
			},
			a: ans{
				str: "IX",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.str, intToRoman(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
