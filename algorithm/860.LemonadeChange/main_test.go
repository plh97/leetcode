package lemonadeChange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	ans bool
}{
	{
		[]int{5, 5, 5, 10, 20},
		true,
	},
	{
		[]int{5, 5, 10},
		true,
	},
	{
		[]int{10, 10},
		false,
	},
	{
		[]int{5, 5, 10, 10, 20},
		false,
	},
	{
		[]int{5, 5, 5, 5, 20, 20, 5, 5, 20, 5},
		false,
	},
	{
		[]int{5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5},
		true,
	},
	{
		[]int{5, 5, 5, 20, 5, 5, 5, 20, 5, 5, 5, 10, 5, 20, 10, 20, 10, 20, 5, 5},
		false,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, lemonadeChange(tc.N1), "输入:%v", tc)
	}
}
