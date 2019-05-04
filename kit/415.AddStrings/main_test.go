package addStrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one string
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
				one: "111",
				two: "222",
			},
			a: ans{
				one: "333",
			},
		},

		question{
			p: para{
				one: "3876620623801494171",
				two: "6529364523802684779",
			},
			a: ans{
				one: "10405985147604178950",
			},
		},
		question{
			p: para{
				one: "999",
				two: "1",
			},
			a: ans{
				one: "1000",
			},
		},
		question{
			p: para{
				one: "1",
				two: "999",
			},
			a: ans{
				one: "1000",
			},
		},
		question{
			p: para{
				one: "408",
				two: "5",
			},
			a: ans{
				one: "413",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, addStrings(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
