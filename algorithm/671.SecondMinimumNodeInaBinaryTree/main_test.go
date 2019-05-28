package findSecondMinimumValue

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one int
}

type ans struct {
	one int
}

type question struct {
	p para
	a ans
}

var tcs = []struct {
	N1  *TreeNode
	ans int
}{
	{
		Helper.Ints2TreeNode([]int{2, 2, 5, -1 << 63, -1 << 63, 5, 7}),
		5,
	},
	{
		Helper.Ints2TreeNode([]int{2, 2, 2}),
		-1,
	},
	{
		Helper.Ints2TreeNode([]int{1, 1, 3, 1, 1, 3, 4, 3, 1, 1, 1, 3, 8, 4, 8, 3, 3, 1, 6, 2, 1}),
		2,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, findSecondMinimumValue(tc.N1), "输入:%v", tc)
	}
}
