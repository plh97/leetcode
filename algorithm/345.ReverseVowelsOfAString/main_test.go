package reverseVowels

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
				one: "hello",
			},
			a: ans{
				one: "holle",
			},
		},
		question{
			p: para{
				one: "aA",
			},
			a: ans{
				one: "Aa",
			},
		},
		question{
			p: para{
				one: "leetcode",
			},
			a: ans{
				one: "leotcede",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, reverseVowels(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
