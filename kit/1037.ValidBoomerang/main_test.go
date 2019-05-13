package isBoomerang

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]int
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
		question{
			p: para{
				one: [][]int{
					[]int{0, 0},
					[]int{1, 2},
					[]int{0, 1},
				},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: [][]int{
					[]int{0, 1},
					[]int{1, 1},
					[]int{2, 1},
				},
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: [][]int{
					[]int{1, 1},
					[]int{2, 3},
					[]int{3, 2},
				},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: [][]int{
					[]int{0, 0},
					[]int{1, 0},
					[]int{2, 2},
				},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: [][]int{
					[]int{1, 0},
					[]int{1, 0},
					[]int{2, 2},
				},
			},
			a: ans{
				one: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isBoomerang(p.one), "输入:%v", p)
	}

}
