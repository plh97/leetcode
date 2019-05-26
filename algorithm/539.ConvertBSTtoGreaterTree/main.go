package diameterOfBinaryTree

import (
	"github.com/pengliheng/leetcode/Helper"
)

type TreeNode = Helper.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return 0
	}
	res := 0
	res = 1 + max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right))
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
