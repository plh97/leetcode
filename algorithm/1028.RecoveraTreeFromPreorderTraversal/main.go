package recoverFromPreorder

import (
	"strconv"

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

func recoverFromPreorder(S string) *TreeNode {
	SS := []string{}
	for i := 0; i < len(S); i++ {
		if i < len(S)-1 && S[i] != '-' && S[i+1] == '-' {
			SS = append(SS, S[:i+1])
			S = S[i+1:]
			i = 0
		} else if i == len(S)-1 {
			SS = append(SS, S)
		}
	}
	root := &TreeNode{}
	currentNode := root
	path := []*TreeNode{root}
	preDepth := -1
	for _, word := range SS {
	loop:
		for i, le := range word {
			if le != '-' {
				depth := i
				num := 0
				if number, err := strconv.Atoi(word[i:]); err == nil {
					num = number
				}
				dx := depth - preDepth
				if dx == 1 {
					currentNode.Left = &TreeNode{Val: int(num)}
					currentNode = currentNode.Left
				} else {
					path = path[:len(path)-1+dx]
					currentNode = path[len(path)-1]
					currentNode.Right = &TreeNode{Val: int(num)}
					currentNode = currentNode.Right
				}
				path = append(path, currentNode)
				preDepth = depth
				break loop
			}
		}
	}
	return root.Left
}
