package solution

//ListNode definiton for singly-linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	cur, count := head, 0
	for cur != nil && count < k {
		cur = cur.Next
		count++
	}
	if count == k {
		curr := reverseKGroup(cur, k)
		for count > 0 {
			head.Next, curr, head = curr, head, head.Next // 不是依此赋值，是一起赋值（注：分开赋值需要tmp变量临时保存head.Next）
			count--
		}
		head = curr
	}
	return head
}
