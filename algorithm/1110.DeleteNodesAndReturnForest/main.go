package delNodes

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

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	newRoot := &TreeNode{
		Val:  0,
		Left: root,
	}
	del := [1000]bool{}
	del[0] = true
	for _, e := range to_delete {
		del[e] = true
	}
	res := []*TreeNode{}
	var helper func(*TreeNode) *TreeNode
	helper = func(node *TreeNode) *TreeNode {
		// 是否是需要删除的节点,
		if node == nil {
			return nil
		}
		node.Left = helper(node.Left)
		node.Right = helper(node.Right)
		if !del[node.Val] {
			return node
		}
		if node.Left != nil {
			res = append(res, node.Left)
		}
		if node.Right != nil {
			res = append(res, node.Right)
		}

		return nil
	}

	helper(newRoot)

	return res
}

func include(nums []int, val int) bool {
	for i := range nums {
		if nums[i] == val {
			return true
		}
	}
	return false
}
