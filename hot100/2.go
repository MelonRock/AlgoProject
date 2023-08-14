package hot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cur := new(ListNode)
	head := cur
	carry := 0
	for l1 != nil || l2 != nil {
		n1 := func() int {
			if l1 != nil {
				return l1.Val
			}
			return 0
		}()
		n2 := func() int {
			if l2 != nil {
				return l2.Val
			}
			return 0
		}()
		sum := n1 + n2 + carry
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: 1}
	}
	return head.Next
}
