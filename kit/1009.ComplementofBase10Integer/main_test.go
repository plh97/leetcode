package bitwiseComplement

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
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
				one: 0,
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: 5,
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: 10,
			},
			a: ans{
				one: 5,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, bitwiseComplement(p.one), "输入:%v", p)
	}

}
