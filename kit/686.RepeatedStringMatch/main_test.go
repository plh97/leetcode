package repeatedStringMatch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one int
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
				one: "abcd",
				two: "abcdabcdabcd",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "abcd",
				two: "cdabcdab",
			},
			a: ans{
				one: 3,
			},
		},
		question{
			p: para{
				one: "aa",
				two: "a",
			},
			a: ans{
				one: 1,
			},
		},
		question{
			p: para{
				one: "abcabcabcabc",
				two: "abac",
			},
			a: ans{
				one: -1,
			},
		},
		question{
			p: para{
				one: "aaaaaaaaaaaaaaaaaaaaaab",
				two: "ba",
			},
			a: ans{
				one: 2,
			},
		},
		question{
			p: para{
				one: "abc",
				two: "cabcabca",
			},
			a: ans{
				one: 4,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, repeatedStringMatch(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
