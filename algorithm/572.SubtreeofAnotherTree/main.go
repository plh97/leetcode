package isSubtree

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

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil {
		return false
	}
	if s.Val == t.Val && isSameTree(s, t) {
		return true
	}
	return isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func isSameTree(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a != nil && b != nil && a.Val == b.Val {
		return isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
	}
	return false
}
