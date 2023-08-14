package codetop2023

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head.Next = head.Next.Next
			}
			// 需要删除第一个重复的节点
			pre.Next = head.Next
		} else {
			pre = pre.Next
		}
		head = head.Next
	}
	return dummy.Next
}
