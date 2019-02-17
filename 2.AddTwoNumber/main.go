package addTwoNumbers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// * -> 取值
// & -> 取址
func l_3(aa, bb, cc int) *ListNode {
	var c *ListNode = &ListNode{
		Val:  cc,
		Next: nil,
	}
	var b *ListNode = &ListNode{
		Val:  bb,
		Next: c,
	}
	return &ListNode{
		Val:  aa,
		Next: b,
	}
}
func l_2(aa, bb int) *ListNode {
	var b *ListNode = &ListNode{
		Val:  bb,
		Next: nil,
	}
	return &ListNode{
		Val:  aa,
		Next: b,
	}
}
func l_1(aa int) *ListNode {
	return &ListNode{
		Val:  aa,
		Next: nil,
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 用于保存头部
	l3 := &ListNode{}
	// 用于保存过程中的 list node
	currentL3 := l3
	add := 0
	for l1 != nil || l2 != nil || add!=0 {
		sum := 0
		fmt.Println(l1, l2, add)
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

