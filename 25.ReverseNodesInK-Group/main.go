package reverseKGroup

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {

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