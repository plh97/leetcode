package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	str string
	row int
}

type ans struct {
	str string
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
				str: "PAYPALISHIRING",
				row: 3,
			},
			a: ans{
				str: "PAHNAPLSIIGYIR",
			},
		},
		question{
			p: para{
				str: "PAYPALISHIRING",
				row: 4,
			},
			a: ans{
				str: "PINALSIGYAHRPI",
			},
		},
		question{
			p: para{
				str: "PAYPALISHIRING",
				row: 1,
			},
			a: ans{
				str: "PAYPALISHIRING",
			},
		},
		question{
			p: para{
				str: "PA",
				row: 2,
			},
			a: ans{
				str: "PA",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.str, convert(p.str,p.row), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
