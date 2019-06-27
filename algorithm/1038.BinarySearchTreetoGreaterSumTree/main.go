package bstToGst

import (
	"github.com/pengliheng/leetcode/Helper"
)

type TreeNode = Helper.TreeNode

func bstToGst(root *TreeNode) *TreeNode {
	helper(root, 0)
	return root
}

func helper(root *TreeNode, preVal int) int {
	if root == nil {
		return preVal
	}
	res := helper(root.Right, preVal)
	root.Val += res
	return helper(root.Left, root.Val)
}
