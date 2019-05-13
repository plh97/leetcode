package isRobotBounded

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
				one: "GL",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "GGLLGG",
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isRobotBounded(p.one), "输入:%v", p)
	}

}
