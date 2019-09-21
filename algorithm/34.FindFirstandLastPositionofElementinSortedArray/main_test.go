package searchRange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	N2  int
	ans []int
}{
	{
		[]int{5, 7, 7, 8, 8, 10},
		8,
		[]int{3, 4},
	},
	{
		[]int{5, 7, 7, 8, 8, 10},
		9,
		[]int{-1, -1},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, searchRange(tc.N1, tc.N2), "输入:%v", tc)
	}
}
