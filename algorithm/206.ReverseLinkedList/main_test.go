package reverseList

import (
	"testing"
	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *ListNode
	ans *ListNode
}{
	{
		Helper.Ints2LinkList([]int{1, 2, 3, 4, 5}),
		Helper.Ints2LinkList([]int{5, 4, 3, 2, 1}),
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, reverseList(tc.N1), "输入:%v", tc)
	}
}
