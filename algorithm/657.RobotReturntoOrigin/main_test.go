package judgeCircle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
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
				one: "UD",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "LL",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "RL",
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, judgeCircle(p.one), "输入:%v", p)
	}

}
