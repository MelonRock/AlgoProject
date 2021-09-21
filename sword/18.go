package sword

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}
	l := head
	cur := head.Next
	for cur != nil && cur.Val != val {
		l = cur
		cur = cur.Next
	}
	if cur != nil {
		l.Next = cur.Next
	}
	return head
}
