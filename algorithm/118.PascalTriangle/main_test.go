package generate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one [][]int
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
				one: 5,
			},
			a: ans{
				one: [][]int{
					[]int{1},
					[]int{1, 1},
					[]int{1, 2, 1},
					[]int{1, 3, 3, 1},
					[]int{1, 4, 6, 4, 1},
				},
			},
		},
		question{
			p: para{
				one: 0,
			},
			a: ans{
				one: [][]int{},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, generate(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
