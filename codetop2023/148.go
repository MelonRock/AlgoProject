package codetop2023

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	// 递归结束条件
	if head == nil || head.Next == nil {
		return head
	}
	// 快慢指针找到中点
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 断开的链表的第一个节点
	tmp := slow.Next
	// 断开慢指针
	slow.Next = nil
	// 排序新的两个链表，按照大小依次加入新的链表
	left := sortList(head)
	right := sortList(tmp)
	// 头部的下个节点
	h := new(ListNode)
	res := h
	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next = left
			left = left.Next
		} else {
			h.Next = right
			right = right.Next
		}
		h = h.Next
	}
	h.Next = func() *ListNode {
		if left != nil {
			return left
		}
		return right
	}()

	return res.Next
}
