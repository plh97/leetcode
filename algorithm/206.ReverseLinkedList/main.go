package reverseList

import (
	"github.com/pengliheng/leetcode/Helper"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode = Helper.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head, head.Next, prev = head.Next, prev, head
	}
	return prev
}

// iteratively
// func reverseList(head *ListNode) *ListNode {
// 	var head, temp *ListNode

// 	for head != nil {
// 		// ================
// 		// head  1,2,3,4,5
// 		// head
// 		// ================
// 		temp = head
// 		head = head.Next // 往后推一格
// 		head.Next = temp
// 		fmt.Println(temp)
// 		// head  2,3,4,5
// 		// tail  1
// 		// ================
// 		// head  3,4,5
// 		// tail  2,1
// 		// ================
// 		// head  4,5
// 		// tail  3,2,1
// 		// ================
// 		// head  5
// 		// tail  4,3,2,1
// 		// ================
// 		// head
// 		// tail  5,4,3,2,1
// 		// ================
// 	}

// 	return head
// }

// func reverseList(head *ListNode) *ListNode {
// 	var tail *ListNode

// 	for head != nil {
// 		head.Next, head, tail = tail, head.Next, head
// 	}

// 	return tail
// }
