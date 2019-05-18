package rotate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
}

type ans struct {
	one []int
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
				one: []int{1, 2, 3, 4, 5, 6, 7},
				two: 3,
			},
			a: ans{
				one: []int{5, 6, 7, 1, 2, 3, 4},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, rotate(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
