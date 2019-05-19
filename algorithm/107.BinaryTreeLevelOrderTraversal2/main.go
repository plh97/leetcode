package levelOrderBottom

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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	dfs(root, 0, &ans)
	return ans
}

func dfs(node *TreeNode, deep int, ans *[][]int) {
	if len(*ans) <= deep {
		*ans = append([][]int{[]int{}}, *ans...)
	}
	(*ans)[len(*ans) - deep-1] = append((*ans)[len(*ans) - deep-1], node.Val)
	if node.Left != nil {
		dfs(node.Left, deep+1, ans)
	}
	if node.Right != nil {
		dfs(node.Right, deep+1, ans)
	}
}
