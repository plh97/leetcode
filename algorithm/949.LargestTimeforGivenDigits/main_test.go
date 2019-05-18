package largestTimeFromDigits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one string
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
				one: []int{2, 0, 6, 6},
			},
			a: ans{
				one: "06:26",
			},
		},
		question{
			p: para{
				one: []int{1, 2, 3, 4},
			},
			a: ans{
				one: "23:41",
			},
		},
		question{
			p: para{
				one: []int{0, 0, 0, 0},
			},
			a: ans{
				one: "00:00",
			},
		},
		question{
			p: para{
				one: []int{0, 4, 0, 0},
			},
			a: ans{
				one: "04:00",
			},
		},
		question{
			p: para{
				one: []int{5, 5, 5, 5},
			},
			a: ans{
				one: "",
			},
		},
		question{
			p: para{
				one: []int{0, 0, 0, 2},
			},
			a: ans{
				one: "20:00",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, largestTimeFromDigits(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
