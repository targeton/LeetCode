package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var intMax, intMin int

func isValidBST(root *TreeNode) bool {
	intMax = int(^uint(0) >> 1)
	intMin = ^int(^uint(0) >> 1)
	return isValid(root, intMin, intMax)
}

func isValid(node *TreeNode, lower int, upper int) bool {
	if node == nil {
		return true
	}
	if lower != intMin && node.Val < lower {
		return false
	}
	if upper != intMax && node.Val > upper {
		return false
	}
	if !isValid(node.Left, lower, node.Val) {
		return false
	}
	if !isValid(node.Right, node.Val, upper) {
		return false
	}
	return true
}
