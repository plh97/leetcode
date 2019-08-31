package bstFromPreorder

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
func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	for i, e := range preorder {
		if i == 0 {
			continue
		}
		current := root
		for {
			if current.Val > e {
				if current.Left == nil {
					current.Left = &TreeNode{Val: e}
					break
				} else {
					current = current.Left
				}
			} else {
				if current.Right == nil {
					current.Right = &TreeNode{Val: e}
					break
				} else {
					current = current.Right
				}
			}
		}
	}
	return root
}
