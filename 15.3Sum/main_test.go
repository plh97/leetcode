package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one [][]int
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
				one: []int{-1, 0, 1, 2, -1, -4},
			},
			a: ans{
				one: [][]int{
					[]int{-1, 0, 1},
					[]int{-1, -1, 2},
				},
			},
		},

		question{
			p: para{
				one: []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0},
			},
			a: ans{
				one: [][]int{
					[]int{-4, 1, 3},
					[]int{-4, 0, 4},
					[]int{-2, 1, 1},
					[]int{-2, -2, 4},
					[]int{-5, 1, 4},
					[]int{0, 0, 0},
				},
			},
		},

		question{
			p: para{
				one: []int{0, 0, 0, 0},
			},
			a: ans{
				one: [][]int{
					[]int{0, 0, 0},
				},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, threeSum(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
