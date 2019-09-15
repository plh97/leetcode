package deleteDuplicates

import (
	"github.com/pengliheng/leetcode/Helper"
)

type ListNode = Helper.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	arr := []int{}
	for head != nil {
		// remove duplicate
		if lastIndexOf(arr, head.Val) == -1 {
			arr = append(arr, head.Val)
		}
		head = head.Next
	}
	node := &ListNode{Val: 0}
	newHead := node
	for len(arr) > 0 {
		node.Next = &ListNode{Val: arr[0]}
		node = node.Next
		arr = arr[1:]
	}
	return newHead.Next
}

func lastIndexOf(nums []int, val int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == val {
			return i
		}
	}
	return -1
}
