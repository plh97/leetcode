package nextGreatestLetter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []byte
	two byte
}

type ans struct {
	one byte
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
				one: []byte{'c', 'f', 'j'},
				two: 'a',
			},
			a: ans{
				one: 'c',
			},
		},
		question{
			p: para{
				one: []byte{'c', 'f', 'j'},
				two: 'j',
			},
			a: ans{
				one: 'c',
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, nextGreatestLetter(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
