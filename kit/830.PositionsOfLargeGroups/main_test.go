package largeGroupPositions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
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
				one: "babaaaabbb",
			},
			a: ans{
				one: [][]int{
					[]int{3, 6},
					[]int{7, 9},
				},
			},
		},
		question{
			p: para{
				one: "bababbabaa",
			},
			a: ans{
				one: [][]int{},
			},
		},
		question{
			p: para{
				one: "aa",
			},
			a: ans{
				one: [][]int{},
			},
		},
		question{
			p: para{
				one: "abbxxxxzzy",
			},
			a: ans{
				one: [][]int{
					[]int{3, 6},
				},
			},
		},
		question{
			p: para{
				one: "abc",
			},
			a: ans{
				one: [][]int{},
			},
		},
		question{
			p: para{
				one: "aaa",
			},
			a: ans{
				one: [][]int{
					[]int{0, 2},
				},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, largeGroupPositions(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
