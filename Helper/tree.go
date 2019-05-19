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
func ArrayIntoTree(a []int) *TreeNode {
	if len(a) == 0 {
		return nil
	} else if (len(a)) == 1 {
		if a[0] == 0 {
			return nil
		} else {
			return &TreeNode{
				Val: a[0],
			}
		}
	} else if (len(a)) == 2 {
		return &TreeNode{
			Val:  a[0],
			Left: ArrayIntoTree(a[1:]),
		}
	} else {
		var l, r []int
		for i := 1; i < len(a); i++ {
			level := []int{1 << uint(i), 2 << uint(i)}
			if i > level[0] && i < level[1]/2 {
				l = append(l, a[i])
			} else {
				r = append(r, a[i])
			}
		}
		return &TreeNode{
			Val:   a[0],
			Left:  ArrayIntoTree(l),
			Right: ArrayIntoTree(r),
		}
	}
}
