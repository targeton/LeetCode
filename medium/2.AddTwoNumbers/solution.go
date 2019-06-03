package solution

/*ListNode Definition for singly-linked list*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	temp := l1.Val + l2.Val
	flag := temp / 10
	result := &ListNode{Val: temp % 10}
	p := result
	l1, l2 = l1.Next, l2.Next
	for {
		if l1 == nil {
			if l2 == nil {
				break
			}
		}
		t1, t2 := 0, 0
		if l1 != nil {
			t1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			t2, l2 = l2.Val, l2.Next
		}
		temp = t1 + t2 + flag
		p.Next = &ListNode{Val: temp % 10}
		p = p.Next
		flag = temp / 10
	}
	if flag != 0 {
		p.Next = &ListNode{Val: flag}
	}
	return result
}
