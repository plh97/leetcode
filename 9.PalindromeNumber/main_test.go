package isPalindrome

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
	bools bool
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
				num: 121,
			},
			a: ans{
				bools: true,
			},
		},
		question{
			p: para{
				num: 122,
			},
			a: ans{
				bools: false,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.bools, isPalindrome(p.num), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
