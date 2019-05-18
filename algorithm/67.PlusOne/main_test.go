package addBinary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one string
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
				one: "11",
				two: "1",
			},
			a: ans{
				one: "100",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, addBinary(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
