package findTarget

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

// 需要一个函数, 可以递归, 遇到一个数字, 将它的答案推入map
func findTarget(root *TreeNode, k int) bool {
	Map := make(map[int]bool)
	return dfs(root, k, Map)
}
func dfs(root *TreeNode, k int, Map map[int]bool) bool {

	if root == nil {
		return false
	}

	if _, ok := Map[root.Val]; ok {
		// 找到
		return true
	}

	Map[k-root.Val] = true

	return dfs(root.Left, k, Map) || dfs(root.Right, k, Map)

}
