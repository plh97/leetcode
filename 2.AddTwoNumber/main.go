package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 用于保存头部
	l3 := &ListNode{}
	// 用于保存过程中的 list node
	currentL3 := l3
	add := 0
	for l1 != nil || l2 != nil || add != 0 {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		currentL3.Next = &ListNode{Val: (sum + add) % 10}
		currentL3 = currentL3.Next
		if (sum + add) > 9 {
			add = 1
		} else {
			add = 0
		}
	}
	return l3.Next
}
