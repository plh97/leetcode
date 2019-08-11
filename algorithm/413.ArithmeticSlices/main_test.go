package topKFrequent

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
		[]int{1, 1, 1, 2, 2, 3},
		2,
		[]int{1, 2},
	},
	{
		[]int{1},
		1,
		[]int{1},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, topKFrequent(tc.N1, tc.N2), "输入:%v", tc)
	}
}
