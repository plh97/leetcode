package validPalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
}

type ans struct {
	one bool
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
				one: "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "ebcbbececabbacecbbcbe",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "eeccccbebaeeabebccceea",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "abc",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "aba",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "cbbcc",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "abca",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "a",
			},
			a: ans{
				one: true,
			},
		},
		// question{
		// 	p: para{
		// 	},
		// 	a: ans{
		// 		one: true,
		// 	},
		// },
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, validPalindrome(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
