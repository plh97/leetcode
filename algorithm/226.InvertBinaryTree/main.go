package invertTree

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

func invertTree(root *TreeNode) *TreeNode {
	mirror(root)
	return root
}

func mirror(root *TreeNode) {
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	mirror(root.Left)
	mirror(root.Right)
}
