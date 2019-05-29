package leafSimilar

import (
	"reflect"

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
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	c1, c2 := []int{}, []int{}
	helper(root1, &c1)
	helper(root2, &c2)
	return reflect.DeepEqual(c1, c2)
}

func helper(root *TreeNode, c *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*c = append(*c, root.Val)
	}
	helper(root.Left, c)
	helper(root.Right, c)
}
