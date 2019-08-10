package largestValues

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

func largestValues(root *TreeNode) []int {
	temp = [][]int{}
	helper(root, 0)
	res := []int{}
	for _, e := range temp {
		res = append(res, findMax(e))
	}
	return res
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

func findMax(A []int) int {
	temp := 0
	for i, e := range A {
		if A[temp] < e {
			temp = i
		}
	}
	return A[temp]
}
