package solution

// ListNode definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var cur, prev *ListNode
	cur, prev = head, nil
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
	}
	return prev
}
