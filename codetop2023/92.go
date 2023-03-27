package codetop2023

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := new(ListNode)
	dummyHead.Next = head
	// p是要反转的第一个节点， g是要反转的节点的前一个节点
	g, p := dummyHead, head
	for i := 0; i < left-1; i++ {
		g = g.Next
		p = p.Next
	}
	// 头插法，把p的下一个节点删除，移动到g的下一个节点
	for i := 0; i < right-left; i++ {
		remove := p.Next
		p.Next = p.Next.Next
		remove.Next = g.Next
		g.Next = remove
	}
	return dummyHead.Next
}
