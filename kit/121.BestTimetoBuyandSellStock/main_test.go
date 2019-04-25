package maxProfit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one   []int
	two   int
	three []int
	four  int
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
				one: []int{7, 1, 5, 3, 6, 4},
			},
			a: ans{
				one: 5,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, maxProfit(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
