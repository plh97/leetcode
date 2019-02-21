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
				num: 123,
			},
			a: ans{
				num: 321,
			},
		},
		question{
			p: para{
				num: -123,
			},
			a: ans{
				num: -321,
			},
		},
		question{
			p: para{
				num: 120,
			},
			a: ans{
				num: 21,
			},
		},
		question{
			p: para{
				num: 1234,
			},
			a: ans{
				num: 4321,
			},
		},
		question{
			p: para{
				num: 10,
			},
			a: ans{
				num: 1,
			},
		},
		question{
			p: para{
				num: -10,
			},
			a: ans{
				num: -1,
			},
		},
		question{
			p: para{
				num: 1534236469,
			},
			a: ans{
				num: 0,
			},
		},
		// question{
		// 	p: para{
		// 		num: -2147483412,
		// 	},
		// 	a: ans{
		// 		num: -2143847412,
		// 	},
		// },
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.num, reverse(p.num), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
