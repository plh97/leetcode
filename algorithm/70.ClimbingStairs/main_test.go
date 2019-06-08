package climbStairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  int
	ans int
}{
	{
		2,
		2,
	},
	{
		3,
		3,
	},
	{
		45,
		1836311903,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, climbStairs(tc.N1), "输入:%v", tc)
	}
}
