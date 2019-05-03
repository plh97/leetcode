package sortArrayByParityII

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
}

type ans struct {
	one []int
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
				one: []int{697, 894, 952, 686, 899, 650, 507, 510, 695, 510, 331, 533, 433, 909, 624, 206, 554, 852, 653, 939},
			},
			a: ans{
				one: []int{894, 697, 952, 899, 686, 507, 650, 695, 510, 331, 510, 533, 206, 909, 624, 433, 554, 653, 852, 939},
			},
		},
		question{
			p: para{
				one: []int{4, 2, 5, 7},
			},
			a: ans{
				one: []int{4, 5, 2, 7},
			},
		},
		question{
			p: para{
				one: []int{2, 3, 1, 1, 4, 0, 0, 4, 3, 3},
			},
			a: ans{
				one: []int{2, 3, 0, 1, 4, 1, 0, 3, 4, 3},
			},
		},
		question{
			p: para{
				one: []int{3, 4},
			},
			a: ans{
				one: []int{4, 3},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, sortArrayByParityII(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
