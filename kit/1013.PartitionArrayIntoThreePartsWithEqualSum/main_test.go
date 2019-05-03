package canThreePartsEqualSum

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
		question{
			p: para{
				one: []int{14, 6, -10, 2, 18, -7, -4, 11},
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: []int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1},
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: []int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1},
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: []int{18, 12, -18, 18, -19, -1, 10, 10},
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, canThreePartsEqualSum(p.one), "输入:%v", p)
	}

}
