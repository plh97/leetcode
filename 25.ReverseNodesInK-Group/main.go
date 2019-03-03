package reverseKGroup

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	arr := []int{}
	res := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	if len(arr) < k {
		return MakeListNode(arr)
	}
	// TODO:
	// 提供reverse函数
	for i := 0; i < len(arr); i += k {
		currArr := []int{}
		if i+k > len(arr) {
			// 
			currArr = arr[i:]
			res = append(res, currArr...)
		} else if len(arr) < k {
			currArr = arr
			res = append(res, reverseGroup(currArr)...)
		} else {
			currArr = arr[i : i+k]
			res = append(res, reverseGroup(currArr)...)
		}
	}
	return MakeListNode(res)
}

// 翻转数组
func reverseGroup(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return arr
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
