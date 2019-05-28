package findSecondMinimumValue

import (
	"sort"

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

func findSecondMinimumValue(root *TreeNode) int {
	c := []int{}
	helper(root, &c)
	sort.Ints(c)
	for i := 1; i < len(c); i++ {
		if c[i] == c[0] {
			c = append(c[:i-1], c[i:]...)
			i--
		}
	}
	if len(c) > 1 {
		return c[1]
	}
	return -1
}

func helper(root *TreeNode, c *[]int) {

	if root == nil {
		return
	}
	*c = append(*c, root.Val)
	helper(root.Left, c)
	helper(root.Right, c)
}
