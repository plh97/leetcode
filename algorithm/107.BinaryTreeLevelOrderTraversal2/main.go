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
	// DFS
	res := [][]int{[]int{}}
	queue := []*TreeNode{root} // 首个进入, 最后一个出来
	seen := map[int]int{root.Val: root.Val}
	seen[root.Val] = root.Val
	for len(queue) > 0 {
		vertex := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		// 所有下一个可能的节点
		nodes := []*TreeNode{vertex.Left, vertex.Right}
		// 选择一个下一个节点,找到未访问过的节点, 推入queue
		for i := range nodes {
			node := nodes[i]
			if node == nil {
				continue
			}
			_, ok := seen[node.Val]
			if !ok {
				// 未曾访问过的节点
				queue = append([]*TreeNode{node}, queue...)
				seen[node.Val] = node.Val
			}
		}
		for i, j := len(res)-1, 0; i >= 0; i-- {
			// 第i层
			if len(res[i]) < 1<<uint(j) {
				// 当前没满则推入, vertex推出到当前i
				res[i] = append(res[i], vertex.Val)
				break
			} else if len(res[i]) == 1<<uint(j) && i == 0 {
				// 当前满了,且是最后一层
				res = append([][]int{
					[]int{vertex.Val},
				}, res...)
				break
			}
			j++
		}
	}
	return res
}
