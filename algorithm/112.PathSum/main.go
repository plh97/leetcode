package hasPathSum

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
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	return pathSum(root, 0, sum)
}

func pathSum(root *TreeNode, calc, sum int) bool {
	if root.Left != nil {
		if pathSum(root.Left, calc+root.Val, sum) {
			return true
		}
	}
	if root.Right != nil {
		if pathSum(root.Right, calc+root.Val, sum) {
			return true
		}
	}
	if root.Right == nil && root.Left == nil {
		// 根节点 p
		return root.Val+calc == sum
	}
	return false
}
