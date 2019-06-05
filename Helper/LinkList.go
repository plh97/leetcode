package Helper

type ListNode struct {
	Val  int
	Next *ListNode
}

// NULL 方便添加测试数据
// var NULL = -1 << 63

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2LinkList(ints []int) *ListNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &ListNode{
		Val: ints[0],
	}
	queue := make([]*ListNode, 1, n*2)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && ints[i] != NULL {
			node.Next = &ListNode{Val: ints[i]}
			queue = append(queue, node.Next)
		}
		i++
	}
	return root
}
