package sumNumbers

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
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	helper(root, &res, 0)
	return res
}

func helper(node *TreeNode, res *int, sum int) {
	sum = 10*sum + node.Val
	if node.Left == nil && node.Right == nil {
		// lofe
		*res += sum
	} else {
		if node.Left != nil {
			helper(node.Left, res, sum)
		}
		if node.Right != nil {
			helper(node.Right, res, sum)
		}
	}
}

//        4
//    9       0
//  5   1         4

// 495  491  40

// 4
// 49+40

// 495 + 491 + 40
