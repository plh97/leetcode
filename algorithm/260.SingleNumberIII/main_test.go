package singleNumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	ans []int
}{
	{
		[]int{1, 2, 1, 3, 2, 5},
		[]int{3, 5},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, singleNumber(tc.N1), "输入:%v", tc)
	}
}
