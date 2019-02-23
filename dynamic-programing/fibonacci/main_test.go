package fibDp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	str, one, two string
	row, num      int
}

type ans struct {
	str     string
	num     int
	boolean bool
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
				num: 5,
			},
			a: ans{
				num: 5,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.num, fibDp(p.num), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
