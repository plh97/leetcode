package sortedArrayToBST

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
	N   []int
	ans *TreeNode
}{
	{
		[]int{-10, -3, 0, 5, 9},
		Helper.IntsPreOrder2TreeNode([]int{0, -3, 9, -10, 5}),
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, sortedArrayToBST(tc.N), "输入:%v", tc)
	}
}

func Benchmark_bitwiseComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sortedArrayToBST(tc.N)
		}
	}
}
