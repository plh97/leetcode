package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
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
				one: "babad",
			},
			a: ans{
				one: "bab",
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findMedianSortedArrays(p.one, p.two), "输入:%v", p)
	}

	ast.Panics(func() { findMedianSortedArrays([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
