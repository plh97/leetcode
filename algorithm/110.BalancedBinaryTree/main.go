package isBalanced

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

// 计算最深, 最浅, max-min<=1 即可??
func isBalanced(root *TreeNode) bool {
	_, isBalanced := getDepthAndBalance(root)
	return isBalanced
}

func getDepthAndBalance(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftDepth, ok1 := getDepthAndBalance(root.Left)
	rightDepth, ok2 := getDepthAndBalance(root.Right)

	if !ok1 || !ok2 {
		return 0, false
	}

	if abs(leftDepth-rightDepth) > 1 {
		return 0, false
	}

	return max(leftDepth, rightDepth) + 1, true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
