package subsets

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  []int
	ans [][]int
}{
	{
		[]int{1, 2, 3},
		[][]int{[]int{}, []int{1}, []int{2}, []int{2, 1}, []int{3}, []int{3, 1}, []int{3, 2}, []int{3, 2, 1}},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, subsets(tc.N1), "输入:%v", tc)
	}
}
