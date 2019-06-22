package solution

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	before := ^int(^uint(0) >> 1)
	return isValid(root, &before)
}

func isValid(root *TreeNode, before *int) bool {
	if root == nil {
		return true
	}
	res := true
	res = res && isValid(root.Left, before)
	if *before < root.Val {
		*before = root.Val
		res = res && true
	} else {
		return false
	}
	res = res && isValid(root.Right, before)
	return res
}
