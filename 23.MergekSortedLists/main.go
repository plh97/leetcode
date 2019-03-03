package mergetwolists

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	top := res
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	for l1 != nil || l2 != nil {
		if l1 == nil {
			res.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		} else if l2 == nil {
			res.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else if l1.Val > l2.Val {
			res.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		} else {
			res.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		}
		res = res.Next
	}
	return top.Next
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
