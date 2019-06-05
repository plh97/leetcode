package swapPairs

import (
	"www/leetcode/Helper"
)

// Definition for singly-linked list.

type ListNode = Helper.ListNode


// MakeListNode 将连续结构的数组转化成 -> 单链结构数据的 -> 取址
func MakeListNode(is []int) *ListNode {
	if len(is) == 0 {
		return nil
	}
	res := &ListNode{
		Val: is[0],
	}
	temp := res
	for i := 1; i < len(is); i++ {
		temp.Next = &ListNode{Val: is[i]}
		temp = temp.Next
	}
	return res
}

func swapPairs(head *ListNode) *ListNode {
	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	for i := 0; i < len(arr)/2; i++ {
		arr[2*i], arr[2*i+1] = arr[2*i+1], arr[2*i]
	}
	return MakeListNode(arr)
}
