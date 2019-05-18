package isSymmetric

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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBranchSame(root.Left, root.Right)
}

func isBranchSame(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	if p.Val != q.Val {
		return false
	}
	return p.Val == q.Val && isBranchSame(q.Left, p.Right) && isBranchSame(p.Left, q.Right)
}
