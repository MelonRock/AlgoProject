package sword

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast := head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	if k > 0 {
		return nil
	}
	slow := head
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
