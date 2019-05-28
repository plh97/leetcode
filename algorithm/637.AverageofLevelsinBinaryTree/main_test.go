package averageOfLevels

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	ans []float64
}{
	{
		Helper.Ints2TreeNode([]int{3, 9, 20, -1 << 63, -1 << 63, 15, 7}),
		[]float64{3, 14.5, 11},
	},
	{
		Helper.Ints2TreeNode([]int{5, 2, -3}),
		[]float64{5.0, -0.5},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, averageOfLevels(tc.N1), "输入:%v", tc)
	}
}
