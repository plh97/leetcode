package myAtoi

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
		question{
			p: para{
				str: "  -42",
			},
			a: ans{
				num: -42,
			},
		},
		question{
			p: para{
				str: "4193 with words",
			},
			a: ans{
				num: 4193,
			},
		},
		question{
			p: para{
				str: "sswords and 987",
			},
			a: ans{
				num: 0,
			},
		},
		question{
			p: para{
				str: "3.14159",
			},
			a: ans{
				num: 3,
			},
		},
		question{
			p: para{
				str: "+1",
			},
			a: ans{
				num: 1,
			},
		},
		question{
			p: para{
				str: "  -0012a42",
			},
			a: ans{
				num: -12,
			},
		},
		question{
			p: para{
				str: "  -0k4",
			},
			a: ans{
				num: 0,
			},
		},
		question{
			p: para{
				str: "  -5-",
			},
			a: ans{
				num: -5,
			},
		},
		question{
			p: para{
				str: "24234324234",
			},
			a: ans{
				num: 2147483647,
			},
		},
		question{
			p: para{
				str: "-24234324234",
			},
			a: ans{
				num: -2147483648,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.num, myAtoi(p.str), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
