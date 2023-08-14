package hot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy

	for head != nil && head.Next != nil {
		// 完成交换
		next := head.Next
		head.Next = next.Next
		next.Next = head
		pre.Next = next
		// 更新指针
		pre = head
		head = head.Next
	}
	return dummy.Next
}
