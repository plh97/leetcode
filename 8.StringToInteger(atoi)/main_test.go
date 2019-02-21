package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	str string
	row int
	num int
}

type ans struct {
	str string
	num int
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
				str: "123",
			},
			a: ans{
				num: 123,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.num, myAtoi(p.str), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
