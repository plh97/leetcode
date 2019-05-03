package numPairsDivisibleBy60

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one int
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
				one: []int{30, 20, 150, 100, 40},
			},
			a: ans{
				one: 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, numPairsDivisibleBy60(p.one), "输入:%v", p)
	}

}
