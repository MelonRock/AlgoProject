package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	var pre, cur *ListNode
	if len(lists) == 0 {
		return nil
	}
	n := len(lists)
	for i := 0; i < n; i++ {
		if i == 0 {
			pre = lists[0]
			continue
		}
		cur = lists[i]
		pre = merge23(pre, cur)
	}
	return pre
}

func merge23(A, B *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for A != nil || B != nil {
		if A != nil && B != nil {
			if A.Val < B.Val {
				cur.Next = A
				A = A.Next
			} else {
				cur.Next = B
				B = B.Next
			}
			cur = cur.Next
		} else if A != nil {
			cur.Next = A
			break
		} else if B != nil {
			cur.Next = B
			break
		}
	}
	return head.Next
}
