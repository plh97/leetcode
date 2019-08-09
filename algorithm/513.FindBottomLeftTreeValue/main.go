package findBottomLeftValue

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

var temp [][]int = make([][]int, 999)

func findBottomLeftValue(root *TreeNode) int {
	temp = [][]int{}
	helper(root, 0)
	return temp[len(temp)-1][0]
}

func helper(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if len(temp) <= depth {
		temp = append(temp, []int{})
	}
	temp[depth] = append(temp[depth], root.Val)
	helper(root.Left, depth+1)
	helper(root.Right, depth+1)
}
