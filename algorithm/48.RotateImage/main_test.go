package deleteDuplicates

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
		Helper.Ints2LinkList([]int{1, 1, 2}),
		Helper.Ints2LinkList([]int{1, 2}),
	},
	{
		Helper.Ints2LinkList([]int{1, 1, 2, 3, 3}),
		Helper.Ints2LinkList([]int{1, 2, 3}),
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, deleteDuplicates(tc.N1), "输入:%v", tc)
	}
}
