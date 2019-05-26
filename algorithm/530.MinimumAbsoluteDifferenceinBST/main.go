package getMinimumDifference

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
 func getMinimumDifference(root *TreeNode) int {
	res := 1 << 62
	c := []int{}
	PreOrder2TreeNode(root, &c)
	for i := 0; i < len(c)-1; i++ {
		if res > c[i+1] - c[i] {
			res = c[i+1] - c[i]
		}
	}
	return res
}

func PreOrder2TreeNode(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	PreOrder2TreeNode(root.Left, res)
	*res = append(*res, root.Val)
	PreOrder2TreeNode(root.Right, res)
}
