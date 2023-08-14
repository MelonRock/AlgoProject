package codetop2023

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 876. 链表的中间结点
func middleListNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 206. 反转链表
func reverse143(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := reverse143(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

/*
L0 → L1 → … Ln/2 → Ln-1 → Ln
请将其重新排列后变为：
L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → … → Ln/2
*/
func reorderList(head *ListNode) {
	middle := middleListNode(head)
	head2 := reverse143(middle)
	for head2.Next != nil {
		nxt := head.Next
		nxt2 := head2.Next
		head.Next = head2
		head2.Next = nxt
		head = nxt
		head2 = nxt2
	}
}
