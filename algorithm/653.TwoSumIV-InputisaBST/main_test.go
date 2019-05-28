package findTarget

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
	N2  int
	ans bool
}{
	{
		Helper.Ints2TreeNode([]int{5, 3, 6, 2, 4, -1 << 63, 7}),
		9,
		true,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, findTarget(tc.N1, tc.N2), "输入:%v", tc)
	}
}
