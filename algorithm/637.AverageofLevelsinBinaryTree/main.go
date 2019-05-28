package averageOfLevels

import "github.com/pengliheng/leetcode/Helper"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode = Helper.TreeNode

var ress [][]int

func averageOfLevels(root *TreeNode) []float64 {
	ress = [][]int{}

	helper(root, 0)
	var res []float64
	for i, nums := range ress {
		for _, num := range nums {
			if len(res) < len(ress) {
				res = append(res, 0.0)
			}
			res[i] += float64(num)
		}
		res[i] /= float64(len(nums))
	}
	return res
}

func helper(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	if len(ress) <= depth {
		ress = append(ress, []int{})
	}
	ress[depth] = append(ress[depth], root.Val)
	helper(root.Left, depth+1)
	helper(root.Right, depth+1)
}
