package inorderTraversal

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
//  先序遍历, 会最先遍历根节点最左边, 然后其次是 根节点最左边的上一个节点, 然后是上一个节点的右节点的子节点的最左边节点.
func inorderTraversal(root *TreeNode) []int {
	s := []*TreeNode{}
	ans := []int{}

	for root != nil || len(s) > 0 {
		// 去到当前节点最左边节点, 并在过程中把路过节点一次推入栈中
		for root != nil {
			s = append(s, root)
			root = root.Left
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]            // 推出栈中的最后一个节点
		ans = append(ans, root.Val) // 推入站
		root = root.Right
	}
	return ans
}

// func inorderTraversal(root *TreeNode) []int {
// 	res := []int{}
// 	var help func(*TreeNode)
// 	help = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}
// 		help(root.Left)
// 		res = append(res, root.Val)
// 		help(root.Right)
// 	}
// 	help(root)
// 	return res
// }
