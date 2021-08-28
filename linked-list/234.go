package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
var left *ListNode

func isPalindrome(head *ListNode) bool {
	left = head
	return traverse234(head)
}

func traverse234(right *ListNode) bool {
	// 后续遍历，然后左右指针比较
	if right == nil {
		return true
	}
	res := traverse234(right.Next)
	res = res && (right.Val == left.Val)
	left = left.Next
	return res
}
