package getIntersectionNode

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var a *Helper.ListNode = Helper.Ints2LinkList([]int{4, 1, 8, 4, 5})
var tcs = []struct {
	N1  *Helper.ListNode
	N2  *Helper.ListNode
	ans *Helper.ListNode
}{
	{
		Helper.Ints2LinkList([]int{4, 1, 8, 4, 5}),
		Helper.Ints2LinkList([]int{5, 0, 1, 8, 4, 5}),
		nil,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, getIntersectionNode(tc.N1, tc.N2), "输入:%v", tc)
	}
}
