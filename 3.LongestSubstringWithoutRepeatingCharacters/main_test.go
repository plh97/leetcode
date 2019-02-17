package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	str string
}

type ans struct {
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
				str: "232432",
			},
			a: ans{
				num: 3,
			},
		},
		question{
			p: para{
				str: "aaa",
			},
			a: ans{
				num: 1,
			},
		},
		question{
			p: para{
				str: "aab",
			},
			a: ans{
				num: 2,
			},
		},
		question{
			p: para{
				str: "pwwkew",
			},
			a: ans{
				num: 3,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.num, lengthOfLongestSubstring(p.str), "输入:%v", p)
	}
}
