package DFS

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one map[string][]string
	two string
}

type ans struct {
	one []string
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
				one: graph,
				two: "A",
			},
			a: ans{
				one: []string{"A", "C", "E", "D", "F", "B"},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, DFS(p.one, p.two), "输入:%v", p)
	}
}
