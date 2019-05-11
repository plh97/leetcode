package selfDividingNumbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
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
				one: 1,
				two: 22,
			},
			a: ans{
				one: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, selfDividingNumbers(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
