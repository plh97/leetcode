package isCousins

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
func isCousins(root *TreeNode, x int, y int) bool {
	xx := -1
	xxx := -1
	yy := -2
	yyy := -2
	var helper func(*TreeNode, int, int)
	helper = func(root *TreeNode, depth int, parent int) {
		if root == nil {
			return
		}
		if xx < 0 && root.Val == x {
			xx = depth
			xxx = parent
		} else if yy < 0 && root.Val == y {
			yy = depth
			yyy = parent
		}
		helper(root.Left, depth+1, root.Val)
		helper(root.Right, depth+1, root.Val)
	}
	helper(root, 0, -3)
	return xx == yy && xxx != yyy
}
