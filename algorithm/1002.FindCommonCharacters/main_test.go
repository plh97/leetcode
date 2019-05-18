package commonChars

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []string
	two int
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
				one: []string{"bella", "label", "roller"},
			},
			a: ans{
				one: []string{"e", "l", "l"},
			},
		},
		question{
			p: para{
				one: []string{"cool", "lock", "cook"},
			},
			a: ans{
				one: []string{"c", "o"},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, commonChars(p.one), "输入:%v", p)
	}

}
