package binaryTreePaths

import (
	"strconv"

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

func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	return helper(root, "", res)
}

func helper(root *TreeNode, str string, res []string) []string {
	if root == nil {
		return res
	}
	val := strconv.Itoa(root.Val)
	if str == "" {
		str = val
	} else {
		str = str + "->" + val
	}
	if root.Left != nil {
		res = helper(root.Left, str, res)
	}
	if root.Right != nil {
		res = helper(root.Right, str, res)
	}
	if root.Right == nil && root.Left == nil {
		// leaf
		res = append(res, str)
	}
	return res
}
