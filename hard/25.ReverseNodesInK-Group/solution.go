package solution

//ListNode definiton for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	var newHead, prev *ListNode
	l, cur, tail := k, head, &ListNode{Val: 0, Next: head}
	for cur != nil {
		cur.Next, prev, cur = prev, cur, cur.Next
		l--
		if l == 0 {
			if newHead == nil {
				newHead = prev
			}
			tail, tail.Next, prev = tail.Next, prev, nil
			tail.Next = cur
			l = k
		}
	}
	if l < k {
		cur, prev = prev, nil
		for prev != tail.Next {
			cur.Next, prev, cur = prev, cur, cur.Next
		}
	}
	if newHead == nil {
		return prev
	}
	return newHead
}
