package transpose

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one   [][]int
	two   int
	three int
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
				one: [][]int{
					[]int{1, 2, 3},
					[]int{4, 5, 6},
					[]int{7, 8, 9},
				},
			},
			a: ans{
				one: [][]int{
					[]int{1, 4, 7},
					[]int{2, 5, 8},
					[]int{3, 6, 9},
				},
			},
		},
		question{
			p: para{
				one: [][]int{
					[]int{1, 2, 3},
					[]int{4, 5, 6},
				},
			},
			a: ans{
				one: [][]int{
					[]int{1, 4},
					[]int{2, 5},
					[]int{3, 6},
				},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, transpose(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
