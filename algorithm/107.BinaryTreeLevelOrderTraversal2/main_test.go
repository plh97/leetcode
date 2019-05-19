package levelOrderBottom

import (
	"testing"
	"www/leetcode/Helper"

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
	N   *TreeNode
	ans [][]int
}{
	{
		&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		},
		[][]int{
			[]int{15, 7},
			[]int{9, 20},
			[]int{3},
		},
	},
	{
		nil,
		[][]int{},
	},
	{
		Helper.ArrayIntoTree([]int{1, 2, nil, 3, nil, 4, nil, 5}),
		[][]int{
			[]int{15, 7},
			[]int{9, 20},
			[]int{3},
		},
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, levelOrderBottom(tc.N), "输入:%v", tc)
	}
}

func Benchmark_bitwiseComplement(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs {
			levelOrderBottom(tc.N)
		}
	}
}
