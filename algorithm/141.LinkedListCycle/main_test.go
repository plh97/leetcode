package hasCycle

import (
	"testing"
	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *ListNode
	ans bool
}{
	{
		Helper.Ints2ListWithCycle([]int{3, 2, 0, -4}, 1),
		true,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, hasCycle(tc.N1), "输入:%v", tc)
	}
}
