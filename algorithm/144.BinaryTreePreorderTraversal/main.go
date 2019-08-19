package preorderTraversal

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

// 		 1
// 		2 3
//  4 5 6 7
// 8

func preorderTraversal(root *TreeNode) []int {
	s := []*TreeNode{}
	res := []int{}
	for root != nil || len(s) > 0 {
		for root != nil {
			res = append(res, root.Val)
			s = append(s, root)
			root = root.Left
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]
		root = root.Right
	}
	return res
}

// func preorderTraversal(root *TreeNode) []int {
//     res:=[]int{}
//     var helper func(*TreeNode)
//     helper = func(root *TreeNode) {
//         if root == nil {
//             return
//         }
//         res = append(res, root.Val)
//         helper(root.Left)
//         helper(root.Right)
//     }
//     helper(root)
//     return res
// }
