package romanToInt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one int
	two bool
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
				one: "III",
			},
			a: ans{
				one: 3,
			},
		},

		question{
			p: para{
				one: "IV",
			},
			a: ans{
				one: 4,
			},
		},

		question{
			p: para{
				one: "IX",
			},
			a: ans{
				one: 9,
			},
		},

		question{
			p: para{
				one: "LVIII",
			},
			a: ans{
				one: 58,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, romanToInt(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
