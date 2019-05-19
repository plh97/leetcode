package minDepth

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
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	} else if root.Left != nil && root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	return 1 + min(minDepth(root.Left), minDepth(root.Right))
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
