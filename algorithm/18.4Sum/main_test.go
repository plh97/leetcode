package foursum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
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
				one: []int{1, 0, -1, 0, -2, 2},
				two: 0,
			},
			a: ans{
				one: [][]int{
					[]int{-2, -1, 1, 2},
					[]int{-2, 0, 0, 2},
					[]int{-1, 0, 0, 1},
				},
			},
		},

		question{
			p: para{
				one: []int{-3, -2, -1, 0, 0, 1, 2, 3},
				two: 0,
			},
			a: ans{
				one: [][]int{
					[]int{-3, -2, 2, 3},
					[]int{-3, -1, 1, 3},
					[]int{-3, 0, 0, 3},
					[]int{-3, 0, 1, 2},
					[]int{-2, -1, 0, 3},
					[]int{-2, -1, 1, 2},
					[]int{-2, 0, 0, 2},
					[]int{-1, 0, 0, 1},
				},
			},
		},

		question{
			p: para{

				one: []int{-5, 5, 4, -3, 0, 0, 4, -2},
				two: 4,
			},
			a: ans{
				one: [][]int{
					[]int{-5, 0, 4, 5},
					[]int{-3, -2, 4, 5},
				},
			},
		},

		question{
			p: para{

				one: []int{0, 0, 0, 0},
				two: 0,
			},
			a: ans{
				one: [][]int{
					[]int{0, 0, 0, 0},
				},
			},
    },
    

		question{
			p: para{

				one: []int{1,-2,-5,-4,-3,3,3,5},
				two: -11,
			},
			a: ans{
				one: [][]int{
					[]int{-5, -4, -3, 1},
				},
			},
    },
    
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, fourSum(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
