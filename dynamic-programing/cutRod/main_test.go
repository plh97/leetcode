package cutRod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	str, one, two      string
	row, num, n, price int
	p                  [8]int
}

type ans struct {
	str     string
	num, n  int
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
				p: [8]int{0, 1, 5, 8, 9, 10, 17, 17},
				n: 4,
			},
			a: ans{
				num: 10,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.num, cutRodDp(p.p, p.n), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
