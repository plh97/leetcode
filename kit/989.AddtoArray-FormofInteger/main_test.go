package addToArrayForm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
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
				one: []int{1, 2, 0, 0},
				two: 34,
			},
			a: ans{
				one: []int{1, 2, 3, 4},
			},
		},
		question{
			p: para{
				one: []int{0},
				two: 0,
			},
			a: ans{
				one: []int{0},
			},
		},
		question{
			p: para{
				one: []int{2, 7, 4},
				two: 181,
			},
			a: ans{
				one: []int{4, 5, 5},
			},
		},
		question{
			p: para{
				one: []int{9, 9, 9},
				two: 1,
			},
			a: ans{
				one: []int{1, 0, 0, 0},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, addToArrayForm(p.one, p.two), "输入:%v", p)
	}

}
