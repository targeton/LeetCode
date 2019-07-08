/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 *
 * https://leetcode.com/problems/binary-tree-level-order-traversal/description/
 *
 * algorithms
 * Medium (49.05%)
 * Likes:    1559
 * Dislikes: 43
 * Total Accepted:    394.4K
 * Total Submissions: 804K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * Given a binary tree, return the level order traversal of its nodes' values.
 * (ie, from left to right, level by level).
 *
 *
 * For example:
 * Given binary tree [3,9,20,null,null,15,7],
 *
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 *
 *
 *
 * return its level order traversal as:
 *
 * [
 * ⁠ [3],
 * ⁠ [9,20],
 * ⁠ [15,7]
 * ]
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	helper(&res, root, 0)
	return res
}

func helper(res *[][]int, node *TreeNode, height int) {
	if node == nil {
		return
	}
	if len(*res) == height {
		*res = append(*res, []int{})
	}
	(*res)[height] = append((*res)[height], node.Val)
	helper(res, node.Left, height+1)
	helper(res, node.Right, height+1)
}
