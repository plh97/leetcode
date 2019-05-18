package matrixReshape

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
					[]int{1, 2},
					[]int{3, 4},
				},
				two:   1,
				three: 4,
			},
			a: ans{
				one: [][]int{
					[]int{1, 2, 3, 4},
				},
			},
		},
		question{
			p: para{
				one: [][]int{
					[]int{1, 2},
					[]int{3, 4},
				},
				two:   2,
				three: 4,
			},
			a: ans{
				one: [][]int{
					[]int{1, 2},
					[]int{3, 4},
				},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, matrixReshape(p.one, p.two, p.three), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
