package Helper

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [
// 	 		       1,
// 			 2,          3,
// 	  4,    5,     6,    7,
// 	8, 9, 10,11, 12,13, 14,15
// ]
// func ArrayIntoTree(a []int) *TreeNode {
// 	if len(a) == 0 {
// 		return nil
// 	} else if (len(a)) == 1 {
// 		if a[0] == 0 {
// 			return nil
// 		} else {
// 			return &TreeNode{
// 				Val: a[0],
// 			}
// 		}
// 	} else if (len(a)) == 2 {
// 		return &TreeNode{
// 			Val:  a[0],
// 			Left: ArrayIntoTree(a[1:]),
// 		}
// 	} else {
// 		var l, r []int
// 		for i, j := 1, 0; i < len(a); j++ {
// 			level := []int{1 << uint(i-1), 2 << uint(i-1)}
// 			if i >= level[0] && i < (level[1]+level[0])/2 {
// 				l = append(l, a[i])
// 			} else {
// 				r = append(r, a[i])
// 			}
// 			i++
// 		}
// 		return &TreeNode{
// 			Val:   a[0],
// 			Left:  ArrayIntoTree(l),
// 			Right: ArrayIntoTree(r),
// 		}
// 	}
// }



// NULL 方便添加测试数据
var NULL = -1 << 63

// Ints2TreeNode 利用 []int 生成 *TreeNode
func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}