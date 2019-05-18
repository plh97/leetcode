package Helper

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ArrayIntoTree(a, b []int) *TreeNode {
	root := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	return root
}
