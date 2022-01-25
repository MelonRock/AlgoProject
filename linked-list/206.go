package linked_list

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}
	// 递归调用，翻转第二个节点开始往后的链表
	last := reverseList(head.Next)
	// 翻转头节点与第二个节点的指向
	head.Next.Next = head
	// 此时的 head 节点为尾节点，next 需要指向 NULL
	head.Next = nil
	return last
}
