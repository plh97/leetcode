package strStr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one, two string
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
				one: "mississippi",
				two: "pi",
			},
			a: ans{
				one: 9,
			},
		},

		question{
			p: para{
				one: "hello",
				two: "ll",
			},
			a: ans{
				one: 2,
			},
		},

		question{
			p: para{
				one: "aaaaa",
				two: "bba",
			},
			a: ans{
				one: -1,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, strStr(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
