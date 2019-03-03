package removenthfromend

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	len := 0
	len2 := 0
	currentNode := head // checked
	for currentNode != nil {
		len++
		currentNode = currentNode.Next
	}
	currentNode = head // checked
	for len2 < len-n-1 {
		len2++
		currentNode = currentNode.Next
	}
	if len == n {
		return head.Next
	} else if currentNode.Next != nil {
		currentNode.Next = currentNode.Next.Next
	} else {
		return nil
	}
	return head
}

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
