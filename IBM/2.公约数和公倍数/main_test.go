package multiple

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
				one: 2,
				two: 45,
			},
			a: ans{
				LCM: 1,
				GCD: 90,
			},
		},
		question{
			p: para{
				one: 4,
				two: 6,
			},
			a: ans{
				LCM: 2,
				GCD: 12,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal([2]int{a.LCM , a.GCD}, multiple(p.one,p.two), "输入:%v", p)
	}

	// ast.Panics(func() { multiple(p.one,p.two) }, "对空切片求中位数，却没有panic")

}
