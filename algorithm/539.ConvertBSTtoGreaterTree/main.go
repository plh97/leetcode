package diameterOfBinaryTree

import (
	"github.com/pengliheng/leetcode/Helper"
)

type TreeNode = Helper.TreeNode

var res int = 0

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	helper(root)
	return res
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := helper(root.Left), helper(root.Right)
	res = max(res, l+r)
	return 1 + max(l, r)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
