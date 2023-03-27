package codetop2023

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	// 先找出[a,b)
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

func reverse(a *ListNode, b *ListNode) *ListNode {
	var pre *ListNode
	cur := a
	for cur != b {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// 前k大数字
// k个一组反转
// 字符串转数字
// go并发编程
// 前缀和
// 二叉树序列化
