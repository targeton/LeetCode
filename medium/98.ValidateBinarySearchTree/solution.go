package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var array []int

func isValidBST(root *TreeNode) bool {
	array = make([]int, 0)
	return isValid(root)
}

func isValid(root *TreeNode) bool {
	if root == nil {
		return true
	}
	res := true
	res = res && isValid(root.Left)
	if len(array) == 0 {
		res = res && true
		array = append(array, root.Val)
	} else {
		before := array[len(array)-1]
		if before < root.Val {
			res = res && true
			array = append(array, root.Val)
		} else {
			return false
		}
	}
	res = res && isValid(root.Right)
	return res
}
