package solution

// ListNode definition for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head
	var before *ListNode
	for i := 1; i < m; i++ {
		before = tmp
		tmp = tmp.Next
	}
	var p, next *ListNode
	for i := m; i <= n; i++ {
		next = tmp.Next
		tmp.Next = p
		p = tmp
		tmp = next
	}
	if before != nil {
		before.Next.Next = tmp
		before.Next = p
		return head
	}
	head.Next = tmp
	return p
}
