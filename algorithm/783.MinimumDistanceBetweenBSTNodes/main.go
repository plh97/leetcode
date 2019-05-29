package minDiffInBST

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
func minDiffInBST(root *TreeNode) int {
	res := 1 << 62
	c := []int{}
	bsf(root, &c)
	for i := 1; i < len(c); i++ {
		res = min(res, c[i]-c[i-1])
	}
	return res
}

func bsf(root *TreeNode, c *[]int) {
	if root == nil {
		return
	}
	bsf(root.Left, c)
	*c = append(*c, root.Val)
	bsf(root.Right, c)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
