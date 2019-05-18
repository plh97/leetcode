package reverseOnlyLetters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one string
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
				one: "ab-cd",
			},
			a: ans{
				one: "dc-ba",
			},
		},
		question{
			p: para{
				one: "a-bC-dEf-ghIj",
			},
			a: ans{
				one: "j-Ih-gfE-dCba",
			},
		},
		question{
			p: para{
				one: "7_28]",
			},
			a: ans{
				one: "7_28]",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, reverseOnlyLetters(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
