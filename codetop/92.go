package codetop

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := new(ListNode) // 虚拟头节点
	dummyHead.Next = head
	cur := dummyHead.Next // 当前节点
	pre := dummyHead      // 前一个节点
	for i := 0; i < left-1; i++ {
		cur = cur.Next // 移动到要反转的节点上
		pre = pre.Next
	}
	// 头插法，把cur.next移动到前面
	for i := 0; i < right-left; i++ {
		removeNode := cur.Next
		cur.Next = cur.Next.Next
		removeNode.Next = pre.Next
		pre.Next = removeNode
	}
	return dummyHead.Next

}
