package solution

//ListNode definiton for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	jump := dummy
	l, r := head, head
	for r != nil {
		count := 0
		for r != nil && count < k {
			r = r.Next
			count++
		}
		if count == k {
			prev, cur := r, l
			for i := 0; i < k; i++ {
				prev, cur, cur.Next = cur, cur.Next, prev
			}
			jump.Next, jump, l = prev, l, r
		}
	}
	return dummy.Next
}
