package swapPairs

import "fmt"

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
	_head := head
	for _head != nil {
		arr = append(arr, _head.Val)
		_head = _head.Next
	}
	fmt.Println(arr)
	// swap if->arr=e=>e%2
	for i := 0; i < len(arr)/2; i ++ {

		arr[2*i], arr[2*i+1] = arr[2*i+1], arr[2*i]
	}
	fmt.Println(arr)

	return MakeListNode(arr)

}
