package searchBST

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
// 二叉搜索树的查找
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	currentVal := root.Val
	if val < currentVal {
		return searchBST(root.Left, val)
	} else if val > currentVal {
		return searchBST(root.Right, val)
	} else {
		return root
	}
}
