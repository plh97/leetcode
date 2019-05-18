package powerfulIntegers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one   int
	two   int
	three int
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
		// question{
		// 	p: para{
		// 		one:   2,
		// 		two:   3,
		// 		three: 10,
		// 	},
		// 	a: ans{
		// 		one: []int{2, 3, 4, 5, 7, 9, 10},
		// 	},
		// },
		// question{
		// 	p: para{
		// 		one:   3,
		// 		two:   5,
		// 		three: 15,
		// 	},
		// 	a: ans{
		// 		one: []int{2, 4, 6, 8, 10, 14},
		// 	},
		// },
		question{
			p: para{
				one:   2,
				two:   1,
				three: 10,
			},
			a: ans{
				one: []int{2, 3,5,9},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, powerfulIntegers(p.one, p.two, p.three), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
