package merge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one   []int
	two   int
	three []int
	four  int
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
				one:   []int{4, 5, 6, 0, 0, 0},
				two:   3,
				three: []int{1, 2, 3},
				four:  3,
			},
			a: ans{
				one: []int{1, 2, 3, 4, 5, 6},
			},
		},
		question{
			p: para{
				one:   []int{1, 2, 3, 0, 0, 0},
				two:   3,
				three: []int{2, 5, 6},
				four:  3,
			},
			a: ans{
				one: []int{1, 2, 2, 3, 5, 6},
			},
		},
		question{
			p: para{
				one:   []int{1},
				two:   1,
				three: []int{1},
				four:  0,
			},
			a: ans{
				one: []int{1},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, merge(p.one, p.two, p.three, p.four), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
