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
	if m > 1 {
		head.Next = reverseBetween(head.Next, m-1, n-1)
		return head
	} else if n > 1 {
		node := reverseBetween(head.Next, m-1, n-1)
		tmp := head.Next
		head.Next = tmp.Next
		tmp.Next = head
		return node
	} else {
		return head
	}
}
