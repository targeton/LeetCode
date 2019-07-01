package solution

// ListNode definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var before *ListNode
	for head.Next != nil {
		tmp := head.Next
		head.Next = before
		before = head
		head = tmp
	}
	head.Next = before
	return head
}
