package lettercombinations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one []string
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
				one: "23",
			},
			a: ans{
				one: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
			},
		},

		question{
			p: para{
				one: "",
			},
			a: ans{
				one: []string{},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, letterCombinations(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
