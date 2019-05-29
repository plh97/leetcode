package sumRootToLeaf

import "github.com/pengliheng/leetcode/Helper"

type TreeNode = Helper.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}
	val = val*2 + root.Val

	if root.Left == root.Right {
		return val
	} else {
		return helper(root.Left, val) + helper(root.Right, val)
	}
}
