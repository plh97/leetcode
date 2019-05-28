package sumOfLeftLeaves

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

var tcs = []struct {
	N   *TreeNode
	ans int
}{
	{
		Helper.Ints2TreeNode([]int{3, 9, 20, -1 << 63, -1 << 63, 15, 7}),
		24,
	},
	{
		Helper.Ints2TreeNode([]int{}),
		0,
	},
	{
		Helper.Ints2TreeNode([]int{1, -1 << 63, 2}),
		0,
	},
	{
		Helper.Ints2TreeNode([]int{3,4,5,-7,-6,-1 <<  63,-1 <<  63,-7,-1 <<  63,-5,-1 <<  63,-1 <<  63,-1 <<  63,-4}),
		-11,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, sumOfLeftLeaves(tc.N), "输入:%v", tc)
	}
}

func Benchmark_bitwiseComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			sumOfLeftLeaves(tc.N)
		}
	}
}
