package codetop2023

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表， 递归
func reverseListRecur(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 遍历到最后一个节点
	if head.Next == nil {
		return head
	}
	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

// 非递归
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
