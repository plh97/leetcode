package validMountainArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one bool
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
		// 		one: []int{3, 4},
		// 	},
		// 	a: ans{
		// 		one: false,
		// 	},
		// },
		// question{
		// 	p: para{
		// 		one: []int{4, 3, 4},
		// 	},
		// 	a: ans{
		// 		one: false,
		// 	},
		// },
		// question{
		// 	p: para{
		// 		one: []int{0, 2, 0},
		// 	},
		// 	a: ans{
		// 		one: true,
		// 	},
		// },
		// question{
		// 	p: para{
		// 		one: []int{4},
		// 	},
		// 	a: ans{
		// 		one: false,
		// 	},
		// },
		// question{
		// 	p: para{
		// 		one: []int{1, 3, 2},
		// 	},
		// 	a: ans{
		// 		one: true,
		// 	},
		// },
		// question{
		// 	p: para{
		// 		one: []int{2, 0, 2},
		// 	},
		// 	a: ans{
		// 		one: false,
		// 	},
		// },
		// question{
		// 	p: para{
		// 		one: []int{0, 1, 2, 4, 2, 1},
		// 	},
		// 	a: ans{
		// 		one: true,
		// 	},
		// },
		// question{
		// 	p: para{
		// 		one: []int{0, 1, 2, 1, 2},
		// 	},
		// 	a: ans{
		// 		one: false,
		// 	},
		// },
		question{
			p: para{
				one: []int{14, 82, 89, 84, 79, 70, 70, 68, 67, 66, 63, 60, 58, 54, 44, 43, 32, 28, 26, 25, 22, 15, 13, 12, 10, 8, 7, 5, 4, 3},
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, validMountainArray(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
