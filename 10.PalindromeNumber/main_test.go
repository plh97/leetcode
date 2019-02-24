package isMatch

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
				one: "mississippi",
				two: "mis*is*p*.",
			},
			a: ans{
				boolean: false,
			},
		},
		question{
			p: para{
				one: "aaa",
				two: "a*a",
			},
			a: ans{
				boolean: true,
			},
		},

		question{
			p: para{
				one: "a",
				two: ".*..a*",
			},
			a: ans{
				boolean: false,
			},
		},
		question{
			p: para{
				one: "aab",
				two: "c*a*b",
			},
			a: ans{
				boolean: true,
			},
		},
		question{
			p: para{
				one: "",
				two: ".",
			},
			a: ans{
				boolean: false,
			},
		},

		question{
			p: para{
				one: "a",
				two: "ab*a",
			},
			a: ans{
				boolean: false,
			},
		},

		question{
			p: para{
				one: "aa",
				two: "a*",
			},
			a: ans{
				boolean: true,
			},
		},

		question{
			p: para{
				one: "a",
				two: "",
			},
			a: ans{
				boolean: false,
			},
		},
		question{
			p: para{
				one: "1111",
				two: "22222",
			},
			a: ans{
				boolean: false,
			},
		},
		question{
			p: para{
				one: "ab",
				two: ".*c",
			},
			a: ans{
				boolean: false,
			},
		},
		question{
			p: para{
				one: "",
				two: ".*",
			},
			a: ans{
				boolean: true,
			},
		},
		question{
			p: para{
				one: "mississippi",
				two: "mis*is*ip*.",
			},
			a: ans{
				boolean: true,
			},
		},
		question{
			p: para{
				one: "ab",
				two: ".*",
			},
			a: ans{
				boolean: true,
			},
		},
		question{
			p: para{
				one: "aaa",
				two: "aaaa",
			},
			a: ans{
				boolean: false,
			},
		},
		question{
			p: para{
				one: "ab",
				two: "a.",
			},
			a: ans{
				boolean: true,
			},
		},
		question{
			p: para{
				one: "aa",
				two: "a",
			},
			a: ans{
				boolean: false,
			},
		},
		question{
			p: para{
				one: "aa",
				two: "a*",
			},
			a: ans{
				boolean: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.boolean, isMatch(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
