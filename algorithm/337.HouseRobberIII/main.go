package rob

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
func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	val := root.Val
	if root.Left != nil {
		val += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		val += rob(root.Right.Left) + rob(root.Right.Right)
	}

	temp := rob(root.Left) + rob(root.Right)

	if val > temp {
		return val
	}
	return temp
}
