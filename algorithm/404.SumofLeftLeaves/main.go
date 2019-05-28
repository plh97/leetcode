package sumOfLeftLeaves

import (
	"github.com/pengliheng/leetcode/Helper"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = Helper.TreeNode

func sumOfLeftLeaves(root *TreeNode) int {
	// 如果左边是根节点 返回左边
	if root == nil || isLeaf(root) {
		return 0
	}
	if root.Left == nil {
		return sumOfLeftLeaves(root.Right)
	}
	if isLeaf(root.Left) {
		return root.Left.Val + sumOfLeftLeaves(root.Right)
	} else {
		return sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
	}
}

func isLeaf(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	return false
}
