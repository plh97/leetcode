package increasingBST

import (
	"fmt"

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

func increasingBST(root *TreeNode) *TreeNode {
	res := helper(root, nil)
	fmt.Println(res)
	return res
}
func helper(root, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}
	res := helper(root.Left, root)
	root.Left = nil
	root.Right = helper(root.Right, tail)
	return res
}
