package convertBST

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
func convertBST(root *TreeNode) *TreeNode {

	sum := 0

	return helper(root, &sum)
}

func helper(root *TreeNode, sum *int) *TreeNode {
	if root == nil {
		return nil
	}
	
	helper(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	helper(root.Left, sum)

	return root
}
