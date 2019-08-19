package kthSmallest

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

var min []int

func kthSmallest(root *TreeNode, k int) int {
	min = []int{}
	helper(root, k)
	return min[k-1]
}

func helper(root *TreeNode, k int) {
	if root == nil {
		return
	}
	helper(root.Left, k-1)
	min = append(min, root.Val)
	helper(root.Right, k-1)
}
