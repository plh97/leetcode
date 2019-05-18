package repeat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	str string
	row int
	num int
	one int
	two int
}

type ans struct {
	str string
	num int
	LCM int
	GCD int
	obj map[string]int
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
				str: "aaabbccc",
			},
			a: ans{
				obj: map[string]int{"a": 3, "b": 2, "c": 3},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.obj, repeat(p.str), "输入:%v", p)
	}

	// ast.Panics(func() { multiple(p.one,p.two) }, "对空切片求中位数，却没有panic")

}
