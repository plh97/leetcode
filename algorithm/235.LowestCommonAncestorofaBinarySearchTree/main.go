package lowestCommonAncestor

import (
	"github.com/pengliheng/leetcode/Helper"
)

type TreeNode = Helper.TreeNode

/**
 * Definition for TreeNode.
 * type TreeNode struct {
	*     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
*/
// O(1)
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if p.Val > q.Val {
// 		q, p = p, q
// 	}
// 	for {
// 		if root.Val < p.Val {
// 			root = root.Right
// 		}else if root.Val > q.Val {
// 			root = root.Left
// 		}else {
// 			return root
// 		}
// 	}
// }

// Recursive
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		q, p = p, q
	}
	if root == nil {
		return nil
	}
	if root.Val < p.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}
