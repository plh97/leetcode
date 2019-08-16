package relativeSortArray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	N2  []int
	ans []int
}{
	{
		[]int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
		[]int{2, 1, 4, 3, 9, 6},
		[]int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, relativeSortArray(tc.N1, tc.N2), "输入:%v", tc)
	}
}
