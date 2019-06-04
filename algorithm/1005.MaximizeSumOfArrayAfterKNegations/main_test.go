package largestSumAfterKNegations

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	N2  int
	ans int
}{
	{
		[]int{4, 2, 3},
		1,
		5,
	},
	{
		[]int{3, -1, 0, 2},
		3,
		6,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, largestSumAfterKNegations(tc.N1, tc.N2), "输入:%v", tc)
	}
}
