package removeElement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
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

		// question{
		// 	p: para{
		// 		one: []int{3, 2, 2, 3},
		// 		two: 2,
		// 	},
		// 	a: ans{
		// 		one: 2,
		// 	},
		// },

		question{
			p: para{
				one: []int{0, 1, 2, 2, 3, 0, 4, 2},
				two: 2,
			},
			a: ans{
				one: 5,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, removeElement(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
