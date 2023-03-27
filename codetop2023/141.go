package codetop2023

func hasCycle(head *ListNode) bool {
	fast := &ListNode{}
	slow := &ListNode{}
	fast, slow = head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
