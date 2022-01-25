package codetop

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	// 先把[a, b)找出来
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	last := reverse(a, b)
	// 继续从头b开始反转k个
	a.Next = reverseKGroup(b, k)
	return last
}

func reverse(a, b *ListNode) *ListNode {
	var pre, cur, next *ListNode
	pre = nil
	cur = a
	next = a
	for cur != b {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
